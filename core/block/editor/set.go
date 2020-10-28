package editor

import (
	"fmt"

	"github.com/anytypeio/go-anytype-middleware/core/block/database"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/basic"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/dataview"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/meta"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/core/block/source"
	"github.com/anytypeio/go-anytype-middleware/core/status"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/google/uuid"
)

func NewSet(
	ms meta.Service,
	dbCtrl database.Ctrl,
	ss status.Service,
) *Set {
	sb := &Set{SmartBlock: smartblock.New(ms, ss)}

	sb.Basic = basic.NewBasic(sb)
	sb.IHistory = basic.NewHistory(sb)
	sb.Dataview = dataview.NewDataview(sb)
	sb.Router = database.New(dbCtrl)

	return sb
}

type Set struct {
	smartblock.SmartBlock
	basic.Basic
	basic.IHistory
	dataview.Dataview
	database.Router
}

func (p *Set) Init(s source.Source, _ bool) (err error) {
	if err = p.SmartBlock.Init(s, true); err != nil {
		return
	}
	return p.init()
}

func (p *Set) init() (err error) {
	s := p.NewState()
	if err = template.InitTemplate(template.WithTitle, s); err != nil {
		return
	}
	root := s.Get(p.RootId())
	setDetails := func() {
		s.SetDetail("name", pbtypes.String("Pages"))
		s.SetDetail("iconEmoji", pbtypes.String("📒"))
	}
	if len(root.Model().ChildrenIds) > 1 {
		return
	}
	// init dataview
	relations := []*model.BlockContentDataviewRelation{{Id: "id", IsVisible: false}, {Id: "name", IsVisible: true}, {Id: "lastOpened", IsVisible: true}, {Id: "lastModified", IsVisible: true}}
	dataview := simple.New(&model.Block{
		Content: &model.BlockContentOfDataview{
			Dataview: &model.BlockContentDataview{
				DatabaseId: "pages",
				SchemaURL:  "https://anytype.io/schemas/page",
				Views: []*model.BlockContentDataviewView{
					{
						Id:   uuid.New().String(),
						Type: model.BlockContentDataviewView_Table,
						Name: "All pages",
						Sorts: []*model.BlockContentDataviewSort{
							{
								RelationId: "name",
								Type:       model.BlockContentDataviewSort_Asc,
							},
						},
						Relations: relations,
						Filters:   nil,
					},
				},
			},
		},
	})

	s.Add(dataview)

	if err = s.InsertTo(template.HeaderLayoutId, model.Block_Bottom, dataview.Model().Id); err != nil {
		return fmt.Errorf("can't insert dataview: %v", err)
	}

	setDetails()
	log.Infof("create default structure for set: %v", s.RootId())
	return p.Apply(s, smartblock.NoEvent, smartblock.NoHistory)
}
