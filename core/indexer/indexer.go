package indexer

import (
	"sync"
	"time"

	"github.com/anytypeio/go-anytype-middleware/core/anytype"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	pbrelation "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/relation"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
	"github.com/cheggaaa/mb"
)

var log = logging.Logger("anytype-doc-indexer")

var (
	ftIndexInterval = time.Minute / 3
)

func NewIndexer(a anytype.Service, searchInfo GetSearchInfo) (Indexer, error) {
	ch, cancel, err := a.SubscribeForNewRecords()
	if err != nil {
		return nil, err
	}
	i := &indexer{
		store:      a.ObjectStore(),
		anytype:    a,
		searchInfo: searchInfo,
		cache:      make(map[string]*doc),
		cancel:     cancel,
		quitWG:     &sync.WaitGroup{},
		quit:       make(chan struct{}),
	}
	i.quitWG.Add(2)
	go i.detailsLoop(ch)
	go i.ftLoop()
	return i, nil
}

type Indexer interface {
	Close()
}

type SearchInfo struct {
	Id      string
	Title   string
	Snippet string
	Text    string
	Links   []string
}

type GetSearchInfo interface {
	GetSearchInfo(id string) (info SearchInfo, err error)
}

type indexer struct {
	store      localstore.ObjectStore
	anytype    anytype.Service
	searchInfo GetSearchInfo
	cache      map[string]*doc
	cancel     func()
	quitWG     *sync.WaitGroup
	quit       chan struct{}

	threadIdsBuf []string
	recBuf       []core.SmartblockRecordWithLogID
}

func (i *indexer) detailsLoop(ch chan core.SmartblockRecordWithThreadID) {
	batch := mb.New(0)
	defer batch.Close()
	go func() {
		defer i.quitWG.Done()
		var records []core.SmartblockRecordWithThreadID
		for {
			msgs := batch.Wait()
			if len(msgs) == 0 {
				return
			}
			records = records[:0]
			for _, msg := range msgs {
				records = append(records, msg.(core.SmartblockRecordWithThreadID))
			}
			i.applyRecords(records)
			// wait 100 millisecond for better batching
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for r := range ch {
		batch.Add(r)
	}
}

func (i *indexer) applyRecords(records []core.SmartblockRecordWithThreadID) {
	threadIds := i.threadIdsBuf[:0]
	// find unique threads
	for _, rec := range records {
		if slice.FindPos(threadIds, rec.ThreadID) == -1 {
			threadIds = append(threadIds, rec.ThreadID)
		}
	}
	// group and apply records by thread
	for _, tid := range threadIds {
		threadRecords := i.recBuf[:0]
		for _, rec := range records {
			threadRecords = append(threadRecords, rec.SmartblockRecordWithLogID)
		}
		i.index(tid, threadRecords)
	}
}

func (i *indexer) getDoc(id string) (d *doc, err error) {
	var ok bool
	if d, ok = i.cache[id]; !ok {
		if d, err = newDoc(id, i.anytype); err != nil {
			return
		}
		i.cache[id] = d
	}
	return
}

func (i *indexer) index(id string, records []core.SmartblockRecordWithLogID) {
	d, err := i.getDoc(id)
	if err != nil {
		log.Warnf("can't get doc '%s': %v", id, err)
		return
	}
	if metaChanged := d.addRecords(records...); metaChanged {
		meta := d.meta()
		if err := i.store.UpdateObject(id, meta.Details, &pbrelation.Relations{Relations: meta.Relations}, nil, ""); err != nil {
			log.With("thread", id).Errorf("can't update object store: %v", err)
		}
	}
	if err := i.store.AddToIndexQueue(id); err != nil {
		log.With("thread", id).Errorf("can't add id to index queue: %v", err)
	}
}

func (i *indexer) ftLoop() {
	defer i.quitWG.Done()
	ticker := time.NewTicker(ftIndexInterval)
	for {
		select {
		case <-i.quit:
			return
		case <-ticker.C:
			i.ftIndex()
		}
	}
}

func (i *indexer) ftIndex() {
	if err := i.store.IndexForEach(i.ftIndexDoc); err != nil {
		log.Errorf("store.IndexForEach error: %v", err)
	}
}

func (i *indexer) ftIndexDoc(id string, tm time.Time) (err error) {
	info, err := i.searchInfo.GetSearchInfo(id)
	if err != nil {
		return
	}
	if err = i.store.UpdateObject(id, nil, nil, info.Links, info.Snippet); err != nil {
		return
	}
	// TODO: add info to FT engine here
	return
}

func (i *indexer) Close() {
	if i.cancel != nil {
		i.cancel()
	}
	i.quitWG.Wait()
}
