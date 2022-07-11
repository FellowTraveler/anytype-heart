package core

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/anytypeio/go-anytype-middleware/core/event"
	"github.com/globalsign/mgo/bson"
	"github.com/golang-jwt/jwt/v4"
	"github.com/miolini/datacounter"
	"google.golang.org/grpc/metadata"

	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/process"
	"github.com/anytypeio/go-anytype-middleware/core/block/source"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/files"
)

func (mw *Middleware) BlockCreate(cctx context.Context, req *pb.RpcBlockCreateRequest) *pb.RpcBlockCreateResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockCreateResponseErrorCode, id string, err error) *pb.RpcBlockCreateResponse {
		m := &pb.RpcBlockCreateResponse{Error: &pb.RpcBlockCreateResponseError{Code: code}, BlockId: id}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var id string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		id, err = bs.CreateBlock(ctx, *req)
		return
	})
	if err != nil {
		return response(pb.RpcBlockCreateResponseError_UNKNOWN_ERROR, "", err)
	}
	return response(pb.RpcBlockCreateResponseError_NULL, id, nil)
}

func (mw *Middleware) BlockLinkCreateWithObject(cctx context.Context, req *pb.RpcBlockLinkCreateWithObjectRequest) *pb.RpcBlockLinkCreateWithObjectResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockLinkCreateWithObjectResponseErrorCode, id, targetId string, err error) *pb.RpcBlockLinkCreateWithObjectResponse {
		m := &pb.RpcBlockLinkCreateWithObjectResponse{Error: &pb.RpcBlockLinkCreateWithObjectResponseError{Code: code}, BlockId: id, TargetId: targetId}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var id, targetId string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		id, targetId, err = bs.CreateLinkToTheNewObject(ctx, "", *req)
		return
	})
	if err != nil {
		return response(pb.RpcBlockLinkCreateWithObjectResponseError_UNKNOWN_ERROR, "", "", err)
	}
	return response(pb.RpcBlockLinkCreateWithObjectResponseError_NULL, id, targetId, nil)
}

