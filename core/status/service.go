package status

import (
	"sort"
	"sync"
	"time"

	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pin"
	"github.com/dgtony/collections/hashset"
	"github.com/dgtony/collections/queue"
	ct "github.com/dgtony/collections/time"
	"github.com/textileio/go-threads/core/net"
	"github.com/textileio/go-threads/core/thread"
)

var log = logging.Logger("anytype-mw-status")

const (
	threadStatusUpdatePeriod     = 5 * time.Second
	threadStatusEventBatchPeriod = 2 * time.Second
	profileInformationLifetime   = 30 * time.Second
	cafeLastPullTimeout          = 10 * time.Minute

	// truncate device names and account IDs to last symbols
	maxNameLength = 8
)

type LogTime struct {
	AccountID string
	DeviceID  string
	LastEdit  int64
}

type Service interface {
	Watch(thread.ID, func() []string) (new bool)
	Unwatch(thread.ID)
	UpdateTimeline(thread.ID, []LogTime)

	Start() error
	Stop()
}

var _ Service = (*service)(nil)

type service struct {
	tInfo       net.SyncInfo
	fInfo       pin.FilePinService
	profile     core.ProfileInfo
	ownDeviceID string
	cafeID      string

	watchers map[thread.ID]func()
	threads  map[thread.ID]*threadStatus

	// deviceID => { thread.ID }
	devThreads map[string]hashset.HashSet
	// deviceID => accountID
	devAccount map[string]string
	// peerID => connected
	connMap map[string]bool

	tsTrigger *queue.BulkQueue
	emitter   func(event *pb.Event)
	mu        sync.Mutex
}

func NewService(
	ts net.SyncInfo,
	fs pin.FilePinService,
	profile core.ProfileInfo,
	emitter func(event *pb.Event),
	cafe string,
	device string,
) *service {
	return &service{
		tInfo:       ts,
		fInfo:       fs,
		profile:     profile,
		emitter:     emitter,
		cafeID:      cafe,
		ownDeviceID: device,
		watchers:    make(map[thread.ID]func()),
		threads:     make(map[thread.ID]*threadStatus),
		devThreads:  make(map[string]hashset.HashSet),
		devAccount:  make(map[string]string),
		connMap:     make(map[string]bool),
		tsTrigger:   queue.NewBulkQueue(threadStatusEventBatchPeriod, 5, 2),
	}
}

func (s *service) Watch(tid thread.ID, fList func() []string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.watchers[tid]; exist {
		// send current status to init stateless caller
		s.tsTrigger.Push(tid)
		return false
	}

	var (
		stop   = make(chan struct{})
		ticker = ct.NewRightAwayTicker(threadStatusUpdatePeriod)
		closer = func() { close(stop); ticker.Stop() }
	)

	s.watchers[tid] = closer

	go func() {
		for {
			select {
			case <-ticker.C:
			case <-stop:
				return
			}

			var (
				tStat, _ = s.tInfo.View(tid)
				pStat    = s.fInfo.PinStatus(fList()...)
			)

			s.mu.Lock()
			ts := s.getThreadStatus(tid)
			s.mu.Unlock()

			ts.Lock()
			for pid, status := range tStat {
				ts.UpdateStatus(pid.String(), status)
			}
			for cid, info := range pStat {
				ts.UpdateFiles(cid, info)
			}
			var modified = ts.modified
			ts.Unlock()

			if modified && !s.tsTrigger.PushTimeout(tid, 2*time.Second) {
				log.Warn("unable to submit thread status notification for more than 2 seconds")
			}
		}
	}()
	return true
}

func (s *service) Unwatch(tid thread.ID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if stop, found := s.watchers[tid]; found {
		delete(s.watchers, tid)
		stop()
	}
}

func (s *service) UpdateTimeline(tid thread.ID, timeline []LogTime) {
	if len(timeline) == 0 {
		return
	}

	s.mu.Lock()
	for _, logTime := range timeline {
		// update account information for devices
		s.devAccount[logTime.DeviceID] = logTime.AccountID

		// update device threads
		dt, exist := s.devThreads[logTime.DeviceID]
		if !exist {
			dt = hashset.New()
			s.devThreads[logTime.DeviceID] = dt
		}
		dt.Add(tid)
	}
	ts := s.getThreadStatus(tid)
	s.mu.Unlock()

	ts.Lock()
	for _, logTime := range timeline {
		ts.UpdateTimeline(logTime.DeviceID, logTime.LastEdit)
	}
	var modified = ts.modified
	ts.Unlock()

	if modified && !s.tsTrigger.PushTimeout(tid, 2*time.Second) {
		log.Warn("unable to submit timeline update notification for more than 2 seconds")
	}
}

//func (s *service) ThreadSummary() net.SyncSummary {
//	ps, _ := s.tInfo.PeerSummary(s.cafe)
//	return ps
//}
//
//func (s *service) FileSummary() core.FilePinSummary {
//	return s.fInfo.FileSummary()
//}

func (s *service) Start() error {
	if err := s.startConnectivityTracking(); err != nil {
		return err
	}
	go s.startSendingThreadStatus()
	return nil
}

