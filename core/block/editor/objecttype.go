package editor

import (
	"strings"

	"github.com/gogo/protobuf/types"

	"github.com/anytypeio/go-anytype-middleware/app"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/basic"
	dataview2 "github.com/anytypeio/go-anytype-middleware/core/block/editor/dataview"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/stext"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/restriction"
	relation2 "github.com/anytypeio/go-anytype-middleware/core/relation"
	"github.com/anytypeio/go-anytype-middleware/core/relation/relationutils"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/addr"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
)

type ObjectType struct {
	*Set

	relationService relation2.Service
}

func NewObjectType(anytype core.Service,
	objectStore objectstore.ObjectStore,
	relationService relation2.Service,
) *ObjectType {
	return &ObjectType{
		Set: NewSet(anytype, objectStore, relationService),

		relationService: relationService,
	}
}

func (p *ObjectType) Init(ctx *smartblock.InitContext) (err error) {
	p.CommonOperations = basic.NewBasic(p.SmartBlock)
	p.IHistory = basic.NewHistory(p.SmartBlock)
	p.Dataview = dataview2.NewDataview(
		p.SmartBlock,
		app.MustComponent[core.Service](ctx.App),
		app.MustComponent[objectstore.ObjectStore](ctx.App),
		app.MustComponent[relation2.Service](ctx.App),
	)
	p.Text = stext.NewText(
		p.SmartBlock,
		app.MustComponent[objectstore.ObjectStore](ctx.App),
	)

	p.relationService = app.MustComponent[relation2.Service](ctx.App)

	if err = p.SmartBlock.Init(ctx); err != nil {
		return
	}

	dataview := model.BlockContentOfDataview{
		Dataview: &model.BlockContentDataview{
			Source: []string{p.Id()},
			Views: []*model.BlockContentDataviewView{
				{
					Id:   "_view1_1",
					Type: model.BlockContentDataviewView_Table,
					Name: "All",
					Sorts: []*model.BlockContentDataviewSort{
						{
							RelationKey: "name",
							Type:        model.BlockContentDataviewSort_Asc,
						},
					},
					Relations: []*model.BlockContentDataviewRelation{},
					Filters:   nil,
				},
			},
		},
	}
	var templatesSource string
	var isBundled bool
	if strings.HasPrefix(p.Id(), addr.BundledObjectTypeURLPrefix) {
		isBundled = true
	}

	if isBundled {
		templatesSource = bundle.TypeKeyTemplate.BundledURL()
	} else {
		templatesSource = bundle.TypeKeyTemplate.URL()
	}

	templatesDataview := model.BlockContentOfDataview{
		Dataview: &model.BlockContentDataview{
			Source: []string{templatesSource},
			Views: []*model.BlockContentDataviewView{
				{
					Id:   "_view2_1",
					Type: model.BlockContentDataviewView_Table,
					Name: "All",
					Sorts: []*model.BlockContentDataviewSort{
						{
							RelationKey: "name",
							Type:        model.BlockContentDataviewSort_Asc,
						},
					},
					Relations: []*model.BlockContentDataviewRelation{},
					Filters: []*model.BlockContentDataviewFilter{
						{
							Operator:    model.BlockContentDataviewFilter_And,
							RelationKey: bundle.RelationKeyTargetObjectType.String(),
							Condition:   model.BlockContentDataviewFilter_Equal,
							Value:       pbtypes.String(p.RootId()),
						}},
				},
			},
		},
	}
	var recommendedRelationsKeys []string
	for _, relId := range pbtypes.GetStringList(ctx.State.Details(), bundle.RelationKeyRecommendedRelations.String()) {
		relKey, err := pbtypes.RelationIdToKey(relId)
		if err != nil {
			log.Errorf("recommendedRelations of %s has incorrect id: %s", p.Id(), relId)
			continue
		}
		if slice.FindPos(recommendedRelationsKeys, relKey) == -1 {
			recommendedRelationsKeys = append(recommendedRelationsKeys, relKey)
		}
	}

	// todo: remove this
	/*
		for _, rel := range bundle.RequiredInternalRelations {
			if slice.FindPos(recommendedRelationsKeys, rel.String()) == -1 {
				recommendedRelationsKeys = append(recommendedRelationsKeys, rel.String())
			}
		}*/

	recommendedLayout := pbtypes.GetString(p.Details(), bundle.RelationKeyRecommendedLayout.String())
	if recommendedLayout == "" {
		recommendedLayout = model.ObjectType_basic.String()
	} else if _, ok := model.ObjectTypeLayout_value[recommendedLayout]; !ok {
		recommendedLayout = model.ObjectType_basic.String()
	}

	recommendedLayoutObj := bundle.MustGetLayout(model.ObjectTypeLayout(model.ObjectTypeLayout_value[recommendedLayout]))
	for _, rel := range recommendedLayoutObj.RequiredRelations {
		if slice.FindPos(recommendedRelationsKeys, rel.Key) == -1 {
			recommendedRelationsKeys = append(recommendedRelationsKeys, rel.Key)
		}
	}

	// filter out internal relations from the recommended
	recommendedRelationsKeys = slice.Filter(recommendedRelationsKeys, func(relKey string) bool {
		for _, k := range bundle.RequiredInternalRelations {
			if k.String() == relKey {
				return false
			}
		}
		return true
	})

	var relIds []string
	var r *relationutils.Relation
	for _, rel := range recommendedRelationsKeys {
		if isBundled {
			relIds = append(relIds, addr.BundledRelationURLPrefix+rel)
		} else {
			relIds = append(relIds, addr.RelationKeyToIdPrefix+rel)
		}

		if r2, _ := bundle.GetRelation(bundle.RelationKey(rel)); r2 != nil {
			if r2.Hidden {
				continue
			}
			r = &relationutils.Relation{Relation: r2}
		} else {
			r, _ = p.relationService.FetchKey(rel)
			if r == nil {
				continue
			}
		}
		// add recommended relation to the dataview
		dataview.Dataview.RelationLinks = append(dataview.Dataview.RelationLinks, r.RelationLink())
		dataview.Dataview.Views[0].Relations = append(dataview.Dataview.Views[0].Relations, &model.BlockContentDataviewRelation{
			Key:       r.Key,
			IsVisible: true,
		})
	}

	defaultValue := &types.Struct{Fields: map[string]*types.Value{bundle.RelationKeyTargetObjectType.String(): pbtypes.String(p.RootId())}}

	if !isBundled {
		var system bool
		for _, o := range bundle.SystemTypes {
			if addr.ObjectTypeKeyToIdPrefix+o.String() == p.RootId() {
				system = true
				break
			}
		}

		var internal bool
		for _, o := range bundle.InternalTypes {
			if addr.ObjectTypeKeyToIdPrefix+o.String() == p.RootId() {
				internal = true
				break
			}
		}

		if system {
			rest := p.Restrictions()
			obj := append(rest.Object.Copy(), []model.RestrictionsObjectRestriction{model.Restrictions_Blocks, model.Restrictions_Details}...)
			dv := rest.Dataview.Copy()
			if internal {
				// internal mean not possible to create the object using the standard ObjectCreate flow
				dv = append(dv, model.RestrictionsDataviewRestrictions{BlockId: template.DataviewBlockId, Restrictions: []model.RestrictionsDataviewRestriction{model.Restrictions_DVCreateObject}})
			}
			p.SetRestrictions(restriction.Restrictions{Object: obj, Dataview: dv})

		}
	}

	fixMissingSmartblockTypes := func(s *state.State) {
		// we have a bug in internal release that was not adding smartblocktype to newly created custom types
		if len(pbtypes.GetIntList(s.Details(), bundle.RelationKeySmartblockTypes.String())) == 0 {
			s.SetDetailAndBundledRelation(bundle.RelationKeySmartblockTypes, pbtypes.IntList(int(model.SmartBlockType_Page)))
		}
	}

	return smartblock.ObjectApplyTemplate(p, ctx.State,
		template.WithObjectTypesAndLayout([]string{bundle.TypeKeyObjectType.URL()}, model.ObjectType_objectType),
		template.WithEmpty,
		template.WithTitle,
		template.WithDefaultFeaturedRelations,
		template.WithDescription,
		template.WithFeaturedRelations,
		template.WithDataviewID("templates", templatesDataview, true),
		template.WithDataview(dataview, true),
		template.WithForcedDetail(bundle.RelationKeyRecommendedRelations, pbtypes.StringList(relIds)),
		template.MigrateRelationValue(bundle.RelationKeySource, bundle.RelationKeySourceObject),
		template.WithChildrenSorter(p.RootId(), func(blockIds []string) {
			i := slice.FindPos(blockIds, "templates")
			j := slice.FindPos(blockIds, template.DataviewBlockId)
			// templates dataview must come before the type dataview
			if i > j {
				blockIds[i], blockIds[j] = blockIds[j], blockIds[i]
			}
		}),
		template.WithCondition(!isBundled, template.WithAddedFeaturedRelation(bundle.RelationKeySourceObject)),
		template.WithObjectTypeLayoutMigration(),
		template.WithRequiredRelations(),
		template.WithBlockField("templates", dataview2.DefaultDetailsFieldName, pbtypes.Struct(defaultValue)),
		fixMissingSmartblockTypes,
	)
}

func (o *ObjectType) SetStruct(st *types.Struct) error {
	o.Lock()
	defer o.Unlock()
	s := o.NewState()
	s.SetDetails(st)
	return o.Apply(s)
}