func (mw *Middleware) ObjectOpen(cctx context.Context, req *pb.RpcObjectOpenRequest) *pb.RpcObjectOpenResponse {
	ctx := mw.newContext(cctx)
	var obj *model.ObjectView
	response := func(code pb.RpcObjectOpenResponseErrorCode, err error) *pb.RpcObjectOpenResponse {
		m := &pb.RpcObjectOpenResponse{Error: &pb.RpcObjectOpenResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.ObjectView = obj
		}
		return m
	}

	err := mw.doBlockService(func(bs block.Service) (err error) {
		obj, err = bs.OpenBlock(ctx, req.ObjectId)
		return err
	})
	if err != nil {
		if err == source.ErrUnknownDataFormat {
			return response(pb.RpcObjectOpenResponseError_ANYTYPE_NEEDS_UPGRADE, err)
		} else if err == source.ErrObjectNotFound {
			return response(pb.RpcObjectOpenResponseError_NOT_FOUND, err)
		}
		return response(pb.RpcObjectOpenResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcObjectOpenResponseError_NULL, nil)
}

func (mw *Middleware) ObjectShow(cctx context.Context, req *pb.RpcObjectShowRequest) *pb.RpcObjectShowResponse {
	ctx := mw.newContext(cctx, state.WithTraceId(req.TraceId))
	var obj *model.ObjectView
	response := func(code pb.RpcObjectShowResponseErrorCode, err error) *pb.RpcObjectShowResponse {
		m := &pb.RpcObjectShowResponse{Error: &pb.RpcObjectShowResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.ObjectView = obj
		}
		return m
	}

	err := mw.doBlockService(func(bs block.Service) (err error) {
		obj, err = bs.ShowBlock(ctx, req.ObjectId)
		return err
	})
	if err != nil {
		if err == source.ErrUnknownDataFormat {
			return response(pb.RpcObjectShowResponseError_ANYTYPE_NEEDS_UPGRADE, err)
		} else if err == source.ErrObjectNotFound {
			return response(pb.RpcObjectShowResponseError_NOT_FOUND, err)
		}
		return response(pb.RpcObjectShowResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcObjectShowResponseError_NULL, nil)
}

func (mw *Middleware) ObjectOpenBreadcrumbs(cctx context.Context, req *pb.RpcObjectOpenBreadcrumbsRequest) *pb.RpcObjectOpenBreadcrumbsResponse {
	ctx := mw.newContext(cctx, state.WithTraceId(req.TraceId))
	response := func(code pb.RpcObjectOpenBreadcrumbsResponseErrorCode, obj *model.ObjectView, id string, err error) *pb.RpcObjectOpenBreadcrumbsResponse {
		m := &pb.RpcObjectOpenBreadcrumbsResponse{Error: &pb.RpcObjectOpenBreadcrumbsResponseError{Code: code}, ObjectId: id}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.ObjectView = obj
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var id string
	var obj *model.ObjectView
	err := mw.doBlockService(func(bs block.Service) (err error) {
		obj, id, err = bs.OpenBreadcrumbsBlock(ctx)
		return
	})
	if err != nil {
		return response(pb.RpcObjectOpenBreadcrumbsResponseError_UNKNOWN_ERROR, nil, "", err)
	}

	return response(pb.RpcObjectOpenBreadcrumbsResponseError_NULL, obj, id, nil)
}

func (mw *Middleware) ObjectSetBreadcrumbs(cctx context.Context, req *pb.RpcObjectSetBreadcrumbsRequest) *pb.RpcObjectSetBreadcrumbsResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectSetBreadcrumbsResponseErrorCode, err error) *pb.RpcObjectSetBreadcrumbsResponse {
		m := &pb.RpcObjectSetBreadcrumbsResponse{Error: &pb.RpcObjectSetBreadcrumbsResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetBreadcrumbs(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcObjectSetBreadcrumbsResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcObjectSetBreadcrumbsResponseError_NULL, nil)
}

func (mw *Middleware) ObjectClose(cctx context.Context, req *pb.RpcObjectCloseRequest) *pb.RpcObjectCloseResponse {
	response := func(code pb.RpcObjectCloseResponseErrorCode, err error) *pb.RpcObjectCloseResponse {
		m := &pb.RpcObjectCloseResponse{Error: &pb.RpcObjectCloseResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.CloseBlock(req.ObjectId)
	})
	if err != nil {
		return response(pb.RpcObjectCloseResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectCloseResponseError_NULL, nil)
}
func (mw *Middleware) BlockCopy(cctx context.Context, req *pb.RpcBlockCopyRequest) *pb.RpcBlockCopyResponse {
	response := func(code pb.RpcBlockCopyResponseErrorCode, textSlot string, htmlSlot string, anySlot []*model.Block, err error) *pb.RpcBlockCopyResponse {
		m := &pb.RpcBlockCopyResponse{
			Error:    &pb.RpcBlockCopyResponseError{Code: code},
			TextSlot: textSlot,
			HtmlSlot: htmlSlot,
			AnySlot:  anySlot,
		}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	var textSlot, htmlSlot string
	var anySlot []*model.Block
	err := mw.doBlockService(func(bs block.Service) (err error) {
		textSlot, htmlSlot, anySlot, err = bs.Copy(*req)
		return
	})
	if err != nil {
		return response(pb.RpcBlockCopyResponseError_UNKNOWN_ERROR, textSlot, htmlSlot, anySlot, err)
	}

	return response(pb.RpcBlockCopyResponseError_NULL, textSlot, htmlSlot, anySlot, nil)
}

func (mw *Middleware) BlockPaste(cctx context.Context, req *pb.RpcBlockPasteRequest) *pb.RpcBlockPasteResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockPasteResponseErrorCode, blockIds []string, caretPosition int32, isSameBlockCaret bool, err error) *pb.RpcBlockPasteResponse {
		m := &pb.RpcBlockPasteResponse{Error: &pb.RpcBlockPasteResponseError{Code: code}, BlockIds: blockIds, CaretPosition: caretPosition, IsSameBlockCaret: isSameBlockCaret}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var (
		blockIds         []string
		caretPosition    int32
		isSameBlockCaret bool
		groupId          = bson.NewObjectId().Hex()
	)
	err := mw.doBlockService(func(bs block.Service) (err error) {
		var uploadArr []pb.RpcBlockUploadRequest
		blockIds, uploadArr, caretPosition, isSameBlockCaret, err = bs.Paste(ctx, *req, groupId)
		if err != nil {
			return
		}
		log.Debug("Image requests to upload after paste:", uploadArr)
		for _, r := range uploadArr {
			r.ContextId = req.ContextId
			if err = bs.UploadBlockFile(nil, r, groupId); err != nil {
				return err
			}
		}
		return
	})
	if err != nil {
		return response(pb.RpcBlockPasteResponseError_UNKNOWN_ERROR, nil, -1, isSameBlockCaret, err)
	}

	return response(pb.RpcBlockPasteResponseError_NULL, blockIds, caretPosition, isSameBlockCaret, nil)
}

func (mw *Middleware) BlockCut(cctx context.Context, req *pb.RpcBlockCutRequest) *pb.RpcBlockCutResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockCutResponseErrorCode, textSlot string, htmlSlot string, anySlot []*model.Block, err error) *pb.RpcBlockCutResponse {
		m := &pb.RpcBlockCutResponse{
			Error:    &pb.RpcBlockCutResponseError{Code: code},
			TextSlot: textSlot,
			HtmlSlot: htmlSlot,
			AnySlot:  anySlot,
		}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var (
		textSlot, htmlSlot string
		anySlot            []*model.Block
	)
	err := mw.doBlockService(func(bs block.Service) (err error) {
		textSlot, htmlSlot, anySlot, err = bs.Cut(ctx, *req)
		return
	})
	if err != nil {
		var emptyAnySlot []*model.Block
		return response(pb.RpcBlockCutResponseError_UNKNOWN_ERROR, "", "", emptyAnySlot, err)
	}

	return response(pb.RpcBlockCutResponseError_NULL, textSlot, htmlSlot, anySlot, nil)
}

func (mw *Middleware) ObjectImportMarkdown(cctx context.Context, req *pb.RpcObjectImportMarkdownRequest) *pb.RpcObjectImportMarkdownResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectImportMarkdownResponseErrorCode, rootLinkIds []string, err error) *pb.RpcObjectImportMarkdownResponse {
		m := &pb.RpcObjectImportMarkdownResponse{
			Error:       &pb.RpcObjectImportMarkdownResponseError{Code: code},
			RootLinkIds: rootLinkIds,
		}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}

	var rootLinkIds []string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		rootLinkIds, err = bs.ImportMarkdown(ctx, *req)
		return err
	})

	if err != nil {
		return response(pb.RpcObjectImportMarkdownResponseError_UNKNOWN_ERROR, rootLinkIds, err)
	}

	return response(pb.RpcObjectImportMarkdownResponseError_NULL, rootLinkIds, nil)
}

