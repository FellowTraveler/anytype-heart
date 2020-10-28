package editor

import (
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/basic"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/bookmark"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/clipboard"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/file"
	_import "github.com/anytypeio/go-anytype-middleware/core/block/editor/import"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/stext"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/meta"
	"github.com/anytypeio/go-anytype-middleware/core/block/source"
	"github.com/anytypeio/go-anytype-middleware/core/status"
	"github.com/anytypeio/go-anytype-middleware/util/linkpreview"
)

func NewPage(
	m meta.Service,
	fileSource file.BlockService,
	bCtrl bookmark.DoBookmark,
	importServices _import.Services,
	lp linkpreview.LinkPreview,
	ss status.Service,
) *Page {
	sb := smartblock.New(m, ss)
	f := file.NewFile(sb, fileSource)
	return &Page{
		SmartBlock: sb,
		Basic:      basic.NewBasic(sb),
		IHistory:   basic.NewHistory(sb),
		Text:       stext.NewText(sb),
		File:       f,
		Clipboard:  clipboard.NewClipboard(sb, f),
		Bookmark:   bookmark.NewBookmark(sb, lp, bCtrl),
		Import:     _import.NewImport(sb, importServices),
	}
}

type Page struct {
	smartblock.SmartBlock
	basic.Basic
	basic.IHistory
	file.File
	stext.Text
	clipboard.Clipboard
	bookmark.Bookmark
	_import.Import
}

func (p *Page) Init(s source.Source, _ bool) (err error) {
	if err = p.SmartBlock.Init(s, true); err != nil {
		return
	}
	return template.ApplyTemplate(p, template.WithTitle, nil)
}
