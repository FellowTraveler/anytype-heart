package editor

import (
	"fmt"
	"github.com/anytypeio/go-anytype-middleware/core/block/database"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/source"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	smartblock2 "github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	database2 "github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/addr"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/threads"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/util"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/gogo/protobuf/types"
	"github.com/google/uuid"
	"github.com/textileio/go-threads/core/thread"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
)

const (
	collectionKeySignature = "signature"
	collectionKeyAccount   = "account"
	collectionKeyProfileId = "profileId"
	collectionKeyAddrs     = "addrs"
	collectionKeyId        = "id"
	collectionKeyKey       = "key"
)

func NewWorkspace(dbCtrl database.Ctrl, dmservice DetailsModifier) *Workspaces {
	s := NewSet(dbCtrl)
	return &Workspaces{
		Set:             NewSet(dbCtrl),
		DetailsModifier: dmservice,
		threadService:   s.Anytype().ThreadsService(),
		threadQueue:     s.Anytype().ThreadsService().ThreadQueue(),
	}
}

type Workspaces struct {
	*Set
	DetailsModifier DetailsModifier
	threadService   threads.Service
	threadQueue     threads.ThreadQueue
}

type WorkspaceParameters struct {
	IsHighlighted bool
}

func (wp *WorkspaceParameters) Equal(other *WorkspaceParameters) bool {
	return wp.IsHighlighted == other.IsHighlighted
}

func (p *Workspaces) CreateObject(sbType smartblock2.SmartBlockType) (core.SmartBlock, error) {
	st := p.NewState()
	threadInfo, err := p.threadQueue.CreateThreadSync(sbType, p.Id())
	if err != nil {
		return nil, err
	}
	st.SetInCollection(source.WorkspaceCollection, p.Id(), p.pbThreadInfoValueFromStruct(threadInfo))

	return core.NewSmartBlock(threadInfo, p.Anytype()), p.Apply(st)
}

func (p *Workspaces) DeleteObject(objectId string) error {
	st := p.NewState()
	err := p.threadQueue.DeleteThreadSync(objectId, p.Id())
	if err != nil {
		return err
	}
	st.RemoveFromCollection(source.WorkspaceCollection, p.Id())
	return p.Apply(st)
}

func (p *Workspaces) GetAllObjects() []string {
	st := p.NewState()
	workspaceCollection := st.GetCollection(source.WorkspaceCollection)
	if workspaceCollection == nil || workspaceCollection.Fields == nil {
		return nil
	}
	objects := make([]string, 0, len(workspaceCollection.Fields))
	for objId, workspaceId := range workspaceCollection.Fields {
		if v, ok := workspaceId.Kind.(*types.Value_StringValue); ok && v.StringValue == p.Id() {
			objects = append(objects, objId)
		}
	}
	return objects
}

func (p *Workspaces) AddCreatorInfoIfNeeded() error {
	st := p.NewState()
	deviceId := p.Anytype().Device()

	creatorCollection := st.GetCollection(source.CreatorCollection)
	if creatorCollection != nil && creatorCollection.Fields != nil && creatorCollection.Fields[deviceId] == nil {
		return nil
	}
	info, err := p.threadService.GetCreatorInfo(p.Id())
	if err != nil {
		return err
	}
	st.SetInCollection(source.CreatorCollection, deviceId, p.pbCreatorInfoValue(info))

	return p.Apply(st)
}

func (p *Workspaces) AddObject(objectId string, key string, addrs []string) error {
	st := p.NewState()
	err := p.threadQueue.AddThreadSync(threads.ThreadInfo{
		ID:    objectId,
		Key:   key,
		Addrs: addrs,
	}, p.Id())
	if err != nil {
		return err
	}
	st.SetInCollection(source.WorkspaceCollection, p.Id(), p.pbThreadInfoValue(objectId, key, addrs))

	return p.Apply(st)
}