func (mw *Middleware) BlockExport(cctx context.Context, req *pb.RpcBlockExportRequest) *pb.RpcBlockExportResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockExportResponseErrorCode, path string, err error) *pb.RpcBlockExportResponse {
		m := &pb.RpcBlockExportResponse{
			Error: &pb.RpcBlockExportResponseError{Code: code},
			Path:  path,
		}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var path string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		path, err = bs.Export(*req)
		return
	})
	if err != nil {
		return response(pb.RpcBlockExportResponseError_UNKNOWN_ERROR, path, err)
	}

	return response(pb.RpcBlockExportResponseError_NULL, path, nil)
}

func (mw *Middleware) BlockUpload(cctx context.Context, req *pb.RpcBlockUploadRequest) *pb.RpcBlockUploadResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockUploadResponseErrorCode, err error) *pb.RpcBlockUploadResponse {
		m := &pb.RpcBlockUploadResponse{Error: &pb.RpcBlockUploadResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.UploadBlockFile(nil, *req, "")
	})
	if err != nil {
		return response(pb.RpcBlockUploadResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockUploadResponseError_NULL, nil)
}

func (mw *Middleware) BlockListDelete(cctx context.Context, req *pb.RpcBlockListDeleteRequest) *pb.RpcBlockListDeleteResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockListDeleteResponseErrorCode, err error) *pb.RpcBlockListDeleteResponse {
		m := &pb.RpcBlockListDeleteResponse{Error: &pb.RpcBlockListDeleteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.UnlinkBlock(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockListDeleteResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockListDeleteResponseError_NULL, nil)
}

func (mw *Middleware) BlockListDuplicate(cctx context.Context, req *pb.RpcBlockListDuplicateRequest) *pb.RpcBlockListDuplicateResponse {
	ctx := mw.newContext(cctx)
	response := func(ids []string, code pb.RpcBlockListDuplicateResponseErrorCode, err error) *pb.RpcBlockListDuplicateResponse {
		m := &pb.RpcBlockListDuplicateResponse{BlockIds: ids, Error: &pb.RpcBlockListDuplicateResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var ids []string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		ids, err = bs.DuplicateBlocks(ctx, *req)
		return
	})
	if err != nil {
		return response(nil, pb.RpcBlockListDuplicateResponseError_UNKNOWN_ERROR, err)
	}
	return response(ids, pb.RpcBlockListDuplicateResponseError_NULL, nil)
}

func (mw *Middleware) BlockSetFields(cctx context.Context, req *pb.RpcBlockSetFieldsRequest) *pb.RpcBlockSetFieldsResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockSetFieldsResponseErrorCode, err error) *pb.RpcBlockSetFieldsResponse {
		m := &pb.RpcBlockSetFieldsResponse{Error: &pb.RpcBlockSetFieldsResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetFields(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockSetFieldsResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockSetFieldsResponseError_NULL, nil)
}

func (mw *Middleware) BlockListSetFields(cctx context.Context, req *pb.RpcBlockListSetFieldsRequest) *pb.RpcBlockListSetFieldsResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockListSetFieldsResponseErrorCode, err error) *pb.RpcBlockListSetFieldsResponse {
		m := &pb.RpcBlockListSetFieldsResponse{Error: &pb.RpcBlockListSetFieldsResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetFieldsList(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockListSetFieldsResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockListSetFieldsResponseError_NULL, nil)
}

func (mw *Middleware) ObjectListDelete(cctx context.Context, req *pb.RpcObjectListDeleteRequest) *pb.RpcObjectListDeleteResponse {
	response := func(code pb.RpcObjectListDeleteResponseErrorCode, err error) *pb.RpcObjectListDeleteResponse {
		m := &pb.RpcObjectListDeleteResponse{Error: &pb.RpcObjectListDeleteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.DeleteArchivedObjects(*req)
	})
	if err != nil {
		return response(pb.RpcObjectListDeleteResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectListDeleteResponseError_NULL, nil)
}

func (mw *Middleware) ObjectListSetIsArchived(cctx context.Context, req *pb.RpcObjectListSetIsArchivedRequest) *pb.RpcObjectListSetIsArchivedResponse {
	response := func(code pb.RpcObjectListSetIsArchivedResponseErrorCode, err error) *pb.RpcObjectListSetIsArchivedResponse {
		m := &pb.RpcObjectListSetIsArchivedResponse{Error: &pb.RpcObjectListSetIsArchivedResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetPagesIsArchived(*req)
	})
	if err != nil {
		return response(pb.RpcObjectListSetIsArchivedResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectListSetIsArchivedResponseError_NULL, nil)
}
func (mw *Middleware) ObjectListSetIsFavorite(cctx context.Context, req *pb.RpcObjectListSetIsFavoriteRequest) *pb.RpcObjectListSetIsFavoriteResponse {
	response := func(code pb.RpcObjectListSetIsFavoriteResponseErrorCode, err error) *pb.RpcObjectListSetIsFavoriteResponse {
		m := &pb.RpcObjectListSetIsFavoriteResponse{Error: &pb.RpcObjectListSetIsFavoriteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetPagesIsFavorite(*req)
	})
	if err != nil {
		return response(pb.RpcObjectListSetIsFavoriteResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectListSetIsFavoriteResponseError_NULL, nil)
}

func (mw *Middleware) BlockReplace(cctx context.Context, req *pb.RpcBlockReplaceRequest) *pb.RpcBlockReplaceResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockReplaceResponseErrorCode, blockId string, err error) *pb.RpcBlockReplaceResponse {
		m := &pb.RpcBlockReplaceResponse{Error: &pb.RpcBlockReplaceResponseError{Code: code}, BlockId: blockId}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var blockId string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		blockId, err = bs.ReplaceBlock(ctx, *req)
		return
	})
	if err != nil {
		return response(pb.RpcBlockReplaceResponseError_UNKNOWN_ERROR, "", err)
	}
	return response(pb.RpcBlockReplaceResponseError_NULL, blockId, nil)
}

func (mw *Middleware) BlockTextSetColor(cctx context.Context, req *pb.RpcBlockTextSetColorRequest) *pb.RpcBlockTextSetColorResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockTextSetColorResponseErrorCode, err error) *pb.RpcBlockTextSetColorResponse {
		m := &pb.RpcBlockTextSetColorResponse{Error: &pb.RpcBlockTextSetColorResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetTextColor(nil, req.ContextId, req.Color, req.BlockId)
	})
	if err != nil {
		return response(pb.RpcBlockTextSetColorResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextSetColorResponseError_NULL, nil)
}