func (s *service) Stop() {
	s.tsTrigger.Stop()

	s.mu.Lock()
	defer s.mu.Unlock()

	// just shutdown all thread status watchers, connectivity tracking
	// will be stopped automatically on closing the network layer
	for tid, stop := range s.watchers {
		delete(s.watchers, tid)
		stop()
	}
}

func (s *service) startConnectivityTracking() error {
	connEvents, err := s.tInfo.Connectivity()
	if err != nil {
		return err
	}

	go func() {
		for event := range connEvents {
			var (
				devID = event.Peer.String()
				ts    = make(map[thread.ID]*threadStatus)
			)

			s.mu.Lock()
			// update peer connectivity
			s.connMap[devID] = event.Connected

			// find threads shared with peer
			if tids, exist := s.devThreads[devID]; exist {
				for _, i := range tids.List() {
					var tid = i.(thread.ID)
					ts[tid] = s.getThreadStatus(tid)
				}
			}
			s.mu.Unlock()

			for tid, t := range ts {
				t.Lock()
				t.UpdateConnectivity(devID, event.Connected)
				var modified = t.modified
				t.Unlock()

				if modified && !s.tsTrigger.PushTimeout(tid, 2*time.Second) {
					log.Warn("unable to submit connectivity update notification for more than 2 seconds")
				}
			}
		}
	}()

	return nil
}

func (s *service) startSendingThreadStatus() {
	var (
		profile        core.Profile
		profileUpdated time.Time
	)

	for is := range s.tsTrigger.RunBulk() {
		var ts = make(map[thread.ID]*threadStatus, len(is))

		s.mu.Lock()
		for i := 0; i < len(is); i++ {
			id := is[i].(thread.ID)
			ts[id] = s.getThreadStatus(id)
		}
		s.mu.Unlock()

		if now := time.Now(); now.Sub(profileUpdated) > profileInformationLifetime {
			profileUpdated = now
			if updated, err := s.profile.LocalProfile(); err != nil {
				log.Errorf("unable to get local profile: %v", err)
			} else {
				profile = updated
			}
		}

		for id, t := range ts {
			event := s.constructEvent(t, profile)
			s.sendEvent(
				id.String(),
				&pb.EventMessageValueOfThreadStatus{ThreadStatus: &event},
			)
		}
	}
}

// Unsafe, use under the global lock!
func (s *service) getThreadStatus(tid thread.ID) *threadStatus {
	ts, exist := s.threads[tid]
	if !exist {
		ts = newThreadStatus(func(devID string) bool {
			// deadlock-safe b/c connectivity resolve should
			// never be running under the global mutex
			s.mu.Lock()
			defer s.mu.Unlock()
			return s.connMap[devID]
		})
		s.threads[tid] = ts
	}
	return ts
}

func (s *service) constructEvent(ts *threadStatus, profile core.Profile) pb.EventStatusThread {
	type devInfo struct {
		id string
		ds deviceStatus
	}

	var (
		accounts = make(map[string][]devInfo)
		cafe     deviceStatus
		dss      []deviceStatus
		event    = pb.EventStatusThread{
			Summary: &pb.EventStatusThreadSummary{},
			Cafe:    &pb.EventStatusThreadCafe{},
		}

		max = func(x, y int64) int64 {
			if x > y {
				return x
			}
			return y
		}

		shorten = func(name string) string {
			if len(name) <= maxNameLength {
				return name
			}
			return name[len(name)-maxNameLength:]
		}
	)

	ts.Lock()
	s.mu.Lock()

	// construct account tree
	for devID, status := range ts.devices {
		if devID == s.cafeID {
			cafe = *status
			continue
		} else if devID == s.ownDeviceID {
			// do not include own device status
			continue
		}

		if accID, found := s.devAccount[devID]; found {
			accountDevices := accounts[accID]
			accounts[accID] = append(accountDevices, devInfo{devID, *status})
		} // omit devices with unmatched account
	}

	// clear modification status
	ts.modified = false

	s.mu.Unlock()
	ts.Unlock()

	// accounts
	for accID, devices := range accounts {
		var accountInfo = pb.EventStatusThreadAccount{Id: shorten(accID)}
		if accID == profile.AccountAddr {
			accountInfo.Name = profile.Name
			accountInfo.ImageHash = profile.IconImage
		}

		for _, device := range devices {
			accountInfo.Devices = append(accountInfo.Devices, &pb.EventStatusThreadDevice{
				Name:       shorten(device.id),
				Online:     device.ds.online,
				LastPulled: device.ds.status.LastPull,
				LastEdited: device.ds.lastEdited,
			})

			// account considered online if any device is online
			accountInfo.Online = accountInfo.Online || device.ds.online
			// the very last edit among all devices
			accountInfo.LastEdited = max(accountInfo.LastEdited, device.ds.lastEdited)
			// the very last pull among all devices
			accountInfo.LastPulled = max(accountInfo.LastPulled, device.ds.status.LastPull)
			// collect individual device statuses for summary
			dss = append(dss, device.ds)
		}

		// devices in the same account ordered by last edit time (desc)
		sort.Slice(accountInfo.Devices, func(i, j int) bool {
			return accountInfo.Devices[i].LastEdited > accountInfo.Devices[j].LastEdited
		})

		event.Accounts = append(event.Accounts, &accountInfo)
	}

	// maintain stable order, with own account in a first position
	sort.Slice(event.Accounts, func(i, j int) bool {
		switch {
		case event.Accounts[i].Id == profile.AccountAddr:
			return true
		case event.Accounts[j].Id == profile.AccountAddr:
			return false
		default:
			return event.Accounts[i].Id < event.Accounts[j].Id
		}
	})

	// cafe
	event.Cafe.Status = cafeStatus(cafe)
	event.Cafe.LastPulled = cafe.status.LastPull
	event.Cafe.LastPushSucceed = cafe.status.Up == net.Success
	if !cafe.online && event.Cafe.Status == pb.EventStatusThread_Failed {
		event.Cafe.Status = pb.EventStatusThread_Offline
	}

	// sync status summary
	event.Summary.Status = summaryStatus(event.Cafe.Status, dss...)
	return event
}

