package editor

import (
	"github.com/anyproto/anytype-heart/core/block/editor/converter"
	"github.com/anyproto/anytype-heart/core/block/editor/file"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/editor/template"
	"github.com/anyproto/anytype-heart/core/files"
	"github.com/anyproto/anytype-heart/core/relation"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/core"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space/typeprovider"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

type RelationOption struct {
	*SubObject
}

func NewRelationOption(
	sb smartblock.SmartBlock,
	objectStore objectstore.ObjectStore,
	fileBlockService file.BlockService,
	anytype core.Service,
	relationService relation.Service,
	tempDirProvider core.TempDirProvider,
	sbtProvider typeprovider.SmartBlockTypeProvider,
	layoutConverter converter.LayoutConverter,
	fileService files.Service,
) *RelationOption {
	return &RelationOption{
		SubObject: NewSubObject(
			sb,
			objectStore,
			fileBlockService,
			anytype,
			relationService,
			tempDirProvider,
			sbtProvider,
			layoutConverter,
			fileService,
		),
	}
}

func (ro *RelationOption) Init(ctx *smartblock.InitContext) error {
	if err := ro.SubObject.Init(ctx); err != nil {
		return err
	}

	st := ctx.State

	relKey := pbtypes.GetString(st.Details(), bundle.RelationKeyRelationKey.String())
	dataview := model.BlockContentOfDataview{
		Dataview: &model.BlockContentDataview{
			Source: []string{relKey},
			Views: []*model.BlockContentDataviewView{
				{
					Id:   "1",
					Type: model.BlockContentDataviewView_Table,
					Name: "All",
					Sorts: []*model.BlockContentDataviewSort{
						{
							RelationKey: "name",
							Type:        model.BlockContentDataviewSort_Asc,
						},
					},
					Relations: []*model.BlockContentDataviewRelation{},
					Filters: []*model.BlockContentDataviewFilter{{
						RelationKey: relKey,
						Condition:   model.BlockContentDataviewFilter_In,
						Value:       pbtypes.String(st.RootId()),
					}},
				},
			},
		},
	}

	return smartblock.ObjectApplyTemplate(ro, st,
		template.WithAllBlocksEditsRestricted,
		template.WithForcedDetail(bundle.RelationKeyLayout, pbtypes.Int64(int64(model.ObjectType_relationOption))),
		template.WithForcedDetail(bundle.RelationKeyIsReadonly, pbtypes.Bool(false)),
		template.WithForcedDetail(bundle.RelationKeyType, pbtypes.String(bundle.TypeKeyRelationOption.URL())),
		template.WithTitle,
		template.WithDefaultFeaturedRelations,
		template.WithDataview(dataview, false))
}
