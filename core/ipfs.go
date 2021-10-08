package core

import (
	"context"
	"fmt"
	pb2 "github.com/anytypeio/go-anytype-middleware/pkg/lib/cafe/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pin"
	"io/ioutil"

	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
)

func (mw *Middleware) ImageGetBlob(req *pb.RpcIpfsImageGetBlobRequest) *pb.RpcIpfsImageGetBlobResponse {
	mw.m.RLock()
	defer mw.m.RUnlock()
	response := func(blob []byte, code pb.RpcIpfsImageGetBlobResponseErrorCode, err error) *pb.RpcIpfsImageGetBlobResponse {
		m := &pb.RpcIpfsImageGetBlobResponse{Blob: blob, Error: &pb.RpcIpfsImageGetBlobResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}

	if mw.app == nil {
		response(nil, pb.RpcIpfsImageGetBlobResponseError_NODE_NOT_STARTED, fmt.Errorf("anytype is nil"))
	}

	at := mw.app.MustComponent(core.CName).(core.Service)

	if !at.IsStarted() {
		response(nil, pb.RpcIpfsImageGetBlobResponseError_NODE_NOT_STARTED, fmt.Errorf("anytype node not started"))
	}

	image, err := at.ImageByHash(context.TODO(), req.GetHash())
	if err != nil {
		if err == core.ErrFileNotFound {
			return response(nil, pb.RpcIpfsImageGetBlobResponseError_NOT_FOUND, err)
		}

		return response(nil, pb.RpcIpfsImageGetBlobResponseError_UNKNOWN_ERROR, err)
	}
	file, err := image.GetFileForWidth(context.TODO(), int(req.WantWidth))
	if err != nil {
		if err == core.ErrFileNotFound {
			return response(nil, pb.RpcIpfsImageGetBlobResponseError_NOT_FOUND, err)
		}

		return response(nil, pb.RpcIpfsImageGetBlobResponseError_UNKNOWN_ERROR, err)
	}

	rd, err := file.Reader()
	if err != nil {
		if err == core.ErrFileNotFound {
			return response(nil, pb.RpcIpfsImageGetBlobResponseError_NOT_FOUND, err)
		}

		return response(nil, pb.RpcIpfsImageGetBlobResponseError_UNKNOWN_ERROR, err)
	}
	data, err := ioutil.ReadAll(rd)
	if err != nil {
		if err == core.ErrFileNotFound {
			return response(nil, pb.RpcIpfsImageGetBlobResponseError_NOT_FOUND, err)
		}

		return response(nil, pb.RpcIpfsImageGetBlobResponseError_UNKNOWN_ERROR, err)
	}
	return response(data, pb.RpcIpfsImageGetBlobResponseError_NULL, nil)
}

func (mw *Middleware) FilesOffloadAll(req *pb.RpcIpfsFileOffloadAllRequest) *pb.RpcIpfsFileOffloadAllResponse {
	mw.m.RLock()
	defer mw.m.RUnlock()
	response := func(filesOffloaded int32, bytesOffloaded uint64, code pb.RpcIpfsFileOffloadAllResponseErrorCode, err error) *pb.RpcIpfsFileOffloadAllResponse {
		m := &pb.RpcIpfsFileOffloadAllResponse{Error: &pb.RpcIpfsFileOffloadAllResponseError{Code: code}, BytesOffloaded: bytesOffloaded, OffloadedFiles: filesOffloaded}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}

	if mw.app == nil {
		response(0, 0, pb.RpcIpfsFileOffloadAllResponseError_NODE_NOT_STARTED, fmt.Errorf("anytype is nil"))
	}

	at := mw.app.MustComponent(core.CName).(core.Service)
	pin := mw.app.MustComponent(pin.CName).(pin.FilePinService)

	if !at.IsStarted() {
		response(0, 0, pb.RpcIpfsFileOffloadAllResponseError_NODE_NOT_STARTED, fmt.Errorf("anytype node not started"))
	}

	files, err := at.FileStore().ListTargets()
	if err != nil {
		return response(0, 0, pb.RpcIpfsFileOffloadAllResponseError_UNKNOWN_ERROR, err)
	}
	pinStatus := pin.PinStatus(files...)
	var (
		totalBytesOffloaded uint64
		totalFilesOffloaded int32
	)
	for fileId, status := range pinStatus {
		if status.Status != pb2.PinStatus_Done && !req.IncludeNotPinned {
			continue
		}
		bytesRemoved, err := at.FileOffload(fileId)
		if err != nil {
			log.Errorf("failed to offload file %s: %s", fileId, err.Error())
			continue
		}
		totalBytesOffloaded += bytesRemoved
		totalFilesOffloaded++
	}

	return response(totalFilesOffloaded, totalBytesOffloaded, pb.RpcIpfsFileOffloadAllResponseError_NULL, nil)
}

func (mw *Middleware) FileOffload(req *pb.RpcIpfsFileOffloadRequest) *pb.RpcIpfsFileOffloadResponse {
	mw.m.RLock()
	defer mw.m.RUnlock()
	response := func(bytesOffloaded uint64, code pb.RpcIpfsFileOffloadResponseErrorCode, err error) *pb.RpcIpfsFileOffloadResponse {
		m := &pb.RpcIpfsFileOffloadResponse{BytesOffloaded: bytesOffloaded, Error: &pb.RpcIpfsFileOffloadResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}

	if mw.app == nil {
		response(0, pb.RpcIpfsFileOffloadResponseError_NODE_NOT_STARTED, fmt.Errorf("anytype is nil"))
	}

	at := mw.app.MustComponent(core.CName).(core.Service)
	pin := mw.app.MustComponent(pin.CName).(pin.FilePinService)

	if !at.IsStarted() {
		response(0, pb.RpcIpfsFileOffloadResponseError_NODE_NOT_STARTED, fmt.Errorf("anytype node not started"))
	}

	pinStatus := pin.PinStatus(req.Id)
	var (
		totalBytesOffloaded uint64
	)
	for fileId, status := range pinStatus {
		if status.Status != pb2.PinStatus_Done && !req.IncludeNotPinned {
			continue
		}
		bytesRemoved, err := at.FileOffload(fileId)
		if err != nil {
			log.Errorf("failed to offload file %s: %s", fileId, err.Error())
			continue
		}
		totalBytesOffloaded += bytesRemoved
	}

	return response(totalBytesOffloaded, pb.RpcIpfsFileOffloadResponseError_NULL, nil)
}
