package block

import (
	"context"
	"errors"
	"fmt"
	"github.com/anytypeio/go-anytype-middleware/util"
	"strings"
	"time"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/table"
	"github.com/anytypeio/go-anytype-middleware/core/session"

	"github.com/anytypeio/go-anytype-middleware/core/block/simple/link"
	"github.com/anytypeio/go-anytype-middleware/core/block/source"
	"github.com/anytypeio/go-anytype-middleware/metrics"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/schema"
	"github.com/anytypeio/go-anytype-middleware/util/internalflag"
	"github.com/anytypeio/go-anytype-middleware/util/ocache"
	ds "github.com/ipfs/go-datastore"
	"github.com/textileio/go-threads/core/thread"

	"github.com/anytypeio/go-anytype-middleware/core/block/doc"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/basic"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/bookmark"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/clipboard"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/dataview"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/file"
	_import "github.com/anytypeio/go-anytype-middleware/core/block/editor/import"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/stext"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/text"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	coresb "github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/gogo/protobuf/types"
)

var ErrOptionUsedByOtherObjects = fmt.Errorf("option is used by other objects")

func (s *service) MarkArchived(id string, archived bool) (err error) {
	return s.Do(id, func(b smartblock.SmartBlock) error {
		return b.SetDetails(nil, []*pb.RpcObjectSetDetailsDetail{
			{
				Key:   "isArchived",
				Value: pbtypes.Bool(archived),
			},
		}, true)
	})
}

func (s *service) SetBreadcrumbs(ctx *session.Context, req pb.RpcObjectSetBreadcrumbsRequest) (err error) {
	return s.Do(req.BreadcrumbsId, func(b smartblock.SmartBlock) error {
		if breadcrumbs, ok := b.(*editor.Breadcrumbs); ok {
			return breadcrumbs.SetCrumbs(req.Ids)
		} else {
			return ErrUnexpectedBlockType
		}
	})
}

func (s *service) CreateBlock(ctx *session.Context, req pb.RpcBlockCreateRequest) (id string, err error) {
	err = s.DoBasic(req.ContextId, func(b basic.Basic) error {
		id, err = b.Create(ctx, "", req)
		return err
	})
	return
}

func (s *service) DuplicateBlocks(ctx *session.Context, req pb.RpcBlockListDuplicateRequest) (newIds []string, err error) {
	if req.ContextId == req.TargetContextId || req.TargetContextId == "" {
		err = s.Do(req.ContextId, func(sb smartblock.SmartBlock) error {
			if sb.Type() == model.SmartBlockType_Set {
				return basic.ErrNotSupported
			}

			st := sb.NewStateCtx(ctx)
			newIds, err = basic.Duplicate(req, st, st)
			if err != nil {
				return fmt.Errorf("duplicate: %w", err)
			}
			return sb.Apply(st)
		})
		return
	}

	err = s.Do(req.ContextId, func(sb smartblock.SmartBlock) error {
		srcState := sb.NewStateCtx(ctx)
		err = s.Do(req.TargetContextId, func(tb smartblock.SmartBlock) error {
			if tb.Type() == model.SmartBlockType_Set {
				return basic.ErrNotSupported
			}

			targetState := tb.NewState()
			newIds, err = basic.Duplicate(req, srcState, targetState)
			if err != nil {
				return fmt.Errorf("duplicate: %w", err)
			}
			return tb.Apply(targetState)
		})
		return sb.Apply(srcState)
	})

	return
}

func (s *service) UnlinkBlock(ctx *session.Context, req pb.RpcBlockListDeleteRequest) (err error) {
	return s.DoBasic(req.ContextId, func(b basic.Basic) error {
		return b.Unlink(ctx, req.BlockIds...)
	})
}

func (s *service) SetDivStyle(ctx *session.Context, contextId string, style model.BlockContentDivStyle, ids ...string) (err error) {
	return s.DoBasic(contextId, func(b basic.Basic) error {
		return b.SetDivStyle(ctx, style, ids...)
	})
}

func (s *service) SplitBlock(ctx *session.Context, req pb.RpcBlockSplitRequest) (blockId string, err error) {
	err = s.DoText(req.ContextId, func(b stext.Text) error {
		blockId, err = b.Split(ctx, req)
		return err
	})
	return
}

func (s *service) MergeBlock(ctx *session.Context, req pb.RpcBlockMergeRequest) (err error) {
	return s.DoText(req.ContextId, func(b stext.Text) error {
		return b.Merge(ctx, req.FirstBlockId, req.SecondBlockId)
	})
}

func (s *service) TurnInto(ctx *session.Context, contextId string, style model.BlockContentTextStyle, ids ...string) error {
	return s.DoText(contextId, func(b stext.Text) error {
		return b.TurnInto(ctx, style, ids...)
	})
}

