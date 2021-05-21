package core

import (
	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/export"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
)

func (mw *Middleware) MakeTemplate(req *pb.RpcMakeTemplateRequest) *pb.RpcMakeTemplateResponse {
	response := func(templateId string, err error) *pb.RpcMakeTemplateResponse {
		m := &pb.RpcMakeTemplateResponse{
			Error: &pb.RpcMakeTemplateResponseError{Code: pb.RpcMakeTemplateResponseError_NULL},
			Id:    templateId,
		}
		if err != nil {
			m.Error.Code = pb.RpcMakeTemplateResponseError_UNKNOWN_ERROR
			m.Error.Description = err.Error()
		}
		return m
	}
	var templateId string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		templateId, err = bs.MakeTemplate(req.ContextId)
		return
	})
	return response(templateId, err)
}

func (mw *Middleware) ExportTemplates(req *pb.RpcExportTemplatesRequest) *pb.RpcExportTemplatesResponse {
	response := func(path string, err error) (res *pb.RpcExportTemplatesResponse) {
		res = &pb.RpcExportTemplatesResponse{
			Error: &pb.RpcExportTemplatesResponseError{
				Code: pb.RpcExportTemplatesResponseError_NULL,
			},
		}
		if err != nil {
			res.Error.Code = pb.RpcExportTemplatesResponseError_UNKNOWN_ERROR
			res.Error.Description = err.Error()
			return
		} else {
			res.Path = path
		}
		return res
	}
	var (
		path string
		err  error
	)
	err = mw.doBlockService(func(_ block.Service) error {
		es := mw.app.MustComponent(export.CName).(export.Export)
		ds := mw.app.MustComponent(objectstore.CName).(objectstore.ObjectStore)
		res, _, err := ds.QueryObjectInfo(database.Query{}, []smartblock.SmartBlockType{smartblock.SmartBlockTypeTemplate})
		if err != nil {
			return err
		}
		var docIds []string
		for _, r := range res {
			docIds = append(docIds, r.Id)
		}
		path, err = es.Export(pb.RpcExportRequest{
			Path:   req.Path,
			DocIds: docIds,
			Format: pb.RpcExport_Protobuf,
			Zip:    true,
		})
		return err
	})
	return response(path, err)
}
