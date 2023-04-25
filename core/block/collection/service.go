package collection

import (
	"fmt"
	"sync"

	"github.com/anytypeio/any-sync/app"
	"github.com/gogo/protobuf/types"

	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/basic"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/session"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	coresb "github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/internalflag"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
)

var log = logging.Logger("collection-service")

type Service struct {
	lock        *sync.RWMutex
	collections map[string]map[string]chan []string

	picker        block.Picker
	objectStore   objectstore.ObjectStore
	objectCreator ObjectCreator
	objectDeleter ObjectDeleter
}

type ObjectCreator interface {
	CreateObject(req block.DetailsGetter, forcedType bundle.TypeKey) (id string, details *types.Struct, err error)
}

type ObjectDeleter interface {
	DeleteObject(id string) (err error)
}

func New(
	picker block.Picker,
	store objectstore.ObjectStore,
	objectCreator ObjectCreator,
	objectDeleter ObjectDeleter,
) *Service {
	return &Service{
		picker:        picker,
		objectStore:   store,
		objectCreator: objectCreator,
		objectDeleter: objectDeleter,
		lock:          &sync.RWMutex{},
		collections:   map[string]map[string]chan []string{},
	}
}

func (s *Service) Init(a *app.App) (err error) {
	return nil
}

func (s *Service) Name() string {
	return "collection"
}

func (s *Service) Add(ctx *session.Context, req *pb.RpcObjectCollectionAddRequest) error {
	return s.updateCollection(ctx, req.ContextId, func(col []string) []string {
		toAdd := slice.Difference(req.ObjectIds, col)
		pos := slice.FindPos(col, req.AfterId)
		if pos >= 0 {
			col = slice.Insert(col, pos+1, toAdd...)
		} else {
			col = append(toAdd, col...)
		}
		return col
	})
}

func (s *Service) Remove(ctx *session.Context, req *pb.RpcObjectCollectionRemoveRequest) error {
	return s.updateCollection(ctx, req.ContextId, func(col []string) []string {
		col = slice.Filter(col, func(id string) bool {
			return slice.FindPos(req.ObjectIds, id) == -1
		})
		return col
	})
}

func (s *Service) Sort(ctx *session.Context, req *pb.RpcObjectCollectionSortRequest) error {
	return s.updateCollection(ctx, req.ContextId, func(col []string) []string {
		exist := map[string]struct{}{}
		for _, id := range col {
			exist[id] = struct{}{}
		}
		col = col[:0]
		for _, id := range req.ObjectIds {
			// Reorder only existing objects
			if _, ok := exist[id]; ok {
				col = append(col, id)
			}
		}
		return col
	})
}

func (s *Service) updateCollection(ctx *session.Context, contextID string, modifier func(src []string) []string) error {
	return block.DoStateCtx(s.picker, ctx, contextID, func(s *state.State, sb smartblock.SmartBlock) error {
		lst := pbtypes.GetStringList(s.Store(), template.CollectionStoreKey)
		lst = modifier(lst)
		s.StoreSlice(template.CollectionStoreKey, lst)
		return nil
	})
}

func (s *Service) RegisterCollection(sb smartblock.SmartBlock) {
	s.lock.Lock()
	col, ok := s.collections[sb.Id()]
	if !ok {
		col = map[string]chan []string{}
		s.collections[sb.Id()] = col
	}
	s.lock.Unlock()

	sb.AddHook(func(info smartblock.ApplyInfo) (err error) {
		for _, ch := range info.Changes {
			if upd := ch.GetStoreSliceUpdate(); upd != nil && upd.Key == template.CollectionStoreKey {
				s.broadcast(sb.Id(), pbtypes.GetStringList(info.State.Store(), template.CollectionStoreKey))
				return nil
			}
		}
		return nil
	}, smartblock.HookAfterApply)
}

func (s *Service) broadcast(collectionID string, objectIDs []string) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, ch := range s.collections[collectionID] {
		ch <- objectIDs
	}
}

type Subscription struct {
	objectsCh chan []string
	closeCh   chan struct{}
}

func (s *Subscription) Chan() <-chan []string {
	return s.objectsCh
}

func (s *Subscription) Close() {
	close(s.closeCh)
}

func (s *Service) SubscribeForCollection(collectionID string, subscriptionID string) ([]string, <-chan []string, error) {
	var initialObjectIDs []string
	// Waking up of collection smart block will automatically add hook used in RegisterCollection
	err := block.DoState(s.picker, collectionID, func(s *state.State, sb smartblock.SmartBlock) error {
		initialObjectIDs = pbtypes.GetStringList(s.Store(), template.CollectionStoreKey)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	col, ok := s.collections[collectionID]
	if !ok {
		return nil, nil, fmt.Errorf("collection is not registered")
	}

	ch, ok := col[subscriptionID]
	if !ok {
		ch = make(chan []string)
		col[subscriptionID] = ch
	}

	return initialObjectIDs, ch, err
}

func (s *Service) UnsubscribeFromCollection(collectionID string, subscriptionID string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	col, ok := s.collections[collectionID]
	if !ok {
		return
	}

	ch := col[subscriptionID]
	close(ch)
	delete(col, subscriptionID)
}

func (s *Service) CreateCollection(details *types.Struct, flags []*model.InternalFlag) (coresb.SmartBlockType, *types.Struct, *state.State, error) {
	details = internalflag.PutToDetails(details, flags)

	newState := state.NewDoc("", nil).NewState()

	tmpls := []template.StateTransformer{
		template.WithRequiredRelations(),
	}

	blockContent := template.MakeCollectionDataviewContent()
	tmpls = append(tmpls,
		template.WithDataview(*blockContent, false),
	)
	template.InitTemplate(newState, tmpls...)

	return coresb.SmartBlockTypeCollection, details, newState, nil
}

// TODO To be removed
func (s *Service) ObjectToCollection(id string) (string, error) {
	var (
		details *types.Struct
	)

	// cleanup details
	delete(details.Fields, bundle.RelationKeyLayout.String())
	delete(details.Fields, bundle.RelationKeyType.String())

	newID, _, err := s.objectCreator.CreateObject(&pb.RpcObjectCreateRequest{
		Details: details,
	}, bundle.TypeKeyCollection)
	if err != nil {
		return "", err
	}

	res, err := s.objectStore.GetWithLinksInfoByID(id)
	if err != nil {
		return "", err
	}
	for _, il := range res.Links.Inbound {
		err = block.Do(s.picker, il.Id, func(b basic.CommonOperations) error {
			return b.ReplaceLink(id, newID)
		})
		if err != nil {
			return "", fmt.Errorf("replace link in %s: %w", il.Id, err)
		}
	}
	err = s.objectDeleter.DeleteObject(id)
	if err != nil {
		// intentionally do not return error here
		log.Errorf("failed to delete object after conversion to set: %s", err.Error())
	}

	return newID, nil
}