func (mw *Middleware) BlockListSetBackgroundColor(cctx context.Context, req *pb.RpcBlockListSetBackgroundColorRequest) *pb.RpcBlockListSetBackgroundColorResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockListSetBackgroundColorResponseErrorCode, err error) *pb.RpcBlockListSetBackgroundColorResponse {
		m := &pb.RpcBlockListSetBackgroundColorResponse{Error: &pb.RpcBlockListSetBackgroundColorResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetBackgroundColor(ctx, req.ContextId, req.Color, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockListSetBackgroundColorResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockListSetBackgroundColorResponseError_NULL, nil)
}

func (mw *Middleware) BlockLinkListSetAppearance(cctx context.Context, req *pb.RpcBlockLinkListSetAppearanceRequest) *pb.RpcBlockLinkListSetAppearanceResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockLinkListSetAppearanceResponseErrorCode, err error) *pb.RpcBlockLinkListSetAppearanceResponse {
		m := &pb.RpcBlockLinkListSetAppearanceResponse{Error: &pb.RpcBlockLinkListSetAppearanceResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetLinkAppearance(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockLinkListSetAppearanceResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockLinkListSetAppearanceResponseError_NULL, nil)
}

func (mw *Middleware) BlockListSetAlign(cctx context.Context, req *pb.RpcBlockListSetAlignRequest) *pb.RpcBlockListSetAlignResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockListSetAlignResponseErrorCode, err error) *pb.RpcBlockListSetAlignResponse {
		m := &pb.RpcBlockListSetAlignResponse{Error: &pb.RpcBlockListSetAlignResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetAlign(ctx, req.ContextId, req.Align, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockListSetAlignResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockListSetAlignResponseError_NULL, nil)
}

func (mw *Middleware) BlockListSetVerticalAlign(cctx context.Context, req *pb.RpcBlockListSetVerticalAlignRequest) *pb.RpcBlockListSetVerticalAlignResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockListSetVerticalAlignResponseErrorCode, err error) *pb.RpcBlockListSetVerticalAlignResponse {
		m := &pb.RpcBlockListSetVerticalAlignResponse{Error: &pb.RpcBlockListSetVerticalAlignResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetVerticalAlign(ctx, req.ContextId, req.VerticalAlign, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockListSetVerticalAlignResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockListSetVerticalAlignResponseError_NULL, nil)
}

func (mw *Middleware) FileDrop(cctx context.Context, req *pb.RpcFileDropRequest) *pb.RpcFileDropResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcFileDropResponseErrorCode, err error) *pb.RpcFileDropResponse {
		m := &pb.RpcFileDropResponse{Error: &pb.RpcFileDropResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.DropFiles(*req)
	})
	if err != nil {
		return response(pb.RpcFileDropResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcFileDropResponseError_NULL, nil)
}