func (s *service) SimplePaste(contextId string, anySlot []*model.Block) (err error) {
	var blocks []simple.Block

	for _, b := range anySlot {
		blocks = append(blocks, simple.New(b))
	}

	return s.DoBasic(contextId, func(b basic.Basic) error {
		return b.PasteBlocks(blocks)
	})
}

func (s *service) ReplaceBlock(ctx *session.Context, req pb.RpcBlockReplaceRequest) (newId string, err error) {
	err = s.DoBasic(req.ContextId, func(b basic.Basic) error {
		newId, err = b.Replace(ctx, req.BlockId, req.Block)
		return err
	})
	return
}

func (s *service) SetFields(ctx *session.Context, req pb.RpcBlockSetFieldsRequest) (err error) {
	return s.DoBasic(req.ContextId, func(b basic.Basic) error {
		return b.SetFields(ctx, &pb.RpcBlockListSetFieldsRequestBlockField{
			BlockId: req.BlockId,
			Fields:  req.Fields,
		})
	})
}

func (s *service) SetDetails(ctx *session.Context, req pb.RpcObjectSetDetailsRequest) (err error) {
	return s.Do(req.ContextId, func(b smartblock.SmartBlock) error {
		return b.SetDetails(ctx, req.Details, true)
	})
}

func (s *service) SetFieldsList(ctx *session.Context, req pb.RpcBlockListSetFieldsRequest) (err error) {
	return s.DoBasic(req.ContextId, func(b basic.Basic) error {
		return b.SetFields(ctx, req.BlockFields...)
	})
}

func (s *service) GetAggregatedRelations(req pb.RpcBlockDataviewRelationListAvailableRequest) (relations []*model.Relation, err error) {
	err = s.DoDataview(req.ContextId, func(b dataview.Dataview) error {
		relations, err = b.GetAggregatedRelations(req.BlockId)
		return err
	})

	return
}

func (s *service) UpdateDataviewView(ctx *session.Context, req pb.RpcBlockDataviewViewUpdateRequest) error {
	return s.DoDataview(req.ContextId, func(b dataview.Dataview) error {
		return b.UpdateView(ctx, req.BlockId, req.ViewId, *req.View, true)
	})
}

func (s *service) UpdateDataviewGroupOrder(ctx *session.Context, req pb.RpcBlockDataviewGroupOrderUpdateRequest) error {
	return s.DoDataview(req.ContextId, func(b dataview.Dataview) error {
		return b.UpdateViewGroupOrder(ctx, req.BlockId, req.GroupOrder)
	})
}

func (s *service) UpdateDataviewObjectOrder(ctx *session.Context, req pb.RpcBlockDataviewObjectOrderUpdateRequest) error {
	return s.DoDataview(req.ContextId, func(b dataview.Dataview) error {
		return b.UpdateViewObjectOrder(ctx, req.BlockId, req.ObjectOrders)
	})
}

func (s *service) DeleteDataviewView(ctx *session.Context, req pb.RpcBlockDataviewViewDeleteRequest) error {
	return s.DoDataview(req.ContextId, func(b dataview.Dataview) error {
		return b.DeleteView(ctx, req.BlockId, req.ViewId, true)
	})
}

func (s *service) SetDataviewActiveView(ctx *session.Context, req pb.RpcBlockDataviewViewSetActiveRequest) error {
	return s.DoDataview(req.ContextId, func(b dataview.Dataview) error {
		return b.SetActiveView(ctx, req.BlockId, req.ViewId, int(req.Limit), int(req.Offset))
	})
}

func (s *service) SetDataviewViewPosition(ctx *session.Context, req pb.RpcBlockDataviewViewSetPositionRequest) error {
	return s.DoDataview(req.ContextId, func(b dataview.Dataview) error {
		return b.SetViewPosition(ctx, req.BlockId, req.ViewId, req.Position)
	})
}

func (s *service) CreateDataviewView(ctx *session.Context, req pb.RpcBlockDataviewViewCreateRequest) (id string, err error) {
	err = s.DoDataview(req.ContextId, func(b dataview.Dataview) error {
		if req.View == nil {
			req.View = &model.BlockContentDataviewView{}
		}
		view, e := b.CreateView(ctx, req.BlockId, *req.View)
		if e != nil {
			return e
		}
		id = view.Id
		return nil
	})
	return
}

func (s *service) AddDataviewRelation(ctx *session.Context, req pb.RpcBlockDataviewRelationAddRequest) (err error) {
	err = s.DoDataview(req.ContextId, func(b dataview.Dataview) error {
		return b.AddRelation(ctx, req.BlockId, req.RelationId, true)
	})

	return
}

