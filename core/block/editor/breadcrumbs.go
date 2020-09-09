package editor

import (
	"github.com/anytypeio/go-anytype-library/logging"
	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/meta"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/core/block/source"
)

var log = logging.Logger("anytype-mw-editor")

func NewBreadcrumbs(m meta.Service) *Breadcrumbs {
	return &Breadcrumbs{
		SmartBlock: smartblock.New(m),
	}
}

type Breadcrumbs struct {
	smartblock.SmartBlock
}

func (p *Breadcrumbs) Init(s source.Source, _ bool) (err error) {
	if err = p.SmartBlock.Init(s, true); err != nil {
		return
	}
	p.SmartBlock.DisableLayouts()
	return
}

func (b *Breadcrumbs) SetCrumbs(ids []string) (err error) {
	s := b.NewState()
	var existingLinks = make(map[string]string)
	s.Iterate(func(b simple.Block) (isContinue bool) {
		if link := b.Model().GetLink(); link != nil {
			existingLinks[link.TargetBlockId] = b.Model().Id
		}
		return true
	})
	root := s.Get(s.RootId()).Model()
	root.ChildrenIds = make([]string, 0, len(ids))
	for _, id := range ids {
		linkId, ok := existingLinks[id]
		if !ok {
			link := simple.New(&model.Block{
				Content: &model.BlockContentOfLink{
					Link: &model.BlockContentLink{
						TargetBlockId: id,
						Style:         model.BlockContentLink_Page,
					},
				},
			})
			s.Add(link)
			linkId = link.Model().Id
		}
		root.ChildrenIds = append(root.ChildrenIds, linkId)
	}
	return b.Apply(s)
}