func (mw *Middleware) BlockListMoveToExistingObject(cctx context.Context, req *pb.RpcBlockListMoveToExistingObjectRequest) *pb.RpcBlockListMoveToExistingObjectResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockListMoveToExistingObjectResponseErrorCode, err error) *pb.RpcBlockListMoveToExistingObjectResponse {
		m := &pb.RpcBlockListMoveToExistingObjectResponse{Error: &pb.RpcBlockListMoveToExistingObjectResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.MoveBlocks(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockListMoveToExistingObjectResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockListMoveToExistingObjectResponseError_NULL, nil)
}

func (mw *Middleware) BlockListMoveToNewObject(cctx context.Context, req *pb.RpcBlockListMoveToNewObjectRequest) *pb.RpcBlockListMoveToNewObjectResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockListMoveToNewObjectResponseErrorCode, linkId string, err error) *pb.RpcBlockListMoveToNewObjectResponse {
		m := &pb.RpcBlockListMoveToNewObjectResponse{Error: &pb.RpcBlockListMoveToNewObjectResponseError{Code: code}, LinkId: linkId}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}

	var linkId string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		linkId, err = bs.MoveBlocksToNewPage(ctx, *req)
		return
	})

	if err != nil {
		return response(pb.RpcBlockListMoveToNewObjectResponseError_UNKNOWN_ERROR, "", err)
	}
	return response(pb.RpcBlockListMoveToNewObjectResponseError_NULL, linkId, nil)
}

func (mw *Middleware) BlockListConvertToObjects(cctx context.Context, req *pb.RpcBlockListConvertToObjectsRequest) *pb.RpcBlockListConvertToObjectsResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockListConvertToObjectsResponseErrorCode, linkIds []string, err error) *pb.RpcBlockListConvertToObjectsResponse {
		m := &pb.RpcBlockListConvertToObjectsResponse{Error: &pb.RpcBlockListConvertToObjectsResponseError{Code: code}, LinkIds: linkIds}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var linkIds []string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		linkIds, err = bs.ListConvertToObjects(ctx, *req)
		return
	})
	if err != nil {
		return response(pb.RpcBlockListConvertToObjectsResponseError_UNKNOWN_ERROR, []string{}, err)
	}
	return response(pb.RpcBlockListConvertToObjectsResponseError_NULL, linkIds, nil)
}

func (mw *Middleware) BlockTextListSetStyle(cctx context.Context, req *pb.RpcBlockTextListSetStyleRequest) *pb.RpcBlockTextListSetStyleResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockTextListSetStyleResponseErrorCode, err error) *pb.RpcBlockTextListSetStyleResponse {
		m := &pb.RpcBlockTextListSetStyleResponse{Error: &pb.RpcBlockTextListSetStyleResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetTextStyle(ctx, req.ContextId, req.Style, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockTextListSetStyleResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextListSetStyleResponseError_NULL, nil)
}

func (mw *Middleware) BlockDivListSetStyle(cctx context.Context, req *pb.RpcBlockDivListSetStyleRequest) *pb.RpcBlockDivListSetStyleResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockDivListSetStyleResponseErrorCode, err error) *pb.RpcBlockDivListSetStyleResponse {
		m := &pb.RpcBlockDivListSetStyleResponse{Error: &pb.RpcBlockDivListSetStyleResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetDivStyle(ctx, req.ContextId, req.Style, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockDivListSetStyleResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockDivListSetStyleResponseError_NULL, nil)
}

func (mw *Middleware) BlockTextListSetColor(cctx context.Context, req *pb.RpcBlockTextListSetColorRequest) *pb.RpcBlockTextListSetColorResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockTextListSetColorResponseErrorCode, err error) *pb.RpcBlockTextListSetColorResponse {
		m := &pb.RpcBlockTextListSetColorResponse{Error: &pb.RpcBlockTextListSetColorResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetTextColor(ctx, req.ContextId, req.Color, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockTextListSetColorResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextListSetColorResponseError_NULL, nil)
}

func (mw *Middleware) BlockTextListSetMark(cctx context.Context, req *pb.RpcBlockTextListSetMarkRequest) *pb.RpcBlockTextListSetMarkResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockTextListSetMarkResponseErrorCode, err error) *pb.RpcBlockTextListSetMarkResponse {
		m := &pb.RpcBlockTextListSetMarkResponse{Error: &pb.RpcBlockTextListSetMarkResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetTextMark(ctx, req.ContextId, req.Mark, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockTextListSetMarkResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextListSetMarkResponseError_NULL, nil)
}

func (mw *Middleware) newContext(cctx context.Context, opts ...state.ContextOption) *state.Context {
	grpcSender, ok := mw.EventSender.(*event.GrpcSender)
	if !ok {
		return state.NewContext()
	}

	md, ok := metadata.FromIncomingContext(cctx)
	if !ok {
		return state.NewContext()
	}

	v := md.Get("token")
	if len(v) != 1 {
		return state.NewContext()
	}

	tok := v[0]

	// TODO deny access
	if tok == "" {
		return state.NewContext()
	}

	token, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		acc, err := core.WalletAccountAt(mw.mnemonic, 0, "")
		if err != nil {
			return nil, err
		}
		priv, err := acc.MarshalBinary()
		if err != nil {
			return nil, err
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return priv, nil
	})
	if err != nil {
		log.Errorf("parse token: %s", err)
	}

	if token != nil && !token.Valid {
		log.Errorf("invalid token")
	}

	return state.NewContext(state.WithSessionId(v[0], grpcSender))
}