func (s *service) DeleteDataviewRelation(ctx *session.Context, req pb.RpcBlockDataviewRelationDeleteRequest) error {
	return s.DoDataview(req.ContextId, func(b dataview.Dataview) error {
		return b.DeleteRelation(ctx, req.BlockId, req.RelationId, true)
	})
}

func (s *service) SetDataviewSource(ctx *session.Context, contextId, blockId string, source []string) (err error) {
	return s.DoDataview(contextId, func(b dataview.Dataview) error {
		return b.SetSource(ctx, blockId, source)
	})
}

func (s *service) Copy(req pb.RpcBlockCopyRequest) (textSlot string, htmlSlot string, anySlot []*model.Block, err error) {
	err = s.DoClipboard(req.ContextId, func(cb clipboard.Clipboard) error {
		textSlot, htmlSlot, anySlot, err = cb.Copy(req)
		return err
	})

	return textSlot, htmlSlot, anySlot, err
}

func (s *service) Paste(ctx *session.Context, req pb.RpcBlockPasteRequest, groupId string) (blockIds []string, uploadArr []pb.RpcBlockUploadRequest, caretPosition int32, isSameBlockCaret bool, err error) {
	err = s.DoClipboard(req.ContextId, func(cb clipboard.Clipboard) error {
		blockIds, uploadArr, caretPosition, isSameBlockCaret, err = cb.Paste(ctx, &req, groupId)
		return err
	})

	return blockIds, uploadArr, caretPosition, isSameBlockCaret, err
}

func (s *service) Cut(ctx *session.Context, req pb.RpcBlockCutRequest) (textSlot string, htmlSlot string, anySlot []*model.Block, err error) {
	err = s.DoClipboard(req.ContextId, func(cb clipboard.Clipboard) error {
		textSlot, htmlSlot, anySlot, err = cb.Cut(ctx, req)
		return err
	})
	return textSlot, htmlSlot, anySlot, err
}

func (s *service) Export(req pb.RpcBlockExportRequest) (path string, err error) {
	err = s.DoClipboard(req.ContextId, func(cb clipboard.Clipboard) error {
		path, err = cb.Export(req)
		return err
	})
	return path, err
}

func (s *service) ImportMarkdown(ctx *session.Context, req pb.RpcObjectImportMarkdownRequest) (rootLinkIds []string, err error) {
	var rootLinks []*model.Block
	err = s.DoImport(req.ContextId, func(imp _import.Import) error {
		rootLinks, err = imp.ImportMarkdown(ctx, req)
		return err
	})
	if err != nil {
		return rootLinkIds, err
	}

	if len(rootLinks) == 1 {
		err = s.SimplePaste(req.ContextId, rootLinks)

		if err != nil {
			return rootLinkIds, err
		}
	} else {
		_, pageId, err := s.CreateLinkToTheNewObject(ctx, "", pb.RpcBlockLinkCreateWithObjectRequest{
			ContextId: req.ContextId,
			Details: &types.Struct{Fields: map[string]*types.Value{
				"name":      pbtypes.String("Import from Notion"),
				"iconEmoji": pbtypes.String("📁"),
			}},
		})

		if err != nil {
			return rootLinkIds, err
		}

		err = s.SimplePaste(pageId, rootLinks)
	}

	for _, r := range rootLinks {
		rootLinkIds = append(rootLinkIds, r.Id)
	}

	return rootLinkIds, err
}

func (s *service) SetTextText(ctx *session.Context, req pb.RpcBlockTextSetTextRequest) error {
	return s.DoText(req.ContextId, func(b stext.Text) error {
		return b.SetText(ctx, req)
	})
}

func (s *service) SetLatexText(ctx *session.Context, req pb.RpcBlockLatexSetTextRequest) error {
	return s.Do(req.ContextId, func(b smartblock.SmartBlock) error {
		return b.(basic.Basic).SetLatexText(ctx, req)
	})
}

func (s *service) SetTextStyle(ctx *session.Context, contextId string, style model.BlockContentTextStyle, blockIds ...string) error {
	return s.DoText(contextId, func(b stext.Text) error {
		return b.UpdateTextBlocks(ctx, blockIds, true, func(t text.Block) error {
			t.SetStyle(style)
			return nil
		})
	})
}

func (s *service) SetTextChecked(ctx *session.Context, req pb.RpcBlockTextSetCheckedRequest) error {
	return s.DoText(req.ContextId, func(b stext.Text) error {
		return b.UpdateTextBlocks(ctx, []string{req.BlockId}, true, func(t text.Block) error {
			t.SetChecked(req.Checked)
			return nil
		})
	})
}

