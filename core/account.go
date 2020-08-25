package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/anytypeio/go-anytype-library/core/smartblock"
	net3 "github.com/anytypeio/go-anytype-library/net"
	util2 "github.com/anytypeio/go-anytype-library/util"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/textileio/go-threads/cbor"
	db2 "github.com/textileio/go-threads/core/db"
	"github.com/textileio/go-threads/core/logstore"
	"github.com/textileio/go-threads/core/net"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
	"github.com/textileio/go-threads/util"
)

const nodeConnectionTimeout = time.Second * 15

// processNewExternalThreadUntilSuccess tries to add the new thread from remote peer until success
// supposed to be run in goroutine
func (a *Anytype) processNewExternalThreadUntilSuccess(tid thread.ID, ti threadInfo) {
	log.Infof("got new thread %s, addrs: %+v", ti.ID.String(), ti.Addrs)

	var attempt int
	for {
		attempt++
		err := a.processNewExternalThread(tid, ti)
		if err != nil {
			log.Errorf("processNewExternalThread %s failed after %d attempt: %s", tid.String(), err.Error())
		} else {
			log.Debugf("processNewExternalThread %s succeed after %d attempt", tid.String())
			return
		}
		select {
		case <-a.shutdownStartsCh:
			return
		case <-time.After(time.Duration(5*attempt) * time.Second):
			continue
		}
	}
}

func (a *Anytype) processNewExternalThread(tid thread.ID, ti threadInfo) error {
	key, err := thread.KeyFromString(ti.Key)
	if err != nil {
		return fmt.Errorf("failed to parse thread keys %s: %s", tid.String(), err.Error())
	}

	threadAdded := false
	hasCafeAddress := false
	var multiAddrs []ma.Multiaddr
	threadComp, err := ma.NewComponent(thread.Name, tid.String())
	if err != nil {
		return err
	}
	var cafeAddrWithThread ma.Multiaddr
	if a.opts.CafeP2PAddr != nil {
		cafeAddrWithThread = a.opts.CafeP2PAddr.Encapsulate(threadComp)
	}

	for _, addrS := range ti.Addrs {
		addr, err := ma.NewMultiaddr(addrS)
		if err != nil {
			log.Errorf("processNewExternalThread: failed to parse addr %s: %s", addrS, err.Error())
			continue
		}

		if addr.Decapsulate(threadComp).Equal(a.opts.CafeP2PAddr) {
			hasCafeAddress = true
		}

		multiAddrs = append(multiAddrs, addr)
	}

	if !hasCafeAddress && cafeAddrWithThread != nil {
		log.Warn("processNewExternalThread %s: cafe addr not found among thread addresses, will add it", ti.ID.String())
		multiAddrs = append(multiAddrs, cafeAddrWithThread)
	}

addrsLoop:
	for _, addr := range multiAddrs {
		for _, ownAddr := range a.t.Host().Addrs() {
			ipOwn, _ := ownAddr.ValueForProtocol(ma.P_IP4)
			ipTarget, _ := addr.ValueForProtocol(ma.P_IP4)

			portOwn, _ := ownAddr.ValueForProtocol(ma.P_TCP)
			portTarget, _ := addr.ValueForProtocol(ma.P_TCP)

			// do not connect to ourselves
			if ipOwn == ipTarget && portOwn == portTarget {
				continue addrsLoop
			}
		}

		// commented-out because we can have a local cafe node for dev purposes
		/*if v, _ := addr.ValueForProtocol(ma.P_IP4); v == "127.0.0.1" {
			// skip localhost
			continue
		}

		if v, _ := addr.ValueForProtocol(ma.P_IP6); v == "::1" {
			// skip localhost
			continue
		}*/

		threadComp, err := ma.NewComponent(thread.Name, tid.String())
		if err != nil {
			log.Errorf("processNewExternalThread %s: failed to parse addr %s: %s", ti.ID.String(), addr.String(), err.Error())
			continue
		}

		peerAddr := addr.Decapsulate(threadComp)
		addri, err := peer.AddrInfoFromP2pAddr(peerAddr)
		if err != nil {
			log.Errorf("processNewExternalThread %s: failed to parse addr %s: %s", ti.ID.String(), addr.String(), err.Error())
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), nodeConnectionTimeout)
		defer cancel()

		if err = a.t.Host().Connect(ctx, *addri); err != nil {
			log.Errorf("processNewExternalThread %s: failed to connect addr %s: %s", ti.ID.String(), addr.String(), err.Error())
			continue
		}

		threadID, _ := addr.ValueForProtocol(thread.Code)
		if threadID == "" {
			// in case we have thread comp missing just add it
			addr = addr.Encapsulate(threadComp)
		}

		_, err = a.t.AddThread(context.Background(), addr, net.WithThreadKey(key), net.WithLogKey(a.opts.Device))
		if err != nil {
			log.Errorf("processNewExternalThread %s: failed to add from %s: %s", ti.ID.String(), addr.String(), err.Error())
			continue
		}

		log.Infof("processNewExternalThread %s: thread successfully added from %s", ti.ID.String(), addri)

		_, err = a.t.AddReplicator(context.Background(), tid, cafeAddrWithThread)
		if err != nil {
			log.Errorf("processNewExternalThread failed to add the replicator for %s: %s", ti.ID.String(), err.Error())
		}

		threadAdded = true
		break
	}

	if !threadAdded {
		return fmt.Errorf("failed to add thread from any provided remote address")
	} else {
		_, err = a.pullThread(context.Background(), tid)
		if err != nil {
			log.Errorf("processNewExternalThread: pull thread %s failed: %s", tid.String(), err.Error())
		}
	}

	err = a.migratePageToChanges(tid)
	if err != nil {
		if err != ErrAlreadyMigrated {
			log.Errorf("processNewExternalThread migratePageToChanges %s failed: %s", tid.String(), err.Error())
		} else {
			log.Debugf("processNewExternalThread migratePageToChanges %s: thread already migrated", tid.String())
		}
	} else {
		log.Debugf("processNewExternalThread: thread %s migrated", tid.String())
	}

	go func() {
		// reindex after migration
		err = a.opts.ReindexFunc(tid.String())
		if err != nil {
			log.Errorf("processNewExternalThread: failed to reindex page %s: %s", tid.String(), err.Error())
		}
	}()

	return nil

}