func (mw *Middleware) BlockTextListClearStyle(cctx context.Context, req *pb.RpcBlockTextListClearStyleRequest) *pb.RpcBlockTextListClearStyleResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockTextListClearStyleResponseErrorCode, err error) *pb.RpcBlockTextListClearStyleResponse {
		m := &pb.RpcBlockTextListClearStyleResponse{Error: &pb.RpcBlockTextListClearStyleResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.ClearTextStyle(ctx, req.ContextId, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockTextListClearStyleResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextListClearStyleResponseError_NULL, nil)
}

func (mw *Middleware) BlockTextListClearContent(cctx context.Context, req *pb.RpcBlockTextListClearContentRequest) *pb.RpcBlockTextListClearContentResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockTextListClearContentResponseErrorCode, err error) *pb.RpcBlockTextListClearContentResponse {
		m := &pb.RpcBlockTextListClearContentResponse{Error: &pb.RpcBlockTextListClearContentResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.ClearTextContent(ctx, req.ContextId, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockTextListClearContentResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextListClearContentResponseError_NULL, nil)
}

func (mw *Middleware) BlockTextSetText(cctx context.Context, req *pb.RpcBlockTextSetTextRequest) *pb.RpcBlockTextSetTextResponse {
	ctx := mw.newContext(cctx)

	response := func(code pb.RpcBlockTextSetTextResponseErrorCode, err error) *pb.RpcBlockTextSetTextResponse {
		m := &pb.RpcBlockTextSetTextResponse{Error: &pb.RpcBlockTextSetTextResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetTextText(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockTextSetTextResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextSetTextResponseError_NULL, nil)
}

func (mw *Middleware) BlockLatexSetText(cctx context.Context, req *pb.RpcBlockLatexSetTextRequest) *pb.RpcBlockLatexSetTextResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockLatexSetTextResponseErrorCode, err error) *pb.RpcBlockLatexSetTextResponse {
		m := &pb.RpcBlockLatexSetTextResponse{Error: &pb.RpcBlockLatexSetTextResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetLatexText(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockLatexSetTextResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockLatexSetTextResponseError_NULL, nil)
}

func (mw *Middleware) BlockTextSetStyle(cctx context.Context, req *pb.RpcBlockTextSetStyleRequest) *pb.RpcBlockTextSetStyleResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockTextSetStyleResponseErrorCode, err error) *pb.RpcBlockTextSetStyleResponse {
		m := &pb.RpcBlockTextSetStyleResponse{Error: &pb.RpcBlockTextSetStyleResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetTextStyle(ctx, req.ContextId, req.Style, req.BlockId)
	})
	if err != nil {
		return response(pb.RpcBlockTextSetStyleResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextSetStyleResponseError_NULL, nil)
}

func (mw *Middleware) BlockTextSetIcon(cctx context.Context, req *pb.RpcBlockTextSetIconRequest) *pb.RpcBlockTextSetIconResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockTextSetIconResponseErrorCode, err error) *pb.RpcBlockTextSetIconResponse {
		m := &pb.RpcBlockTextSetIconResponse{Error: &pb.RpcBlockTextSetIconResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetTextIcon(ctx, req.ContextId, req.IconImage, req.IconEmoji, req.BlockId)
	})
	if err != nil {
		return response(pb.RpcBlockTextSetIconResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextSetIconResponseError_NULL, nil)
}

func (mw *Middleware) BlockTextSetChecked(cctx context.Context, req *pb.RpcBlockTextSetCheckedRequest) *pb.RpcBlockTextSetCheckedResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockTextSetCheckedResponseErrorCode, err error) *pb.RpcBlockTextSetCheckedResponse {
		m := &pb.RpcBlockTextSetCheckedResponse{Error: &pb.RpcBlockTextSetCheckedResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetTextChecked(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockTextSetCheckedResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextSetCheckedResponseError_NULL, nil)
}

func (mw *Middleware) BlockFileSetName(cctx context.Context, req *pb.RpcBlockFileSetNameRequest) *pb.RpcBlockFileSetNameResponse {
	response := func(code pb.RpcBlockFileSetNameResponseErrorCode, err error) *pb.RpcBlockFileSetNameResponse {
		m := &pb.RpcBlockFileSetNameResponse{Error: &pb.RpcBlockFileSetNameResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}
	// TODO
	return response(pb.RpcBlockFileSetNameResponseError_NULL, nil)
}

func (mw *Middleware) BlockFileListSetStyle(cctx context.Context, req *pb.RpcBlockFileListSetStyleRequest) *pb.RpcBlockFileListSetStyleResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockFileListSetStyleResponseErrorCode, err error) *pb.RpcBlockFileListSetStyleResponse {
		m := &pb.RpcBlockFileListSetStyleResponse{Error: &pb.RpcBlockFileListSetStyleResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}

		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetFileStyle(ctx, req.ContextId, req.Style, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockFileListSetStyleResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcBlockFileListSetStyleResponseError_NULL, nil)
}

func (mw *Middleware) BlockImageSetName(cctx context.Context, req *pb.RpcBlockImageSetNameRequest) *pb.RpcBlockImageSetNameResponse {
	response := func(code pb.RpcBlockImageSetNameResponseErrorCode, err error) *pb.RpcBlockImageSetNameResponse {
		m := &pb.RpcBlockImageSetNameResponse{Error: &pb.RpcBlockImageSetNameResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}
	// TODO
	return response(pb.RpcBlockImageSetNameResponseError_NULL, nil)
}

func (mw *Middleware) BlockVideoSetName(cctx context.Context, req *pb.RpcBlockVideoSetNameRequest) *pb.RpcBlockVideoSetNameResponse {
	response := func(code pb.RpcBlockVideoSetNameResponseErrorCode, err error) *pb.RpcBlockVideoSetNameResponse {
		m := &pb.RpcBlockVideoSetNameResponse{Error: &pb.RpcBlockVideoSetNameResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}
	// TODO
	return response(pb.RpcBlockVideoSetNameResponseError_NULL, nil)
}

func (mw *Middleware) BlockSplit(cctx context.Context, req *pb.RpcBlockSplitRequest) *pb.RpcBlockSplitResponse {
	ctx := mw.newContext(cctx)
	response := func(blockId string, code pb.RpcBlockSplitResponseErrorCode, err error) *pb.RpcBlockSplitResponse {
		m := &pb.RpcBlockSplitResponse{BlockId: blockId, Error: &pb.RpcBlockSplitResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var id string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		id, err = bs.SplitBlock(ctx, *req)
		return
	})
	if err != nil {
		return response("", pb.RpcBlockSplitResponseError_UNKNOWN_ERROR, err)
	}
	return response(id, pb.RpcBlockSplitResponseError_NULL, nil)
}

func (mw *Middleware) BlockMerge(cctx context.Context, req *pb.RpcBlockMergeRequest) *pb.RpcBlockMergeResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockMergeResponseErrorCode, err error) *pb.RpcBlockMergeResponse {
		m := &pb.RpcBlockMergeResponse{Error: &pb.RpcBlockMergeResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.MergeBlock(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockMergeResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockMergeResponseError_NULL, nil)
}

func (mw *Middleware) BlockBookmarkFetch(cctx context.Context, req *pb.RpcBlockBookmarkFetchRequest) *pb.RpcBlockBookmarkFetchResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockBookmarkFetchResponseErrorCode, err error) *pb.RpcBlockBookmarkFetchResponse {
		m := &pb.RpcBlockBookmarkFetchResponse{Error: &pb.RpcBlockBookmarkFetchResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.BookmarkFetch(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockBookmarkFetchResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockBookmarkFetchResponseError_NULL, nil)
}

func (mw *Middleware) FileUpload(cctx context.Context, req *pb.RpcFileUploadRequest) *pb.RpcFileUploadResponse {
	response := func(hash string, code pb.RpcFileUploadResponseErrorCode, err error) *pb.RpcFileUploadResponse {
		m := &pb.RpcFileUploadResponse{Error: &pb.RpcFileUploadResponseError{Code: code}, Hash: hash}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	var hash string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		hash, err = bs.UploadFile(*req)
		return
	})
	if err != nil {
		return response("", pb.RpcFileUploadResponseError_UNKNOWN_ERROR, err)
	}
	return response(hash, pb.RpcFileUploadResponseError_NULL, nil)
}

func (mw *Middleware) FileDownload(cctx context.Context, req *pb.RpcFileDownloadRequest) *pb.RpcFileDownloadResponse {
	response := func(path string, code pb.RpcFileDownloadResponseErrorCode, err error) *pb.RpcFileDownloadResponse {
		m := &pb.RpcFileDownloadResponse{Error: &pb.RpcFileDownloadResponseError{Code: code}, LocalPath: path}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	if req.Path == "" {
		req.Path = mw.GetAnytype().TempDir() + string(os.PathSeparator) + "anytype-download"
	}

	err := os.MkdirAll(req.Path, 0755)
	if err != nil {
		return response("", pb.RpcFileDownloadResponseError_BAD_INPUT, err)
	}
	progress := process.NewProgress(pb.ModelProcess_SaveFile)
	defer progress.Finish()

	err = mw.doBlockService(func(bs block.Service) (err error) {
		return bs.ProcessAdd(progress)
	})
	if err != nil {
		return response("", pb.RpcFileDownloadResponseError_BAD_INPUT, err)
	}

	progress.SetProgressMessage("saving file")
	var countReader *datacounter.ReaderCounter
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-progress.Canceled():
				cancel()
			case <-time.After(time.Second):
				if countReader != nil {
					progress.SetDone(int64(countReader.Count()))
				}
			}
		}
	}()

	f, err := mw.getFileOrLargestImage(ctx, req.Hash)
	if err != nil {
		return response("", pb.RpcFileDownloadResponseError_BAD_INPUT, err)
	}

	progress.SetTotal(f.Meta().Size)

	r, err := f.Reader()
	if err != nil {
		return response("", pb.RpcFileDownloadResponseError_BAD_INPUT, err)
	}
	countReader = datacounter.NewReaderCounter(r)
	fileName := f.Meta().Name
	if fileName == "" {
		fileName = f.Info().Name
	}

	path, err := files.WriteReaderIntoFileReuseSameExistingFile(req.Path+string(os.PathSeparator)+fileName, countReader)
	if err != nil {
		return response("", pb.RpcFileDownloadResponseError_UNKNOWN_ERROR, err)
	}

	progress.SetDone(f.Meta().Size)

	return response(path, pb.RpcFileDownloadResponseError_NULL, nil)
}