func (s *service) SetTextColor(ctx *session.Context, contextId string, color string, blockIds ...string) error {
	return s.DoText(contextId, func(b stext.Text) error {
		return b.UpdateTextBlocks(ctx, blockIds, true, func(t text.Block) error {
			t.SetTextColor(color)
			return nil
		})
	})
}

func (s *service) ClearTextStyle(ctx *session.Context, contextId string, blockIds ...string) error {
	return s.DoText(contextId, func(b stext.Text) error {
		return b.UpdateTextBlocks(ctx, blockIds, true, func(t text.Block) error {
			t.Model().BackgroundColor = ""
			t.Model().Align = model.Block_AlignLeft
			t.Model().VerticalAlign = model.Block_VerticalAlignTop
			t.SetTextColor("")
			t.SetStyle(model.BlockContentText_Paragraph)

			marks := t.Model().GetText().Marks.Marks[:0]
			for _, m := range t.Model().GetText().Marks.Marks {
				switch m.Type {
				case model.BlockContentTextMark_Strikethrough,
					model.BlockContentTextMark_Keyboard,
					model.BlockContentTextMark_Italic,
					model.BlockContentTextMark_Bold,
					model.BlockContentTextMark_Underscored,
					model.BlockContentTextMark_TextColor,
					model.BlockContentTextMark_BackgroundColor:
				default:
					marks = append(marks, m)
				}
			}
			t.Model().GetText().Marks.Marks = marks

			return nil
		})
	})
}

func (s *service) ClearTextContent(ctx *session.Context, contextId string, blockIds ...string) error {
	return s.DoText(contextId, func(b stext.Text) error {
		return b.UpdateTextBlocks(ctx, blockIds, true, func(t text.Block) error {
			return t.SetText("", nil)
		})
	})
}

func (s *service) SetTextMark(ctx *session.Context, contextId string, mark *model.BlockContentTextMark, blockIds ...string) error {
	return s.DoText(contextId, func(b stext.Text) error {
		return b.SetMark(ctx, mark, blockIds...)
	})
}

func (s *service) SetTextIcon(ctx *session.Context, contextId, image, emoji string, blockIds ...string) error {
	return s.DoText(contextId, func(b stext.Text) error {
		return b.SetIcon(ctx, image, emoji, blockIds...)
	})
}

func (s *service) SetBackgroundColor(ctx *session.Context, contextId string, color string, blockIds ...string) (err error) {
	return s.DoBasic(contextId, func(b basic.Basic) error {
		return b.Update(ctx, func(b simple.Block) error {
			b.Model().BackgroundColor = color
			return nil
		}, blockIds...)
	})
}

func (s *service) SetLinkAppearance(ctx *session.Context, req pb.RpcBlockLinkListSetAppearanceRequest) (err error) {
	return s.DoBasic(req.ContextId, func(b basic.Basic) error {
		return b.Update(ctx, func(b simple.Block) error {
			if linkBlock, ok := b.(link.Block); ok {
				return linkBlock.SetAppearance(&model.BlockContentLink{
					IconSize:    req.IconSize,
					CardStyle:   req.CardStyle,
					Description: req.Description,
					Relations:   req.Relations,
				})
			}
			return nil
		}, req.BlockIds...)
	})
}

func (s *service) SetAlign(ctx *session.Context, contextId string, align model.BlockAlign, blockIds ...string) (err error) {
	return s.Do(contextId, func(sb smartblock.SmartBlock) error {
		return sb.SetAlign(ctx, align, blockIds...)
	})
}

func (s *service) SetVerticalAlign(ctx *session.Context, contextId string, align model.BlockVerticalAlign, blockIds ...string) (err error) {
	return s.Do(contextId, func(sb smartblock.SmartBlock) error {
		return sb.SetVerticalAlign(ctx, align, blockIds...)
	})
}

func (s *service) SetLayout(ctx *session.Context, contextId string, layout model.ObjectTypeLayout) (err error) {
	return s.Do(contextId, func(sb smartblock.SmartBlock) error {
		return sb.SetLayout(ctx, layout)
	})
}

func (s *service) FeaturedRelationAdd(ctx *session.Context, contextId string, relations ...string) error {
	return s.DoBasic(contextId, func(b basic.Basic) error {
		return b.FeaturedRelationAdd(ctx, relations...)
	})
}

func (s *service) FeaturedRelationRemove(ctx *session.Context, contextId string, relations ...string) error {
	return s.DoBasic(contextId, func(b basic.Basic) error {
		return b.FeaturedRelationRemove(ctx, relations...)
	})
}

