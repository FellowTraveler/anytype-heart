package block

import (
	"context"
	"fmt"
	"time"

	"github.com/anyproto/anytype-heart/core/block/editor/basic"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/editor/state"
	"github.com/anyproto/anytype-heart/core/block/object/objectcreator"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/session"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/internalflag"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

func (s *Service) ObjectDuplicate(ctx context.Context, id string) (objectID string, err error) {
	var (
		st             *state.State
		objectTypeKeys []domain.TypeKey
	)
	if err = Do(s, id, func(b smartblock.SmartBlock) error {
		objectTypeKeys = b.ObjectTypeKeys()
		if err = b.Restrictions().Object.Check(model.Restrictions_Duplicate); err != nil {
			return err
		}
		st = b.NewState().Copy()
		st.SetLocalDetails(nil)
		st.SetDetail(bundle.RelationKeySourceObject.String(), pbtypes.String(id))
		return nil
	}); err != nil {
		return
	}

	spaceID, err := s.resolver.ResolveSpaceID(id)
	if err != nil {
		return "", fmt.Errorf("resolve spaceID: %w", err)
	}
	objectID, _, err = s.objectCreator.CreateSmartBlockFromState(ctx, spaceID, objectTypeKeys, st)
	if err != nil {
		return
	}
	return
}

func (s *Service) CreateWorkspace(ctx context.Context, req *pb.RpcWorkspaceCreateRequest) (spaceID string, err error) {
	newSpace, err := s.spaceService.Create(ctx)
	if err != nil {
		return "", fmt.Errorf("error creating space: %w", err)
	}
	predefinedObjectIDs := newSpace.DerivedIDs()

	err = Do(s, predefinedObjectIDs.Workspace, func(b basic.DetailsSettable) error {
		details := make([]*pb.RpcObjectSetDetailsDetail, 0, len(req.Details.GetFields()))
		for k, v := range req.Details.GetFields() {
			details = append(details, &pb.RpcObjectSetDetailsDetail{
				Key:   k,
				Value: v,
			})
		}
		return b.SetDetails(nil, details, true)
	})
	if err != nil {
		return "", fmt.Errorf("set details for space %s: %w", newSpace.Id(), err)
	}
	_, err = s.builtinObjectService.CreateObjectsForUseCase(nil, newSpace.Id(), req.UseCase)
	if err != nil {
		return "", fmt.Errorf("import use-case: %w", err)
	}
	return newSpace.Id(), err
}

// CreateLinkToTheNewObject creates an object and stores the link to it in the context block
func (s *Service) CreateLinkToTheNewObject(
	ctx context.Context,
	sctx session.Context,
	req *pb.RpcBlockLinkCreateWithObjectRequest,
) (linkID string, objectID string, err error) {
	if req.ContextId == req.TemplateId && req.ContextId != "" {
		err = fmt.Errorf("unable to create link to template from this template")
		return
	}

	objectTypeKey, err := domain.GetTypeKeyFromRawUniqueKey(req.ObjectTypeUniqueKey)
	if err != nil {
		return "", "", fmt.Errorf("get type key from raw unique key: %w", err)
	}

	createReq := objectcreator.CreateObjectRequest{
		Details:       req.Details,
		InternalFlags: req.InternalFlags,
		ObjectTypeKey: objectTypeKey,
		TemplateId:    req.TemplateId,
	}
	objectID, _, err = s.objectCreator.CreateObject(ctx, req.SpaceId, createReq)
	if err != nil {
		return
	}

	if err = s.UpdateLastUsedDate(req.SpaceId, objectTypeKey); err != nil {
		log.Errorf("failed to update lastUsedDate of type object '%s': %w", objectTypeKey, err)
	}

	if req.ContextId == "" {
		return
	}

	err = DoStateCtx(s, sctx, req.ContextId, func(st *state.State, sb basic.Creatable) error {
		linkID, err = sb.CreateBlock(st, pb.RpcBlockCreateRequest{
			TargetId: req.TargetId,
			Block: &model.Block{
				Content: &model.BlockContentOfLink{
					Link: &model.BlockContentLink{
						TargetBlockId: objectID,
						Style:         model.BlockContentLink_Page,
					},
				},
				Fields: req.Fields,
			},
			Position: req.Position,
		})
		if err != nil {
			return fmt.Errorf("link create error: %w", err)
		}
		return nil
	})
	return
}

func (s *Service) ObjectToSet(id string, source []string) error {
	if err := Do(s, id, func(b smartblock.SmartBlock) error {
		commonOperations, ok := b.(basic.CommonOperations)
		if !ok {
			return fmt.Errorf("invalid smartblock impmlementation: %T", b)
		}
		st := b.NewState()
		st.SetDetail(bundle.RelationKeySetOf.String(), pbtypes.StringList(source))
		err := commonOperations.SetLayoutInStateAndIgnoreRestriction(st, model.ObjectType_set)
		if err != nil {
			return fmt.Errorf("set layout: %w", err)
		}
		st.SetObjectTypeKey(bundle.TypeKeySet)
		flags := internalflag.NewFromState(st)
		flags.Remove(model.InternalFlag_editorSelectType)
		flags.Remove(model.InternalFlag_editorDeleteEmpty)
		flags.AddToState(st)

		return b.Apply(st)
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateLastUsedDate(spaceId string, key domain.TypeKey) error {
	uk, err := domain.UnmarshalUniqueKey(key.URL())
	if err != nil {
		return fmt.Errorf("failed to unmarshall type key '%s': %w", key.String(), err)
	}
	details, err := s.objectStore.GetObjectByUniqueKey(spaceId, uk)
	if err != nil {
		return fmt.Errorf("failed to get details of type object '%s': %w", key.String(), err)
	}
	id := pbtypes.GetString(details.Details, bundle.RelationKeyId.String())
	if id == "" {
		return fmt.Errorf("failed to get id from details of type object '%s': %w", key.String(), err)
	}

	return DoState(s, id, func(st *state.State, sb smartblock.SmartBlock) error {
		st.SetLocalDetail(bundle.RelationKeyLastUsedDate.String(), pbtypes.Int64(time.Now().Unix()))
		return nil
	})
}