func (p *Workspaces) GetObjectKeyAddrs(objectId string) (string, []string, error) {
	st := p.NewState()
	if !st.ContainsInCollection(source.WorkspaceCollection, objectId) {
		return "", nil, fmt.Errorf("%s is not contained in workspace %s", objectId, p.Id())
	}
	threadId, err := thread.Decode(objectId)
	if err != nil {
		return "", nil, fmt.Errorf("failed to decode object %s: %w", objectId, err)
	}

	threadInfo, err := p.threadService.GetThreadInfo(threadId)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get info on the thread %s: %w", objectId, err)
	}

	return threadInfo.Key.String(), util.MultiAddressesToStrings(threadInfo.Addrs), nil
}

func (p *Workspaces) SetIsHighlighted(objectId string, value bool) error {
	// TODO: this should be removed probably in the future?
	st := p.NewState()
	st.SetInCollection(source.HighlightedCollection, objectId, pbtypes.Bool(value))
	return p.Apply(st)
}

func (p *Workspaces) Init(ctx *smartblock.InitContext) (err error) {
	if ctx.Source.Type() != model.SmartBlockType_Workspace && ctx.Source.Type() != model.SmartBlockType_AccountOld {
		return fmt.Errorf("source type should be a workspace or an old account")
	}

	if err = p.SmartBlock.Init(ctx); err != nil {
		return
	}

	dataviewAllHighlightedObjects := model.BlockContentOfDataview{
		Dataview: &model.BlockContentDataview{
			Source:    []string{addr.BundledRelationURLPrefix + bundle.RelationKeyName.String()},
			Relations: []*model.Relation{bundle.MustGetRelation(bundle.RelationKeyName)},
			Views: []*model.BlockContentDataviewView{
				{
					Id:   uuid.New().String(),
					Type: model.BlockContentDataviewView_Gallery,
					Name: "Highlighted",
					Sorts: []*model.BlockContentDataviewSort{
						{
							RelationKey: "name",
							Type:        model.BlockContentDataviewSort_Asc,
						},
					},
					Relations: []*model.BlockContentDataviewRelation{
						{
							Key:       bundle.RelationKeyName.String(),
							IsVisible: true,
						},
						{
							Key:       bundle.RelationKeyCreator.String(),
							IsVisible: true,
						},
					},
					Filters: []*model.BlockContentDataviewFilter{{
						RelationKey: bundle.RelationKeyWorkspaceId.String(),
						Condition:   model.BlockContentDataviewFilter_Equal,
						Value:       pbtypes.String(p.Id()),
					}, {
						RelationKey: bundle.RelationKeyId.String(),
						Condition:   model.BlockContentDataviewFilter_NotEqual,
						Value:       pbtypes.String(p.Id()),
					}, {
						RelationKey: bundle.RelationKeyIsHighlighted.String(),
						Condition:   model.BlockContentDataviewFilter_Equal,
						Value:       pbtypes.Bool(true),
					}},
				},
			},
		},
	}

	dataviewAllWorkspaceObjects := model.BlockContentOfDataview{
		Dataview: &model.BlockContentDataview{
			Source:    []string{addr.BundledRelationURLPrefix + bundle.RelationKeyName.String()},
			Relations: []*model.Relation{bundle.MustGetRelation(bundle.RelationKeyName), bundle.MustGetRelation(bundle.RelationKeyCreator)},
			Views: []*model.BlockContentDataviewView{
				{
					Id:   uuid.New().String(),
					Type: model.BlockContentDataviewView_Table,
					Name: "All",
					Sorts: []*model.BlockContentDataviewSort{
						{
							RelationKey: "name",
							Type:        model.BlockContentDataviewSort_Asc,
						},
					},
					Relations: []*model.BlockContentDataviewRelation{
						{
							Key:       bundle.RelationKeyName.String(),
							IsVisible: true,
						},
						{
							Key:       bundle.RelationKeyCreator.String(),
							IsVisible: true,
						},
					},
					Filters: []*model.BlockContentDataviewFilter{{
						RelationKey: bundle.RelationKeyWorkspaceId.String(),
						Condition:   model.BlockContentDataviewFilter_Equal,
						Value:       pbtypes.String(p.Id()),
					}, {
						RelationKey: bundle.RelationKeyId.String(),
						Condition:   model.BlockContentDataviewFilter_NotEqual,
						Value:       pbtypes.String(p.Id()),
					}},
				},
			},
		},
	}

	p.AddHook(p.updateObjects, smartblock.HookAfterApply)
	err = smartblock.ApplyTemplate(p, ctx.State,
		template.WithEmpty,
		template.WithTitle,
		template.WithFeaturedRelations,
		template.WithForcedDetail(bundle.RelationKeyFeaturedRelations, pbtypes.StringList([]string{bundle.RelationKeyType.String(), bundle.RelationKeyCreator.String()})),
		template.WithDataviewID("highlighted", dataviewAllHighlightedObjects, true),
		template.WithDataviewID("dataview", dataviewAllWorkspaceObjects, true),
	)
	if err != nil {
		return err
	}
	defaultValue := &types.Struct{Fields: map[string]*types.Value{bundle.RelationKeyWorkspaceId.String(): pbtypes.String(p.Id())}}
	return p.Set.SetNewRecordDefaultFields("dataview", defaultValue)
}