func (s *service) UploadBlockFile(ctx *session.Context, req pb.RpcBlockUploadRequest, groupId string) (err error) {
	return s.DoFile(req.ContextId, func(b file.File) error {
		err = b.Upload(ctx, req.BlockId, file.FileSource{
			Path:    req.FilePath,
			Url:     req.Url,
			GroupId: groupId,
		}, false)
		return err
	})
}

func (s *service) UploadBlockFileSync(ctx *session.Context, req pb.RpcBlockUploadRequest) (err error) {
	return s.DoFile(req.ContextId, func(b file.File) error {
		err = b.Upload(ctx, req.BlockId, file.FileSource{
			Path: req.FilePath,
			Url:  req.Url,
		}, true)
		return err
	})
}

func (s *service) CreateAndUploadFile(ctx *session.Context, req pb.RpcBlockFileCreateAndUploadRequest) (id string, err error) {
	err = s.DoFile(req.ContextId, func(b file.File) error {
		id, err = b.CreateAndUpload(ctx, req)
		return err
	})
	return
}

func (s *service) UploadFile(req pb.RpcFileUploadRequest) (hash string, err error) {
	upl := file.NewUploader(s)
	if req.DisableEncryption {
		log.Errorf("DisableEncryption is deprecated and has no effect")
	}

	upl.SetStyle(req.Style)
	if req.Type != model.BlockContentFile_None {
		upl.SetType(req.Type)
	} else {
		upl.AutoType(true)
	}
	res := upl.SetFile(req.LocalPath).Upload(context.TODO())
	if res.Err != nil {
		return "", res.Err
	}
	return res.Hash, nil
}

func (s *service) DropFiles(req pb.RpcFileDropRequest) (err error) {
	return s.DoFileNonLock(req.ContextId, func(b file.File) error {
		return b.DropFiles(req)
	})
}

func (s *service) SetFileStyle(ctx *session.Context, contextId string, style model.BlockContentFileStyle, blockIds ...string) error {
	return s.DoFile(contextId, func(b file.File) error {
		return b.SetFileStyle(ctx, style, blockIds...)
	})
}

func (s *service) Undo(ctx *session.Context, req pb.RpcObjectUndoRequest) (counters pb.RpcObjectUndoRedoCounter, err error) {
	err = s.DoHistory(req.ContextId, func(b basic.IHistory) error {
		counters, err = b.Undo(ctx)
		return err
	})
	return
}

func (s *service) Redo(ctx *session.Context, req pb.RpcObjectRedoRequest) (counters pb.RpcObjectUndoRedoCounter, err error) {
	err = s.DoHistory(req.ContextId, func(b basic.IHistory) error {
		counters, err = b.Redo(ctx)
		return err
	})
	return
}

func (s *service) BookmarkFetch(ctx *session.Context, req pb.RpcBlockBookmarkFetchRequest) (err error) {
	return s.DoBookmark(req.ContextId, func(b bookmark.Bookmark) error {
		return b.Fetch(ctx, req.BlockId, req.Url, false)
	})
}

func (s *service) BookmarkFetchSync(ctx *session.Context, req pb.RpcBlockBookmarkFetchRequest) (err error) {
	return s.DoBookmark(req.ContextId, func(b bookmark.Bookmark) error {
		return b.Fetch(ctx, req.BlockId, req.Url, true)
	})
}

func (s *service) BookmarkCreateAndFetch(ctx *session.Context, req pb.RpcBlockBookmarkCreateAndFetchRequest) (id string, err error) {
	err = s.DoBookmark(req.ContextId, func(b bookmark.Bookmark) error {
		id, err = b.CreateAndFetch(ctx, req)
		return err
	})
	return
}

func (s *service) SetRelationKey(ctx *session.Context, req pb.RpcBlockRelationSetKeyRequest) error {
	return s.DoBasic(req.ContextId, func(b basic.Basic) error {
		rel, err := s.relationService.FetchKey(req.Key)
		if err != nil {
			return err
		}
		return b.AddRelationAndSet(ctx, pb.RpcBlockRelationAddRequest{RelationId: rel.Id, BlockId: req.BlockId, ContextId: req.ContextId})
	})
}

func (s *service) AddRelationBlock(ctx *session.Context, req pb.RpcBlockRelationAddRequest) error {
	return s.DoBasic(req.ContextId, func(b basic.Basic) error {
		return b.AddRelationAndSet(ctx, req)
	})
}

func (s *service) GetDocInfo(ctx context.Context, id string) (info doc.DocInfo, err error) {
	if err = s.DoWithContext(ctx, id, func(b smartblock.SmartBlock) error {
		info, err = b.GetDocInfo()
		return err
	}); err != nil {
		return
	}
	return
}

func (s *service) Wakeup(id string) (err error) {
	return s.Do(id, func(b smartblock.SmartBlock) error {
		return nil
	})
}

