package core

import (
	"context"

	"github.com/anyproto/any-sync/app"
	"github.com/globalsign/mgo/bson"
	"google.golang.org/grpc/metadata"

	"github.com/anyproto/anytype-heart/core/block"
	"github.com/anyproto/anytype-heart/core/block/source"
	"github.com/anyproto/anytype-heart/core/session"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space"
)

func (mw *Middleware) BlockCreate(cctx context.Context, req *pb.RpcBlockCreateRequest) *pb.RpcBlockCreateResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockCreateResponseErrorCode, id string, err error) *pb.RpcBlockCreateResponse {
		m := &pb.RpcBlockCreateResponse{Error: &pb.RpcBlockCreateResponseError{Code: code}, BlockId: id}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var id string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var id, targetId string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		id, targetId, err = bs.CreateLinkToTheNewObject(ctx, req)
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

	err := mw.doBlockService(func(bs *block.Service) (err error) {
		obj, err = bs.OpenBlock(ctx, req.ObjectId, req.IncludeRelationsAsDependentObjects)
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
	ctx := mw.newContext(cctx, session.WithTraceId(req.TraceId))
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

	err := mw.doBlockService(func(bs *block.Service) (err error) {
		obj, err = bs.ShowBlock(ctx, req.ObjectId, req.IncludeRelationsAsDependentObjects)
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

func (mw *Middleware) ObjectClose(cctx context.Context, req *pb.RpcObjectCloseRequest) *pb.RpcObjectCloseResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectCloseResponseErrorCode, err error) *pb.RpcObjectCloseResponse {
		m := &pb.RpcObjectCloseResponse{Error: &pb.RpcObjectCloseResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.CloseBlock(ctx, req.ObjectId)
	})
	if err != nil {
		return response(pb.RpcObjectCloseResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectCloseResponseError_NULL, nil)
}
func (mw *Middleware) BlockCopy(cctx context.Context, req *pb.RpcBlockCopyRequest) *pb.RpcBlockCopyResponse {
	ctx := mw.newContext(cctx)
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
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		textSlot, htmlSlot, anySlot, err = bs.Copy(ctx, *req)
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var (
		blockIds         []string
		caretPosition    int32
		isSameBlockCaret bool
		groupId          = bson.NewObjectId().Hex()
	)
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var (
		textSlot, htmlSlot string
		anySlot            []*model.Block
	)
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		textSlot, htmlSlot, anySlot, err = bs.Cut(ctx, *req)
		return
	})
	if err != nil {
		var emptyAnySlot []*model.Block
		return response(pb.RpcBlockCutResponseError_UNKNOWN_ERROR, "", "", emptyAnySlot, err)
	}

	return response(pb.RpcBlockCutResponseError_NULL, textSlot, htmlSlot, anySlot, nil)
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var path string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		path, err = bs.Export(ctx, *req)
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var ids []string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.SetFieldsList(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockListSetFieldsResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockListSetFieldsResponseError_NULL, nil)
}

func (mw *Middleware) ObjectListDelete(cctx context.Context, req *pb.RpcObjectListDeleteRequest) *pb.RpcObjectListDeleteResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectListDeleteResponseErrorCode, err error) *pb.RpcObjectListDeleteResponse {
		m := &pb.RpcObjectListDeleteResponse{Error: &pb.RpcObjectListDeleteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.DeleteArchivedObjects(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcObjectListDeleteResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectListDeleteResponseError_NULL, nil)
}

func (mw *Middleware) ObjectListSetIsArchived(cctx context.Context, req *pb.RpcObjectListSetIsArchivedRequest) *pb.RpcObjectListSetIsArchivedResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectListSetIsArchivedResponseErrorCode, err error) *pb.RpcObjectListSetIsArchivedResponse {
		m := &pb.RpcObjectListSetIsArchivedResponse{Error: &pb.RpcObjectListSetIsArchivedResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.SetPagesIsArchived(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcObjectListSetIsArchivedResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectListSetIsArchivedResponseError_NULL, nil)
}
func (mw *Middleware) ObjectListSetIsFavorite(cctx context.Context, req *pb.RpcObjectListSetIsFavoriteRequest) *pb.RpcObjectListSetIsFavoriteResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectListSetIsFavoriteResponseErrorCode, err error) *pb.RpcObjectListSetIsFavoriteResponse {
		m := &pb.RpcObjectListSetIsFavoriteResponse{Error: &pb.RpcObjectListSetIsFavoriteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.SetPagesIsFavorite(ctx, *req)
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var blockId string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.SetVerticalAlign(ctx, req.ContextId, req.VerticalAlign, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockListSetVerticalAlignResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockListSetVerticalAlignResponseError_NULL, nil)
}

func (mw *Middleware) BlockListMoveToExistingObject(cctx context.Context, req *pb.RpcBlockListMoveToExistingObjectRequest) *pb.RpcBlockListMoveToExistingObjectResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockListMoveToExistingObjectResponseErrorCode, err error) *pb.RpcBlockListMoveToExistingObjectResponse {
		m := &pb.RpcBlockListMoveToExistingObjectResponse{Error: &pb.RpcBlockListMoveToExistingObjectResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}

	var linkId string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var linkIds []string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.SetTextMark(ctx, req.ContextId, req.Mark, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockTextListSetMarkResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockTextListSetMarkResponseError_NULL, nil)
}

func (mw *Middleware) newContext(cctx context.Context, opts ...session.ContextOption) session.Context {
	var spaceID string
	tok, ok := getSessionToken(cctx)
	if ok {
		var err error
		spaceID, err = mw.sessions.GetSpaceID(tok)
		if err != nil {
			log.Errorf("failed to get space id from token: %v", err)
		}
	}
	if spaceID == "" {
		// spaceID = "bafyreian4fawywbjksy4qhv7kp4374gjs3k4krdkjrpl5gbji6t4t6k3ji.0"
		log.Errorf("newContext: set spaceID to accountID")
		spaceID = getService[space.Service](mw).AccountId()
	}

	return mw.newContextWithSpace(cctx, spaceID, opts...)
}

func (mw *Middleware) newContextNoLock(cctx context.Context, opts ...session.ContextOption) session.Context {
	var spaceID string
	tok, ok := getSessionToken(cctx)
	if ok {
		var err error
		spaceID, err = mw.sessions.GetSpaceID(tok)
		if err != nil {
			log.Errorf("failed to get space id from token: %v", err)
		}
	}
	if spaceID == "" {
		// spaceID = "bafyreian4fawywbjksy4qhv7kp4374gjs3k4krdkjrpl5gbji6t4t6k3ji.0"
		log.Errorf("newContextNoLock: set spaceID to accountID")
		spaceID = app.MustComponent[space.Service](mw.app).AccountId()
	}

	return mw.newContextWithSpace(cctx, spaceID, opts...)
}

func getSessionToken(cctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(cctx)
	if !ok {
		return "", false
	}
	v := md.Get("token")
	if len(v) != 1 {
		return "", false
	}

	tok := v[0]
	if tok == "" {
		return "", false
	}
	return tok, true
}

func (mw *Middleware) newContextWithSpace(cctx context.Context, spaceID string, opts ...session.ContextOption) session.Context {
	tok, ok := getSessionToken(cctx)
	if !ok {
		session.NewContext(cctx, spaceID, append(opts, session.WithSession(tok))...)
	}
	return session.NewContext(cctx, spaceID, opts...)
}

func (mw *Middleware) BlockTextListClearStyle(cctx context.Context, req *pb.RpcBlockTextListClearStyleRequest) *pb.RpcBlockTextListClearStyleResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockTextListClearStyleResponseErrorCode, err error) *pb.RpcBlockTextListClearStyleResponse {
		m := &pb.RpcBlockTextListClearStyleResponse{Error: &pb.RpcBlockTextListClearStyleResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}

		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var id string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.BookmarkFetch(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcBlockBookmarkFetchResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockBookmarkFetchResponseError_NULL, nil)
}

func (mw *Middleware) BlockBookmarkCreateAndFetch(cctx context.Context, req *pb.RpcBlockBookmarkCreateAndFetchRequest) *pb.RpcBlockBookmarkCreateAndFetchResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockBookmarkCreateAndFetchResponseErrorCode, id string, err error) *pb.RpcBlockBookmarkCreateAndFetchResponse {
		m := &pb.RpcBlockBookmarkCreateAndFetchResponse{Error: &pb.RpcBlockBookmarkCreateAndFetchResponseError{Code: code}, BlockId: id}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var id string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	var id string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}

	if err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}

	if err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}

	if err := mw.doBlockService(func(bs *block.Service) (err error) {
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
			m.Event = mw.getResponseEvent(ctx)
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.TurnInto(ctx, req.ContextId, req.Style, req.BlockIds...)
	})
	if err != nil {
		return response(pb.RpcBlockListTurnIntoResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcBlockListTurnIntoResponseError_NULL, nil)
}