// TODO: try to save results from processing of previous state and get changes from apply for performance
func (p *Workspaces) updateObjects() {
	st := p.NewState()

	objects, parameters := p.workspaceObjectsAndParametersFromState(st)
	p.threadQueue.ProcessThreadsAsync(objects, p.Id())
	// TODO: Also include old account
	if p.Id() != p.Anytype().PredefinedBlocks().Account {
		storedParameters := p.workspaceParametersFromRecords(p.storedRecordsForWorkspace())
		// we ignore the workspace object itself
		delete(storedParameters, p.Id())
		p.updateDetailsIfParametersChanged(storedParameters, parameters)
	}
}

func (p *Workspaces) storedRecordsForWorkspace() []database2.Record {
	records, _, err := p.ObjectStore().Query(nil, database2.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyWorkspaceId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(p.Id()),
			},
		},
	})
	if err != nil {
		log.Errorf("workspace: can't get store workspace ids: %v", err)
		return nil
	}
	return records
}

func (p *Workspaces) workspaceParametersFromRecords(records []database2.Record) map[string]*WorkspaceParameters {
	var storeObjectInWorkspace = make(map[string]*WorkspaceParameters, len(records))
	for _, rec := range records {
		id := pbtypes.GetString(rec.Details, bundle.RelationKeyId.String())
		storeObjectInWorkspace[id] = &WorkspaceParameters{
			IsHighlighted: pbtypes.GetBool(rec.Details, bundle.RelationKeyIsHighlighted.String()),
		}
	}
	return storeObjectInWorkspace
}

func (p *Workspaces) workspaceObjectsAndParametersFromState(st *state.State) ([]threads.ThreadInfo, map[string]*WorkspaceParameters) {
	workspaceCollection := st.GetCollection(source.WorkspaceCollection)
	if workspaceCollection == nil || workspaceCollection.Fields == nil {
		return nil, nil
	}
	parameters := make(map[string]*WorkspaceParameters, len(workspaceCollection.Fields))
	objects := make([]threads.ThreadInfo, 0, len(workspaceCollection.Fields))
	for objId, value := range workspaceCollection.Fields {
		if value == nil {
			continue
		}
		parameters[objId] = &WorkspaceParameters{IsHighlighted: false}
		objects = append(objects, p.threadInfoFromWorkspacePB(value))
	}

	creatorCollection := st.GetCollection(source.CreatorCollection)
	if creatorCollection != nil {
		for _, value := range creatorCollection.Fields {
			info, err := p.threadInfoFromCreatorPB(value)
			if err != nil {
				continue
			}
			objects = append(objects, info)
		}
	}
	highlightedCollection := st.GetCollection(source.HighlightedCollection)
	if highlightedCollection != nil {
		for objId, isHighlighted := range highlightedCollection.Fields {
			if isHighlighted == nil {
				continue
			}
			if v, ok := isHighlighted.Kind.(*types.Value_BoolValue); ok && v.BoolValue {
				if _, exists := parameters[objId]; exists {
					parameters[objId].IsHighlighted = true
				}
			}
		}
	}

	return objects, parameters
}