func (s *service) GetRelations(objectId string) (relations []*model.Relation, err error) {
	err = s.Do(objectId, func(b smartblock.SmartBlock) error {
		relations = b.Relations(nil).Models()
		return nil
	})
	return
}

// ModifyDetails performs details get and update under the sb lock to make sure no modifications are done in the middle
func (s *service) ModifyDetails(objectId string, modifier func(current *types.Struct) (*types.Struct, error)) (err error) {
	if modifier == nil {
		return fmt.Errorf("modifier is nil")
	}
	return s.Do(objectId, func(b smartblock.SmartBlock) error {
		dets, err := modifier(b.CombinedDetails())
		if err != nil {
			return err
		}

		return b.Apply(b.NewState().SetDetails(dets))
	})
}

// ModifyLocalDetails modifies local details of the object in cache, and if it is not found, sets pending details in object store
func (s *service) ModifyLocalDetails(objectId string, modifier func(current *types.Struct) (*types.Struct, error)) (err error) {
	if modifier == nil {
		return fmt.Errorf("modifier is nil")
	}
	// we set pending details if object is not in cache
	// we do this under lock to prevent races if the object is created in parallel
	// because in that case we can lose changes
	err = s.cache.DoLockedIfNotExists(objectId, func() error {
		objectDetails, err := s.objectStore.GetPendingLocalDetails(objectId)
		if err != nil && err != ds.ErrNotFound {
			return err
		}
		var details *types.Struct
		if objectDetails != nil {
			details = objectDetails.GetDetails()
		}
		modifiedDetails, err := modifier(details)
		if err != nil {
			return err
		}
		return s.objectStore.UpdatePendingLocalDetails(objectId, modifiedDetails)
	})
	if err != nil && err != ocache.ErrExists {
		return err
	}
	err = s.Do(objectId, func(b smartblock.SmartBlock) error {
		// we just need to invoke the smartblock so it reads from pending details
		// no need to call modify twice
		if err == nil {
			return nil
		}

		dets, err := modifier(b.CombinedDetails())
		if err != nil {
			return err
		}

		return b.Apply(b.NewState().SetDetails(dets))
	})
	// that means that we will apply the change later as soon as the block is loaded by thread queue
	if err == source.ErrObjectNotFound {
		return nil
	}
	return err
}

func (s *service) AddExtraRelations(ctx *session.Context, objectId string, relationIds []string) (err error) {
	return s.Do(objectId, func(b smartblock.SmartBlock) error {
		return b.AddExtraRelations(ctx, relationIds...)
	})
}

func (s *service) SetObjectTypes(ctx *session.Context, objectId string, objectTypes []string) (err error) {
	return s.Do(objectId, func(b smartblock.SmartBlock) error {
		return b.SetObjectTypes(ctx, objectTypes)
	})
}

