package editor

import (
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/basic"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/widget"
	"github.com/anytypeio/go-anytype-middleware/core/block/migration"
	"github.com/anytypeio/go-anytype-middleware/core/session"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
)

const BlockAdditionError = "failed to add widget '%s': %v"

type WidgetObject struct {
	smartblock.SmartBlock
	basic.IHistory
	basic.Movable
	basic.Unlinkable
	basic.Updatable
	widget.Widget
}

func NewWidgetObject() *WidgetObject {
	sb := smartblock.New()
	bs := basic.NewBasic(sb)
	return &WidgetObject{
		SmartBlock: sb,
		Movable:    bs,
		Updatable:  bs,
		IHistory:   basic.NewHistory(sb),
		Widget:     widget.NewWidget(sb),
	}
}

func (w *WidgetObject) Init(ctx *smartblock.InitContext) (err error) {
	if err = w.SmartBlock.Init(ctx); err != nil {
		return
	}

	return nil
}

func (w *WidgetObject) CreationStateMigration(ctx *smartblock.InitContext) migration.Migration {
	return migration.Migration{
		Version: 1,
		Proc: func(st *state.State) {
			template.InitTemplate(st,
				template.WithEmpty,
				template.WithObjectTypesAndLayout([]string{bundle.TypeKeyDashboard.URL()}, model.ObjectType_basic),
				template.WithDetail(bundle.RelationKeyIsHidden, pbtypes.Bool(true)),
				w.withDefaultWidgets,
			)
		},
	}
}

func (w *WidgetObject) withDefaultWidgets(st *state.State) {
	for _, id := range []string{
		widget.DefaultWidgetFavorite,
		widget.DefaultWidgetSet,
		widget.DefaultWidgetRecent,
	} {
		if _, err := w.CreateBlock(st, &pb.RpcBlockCreateWidgetRequest{
			TargetId:     "",
			Position:     model.Block_Bottom,
			WidgetLayout: widget.LayoutList,
			Block: &model.Block{
				Id:          "",
				ChildrenIds: nil,
				Content: &model.BlockContentOfLink{
					Link: &model.BlockContentLink{
						TargetBlockId: id,
						Style:         model.BlockContentLink_Page,
						IconSize:      model.BlockContentLink_SizeNone,
						CardStyle:     model.BlockContentLink_Text,
						Description:   model.BlockContentLink_None,
					},
				},
			},
		}); err != nil {
			log.Errorf(BlockAdditionError, widget.DefaultWidgetFavorite, err)
		}
	}
}

func (w *WidgetObject) StateMigrations() migration.Migrations {
	return migration.MakeMigrations(nil)
}

func (w *WidgetObject) Unlink(ctx *session.Context, ids ...string) (err error) {
	st := w.NewStateCtx(ctx)
	for _, id := range ids {
		if p := st.PickParentOf(id); p != nil && p.Model().GetWidget() != nil {
			st.Unlink(p.Model().Id)
		}
		st.Unlink(id)
	}
	return w.Apply(st)
}
