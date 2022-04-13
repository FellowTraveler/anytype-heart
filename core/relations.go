package core

import (
	"context"
	"fmt"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
	"strings"

	"github.com/anytypeio/go-anytype-middleware/core/block/source"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/addr"
	"github.com/globalsign/mgo/bson"

	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/gogo/protobuf/types"
)

func (mw *Middleware) ObjectTypeRelationList(req *pb.RpcObjectTypeRelationListRequest) *pb.RpcObjectTypeRelationListResponse {
	response := func(code pb.RpcObjectTypeRelationListResponseErrorCode, relations []*model.Relation, err error) *pb.RpcObjectTypeRelationListResponse {
		m := &pb.RpcObjectTypeRelationListResponse{Relations: relations, Error: &pb.RpcObjectTypeRelationListResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	at := mw.GetAnytype()
	if at == nil {
		return response(pb.RpcObjectTypeRelationListResponseError_BAD_INPUT, nil, fmt.Errorf("account must be started"))
	}

	objType, err := mw.getObjectType(at, req.ObjectTypeUrl)
	if err != nil {
		if err == block.ErrUnknownObjectType {
			return response(pb.RpcObjectTypeRelationListResponseError_UNKNOWN_OBJECT_TYPE_URL, nil, err)
		}
		return response(pb.RpcObjectTypeRelationListResponseError_UNKNOWN_ERROR, nil, err)
	}

	// todo: AppendRelationsFromOtherTypes case
	return response(pb.RpcObjectTypeRelationListResponseError_NULL, objType.Relations, nil)
}

func (mw *Middleware) ObjectTypeRelationAdd(req *pb.RpcObjectTypeRelationAddRequest) *pb.RpcObjectTypeRelationAddResponse {
	response := func(code pb.RpcObjectTypeRelationAddResponseErrorCode, relations []*model.Relation, err error) *pb.RpcObjectTypeRelationAddResponse {
		m := &pb.RpcObjectTypeRelationAddResponse{Relations: relations, Error: &pb.RpcObjectTypeRelationAddResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	at := mw.GetAnytype()
	if at == nil {
		return response(pb.RpcObjectTypeRelationAddResponseError_BAD_INPUT, nil, fmt.Errorf("account must be started"))
	}

	objType, err := mw.getObjectType(at, req.ObjectTypeUrl)
	if err != nil {
		if err == block.ErrUnknownObjectType {
			return response(pb.RpcObjectTypeRelationAddResponseError_UNKNOWN_OBJECT_TYPE_URL, nil, err)
		}

		return response(pb.RpcObjectTypeRelationAddResponseError_UNKNOWN_ERROR, nil, err)
	}

	if strings.HasPrefix(objType.Url, bundle.TypePrefix) {
		return response(pb.RpcObjectTypeRelationAddResponseError_READONLY_OBJECT_TYPE, nil, fmt.Errorf("can't modify bundled object type"))
	}

	var relations []*model.Relation

	err = mw.doBlockService(func(bs block.Service) (err error) {
		relations, err = bs.AddExtraRelations(nil, objType.Url, req.Relations)
		if err != nil {
			return err
		}
		err = bs.ModifyDetails(objType.Url, func(current *types.Struct) (*types.Struct, error) {
			list := pbtypes.GetStringList(current, bundle.RelationKeyRecommendedRelations.String())
			for _, rel := range relations {
				var relId string
				if bundle.HasRelation(rel.Key) {
					relId = addr.BundledRelationURLPrefix + rel.Key
				} else {
					relId = addr.CustomRelationURLPrefix + rel.Key
				}

				if slice.FindPos(list, relId) == -1 {
					list = append(list, relId)
				}
			}
			detCopy := pbtypes.CopyStruct(current)
			detCopy.Fields[bundle.RelationKeyRecommendedRelations.String()] = pbtypes.StringList(list)
			return detCopy, nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return response(pb.RpcObjectTypeRelationAddResponseError_UNKNOWN_ERROR, nil, err)
	}

	return response(pb.RpcObjectTypeRelationAddResponseError_NULL, relations, nil)
}

func (mw *Middleware) ObjectTypeRelationUpdate(req *pb.RpcObjectTypeRelationUpdateRequest) *pb.RpcObjectTypeRelationUpdateResponse {
	response := func(code pb.RpcObjectTypeRelationUpdateResponseErrorCode, err error) *pb.RpcObjectTypeRelationUpdateResponse {
		m := &pb.RpcObjectTypeRelationUpdateResponse{Error: &pb.RpcObjectTypeRelationUpdateResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	at := mw.GetAnytype()
	if at == nil {
		return response(pb.RpcObjectTypeRelationUpdateResponseError_BAD_INPUT, fmt.Errorf("account must be started"))
	}

	objType, err := mw.getObjectType(at, req.ObjectTypeUrl)
	if err != nil {
		if err == block.ErrUnknownObjectType {
			return response(pb.RpcObjectTypeRelationUpdateResponseError_UNKNOWN_OBJECT_TYPE_URL, err)
		}

		return response(pb.RpcObjectTypeRelationUpdateResponseError_UNKNOWN_ERROR, err)
	}

	if strings.HasPrefix(objType.Url, bundle.TypePrefix) {
		return response(pb.RpcObjectTypeRelationUpdateResponseError_READONLY_OBJECT_TYPE, fmt.Errorf("can't modify bundled object type"))
	}

	err = mw.doBlockService(func(bs block.Service) (err error) {
		err = bs.UpdateExtraRelations(nil, objType.Url, []*model.Relation{req.Relation}, false)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return response(pb.RpcObjectTypeRelationUpdateResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcObjectTypeRelationUpdateResponseError_NULL, nil)
}

func (mw *Middleware) ObjectTypeRelationRemove(req *pb.RpcObjectTypeRelationRemoveRequest) *pb.RpcObjectTypeRelationRemoveResponse {
	response := func(code pb.RpcObjectTypeRelationRemoveResponseErrorCode, err error) *pb.RpcObjectTypeRelationRemoveResponse {
		m := &pb.RpcObjectTypeRelationRemoveResponse{Error: &pb.RpcObjectTypeRelationRemoveResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	at := mw.GetAnytype()
	if at == nil {
		return response(pb.RpcObjectTypeRelationRemoveResponseError_BAD_INPUT, fmt.Errorf("account must be started"))
	}

	objType, err := mw.getObjectType(at, req.ObjectTypeUrl)
	if err != nil {
		if err == block.ErrUnknownObjectType {
			return response(pb.RpcObjectTypeRelationRemoveResponseError_UNKNOWN_OBJECT_TYPE_URL, err)
		}

		return response(pb.RpcObjectTypeRelationRemoveResponseError_UNKNOWN_ERROR, err)
	}

	if strings.HasPrefix(objType.Url, bundle.TypePrefix) {
		return response(pb.RpcObjectTypeRelationRemoveResponseError_READONLY_OBJECT_TYPE, fmt.Errorf("can't modify bundled object type"))
	}

	err = mw.doBlockService(func(bs block.Service) (err error) {
		err = bs.ModifyDetails(objType.Url, func(current *types.Struct) (*types.Struct, error) {
			list := pbtypes.GetStringList(current, bundle.RelationKeyRecommendedRelations.String())
			var relId string
			if bundle.HasRelation(req.RelationKey) {
				relId = addr.BundledRelationURLPrefix + req.RelationKey
			} else {
				relId = addr.CustomRelationURLPrefix + req.RelationKey
			}

			list = slice.Remove(list, relId)
			detCopy := pbtypes.CopyStruct(current)
			detCopy.Fields[bundle.RelationKeyRecommendedRelations.String()] = pbtypes.StringList(list)
			return detCopy, nil
		})
		if err != nil {
			return err
		}
		err = bs.RemoveExtraRelations(nil, objType.Url, []string{req.RelationKey})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return response(pb.RpcObjectTypeRelationRemoveResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcObjectTypeRelationRemoveResponseError_NULL, nil)
}

func (mw *Middleware) ObjectTypeCreate(req *pb.RpcObjectTypeCreateRequest) *pb.RpcObjectTypeCreateResponse {
	response := func(code pb.RpcObjectTypeCreateResponseErrorCode, otype *model.ObjectType, err error) *pb.RpcObjectTypeCreateResponse {
		m := &pb.RpcObjectTypeCreateResponse{ObjectType: otype, Error: &pb.RpcObjectTypeCreateResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	var sbId string
	var recommendedRelationKeys []string
	var relations = make([]*model.Relation, 0, len(req.ObjectType.Relations)+len(bundle.RequiredInternalRelations))

	layout, _ := bundle.GetLayout(req.ObjectType.Layout)
	if layout == nil {
		return response(pb.RpcObjectTypeCreateResponseError_BAD_INPUT, nil, fmt.Errorf("invalid layout"))
	}

	for _, rel := range bundle.RequiredInternalRelations {
		relations = append(relations, bundle.MustGetRelation(rel))
		recommendedRelationKeys = append(recommendedRelationKeys, addr.BundledRelationURLPrefix+rel.String())
	}

	for _, rel := range layout.RequiredRelations {
		if pbtypes.HasRelation(relations, rel.Key) {
			continue
		}
		relations = append(relations, pbtypes.CopyRelation(rel))
		recommendedRelationKeys = append(recommendedRelationKeys, addr.BundledRelationURLPrefix+rel.Key)
	}

	for i, rel := range req.ObjectType.Relations {
		if v := pbtypes.GetRelation(relations, rel.Key); v != nil {
			if !pbtypes.RelationEqual(v, rel) {
				return response(pb.RpcObjectTypeCreateResponseError_BAD_INPUT, nil, fmt.Errorf("required relation %s not equals the bundled one", rel.Key))
			}
		} else {
			if rel.Key == "" {
				rel.Key = bson.NewObjectId().Hex()
			}
			rel.Creator = mw.GetAnytype().ProfileID()
			if bundle.HasRelation(rel.Key) {
				recommendedRelationKeys = append(recommendedRelationKeys, addr.BundledRelationURLPrefix+rel.Key)
			} else {
				recommendedRelationKeys = append(recommendedRelationKeys, addr.CustomRelationURLPrefix+rel.Key)
			}
			relations = append(relations, req.ObjectType.Relations[i])
		}
	}

	err := mw.doBlockService(func(bs block.Service) (err error) {
		sbId, _, err = bs.CreateSmartBlock(context.TODO(), smartblock.SmartBlockTypeObjectType, &types.Struct{
			Fields: map[string]*types.Value{
				bundle.RelationKeyName.String():                 pbtypes.String(req.ObjectType.Name),
				bundle.RelationKeyIconEmoji.String():            pbtypes.String(req.ObjectType.IconEmoji),
				bundle.RelationKeyType.String():                 pbtypes.String(bundle.TypeKeyObjectType.URL()),
				bundle.RelationKeyLayout.String():               pbtypes.Float64(float64(model.ObjectType_objectType)),
				bundle.RelationKeyRecommendedLayout.String():    pbtypes.Float64(float64(req.ObjectType.Layout)),
				bundle.RelationKeyRecommendedRelations.String(): pbtypes.StringList(recommendedRelationKeys),
				bundle.RelationKeyIsArchived.String():           pbtypes.Bool(req.ObjectType.IsArchived),
			},
		}, req.ObjectType.Relations)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return response(pb.RpcObjectTypeCreateResponseError_UNKNOWN_ERROR, nil, err)
	}

	otype := req.ObjectType
	otype.Relations = relations
	otype.Url = sbId
	otype.Types = []model.SmartBlockType{model.SmartBlockType_Page}
	return response(pb.RpcObjectTypeCreateResponseError_NULL, otype, nil)
}

func (mw *Middleware) ObjectTypeList(_ *pb.RpcObjectTypeListRequest) *pb.RpcObjectTypeListResponse {
	response := func(code pb.RpcObjectTypeListResponseErrorCode, otypes []*model.ObjectType, err error) *pb.RpcObjectTypeListResponse {
		m := &pb.RpcObjectTypeListResponse{ObjectTypes: otypes, Error: &pb.RpcObjectTypeListResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	at := mw.GetAnytype()
	if at == nil {
		return response(pb.RpcObjectTypeListResponseError_BAD_INPUT, nil, fmt.Errorf("account must be started"))
	}

	var (
		ids    []string
		otypes []*model.ObjectType
	)
	for _, t := range []smartblock.SmartBlockType{smartblock.SmartBlockTypeObjectType, smartblock.SmartBlockTypeBundledObjectType} {
		st, err := mw.GetApp().MustComponent(source.CName).(source.Service).SourceTypeBySbType(t)
		if err != nil {
			return response(pb.RpcObjectTypeListResponseError_UNKNOWN_ERROR, nil, err)
		}
		idsT, err := st.ListIds()
		if err != nil {
			return response(pb.RpcObjectTypeListResponseError_UNKNOWN_ERROR, nil, err)
		}
		ids = append(ids, idsT...)
	}

	for _, id := range ids {
		otype, err := mw.getObjectType(at, id)
		if err != nil {
			log.Errorf("failed to get objectType %s info: %s", id, err.Error())
			continue
		}
		otypes = append(otypes, otype)
	}

	return response(pb.RpcObjectTypeListResponseError_NULL, otypes, nil)
}

func (mw *Middleware) ObjectCreateSet(req *pb.RpcObjectCreateSetRequest) *pb.RpcObjectCreateSetResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectCreateSetResponseErrorCode, id string, err error) *pb.RpcObjectCreateSetResponse {
		m := &pb.RpcObjectCreateSetResponse{Error: &pb.RpcObjectCreateSetResponseError{Code: code}, Id: id}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}

	var id string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		req.Details.Fields[bundle.RelationKeySetOf.String()] = pbtypes.StringList(req.Source)
		id, _, err = bs.CreateSmartBlock(context.Background(), smartblock.SmartBlockTypeSet, req.Details, nil)
		return err
	})

	if err != nil {
		if err == block.ErrUnknownObjectType {
			return response(pb.RpcObjectCreateSetResponseError_UNKNOWN_OBJECT_TYPE_URL, "", err)
		}
		return response(pb.RpcObjectCreateSetResponseError_UNKNOWN_ERROR, "", err)
	}

	return response(pb.RpcObjectCreateSetResponseError_NULL, id, nil)
}

func (mw *Middleware) getObjectType(at core.Service, url string) (*model.ObjectType, error) {
	return objectstore.GetObjectType(at.ObjectStore(), url)
}