// todo: rewrite with options
// withId may me empty
func (s *service) CreateObjectInWorkspace(ctx context.Context, workspaceId string, withId thread.ID, sbType coresb.SmartBlockType) (csm core.SmartBlock, err error) {
	startTime := time.Now()
	ev, exists := ctx.Value(ObjectCreateEvent).(*metrics.CreateObjectEvent)
	err = s.DoWithContext(ctx, workspaceId, func(b smartblock.SmartBlock) error {
		if exists {
			ev.GetWorkspaceBlockWaitMs = time.Now().Sub(startTime).Milliseconds()
		}
		workspace, ok := b.(*editor.Workspaces)
		if !ok {
			return fmt.Errorf("incorrect object with workspace id")
		}
		csm, err = workspace.CreateObject(withId, sbType)
		if exists {
			ev.WorkspaceCreateMs = time.Now().Sub(startTime).Milliseconds() - ev.GetWorkspaceBlockWaitMs
		}
		if err != nil {
			return fmt.Errorf("anytype.CreateBlock error: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return csm, nil
}

func (s *service) DeleteObjectFromWorkspace(workspaceId string, objectId string) error {
	return s.Do(workspaceId, func(b smartblock.SmartBlock) error {
		workspace, ok := b.(*editor.Workspaces)
		if !ok {
			return fmt.Errorf("incorrect object with workspace id")
		}

		if parts := strings.Split(objectId, util.SubIdSeparator); len(parts) > 1 {
			return workspace.DeleteSubObject(objectId)
		}

		return workspace.DeleteObject(objectId)
	})
}

func (s *service) CreateSet(req pb.RpcObjectCreateSetRequest) (setId string, err error) {
	req.Details = internalflag.AddToDetails(req.Details, req.InternalFlags)

	var dvContent model.BlockContentOfDataview
	var dvSchema schema.Schema
	if len(req.Source) != 0 {
		if dvContent, dvSchema, err = dataview.DataviewBlockBySource(s.anytype.ObjectStore(), req.Source); err != nil {
			return
		}
	}
	var workspaceId string
	if req.GetDetails().GetFields() != nil {
		detailsWorkspaceId := req.GetDetails().Fields[bundle.RelationKeyWorkspaceId.String()]
		if detailsWorkspaceId != nil && detailsWorkspaceId.GetStringValue() != "" {
			workspaceId = detailsWorkspaceId.GetStringValue()
		}
	}

	// if we don't have anything in details then check the object store
	if workspaceId == "" {
		workspaceId = s.anytype.PredefinedBlocks().Account
	}

	// TODO: here can be a deadlock if this is somehow created from workspace (as set)
	csm, err := s.CreateObjectInWorkspace(context.TODO(), workspaceId, thread.Undef, coresb.SmartBlockTypeSet)
	if err != nil {
		return "", err
	}

	setId = csm.ID()

	state := state.NewDoc(csm.ID(), nil).NewState()
	if workspaceId != "" {
		state.SetDetailAndBundledRelation(bundle.RelationKeyWorkspaceId, pbtypes.String(workspaceId))
	}

	name := pbtypes.GetString(req.Details, bundle.RelationKeyName.String())
	icon := pbtypes.GetString(req.Details, bundle.RelationKeyIconEmoji.String())

	tmpls := []template.StateTransformer{
		template.WithForcedDetail(bundle.RelationKeyName, pbtypes.String(name)),
		template.WithForcedDetail(bundle.RelationKeyIconEmoji, pbtypes.String(icon)),
		template.WithRequiredRelations(),
	}
	var blockContent *model.BlockContentOfDataview
	if dvSchema != nil {
		blockContent = &dvContent
	}
	if blockContent != nil {
		for i, view := range blockContent.Dataview.Views {
			if view.Relations == nil {
				blockContent.Dataview.Views[i].Relations = editor.GetDefaultViewRelations(blockContent.Dataview.Relations)
			}
		}
		tmpls = append(tmpls,
			template.WithForcedDetail(bundle.RelationKeySetOf, pbtypes.StringList(blockContent.Dataview.Source)),
			template.WithDataview(*blockContent, false),
		)
	}

	if err = template.InitTemplate(state, tmpls...); err != nil {
		return "", err
	}

	sb, err := s.newSmartBlock(setId, &smartblock.InitContext{
		State: state,
	})
	if err != nil {
		return "", err
	}
	_, ok := sb.(*editor.Set)
	if !ok {
		return setId, fmt.Errorf("unexpected set block type: %T", sb)
	}
	return setId, err
}

func (s *service) ObjectToSet(id string, source []string) (newId string, err error) {
	if s.app == nil {
		err = errors.New("app can't be nil")
		return
	}
	var details *types.Struct
	if err = s.Do(id, func(b smartblock.SmartBlock) error {
		details = pbtypes.CopyStruct(b.Details())

		s := b.NewState()
		if layout, ok := s.Layout(); ok && layout == model.ObjectType_note {
			textBlock, err := s.GetFirstTextBlock()
			if err != nil {
				return err
			}
			if textBlock != nil {
				details.Fields[bundle.RelationKeyName.String()] = pbtypes.String(textBlock.Text.Text)
			}
		}

		return nil
	}); err != nil {
		return
	}

	details.Fields[bundle.RelationKeySetOf.String()] = pbtypes.StringList(source)
	newId, err = s.CreateSet(pb.RpcObjectCreateSetRequest{
		Source:  source,
		Details: details,
	})
	if err != nil {
		return
	}

	oStore := s.app.MustComponent(objectstore.CName).(objectstore.ObjectStore)
	res, err := oStore.GetWithLinksInfoByID(id)
	if err != nil {
		return
	}
	for _, il := range res.Links.Inbound {
		if err = s.replaceLink(il.Id, id, newId); err != nil {
			return
		}
	}
	err = s.DeleteObject(id)
	if err != nil {
		// intentionally do not return error here
		log.Errorf("failed to delete object after conversion to set: %s", err.Error())
	}

	return
}

func (s *service) RemoveExtraRelations(ctx *session.Context, objectTypeId string, relationKeys []string) (err error) {
	return s.Do(objectTypeId, func(b smartblock.SmartBlock) error {
		return b.RemoveExtraRelations(ctx, relationKeys)
	})
}

func (s *service) ListAvailableRelations(objectId string) (aggregatedRelations []*model.Relation, err error) {
	err = s.Do(objectId, func(b smartblock.SmartBlock) error {
		// TODO: not implemented
		return nil
	})
	return
}

func (s *service) ListConvertToObjects(ctx *session.Context, req pb.RpcBlockListConvertToObjectsRequest) (linkIds []string, err error) {
	err = s.DoBasic(req.ContextId, func(b basic.Basic) error {
		linkIds, err = b.ExtractBlocksToObjects(ctx, s, req)
		return err
	})
	return
}

func (s *service) MoveBlocksToNewPage(ctx *session.Context, req pb.RpcBlockListMoveToNewObjectRequest) (linkId string, err error) {
	// 1. Create new page, link
	linkId, pageId, err := s.CreateLinkToTheNewObject(ctx, "", pb.RpcBlockLinkCreateWithObjectRequest{
		ContextId: req.ContextId,
		TargetId:  req.DropTargetId,
		Position:  req.Position,
		Details:   req.Details,
	})

	if err != nil {
		return linkId, err
	}

	// 2. Move blocks to new page
	err = s.MoveBlocks(nil, pb.RpcBlockListMoveToExistingObjectRequest{
		ContextId:       req.ContextId,
		BlockIds:        req.BlockIds,
		TargetContextId: pageId,
		DropTargetId:    "",
		Position:        0,
	})

	if err != nil {
		return linkId, err
	}

	return linkId, err
}

func (s *service) MoveBlocks(ctx *session.Context, req pb.RpcBlockListMoveToExistingObjectRequest) error {
	if req.ContextId == req.TargetContextId {
		return s.DoBasic(req.ContextId, func(b basic.Basic) error {
			return b.Move(ctx, req)
		})
	}
	return s.Do(req.ContextId, func(cb smartblock.SmartBlock) error {
		return s.DoClipboard(req.TargetContextId, func(tb clipboard.Clipboard) error {
			cs := cb.NewState()
			bs := basic.CutBlocks(cs, req.BlockIds)
			blocks := make([]*model.Block, 0, len(bs))
			for _, b := range bs {
				blocks = append(blocks, b.Model())
			}
			_, _, _, _, err := tb.Paste(ctx, &pb.RpcBlockPasteRequest{
				FocusedBlockId: req.DropTargetId,
				AnySlot:        blocks,
			}, "")
			if err != nil {
				return fmt.Errorf("paste: %w", err)
			}
			return cb.Apply(cs)
		})
	})
}

func (s *service) CreateTableBlock(ctx *session.Context, req pb.RpcBlockTableCreateRequest) (id string, err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		id, err = t.TableCreate(st, req)
		return err
	})
	return
}

func (s *service) TableRowCreate(ctx *session.Context, req pb.RpcBlockTableRowCreateRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.RowCreate(st, req)
	})
	return
}

