package importer

import (
	"context"
	"fmt"
	"sync"

	"github.com/anyproto/any-sync/commonspace/object/tree/treestorage"
	"github.com/gogo/protobuf/types"
	"github.com/samber/lo"
	"go.uber.org/zap"

	"github.com/anyproto/anytype-heart/core/block"
	"github.com/anyproto/anytype-heart/core/block/editor/basic"
	sb "github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/editor/state"
	"github.com/anyproto/anytype-heart/core/block/editor/template"
	"github.com/anyproto/anytype-heart/core/block/history"
	"github.com/anyproto/anytype-heart/core/block/import/converter"
	"github.com/anyproto/anytype-heart/core/block/import/syncer"
	"github.com/anyproto/anytype-heart/core/block/simple"
	simpleDataview "github.com/anyproto/anytype-heart/core/block/simple/dataview"
	"github.com/anyproto/anytype-heart/core/filestorage/filesync"
	"github.com/anyproto/anytype-heart/core/session"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/core"
	coresb "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/filestore"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

const relationsLimit = 10

type objectCreator interface {
	CreateSmartBlockFromState(ctx context.Context, sbType coresb.SmartBlockType, details *types.Struct, createState *state.State) (id string, newDetails *types.Struct, err error)
	CreateSubObjectInWorkspace(details *types.Struct, workspaceID string) (id string, newDetails *types.Struct, err error)
	CreateSubObjectsInWorkspace(details []*types.Struct) (ids []string, objects []*types.Struct, err error)
}

type ObjectCreator struct {
	service         *block.Service
	objCreator      objectCreator
	core            core.Service
	objectStore     objectstore.ObjectStore
	fileStore       filestore.FileStore
	relationCreator RelationCreator
	syncFactory     *syncer.Factory
	spaceService    space.Service
	fileSync        filesync.FileSync
	mu              sync.Mutex
}

func NewCreator(service *block.Service,
	objCreator objectCreator,
	core core.Service,
	syncFactory *syncer.Factory,
	relationCreator RelationCreator,
	objectStore objectstore.ObjectStore,
	fileStore filestore.FileStore,
) Creator {
	return &ObjectCreator{
		service:         service,
		objCreator:      objCreator,
		core:            core,
		syncFactory:     syncFactory,
		relationCreator: relationCreator,
		objectStore:     objectStore,
		fileStore:       fileStore,
	}
}

// Create creates smart blocks from given snapshots
func (oc *ObjectCreator) Create(ctx *session.Context,
	sn *converter.Snapshot,
	relations []*converter.Relation,
	oldIDtoNew map[string]string,
	createPayloads map[string]treestorage.TreeStorageCreatePayload,
	existing bool) (*types.Struct, string, error) {
	snapshot := sn.Snapshot.Data

	var err error
	newID := oldIDtoNew[sn.Id]
	oc.updateRootBlock(snapshot, newID)

	oc.setWorkspaceID(err, newID, snapshot)

	filesToDelete, _, createdRelations, err := oc.relationCreator.CreateRelations(ctx, snapshot, newID, relations)
	if err != nil {
		return nil, "", fmt.Errorf("relation create '%s'", err)
	}

	st := state.NewDocFromSnapshot(newID, sn.Snapshot).(*state.State)
	st.SetRootId(newID)
	// explicitly set last modified date, because all local details removed in NewDocFromSnapshot; createdDate covered in the object header
	st.SetLastModified(pbtypes.GetInt64(sn.Snapshot.Data.Details, bundle.RelationKeyLastModifiedDate.String()), oc.core.ProfileID())
	defer func() {
		// delete file in ipfs if there is error after creation
		oc.onFinish(err, st, filesToDelete)
	}()

	converter.UpdateRelationsIDs(st, oldIDtoNew)
	if sn.SbType == coresb.SmartBlockTypeSubObject {
		oc.handleSubObject(st, newID)
		return nil, newID, nil
	}

	if err = converter.UpdateLinksToObjects(st, oldIDtoNew, newID); err != nil {
		log.With("object", newID).Errorf("failed to update objects ids: %s", err.Error())
	}

	if sn.SbType == coresb.SmartBlockTypeWorkspace {
		oc.setSpaceDashboardID(st, oldIDtoNew)
		return nil, newID, nil
	}

	var respDetails *types.Struct
	converter.UpdateObjectType(oldIDtoNew, st)
	if payload := createPayloads[newID]; payload.RootRawChange != nil {
		sb, err := oc.service.CreateTreeObjectWithPayload(context.Background(), payload, func(id string) *sb.InitContext {
			return &sb.InitContext{
				IsNewObject: true,
				State:       st,
			}
		})
		if err != nil {
			log.With("object", newID).Errorf("failed to create %s: %s", newID, err.Error())
			return nil, "", err
		}
		respDetails = sb.Details()
		// update collection after we create it
		if st.Store() != nil {
			oc.updateCollectionState(st, oldIDtoNew, relations, createdRelations)
			oc.resetState(ctx, newID, st)
		}
	} else {
		if st.Store() != nil {
			oc.updateCollectionState(st, oldIDtoNew, relations, createdRelations)
		}
		respDetails = oc.resetState(ctx, newID, st)
	}
	oc.setFavorite(snapshot, newID)

	oc.setArchived(snapshot, newID)

	// we do not change relation ids during the migration
	// oc.relationCreator.ReplaceRelationBlock(ctx, oldRelationBlocksToNew, newID)

	syncErr := oc.syncFilesAndLinks(ctx, st, newID)
	if syncErr != nil {
		log.With(zap.String("object id", newID)).Errorf("failed to sync %s: %s", newID, err.Error())
	}

	return respDetails, newID, nil
}

func (oc *ObjectCreator) updateCollectionState(st *state.State,
	oldIDtoNew map[string]string,
	relations []*converter.Relation,
	createdRelations map[string]RelationsIDToFormat) {
	oc.updateLinksInCollections(st, oldIDtoNew)
	if err := oc.addRelationsToCollectionDataView(st, relations, createdRelations); err != nil {
		log.With("object", st.RootId()).Errorf("failed to add relations to object view: %s", err.Error())
	}
}

func (oc *ObjectCreator) setArchived(snapshot *model.SmartBlockSnapshotBase, newID string) {
	isArchive := pbtypes.GetBool(snapshot.Details, bundle.RelationKeyIsArchived.String())
	if isArchive {
		err := oc.service.SetPageIsArchived(pb.RpcObjectSetIsArchivedRequest{ContextId: newID, IsArchived: true})
		if err != nil {
			log.With(zap.String("object id", newID)).
				Errorf("failed to set isFavorite when importing object %s: %s", newID, err.Error())
		}
	}
}

func (oc *ObjectCreator) setFavorite(snapshot *model.SmartBlockSnapshotBase, newID string) {
	isFavorite := pbtypes.GetBool(snapshot.Details, bundle.RelationKeyIsFavorite.String())
	if isFavorite {
		err := oc.service.SetPageIsFavorite(pb.RpcObjectSetIsFavoriteRequest{ContextId: newID, IsFavorite: true})
		if err != nil {
			log.With(zap.String("object id", newID)).Errorf("failed to set isFavorite when importing object: %s", err.Error())
		}
	}
}

func (oc *ObjectCreator) setWorkspaceID(err error, newID string, snapshot *model.SmartBlockSnapshotBase) {
	if oc.core.PredefinedBlocks().Account == newID {
		return
	}
	workspaceID, err := oc.core.GetWorkspaceIdForObject(newID)
	if err != nil {
		// todo: GO-1304 I catch this during the import, we need find the root cause and fix it
		log.With(zap.String("object id", newID)).Errorf("failed to get workspace id %s: %s", newID, err.Error())
	}

	if snapshot.Details != nil && snapshot.Details.Fields != nil {
		snapshot.Details.Fields[bundle.RelationKeyWorkspaceId.String()] = pbtypes.String(workspaceID)
	}
}

func (oc *ObjectCreator) updateRootBlock(snapshot *model.SmartBlockSnapshotBase, newID string) {
	var found bool
	for _, b := range snapshot.Blocks {
		if b.Id == newID {
			found = true
			break
		}
	}
	if !found {
		oc.addRootBlock(snapshot, newID)
	}
}

func (oc *ObjectCreator) syncFilesAndLinks(ctx *session.Context, st *state.State, newID string) error {
	return st.Iterate(func(bl simple.Block) (isContinue bool) {
		s := oc.syncFactory.GetSyncer(bl)
		if s != nil {
			if sErr := s.Sync(ctx, newID, bl); sErr != nil {
				log.With(zap.String("object id", newID)).Errorf("sync: %s", sErr)
			}
		}
		return true
	})
}

func (oc *ObjectCreator) onFinish(err error, st *state.State, filesToDelete []string) {
	if err != nil {
		for _, bl := range st.Blocks() {
			if f := bl.GetFile(); f != nil {
				oc.deleteFile(f.Hash)
			}
			for _, hash := range filesToDelete {
				oc.deleteFile(hash)
			}
		}
	}
}

func (oc *ObjectCreator) setSpaceDashboardID(st *state.State, oldIDtoNew map[string]string) {
	// hand-pick relation because space is a special case
	var details []*pb.RpcObjectSetDetailsDetail
	spaceDashBoardID := pbtypes.GetString(st.CombinedDetails(), bundle.RelationKeySpaceDashboardId.String())
	if spaceDashBoardID != "" {
		details = append(details, &pb.RpcObjectSetDetailsDetail{
			Key:   bundle.RelationKeySpaceDashboardId.String(),
			Value: pbtypes.String(spaceDashBoardID),
		})
	}

	spaceName := pbtypes.GetString(st.CombinedDetails(), bundle.RelationKeyName.String())
	if spaceName != "" {
		details = append(details, &pb.RpcObjectSetDetailsDetail{
			Key:   bundle.RelationKeyName.String(),
			Value: pbtypes.String(spaceName),
		})
	}

	iconOption := pbtypes.GetInt64(st.CombinedDetails(), bundle.RelationKeyIconOption.String())
	if iconOption != 0 {
		details = append(details, &pb.RpcObjectSetDetailsDetail{
			Key:   bundle.RelationKeyIconOption.String(),
			Value: pbtypes.Int64(iconOption),
		})
	}
	if len(details) > 0 {
		err := block.Do(oc.service, oc.core.PredefinedBlocks().Account, func(ws basic.CommonOperations) error {
			if err := ws.SetDetails(nil, details, false); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Errorf("failed to set spaceDashBoardID, %s", err.Error())
		}
	}
}

func (oc *ObjectCreator) resetState(ctx *session.Context, newID string, st *state.State) *types.Struct {
	var respDetails *types.Struct
	err := oc.service.Do(newID, func(b sb.SmartBlock) error {
		err := history.ResetToVersion(b, st)
		if err != nil {
			log.With(zap.String("object id", newID)).Errorf("failed to set state %s: %s", newID, err.Error())
		}
		commonOperations, ok := b.(basic.CommonOperations)
		if !ok {
			return fmt.Errorf("common operations is not allowed for this object")
		}
		err = commonOperations.FeaturedRelationAdd(ctx, bundle.RelationKeyType.String())
		if err != nil {
			log.With(zap.String("object id", newID)).Errorf("failed to set featuredRelations %s: %s", newID, err.Error())
		}
		respDetails = b.CombinedDetails()
		return nil
	})
	if err != nil {
		log.With(zap.String("object id", newID)).Errorf("failed to reset state %s: %s", newID, err.Error())
	}
	return respDetails
}

func (oc *ObjectCreator) addRootBlock(snapshot *model.SmartBlockSnapshotBase, pageID string) {
	for i, b := range snapshot.Blocks {
		if _, ok := b.Content.(*model.BlockContentOfSmartblock); ok {
			snapshot.Blocks[i].Id = pageID
			return
		}
	}
	notRootBlockChild := make(map[string]bool, 0)
	for _, b := range snapshot.Blocks {
		for _, id := range b.ChildrenIds {
			notRootBlockChild[id] = true
		}
	}
	childrenIds := make([]string, 0)
	for _, b := range snapshot.Blocks {
		if _, ok := notRootBlockChild[b.Id]; !ok {
			childrenIds = append(childrenIds, b.Id)
		}
	}
	snapshot.Blocks = append(snapshot.Blocks, &model.Block{
		Id: pageID,
		Content: &model.BlockContentOfSmartblock{
			Smartblock: &model.BlockContentSmartblock{},
		},
		ChildrenIds: childrenIds,
	})
}

func (oc *ObjectCreator) deleteFile(hash string) {
	inboundLinks, err := oc.objectStore.GetOutboundLinksById(hash)
	if err != nil {
		log.With("file", hash).Errorf("failed to get inbound links for file: %s", err)
	}
	if len(inboundLinks) == 0 {
		err = oc.service.DeleteObject(hash)
		if err != nil {
			log.With("file", hash).Errorf("failed to delete file: %s", err)
		}
	}
}

func (oc *ObjectCreator) handleSubObject(st *state.State, newID string) {
	defer oc.mu.Unlock()
	oc.mu.Lock()
	if deleted := pbtypes.GetBool(st.CombinedDetails(), bundle.RelationKeyIsDeleted.String()); deleted {
		err := oc.service.RemoveSubObjectsInWorkspace([]string{newID}, oc.core.PredefinedBlocks().Account, true)
		if err != nil {
			log.With(zap.String("object id", newID)).Errorf("failed to remove from collections %s: %s", newID, err.Error())
		}
		return
	}

	// RQ: the rest handling were removed
}

func (oc *ObjectCreator) addRelationsToCollectionDataView(st *state.State, rels []*converter.Relation, createdRelations map[string]RelationsIDToFormat) error {
	return st.Iterate(func(bl simple.Block) (isContinue bool) {
		if dv, ok := bl.(simpleDataview.Block); ok {
			return oc.handleDataviewBlock(bl, rels, createdRelations, dv)
		}
		return true
	})
}

func (oc *ObjectCreator) handleDataviewBlock(bl simple.Block, rels []*converter.Relation, createdRelations map[string]RelationsIDToFormat, dv simpleDataview.Block) bool {
	for i, rel := range rels {
		if relation, exist := createdRelations[rel.Name]; exist {
			isVisible := i <= relationsLimit
			if err := oc.addRelationToView(bl, relation, rel, dv, isVisible); err != nil {
				log.Errorf("can't add relations to view: %s", err.Error())
			}
		}
	}
	return false
}

func (oc *ObjectCreator) addRelationToView(bl simple.Block, relation RelationsIDToFormat, rel *converter.Relation, dv simpleDataview.Block, visible bool) error {
	for _, relFormat := range relation {
		if relFormat.Format == rel.Format {
			if len(bl.Model().GetDataview().GetViews()) == 0 {
				return nil
			}
			for _, view := range bl.Model().GetDataview().GetViews() {
				err := dv.AddViewRelation(view.GetId(), &model.BlockContentDataviewRelation{
					Key:       relFormat.ID,
					IsVisible: visible,
					Width:     192,
				})
				if err != nil {
					return err
				}
			}
			err := dv.AddRelation(&model.RelationLink{
				Key:    relFormat.ID,
				Format: relFormat.Format,
			})
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func (oc *ObjectCreator) updateLinksInCollections(st *state.State, oldIDtoNew map[string]string) {
	err := block.DoStateCtx(oc.service, nil, st.RootId(), func(s *state.State, b sb.SmartBlock) error {
		oc.mergeCollections(s.GetStoreSlice(template.CollectionStoreKey), st, oldIDtoNew)
		return nil
	})
	if err != nil {
		log.Errorf("failed to get existed objects in collection, %s", err)
	}
}

func (oc *ObjectCreator) mergeCollections(existedObjects []string, st *state.State, oldIDtoNew map[string]string) {
	objectsInCollections := st.GetStoreSlice(template.CollectionStoreKey)
	for i, id := range objectsInCollections {
		if newID, ok := oldIDtoNew[id]; ok {
			objectsInCollections[i] = newID
		}
	}
	result := lo.Union(existedObjects, objectsInCollections)
	st.UpdateStoreSlice(template.CollectionStoreKey, result)
}