func (p *Workspaces) updateDetailsIfParametersChanged(
	oldParameters map[string]*WorkspaceParameters,
	newParameters map[string]*WorkspaceParameters) {
	for id, params := range newParameters {
		if oldParams, exists := oldParameters[id]; exists && oldParams.Equal(params) {
			continue
		}

		if err := p.DetailsModifier.ModifyDetails(id, func(current *types.Struct) (*types.Struct, error) {
			if current == nil || current.Fields == nil {
				current = &types.Struct{
					Fields: map[string]*types.Value{},
				}
			}
			current.Fields[bundle.RelationKeyWorkspaceId.String()] = pbtypes.String(p.Id())
			current.Fields[bundle.RelationKeyIsHighlighted.String()] = pbtypes.Bool(params.IsHighlighted)

			return current, nil
		}); err != nil {
			log.Errorf("workspace: can't set detail to object: %v", err)
		}
	}
}

func (p *Workspaces) pbCreatorInfoValue(info threads.CreatorInfo) *types.Value {
	return &types.Value{
		Kind: &types.Value_StructValue{
			StructValue: &types.Struct{
				Fields: map[string]*types.Value{
					collectionKeyAccount:   pbtypes.String(info.AccountPubKey),
					collectionKeySignature: pbtypes.String(string(info.WorkspaceSig)),
					collectionKeyProfileId: pbtypes.String(info.ProfileId),
					collectionKeyAddrs:     pbtypes.StringList(info.Addrs),
				},
			},
		},
	}
}

func (p *Workspaces) pbThreadInfoValue(id string, key string, addrs []string) *types.Value {
	return &types.Value{
		Kind: &types.Value_StructValue{
			StructValue: &types.Struct{
				Fields: map[string]*types.Value{
					collectionKeyId:    pbtypes.String(id),
					collectionKeyKey:   pbtypes.String(key),
					collectionKeyAddrs: pbtypes.StringList(addrs),
				},
			},
		},
	}
}

func (p *Workspaces) pbThreadInfoValueFromStruct(ti thread.Info) *types.Value {
	return p.pbThreadInfoValue(ti.ID.String(), ti.Key.String(), util.MultiAddressesToStrings(ti.Addrs))
}

func (p *Workspaces) threadInfoFromWorkspacePB(val *types.Value) threads.ThreadInfo {
	fields := val.Kind.(*types.Value_StructValue).StructValue.Fields
	return threads.ThreadInfo{
		ID:    fields[collectionKeyId].Kind.(*types.Value_StringValue).String(),
		Key:   fields[collectionKeyKey].Kind.(*types.Value_StringValue).String(),
		Addrs: pbtypes.GetStringListValue(fields[collectionKeyAddrs]),
	}
}

func (p *Workspaces) threadInfoFromCreatorPB(val *types.Value) (threads.ThreadInfo, error) {
	fields := val.Kind.(*types.Value_StructValue).StructValue.Fields
	account := fields[collectionKeyAccount].Kind.(*types.Value_StringValue).String()
	profileId, err := threads.ProfileThreadIDFromAccountAddress(account)
	if err != nil {
		return threads.ThreadInfo{}, err
	}
	sk, pk, err := threads.ProfileThreadKeysFromAccountAddress(account)
	if err != nil {
		return threads.ThreadInfo{}, err
	}
	return threads.ThreadInfo{
		ID:    profileId.String(),
		Key:   thread.NewKey(sk, pk).String(),
		Addrs: pbtypes.GetStringListValue(fields[collectionKeyAddrs]),
	}, nil
}