func (s *service) TableColumnCreate(ctx *session.Context, req pb.RpcBlockTableColumnCreateRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.ColumnCreate(st, req)
	})
	return
}

func (s *service) TableRowDelete(ctx *session.Context, req pb.RpcBlockTableRowDeleteRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.RowDelete(st, req)
	})
	return
}

func (s *service) TableColumnDelete(ctx *session.Context, req pb.RpcBlockTableColumnDeleteRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.ColumnDelete(st, req)
	})
	return
}

func (s *service) TableColumnMove(ctx *session.Context, req pb.RpcBlockTableColumnMoveRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.ColumnMove(st, req)
	})
	return
}

func (s *service) TableRowDuplicate(ctx *session.Context, req pb.RpcBlockTableRowDuplicateRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.RowDuplicate(st, req)
	})
	return
}

func (s *service) TableColumnDuplicate(ctx *session.Context, req pb.RpcBlockTableColumnDuplicateRequest) (id string, err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		id, err = t.ColumnDuplicate(st, req)
		return err
	})
	return id, err
}

func (s *service) TableExpand(ctx *session.Context, req pb.RpcBlockTableExpandRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.Expand(st, req)
	})
	return err
}

func (s *service) TableRowListFill(ctx *session.Context, req pb.RpcBlockTableRowListFillRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.RowListFill(st, req)
	})
	return err
}

func (s *service) TableRowListClean(ctx *session.Context, req pb.RpcBlockTableRowListCleanRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.RowListClean(st, req)
	})
	return err
}

func (s *service) TableRowSetHeader(ctx *session.Context, req pb.RpcBlockTableRowSetHeaderRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.RowSetHeader(st, req)
	})
	return err
}

func (s *service) TableSort(ctx *session.Context, req pb.RpcBlockTableSortRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.Sort(st, req)
	})
	return err
}

func (s *service) TableColumnListFill(ctx *session.Context, req pb.RpcBlockTableColumnListFillRequest) (err error) {
	err = s.DoTable(req.ContextId, ctx, func(st *state.State, t table.Editor) error {
		return t.ColumnListFill(st, req)
	})
	return err
}