func (mw *Middleware) getFileOrLargestImage(ctx context.Context, hash string) (core.File, error) {
	image, err := mw.GetAnytype().ImageByHash(ctx, hash)
	if err != nil {
		return mw.GetAnytype().FileByHash(ctx, hash)
	}

	f, err := image.GetOriginalFile(ctx)
	if err != nil {
		return mw.GetAnytype().FileByHash(ctx, hash)
	}

	return f, nil
}

func (mw *Middleware) BlockBookmarkCreateAndFetch(cctx context.Context, req *pb.RpcBlockBookmarkCreateAndFetchRequest) *pb.RpcBlockBookmarkCreateAndFetchResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockBookmarkCreateAndFetchResponseErrorCode, id string, err error) *pb.RpcBlockBookmarkCreateAndFetchResponse {
		m := &pb.RpcBlockBookmarkCreateAndFetchResponse{Error: &pb.RpcBlockBookmarkCreateAndFetchResponseError{Code: code}, BlockId: id}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var id string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		id, err = bs.BookmarkCreateAndFetch(ctx, *req)
		return
	})
	if err != nil {
		return response(pb.RpcBlockBookmarkCreateAndFetchResponseError_UNKNOWN_ERROR, "", err)
	}
	return response(pb.RpcBlockBookmarkCreateAndFetchResponseError_NULL, id, nil)
}