type threadRecord struct {
	net.Record
	threadID thread.ID
	logID    peer.ID
}

func (t threadRecord) Value() net.Record {
	return t.Record
}

func (t threadRecord) ThreadID() thread.ID {
	return t.threadID
}

func (t threadRecord) LogID() peer.ID {
	return t.logID
}

func (a *Anytype) addMissingThreadsToCollection() error {
	threadsCollection := a.ThreadsCollection()
	instancesBytes, err := threadsCollection.Find(&db.Query{})
	if err != nil {
		return err
	}

	var threadsInCollection = make(map[string]struct{})
	for _, instanceBytes := range instancesBytes {
		ti := threadInfo{}
		util.InstanceFromJSON(instanceBytes, &ti)

		tid, err := thread.Decode(ti.ID.String())
		if err != nil {
			log.Errorf("failed to parse thread id %s: %s", ti.ID, err.Error())
			continue
		}
		threadsInCollection[tid.String()] = struct{}{}
	}

	log.Debugf("%d threads in collection", len(threadsInCollection))

	threadsIds, err := a.ThreadsNet().Logstore().Threads()
	if err != nil {
		return err
	}

	var missingThreads int
	for _, threadId := range threadsIds {
		t, _ := smartblock.SmartBlockTypeFromThreadID(threadId)
		if t != smartblock.SmartBlockTypePage {
			continue
		}

		if _, exists := threadsInCollection[threadId.String()]; !exists {
			thrd, err := a.ThreadsNet().GetThread(context.Background(), threadId)
			if err != nil {
				log.Errorf("addMissingThreadsToCollection migration: error getting info: %s\n", err.Error())
				continue
			}
			threadInfo := threadInfo{
				ID:    db2.InstanceID(thrd.ID.String()),
				Key:   thrd.Key.String(),
				Addrs: util2.MultiAddressesToStrings(thrd.Addrs),
			}

			_, err = a.threadsCollection.Create(util.JSONFromInstance(threadInfo))
			if err != nil {
				log.With("thread", thrd.ID.String()).Errorf("failed to create thread at collection: %s: ", err.Error())
			} else {
				missingThreads++
			}
		}
	}

	if missingThreads > 0 {
		log.Errorf("addMissingThreadsToCollection migration: added %d missing threads", missingThreads)
	} else {
		log.Debugf("addMissingThreadsToCollection migration: no missing threads found")
	}

	return nil
}

func (a *Anytype) addMissingThreadsFromCollection() error {
	threadsCollection := a.ThreadsCollection()
	instancesBytes, err := threadsCollection.Find(&db.Query{})
	if err != nil {
		return err
	}

	var missingThreadsAdded int
	for _, instanceBytes := range instancesBytes {
		ti := threadInfo{}
		util.InstanceFromJSON(instanceBytes, &ti)

		tid, err := thread.Decode(ti.ID.String())
		if err != nil {
			log.Errorf("failed to parse thread id %s: %s", ti.ID, err.Error())
			continue
		}

		if _, err = a.ThreadsNet().GetThread(context.Background(), tid); err != nil && errors.Is(err, logstore.ErrThreadNotFound) {
			missingThreadsAdded++
			go func() {
				a.processNewExternalThreadUntilSuccess(tid, ti)
			}()
		}
	}

	if missingThreadsAdded > 0 {
		log.Errorf("addMissingThreadsFromCollection: adding %d missing threads in background...", missingThreadsAdded)
	}

	return nil
}