func (s *service) sendEvent(ctx string, event pb.IsEventMessageValue) {
	s.emitter(&pb.Event{
		Messages:  []*pb.EventMessage{{Value: event}},
		ContextId: ctx,
	})
}

// Infer cafe status from net-level information
func cafeStatus(cafe deviceStatus) pb.EventStatusThreadSyncStatus {
	switch {
	case cafe.status.Up == net.Failure || (cafe.status.Down == net.Failure &&
		time.Since(time.Unix(cafe.status.LastPull, 0)) > cafeLastPullTimeout):
		return pb.EventStatusThread_Failed
	case cafe.status.Up == net.InProgress || cafe.status.Down == net.InProgress:
		return pb.EventStatusThread_Syncing
	case cafe.status.Up == net.Success ||
		(cafe.status.Up == net.Unknown && cafe.status.Down == net.Success):
		return pb.EventStatusThread_Synced
	default:
		return pb.EventStatusThread_Unknown
	}
}

// Infer sync status summary from individual devices and cafe
func summaryStatus(cafe pb.EventStatusThreadSyncStatus, devices ...deviceStatus) pb.EventStatusThreadSyncStatus {
	var unknown, offline, inProgress, synced, failed int
	for _, device := range devices {
		switch device.status.Down {
		case net.Unknown:
			unknown += 1
		case net.InProgress:
			inProgress += 1
		case net.Success:
			synced += 1
		case net.Failure:
			failed += 1
		}
		if !device.online {
			offline += 1
		}
	}

	switch {
	case synced > 0 || cafe == pb.EventStatusThread_Synced:
		// if thread was synced with cafe or at least one device,
		// it could be considered as a successfully synchronised
		return pb.EventStatusThread_Synced
	case inProgress > 0 || cafe == pb.EventStatusThread_Syncing:
		// sync with some devices or cafe is in progress
		return pb.EventStatusThread_Syncing
	case len(devices) == offline && cafe == pb.EventStatusThread_Offline:
		// no connection with cafe/devices
		return pb.EventStatusThread_Offline
	case synced == 0 && cafe == pb.EventStatusThread_Failed:
		// not synced at all
		return pb.EventStatusThread_Failed
	case unknown > 0 && cafe == pb.EventStatusThread_Unknown:
		// not enough status information
		return pb.EventStatusThread_Unknown
	default:
		return pb.EventStatusThread_Failed
	}
}

type (
	deviceStatus struct {
		status     net.SyncStatus
		lastEdited int64
		online     bool
	}

	threadStatus struct {
		devices  map[string]*deviceStatus
		devConn  func(devID string) bool
		files    map[string]pin.FilePinInfo
		modified bool
		sync.Mutex
	}
)

func newThreadStatus(conn func(devID string) bool) *threadStatus {
	return &threadStatus{
		devConn: conn,
		devices: make(map[string]*deviceStatus),
		files:   make(map[string]pin.FilePinInfo),
	}
}

func (s *threadStatus) UpdateStatus(devID string, ss net.SyncStatus) {
	var dev = s.getDevice(devID)
	if dev.status != ss {
		dev.status = ss
		s.modified = true
	}
}

func (s *threadStatus) UpdateTimeline(devID string, lastEdit int64) {
	var dev = s.getDevice(devID)
	if dev.lastEdited < lastEdit {
		dev.lastEdited = lastEdit
		s.modified = true
	}
}

func (s *threadStatus) UpdateConnectivity(devID string, online bool) {
	var dev = s.getDevice(devID)
	if dev.online != online {
		dev.online = online
		s.modified = true
	}
}

func (s *threadStatus) UpdateFiles(cid string, info pin.FilePinInfo) {
	if fi, ok := s.files[cid]; !ok || fi != info {
		s.files[cid] = info
		s.modified = true
	}
}

func (s *threadStatus) getDevice(id string) *deviceStatus {
	dev, found := s.devices[id]
	if !found {
		dev = &deviceStatus{online: s.devConn(id)}
		s.devices[id] = dev
		s.modified = true
	}

	return dev
}