func (mw *Middleware) BlockFileCreateAndUpload(cctx context.Context, req *pb.RpcBlockFileCreateAndUploadRequest) *pb.RpcBlockFileCreateAndUploadResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockFileCreateAndUploadResponseErrorCode, id string, err error) *pb.RpcBlockFileCreateAndUploadResponse {
		m := &pb.RpcBlockFileCreateAndUploadResponse{Error: &pb.RpcBlockFileCreateAndUploadResponseError{Code: code}, BlockId: id}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var id string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		id, err = bs.CreateAndUploadFile(ctx, *req)
		return
	})
	if err != nil {
		return response(pb.RpcBlockFileCreateAndUploadResponseError_UNKNOWN_ERROR, "", err)
	}
	return response(pb.RpcBlockFileCreateAndUploadResponseError_NULL, id, nil)
}

func (mw *Middleware) ObjectSetObjectType(cctx context.Context, req *pb.RpcObjectSetObjectTypeRequest) *pb.RpcObjectSetObjectTypeResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectSetObjectTypeResponseErrorCode, err error) *pb.RpcObjectSetObjectTypeResponse {
		m := &pb.RpcObjectSetObjectTypeResponse{Error: &pb.RpcObjectSetObjectTypeResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}

	if err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetObjectTypes(ctx, req.ContextId, []string{req.ObjectTypeUrl})
	}); err != nil {
		return response(pb.RpcObjectSetObjectTypeResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcObjectSetObjectTypeResponseError_NULL, nil)
}

func (mw *Middleware) BlockRelationSetKey(cctx context.Context, req *pb.RpcBlockRelationSetKeyRequest) *pb.RpcBlockRelationSetKeyResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockRelationSetKeyResponseErrorCode, err error) *pb.RpcBlockRelationSetKeyResponse {
		m := &pb.RpcBlockRelationSetKeyResponse{Error: &pb.RpcBlockRelationSetKeyResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}

	if err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetRelationKey(ctx, *req)
	}); err != nil {
		return response(pb.RpcBlockRelationSetKeyResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcBlockRelationSetKeyResponseError_NULL, nil)
}

func (mw *Middleware) BlockRelationAdd(cctx context.Context, req *pb.RpcBlockRelationAddRequest) *pb.RpcBlockRelationAddResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockRelationAddResponseErrorCode, err error) *pb.RpcBlockRelationAddResponse {
		m := &pb.RpcBlockRelationAddResponse{Error: &pb.RpcBlockRelationAddResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}

	if err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.AddRelationBlock(ctx, *req)
	}); err != nil {
		return response(pb.RpcBlockRelationAddResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcBlockRelationAddResponseError_NULL, nil)
}

func (mw *Middleware) BlockListTurnInto(cctx context.Context, req *pb.RpcBlockListTurnIntoRequest) *pb.RpcBlockListTurnIntoResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockListTurnIntoResponseErrorCode, err error) *pb.RpcBlockListTurnIntoResponse {
		m := &pb.RpcBlockListTurnIntoResponse{Error: &pb.RpcBlockListTurnIntoResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.TurnInto(ctx, req.ContextId, req.Style, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockListTurnIntoResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockListTurnIntoResponseError_NULL, nil)
}
