package converter

import (
	"github.com/gogo/protobuf/types"
	"github.com/google/uuid"

	"github.com/anytypeio/go-anytype-middleware/core/block/collection"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	simpleDataview "github.com/anytypeio/go-anytype-middleware/core/block/simple/dataview"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	sb "github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
)

type RootCollection struct {
	service *collection.Service
}

func NewRootCollection(service *collection.Service) *RootCollection {
	return &RootCollection{service: service}
}

func (r *RootCollection) AddObjects(collectionName string, targetObjects []string) (*Snapshot, error) {
	detailsStruct := r.getCreateCollectionRequest(collectionName)
	_, _, st, err := r.service.CreateCollection(detailsStruct, []*model.InternalFlag{{
		Value: model.InternalFlag_collectionDontIndex,
	}})
	if err != nil {
		return nil, err
	}

	err = r.addRelations(st)
	if err != nil {
		return nil, err
	}

	detailsStruct = pbtypes.StructMerge(st.CombinedDetails(), detailsStruct, false)
	st.UpdateStoreSlice(template.CollectionStoreKey, targetObjects)

	return r.getRootCollectionSnapshot(collectionName, st, detailsStruct), nil
}

func (r *RootCollection) getRootCollectionSnapshot(collectionName string, st *state.State, detailsStruct *types.Struct) *Snapshot {
	if detailsStruct.GetFields() == nil {
		detailsStruct = &types.Struct{Fields: map[string]*types.Value{}}
	}
	detailsStruct.Fields[bundle.RelationKeyLayout.String()] = pbtypes.Int64(int64(model.ObjectType_collection))
	return &Snapshot{
		Id:       uuid.New().String(),
		FileName: collectionName,
		SbType:   sb.SmartBlockTypePage,
		Snapshot: &pb.ChangeSnapshot{
			Data: &model.SmartBlockSnapshotBase{
				Blocks:        st.Blocks(),
				Details:       detailsStruct,
				ObjectTypes:   []string{bundle.TypeKeyCollection.URL()},
				RelationLinks: st.GetRelationLinks(),
				Collections:   st.Store(),
			},
		},
	}
}

func (r *RootCollection) addRelations(st *state.State) error {
	for _, relation := range []*model.Relation{
		{
			Key:    bundle.RelationKeyTag.String(),
			Format: model.RelationFormat_tag,
		},
		{
			Key:    bundle.RelationKeyCreatedDate.String(),
			Format: model.RelationFormat_date,
		},
	} {
		err := addRelationsToCollectionDataView(st, relation)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RootCollection) getCreateCollectionRequest(collectionName string) *types.Struct {
	details := make(map[string]*types.Value, 0)
	details[bundle.RelationKeySourceFilePath.String()] = pbtypes.String(collectionName)
	details[bundle.RelationKeyName.String()] = pbtypes.String(collectionName)
	details[bundle.RelationKeyIsFavorite.String()] = pbtypes.Bool(true)
	details[bundle.RelationKeyLayout.String()] = pbtypes.Float64(float64(model.ObjectType_collection))

	detailsStruct := &types.Struct{Fields: details}
	return detailsStruct
}

func addRelationsToCollectionDataView(st *state.State, rel *model.Relation) error {
	return st.Iterate(func(bl simple.Block) (isContinue bool) {
		if dv, ok := bl.(simpleDataview.Block); ok {
			if len(bl.Model().GetDataview().GetViews()) == 0 {
				return false
			}
			for _, view := range bl.Model().GetDataview().GetViews() {
				err := dv.AddViewRelation(view.GetId(), &model.BlockContentDataviewRelation{
					Key:       rel.Key,
					IsVisible: true,
					Width:     192,
				})
				if err != nil {
					return false
				}
			}
			err := dv.AddRelation(&model.RelationLink{
				Key:    rel.Key,
				Format: rel.Format,
			})
			if err != nil {
				return false
			}
		}
		return true
	})
}