func (a *Anytype) addMissingReplicators() error {
	threadsIds, err := a.ThreadsNet().Logstore().Threads()
	if err != nil {
		return fmt.Errorf("failed to list threads: %s", err.Error())
	}

	if a.opts.CafeP2PAddr == nil {
		return nil
	}

	cafeP2PAddr, err := a.opts.CafeP2PAddr.ValueForProtocol(ma.P_P2P)
	if err != nil {
		return err
	}

	for _, threadId := range threadsIds {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)

		thrd, err := a.ThreadsNet().GetThread(ctx, threadId)
		if err != nil {
			log.Errorf("error getting thread info: %s", err.Error())
			continue
		}

		var exists = false
		for _, addr := range thrd.Addrs {
			if p2paddr, err := addr.ValueForProtocol(ma.P_P2P); err == nil {
				if p2paddr == cafeP2PAddr {
					exists = true
					break
				}
			}
		}

		if !exists {
			threadComp, err := ma.NewComponent(thread.Name, threadId.String())
			if err != nil {
				return err
			}

			if _, err = a.ThreadsNet().AddReplicator(ctx, thrd.ID, a.opts.CafeP2PAddr.Encapsulate(threadComp)); err != nil {
				log.Errorf("failed to add missing replicator for %s: %s", thrd.ID, err.Error())
			} else {
				log.Infof("added missing replicator for %s", thrd.ID)
			}
		}
	}

	return nil
}

func (a *Anytype) handleAllMissingDbRecords(threadId string) error {
	tid, err := thread.Decode(threadId)
	if err != nil {
		return fmt.Errorf("failed to parse thread id %s: %s", threadId, err.Error())
	}

	thrd, err := a.ThreadsNet().GetThread(context.Background(), tid)
	if err != nil {
		return fmt.Errorf("failed to get thread info: %s", err.Error())
	}

	for _, logInfo := range thrd.Logs {
		log.Debugf("traversing %s log from head %s", logInfo.ID, logInfo.Head)
		handleAllRecordsInLog(a.ThreadsDB(), a.ThreadsNet(), thrd, logInfo)
	}
	return nil
}

func getRecord(net net3.NetBoostrapper, thrd thread.Info, rid cid.Cid) (net.Record, *cbor.Event, format.Node, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if thrd.ID == thread.Undef {
		return nil, nil, nil, fmt.Errorf("undef id")
	}

	rec, err := net.GetRecord(ctx, thrd.ID, rid)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to load record: %s", err.Error())
	}

	event, err := cbor.EventFromRecord(ctx, net, rec)
	if err != nil {
		return rec, nil, nil, fmt.Errorf("failed to load event: %s", err.Error())
	}

	node, err := event.GetBody(context.TODO(), net, thrd.Key.Read())
	if err != nil {
		return rec, event, nil, fmt.Errorf("failed to get record body: %w", err)
	}

	return rec, event, node, nil
}

func handleAllRecordsInLog(tdb *db.DB, net net3.NetBoostrapper, thrd thread.Info, li thread.LogInfo) {
	var (
		rid     = li.Head
		total   = 0
		records []threadRecord
	)

	defer func() {
		for i := len(records) - 1; i >= 0; i-- {
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			err := tdb.HandleNetRecord(ctx, records[i], thrd.Key)
			if err != nil {
				log.Errorf("failed to handle record: %s", err.Error())
			}
		}
	}()

	for {
		if !rid.Defined() {
			return
		}
		total++
		rec, _, _, err := getRecord(net, thrd, rid)
		if rec != nil {
			trec := threadRecord{
				Record:   rec,
				threadID: thrd.ID,
				logID:    li.ID,
			}

			records = append(records, trec)
			rid = rec.PrevID()
		} else {
			log.Errorf("can't continue the traverse, failed to load a record: %s", err.Error())
			return
		}
	}
}

func (a *Anytype) listenExternalNewThreads() error {
	l, err := a.db.Listen()
	if err != nil {
		return err
	}

	go func() {
		defer l.Close()
		for {
			select {
			case <-a.shutdownStartsCh:
				log.Infof("shutting down external changes listener")
				return
			case c := <-l.Channel():
				switch c.Type {
				case db.ActionCreate:
					instanceBytes, err := a.threadsCollection.FindByID(c.ID)
					if err != nil {
						log.Errorf("failed to find thread info for id %s: %w", c.ID.String(), err)
						continue
					}

					ti := threadInfo{}
					util.InstanceFromJSON(instanceBytes, &ti)
					tid, err := thread.Decode(ti.ID.String())
					if err != nil {
						log.Errorf("failed to parse thread id %s: %s", ti.ID, err.Error())
						continue
					}

					info, _ := a.t.GetThread(context.Background(), tid)
					if info.ID != thread.Undef {
						// our own event
						continue
					}

					go func() {
						a.processNewExternalThreadUntilSuccess(tid, ti)
					}()
				}
			}
		}
	}()

	return nil
}
