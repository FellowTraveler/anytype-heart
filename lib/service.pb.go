// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/protos/service/service.proto

package lib

import (
	fmt "fmt"
	pb "github.com/anytypeio/go-anytype-middleware/pb"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("pb/protos/service/service.proto", fileDescriptor_93a29dc403579097) }

var fileDescriptor_93a29dc403579097 = []byte{
	// 1136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x98, 0x4f, 0x6f, 0xdc, 0x44,
	0x18, 0xc6, 0xb3, 0x02, 0x51, 0x31, 0x84, 0x40, 0xa7, 0x80, 0x68, 0x50, 0x37, 0x25, 0x4d, 0x20,
	0x49, 0xc1, 0x2d, 0x2d, 0x12, 0x5c, 0x38, 0x34, 0x9b, 0x34, 0x8d, 0x94, 0x94, 0x90, 0x0d, 0x8d,
	0x54, 0x81, 0xc0, 0xf1, 0xbe, 0xd9, 0x0c, 0xeb, 0x9d, 0x71, 0xed, 0xc9, 0xd2, 0xbd, 0x72, 0x40,
	0x1c, 0x39, 0xf0, 0xa1, 0x38, 0xf6, 0xc8, 0x11, 0x25, 0x5f, 0xa4, 0x9a, 0xf1, 0xeb, 0x3f, 0x63,
	0x7b, 0x66, 0x9d, 0x43, 0x12, 0xc5, 0xcf, 0xef, 0x7d, 0x9e, 0x19, 0xbf, 0x33, 0x63, 0xef, 0x92,
	0xa5, 0xe8, 0xe4, 0x5e, 0x14, 0x0b, 0x29, 0x92, 0x7b, 0x09, 0xc4, 0x13, 0x16, 0x40, 0xf6, 0xd7,
	0xd3, 0x97, 0xe9, 0x35, 0x9f, 0x4f, 0xe5, 0x34, 0x82, 0xc5, 0x8f, 0x0b, 0x32, 0x10, 0xe3, 0xb1,
	0xcf, 0x07, 0x49, 0x8a, 0x2c, 0x7e, 0x54, 0x28, 0x30, 0x01, 0x2e, 0xf1, 0xfa, 0x83, 0x3f, 0x57,
	0xc9, 0x42, 0x2f, 0x64, 0xc0, 0x65, 0x0f, 0x0b, 0xe8, 0x31, 0x99, 0x3f, 0xf6, 0xc3, 0x10, 0x64,
	0x2f, 0x06, 0x5f, 0x02, 0x5d, 0xf6, 0xd0, 0xde, 0x3b, 0x8c, 0x02, 0x2f, 0x95, 0xbc, 0x54, 0xf3,
	0x0e, 0xe1, 0xc5, 0x39, 0x24, 0x72, 0xf1, 0x8e, 0x93, 0x49, 0x22, 0xc1, 0x13, 0xa0, 0xcf, 0xc9,
	0xbb, 0xa9, 0x72, 0x08, 0x81, 0x98, 0x40, 0x4c, 0x1b, 0xab, 0x50, 0xcc, 0xad, 0x57, 0xdc, 0x10,
	0x7a, 0xff, 0x4c, 0x16, 0x1e, 0x05, 0x81, 0x38, 0xe7, 0xb9, 0xb9, 0x59, 0x87, 0x62, 0xcd, 0x7d,
	0x75, 0x06, 0x55, 0x0c, 0x1d, 0x35, 0xbc, 0x29, 0x77, 0x1a, 0xeb, 0x2a, 0x77, 0x65, 0xc5, 0x0d,
	0xd5, 0xbc, 0xfb, 0x10, 0x42, 0x20, 0x2d, 0xde, 0xa9, 0x38, 0xc3, 0x3b, 0x87, 0xd0, 0x3b, 0x20,
	0xf3, 0xbb, 0x63, 0x7f, 0x08, 0x3b, 0x20, 0x37, 0x43, 0x71, 0x42, 0xd7, 0x8c, 0xaa, 0xdd, 0xe8,
	0x34, 0xf1, 0xb4, 0xee, 0xed, 0x80, 0xf4, 0x14, 0x91, 0xfb, 0xaf, 0xb7, 0x20, 0x31, 0xe4, 0x07,
	0x42, 0x9e, 0x41, 0x9c, 0x30, 0xc1, 0x77, 0x40, 0xd2, 0xdb, 0x46, 0x21, 0x0a, 0xba, 0x2a, 0xb3,
	0xfe, 0xd4, 0x41, 0xa0, 0xe5, 0x13, 0x72, 0x6d, 0x4f, 0x0c, 0xfb, 0xc0, 0x07, 0xf4, 0x96, 0x41,
	0xef, 0x89, 0xa1, 0xa7, 0x2e, 0xe7, 0x66, 0x5d, 0x9b, 0x8c, 0x4e, 0x4f, 0xc9, 0xdb, 0x3d, 0xc1,
	0x4f, 0xd9, 0x50, 0x8d, 0x6d, 0xc9, 0x80, 0xd3, 0xeb, 0xc6, 0xd0, 0x6e, 0xdb, 0x01, 0xf4, 0x3b,
	0x25, 0xd7, 0xb7, 0x5f, 0x4a, 0x88, 0xb9, 0x1f, 0x6e, 0xc5, 0x22, 0x7a, 0xcc, 0x42, 0x48, 0xe8,
	0xe7, 0x46, 0x59, 0x59, 0xf7, 0x34, 0x90, 0xfb, 0xaf, 0xcd, 0x06, 0x31, 0x27, 0x24, 0x37, 0xca,
	0x72, 0x4f, 0x70, 0x09, 0x5c, 0xd2, 0x75, 0xbb, 0x01, 0x22, 0x79, 0xd6, 0x46, 0x1b, 0x14, 0xd3,
	0x8e, 0xc8, 0x3b, 0x9b, 0xa1, 0x08, 0x46, 0x3f, 0x46, 0xa1, 0xf0, 0x07, 0xd4, 0xec, 0x90, 0x56,
	0xbc, 0x54, 0xca, 0xdd, 0x97, 0x5d, 0x08, 0xba, 0x1e, 0x93, 0x79, 0x2d, 0x1c, 0x42, 0x14, 0xfa,
	0x41, 0xf5, 0x24, 0x49, 0x6b, 0x50, 0xb3, 0x9c, 0x24, 0x55, 0xa6, 0x68, 0xaa, 0x56, 0xbe, 0x8f,
	0x80, 0x57, 0x9a, 0x9a, 0x56, 0x28, 0xc1, 0xd2, 0x54, 0x03, 0xa8, 0x4c, 0x1f, 0x37, 0x77, 0xd3,
	0xf4, 0x2b, 0x5b, 0x7b, 0xd9, 0x85, 0xa0, 0xeb, 0xaf, 0xe4, 0xbd, 0x92, 0xeb, 0x81, 0x3f, 0x04,
	0xba, 0x6a, 0x2d, 0x53, 0x72, 0xee, 0xfe, 0xd9, 0x2c, 0xac, 0xda, 0x36, 0x1e, 0x32, 0x3e, 0x6a,
	0x6e, 0x9b, 0x96, 0xdc, 0x6d, 0xcb, 0x90, 0x62, 0x3f, 0xa7, 0xe3, 0x0e, 0x45, 0x02, 0xb4, 0xe9,
	0xee, 0x69, 0xc5, 0xb2, 0x9f, 0x4d, 0xa2, 0x38, 0xe3, 0xf4, 0xf5, 0x2d, 0xf1, 0x3b, 0xd7, 0x2b,
	0xac, 0xa9, 0xcd, 0x99, 0x68, 0x39, 0xe3, 0x6a, 0x10, 0x7a, 0xff, 0x84, 0xde, 0x3b, 0x20, 0xf7,
	0xfd, 0x78, 0x94, 0xd0, 0xa6, 0x32, 0xb5, 0x87, 0xb5, 0x6a, 0x39, 0xf9, 0xeb, 0x14, 0xba, 0x03,
	0x79, 0x5f, 0x6b, 0x4f, 0x58, 0x22, 0x45, 0x3c, 0xdd, 0x17, 0x13, 0xa8, 0x6c, 0xf7, 0xb4, 0x14,
	0x75, 0x4f, 0x01, 0x96, 0xed, 0xde, 0x08, 0x62, 0xcc, 0x2f, 0x64, 0x41, 0xcb, 0x7d, 0x90, 0x8f,
	0x19, 0x84, 0x83, 0xa4, 0x71, 0xa9, 0xf4, 0x41, 0x7a, 0xa9, 0xec, 0x5c, 0x2a, 0x06, 0x86, 0x01,
	0x2f, 0xc8, 0x07, 0x59, 0xc0, 0x21, 0x24, 0x32, 0x66, 0x81, 0x64, 0x82, 0x27, 0xf4, 0xae, 0xa5,
	0xbe, 0x0c, 0xe5, 0x61, 0x5f, 0xb4, 0x83, 0x31, 0x72, 0x44, 0x68, 0x16, 0xb9, 0x9b, 0x3c, 0x8a,
	0x83, 0x33, 0x36, 0x81, 0x41, 0xe5, 0x04, 0x2b, 0x3c, 0x0a, 0xc4, 0x72, 0x82, 0x59, 0xd0, 0xca,
	0x0a, 0xdb, 0x63, 0x89, 0xd4, 0x4d, 0x6a, 0x58, 0x61, 0x4a, 0x33, 0x1b, 0xb4, 0xe2, 0x86, 0x2a,
	0x13, 0x51, 0x52, 0xd1, 0xa0, 0x75, 0x4b, 0x6d, 0x43, 0x93, 0x36, 0xda, 0xa0, 0x18, 0x36, 0x21,
	0x1f, 0x96, 0xc3, 0x8e, 0xe0, 0xa5, 0xec, 0xcb, 0x69, 0x08, 0xf4, 0x4b, 0x87, 0x89, 0xa2, 0x3c,
	0x8d, 0xe5, 0x99, 0x5e, 0x5b, 0x1c, 0x73, 0x59, 0x69, 0x92, 0x5b, 0xe7, 0x51, 0xc8, 0x02, 0x75,
	0x14, 0xae, 0x59, 0x5c, 0x72, 0xc2, 0xf2, 0xc2, 0xd0, 0x4c, 0x62, 0xd4, 0x10, 0xf7, 0x14, 0x4e,
	0x4f, 0xfd, 0x34, 0x05, 0x15, 0x43, 0xd5, 0xbf, 0x1c, 0x41, 0x35, 0x12, 0x83, 0x7e, 0x23, 0xd7,
	0xcb, 0x41, 0x3d, 0x11, 0x8a, 0x98, 0x3a, 0xeb, 0x35, 0x32, 0x73, 0x01, 0x1a, 0xa8, 0xbd, 0x6f,
	0x69, 0xde, 0xcc, 0xbe, 0x99, 0x99, 0x5e, 0x5b, 0x1c, 0x73, 0xff, 0xea, 0x90, 0x4f, 0xca, 0x93,
	0xdc, 0xf4, 0x83, 0xd1, 0x30, 0x16, 0xe7, 0x7c, 0x90, 0xc6, 0x3f, 0x74, 0xcd, 0xa1, 0x02, 0xe7,
	0x83, 0xf8, 0xfa, 0x6a, 0x45, 0x38, 0x94, 0x7f, 0x3a, 0x64, 0xa9, 0x7a, 0x0f, 0xaa, 0xc3, 0xf9,
	0x66, 0xd6, 0xf4, 0x6c, 0x43, 0xfa, 0xf6, 0xea, 0x85, 0xcd, 0xab, 0x20, 0xdd, 0x4d, 0xce, 0x55,
	0x60, 0xee, 0xa4, 0x8d, 0x36, 0x28, 0x66, 0x45, 0xe4, 0x86, 0xb1, 0xe2, 0xce, 0x20, 0x18, 0xc1,
	0xc0, 0x7a, 0xca, 0xa6, 0x0d, 0x4d, 0xa1, 0x99, 0xa7, 0x6c, 0x05, 0xae, 0x3c, 0xad, 0xfb, 0x51,
	0xc8, 0x64, 0xe3, 0xd3, 0x5a, 0x2b, 0xce, 0xa7, 0x75, 0x46, 0x54, 0x2c, 0xf7, 0x21, 0x1e, 0x36,
	0xbf, 0x00, 0x68, 0xc5, 0x69, 0x99, 0x11, 0x95, 0x37, 0xb6, 0x9e, 0x88, 0xa6, 0x8d, 0x6f, 0x6c,
	0x4a, 0x70, 0xbe, 0xb1, 0x21, 0x50, 0x19, 0xe2, 0x81, 0x9f, 0xc8, 0xe6, 0x21, 0x6a, 0xc5, 0x39,
	0xc4, 0x8c, 0xa8, 0x9f, 0x4a, 0xea, 0x5d, 0xfc, 0xa9, 0x3f, 0x06, 0xeb, 0xa9, 0xa4, 0x00, 0x4f,
	0x11, 0x33, 0x4f, 0xa5, 0x32, 0x59, 0x5f, 0x8f, 0xfa, 0x23, 0x95, 0x4e, 0xb2, 0x3e, 0x16, 0xf5,
	0x87, 0x2e, 0x23, 0x6a, 0xa3, 0x0d, 0x8a, 0x59, 0xe3, 0xd2, 0x33, 0x58, 0xc9, 0xc7, 0x6c, 0x20,
	0xcf, 0xa8, 0xdb, 0x41, 0x33, 0x79, 0xda, 0xdd, 0x56, 0x6c, 0x7d, 0x6a, 0xcf, 0xd8, 0x00, 0x84,
	0x73, 0x6a, 0x9a, 0x68, 0x37, 0x35, 0x03, 0xad, 0x4f, 0x4d, 0xcb, 0xee, 0xa9, 0xa5, 0x0e, 0xed,
	0xa6, 0x66, 0xb2, 0xf5, 0xe5, 0xb1, 0x1b, 0x08, 0xee, 0x5c, 0x1e, 0x0a, 0x68, 0xb7, 0x3c, 0xca,
	0x24, 0x06, 0xfd, 0xd1, 0x21, 0x37, 0xb3, 0xa4, 0x3d, 0xc6, 0x47, 0x47, 0x7e, 0x3c, 0xd4, 0x1f,
	0xdf, 0x83, 0xd1, 0xee, 0x80, 0x7e, 0x65, 0x31, 0x52, 0xa4, 0x67, 0xa0, 0x79, 0xf6, 0x83, 0xab,
	0x94, 0xe0, 0x20, 0xbe, 0x23, 0x6f, 0x1e, 0x30, 0x3e, 0xa4, 0x37, 0x8d, 0x5a, 0x75, 0x29, 0xb7,
	0x5d, 0x6c, 0x92, 0xb0, 0xfc, 0x3e, 0x99, 0x57, 0x87, 0x33, 0xf0, 0x6d, 0xfd, 0x65, 0x13, 0x5d,
	0xc8, 0xd9, 0xed, 0x71, 0x24, 0xa7, 0x8b, 0xa5, 0xff, 0x15, 0x70, 0xbf, 0xb3, 0x79, 0xeb, 0xdf,
	0x8b, 0x6e, 0xe7, 0xd5, 0x45, 0xb7, 0xf3, 0xff, 0x45, 0xb7, 0xf3, 0xf7, 0x65, 0x77, 0xee, 0xd5,
	0x65, 0x77, 0xee, 0xbf, 0xcb, 0xee, 0xdc, 0xf3, 0x37, 0x42, 0x76, 0x72, 0xf2, 0x96, 0xfe, 0xba,
	0xea, 0xe1, 0xeb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x8d, 0xf1, 0xb9, 0x0c, 0x13, 0x00, 0x00,
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the gomobile package it is being compiled against.

// ClientCommandsHandler is the handler API for ClientCommands service.
var clientCommandsHandler ClientCommandsHandler

type ClientCommandsHandler interface {
	WalletCreate(*pb.RpcWalletCreateRequest) *pb.RpcWalletCreateResponse
	WalletRecover(*pb.RpcWalletRecoverRequest) *pb.RpcWalletRecoverResponse
	AccountRecover(*pb.RpcAccountRecoverRequest) *pb.RpcAccountRecoverResponse
	AccountCreate(*pb.RpcAccountCreateRequest) *pb.RpcAccountCreateResponse
	AccountSelect(*pb.RpcAccountSelectRequest) *pb.RpcAccountSelectResponse
	ImageGetBlob(*pb.RpcIpfsImageGetBlobRequest) *pb.RpcIpfsImageGetBlobResponse
	VersionGet(*pb.RpcVersionGetRequest) *pb.RpcVersionGetResponse
	LogSend(*pb.RpcLogSendRequest) *pb.RpcLogSendResponse
	ConfigGet(*pb.RpcConfigGetRequest) *pb.RpcConfigGetResponse
	ExternalDropFiles(*pb.RpcExternalDropFilesRequest) *pb.RpcExternalDropFilesResponse
	ExternalDropContent(*pb.RpcExternalDropContentRequest) *pb.RpcExternalDropContentResponse
	BlockUpload(*pb.RpcBlockUploadRequest) *pb.RpcBlockUploadResponse
	BlockReplace(*pb.RpcBlockReplaceRequest) *pb.RpcBlockReplaceResponse
	BlockOpen(*pb.RpcBlockOpenRequest) *pb.RpcBlockOpenResponse
	BlockCreate(*pb.RpcBlockCreateRequest) *pb.RpcBlockCreateResponse
	BlockCreatePage(*pb.RpcBlockCreatePageRequest) *pb.RpcBlockCreatePageResponse
	BlockUnlink(*pb.RpcBlockUnlinkRequest) *pb.RpcBlockUnlinkResponse
	BlockClose(*pb.RpcBlockCloseRequest) *pb.RpcBlockCloseResponse
	BlockDownload(*pb.RpcBlockDownloadRequest) *pb.RpcBlockDownloadResponse
	BlockGetMarks(*pb.RpcBlockGetMarksRequest) *pb.RpcBlockGetMarksResponse
	BlockHistoryMove(*pb.RpcBlockHistoryMoveRequest) *pb.RpcBlockHistoryMoveResponse
	BlockSetFields(*pb.RpcBlockSetFieldsRequest) *pb.RpcBlockSetFieldsResponse
	BlockSetRestrictions(*pb.RpcBlockSetRestrictionsRequest) *pb.RpcBlockSetRestrictionsResponse
	BlockSetIsArchived(*pb.RpcBlockSetIsArchivedRequest) *pb.RpcBlockSetIsArchivedResponse
	BlockListMove(*pb.RpcBlockListMoveRequest) *pb.RpcBlockListMoveResponse
	BlockListSetFields(*pb.RpcBlockListSetFieldsRequest) *pb.RpcBlockListSetFieldsResponse
	BlockListSetTextStyle(*pb.RpcBlockListSetTextStyleRequest) *pb.RpcBlockListSetTextStyleResponse
	BlockListDuplicate(*pb.RpcBlockListDuplicateRequest) *pb.RpcBlockListDuplicateResponse
	BlockSetTextText(*pb.RpcBlockSetTextTextRequest) *pb.RpcBlockSetTextTextResponse
	BlockSetTextColor(*pb.RpcBlockSetTextColorRequest) *pb.RpcBlockSetTextColorResponse
	BlockListSetTextColor(*pb.RpcBlockListSetTextColorRequest) *pb.RpcBlockListSetTextColorResponse
	BlockSetTextBackgroundColor(*pb.RpcBlockSetTextBackgroundColorRequest) *pb.RpcBlockSetTextBackgroundColorResponse
	BlockListSetTextBackgroundColor(*pb.RpcBlockListSetTextBackgroundColorRequest) *pb.RpcBlockListSetTextBackgroundColorResponse
	BlockSetTextStyle(*pb.RpcBlockSetTextStyleRequest) *pb.RpcBlockSetTextStyleResponse
	BlockSetTextChecked(*pb.RpcBlockSetTextCheckedRequest) *pb.RpcBlockSetTextCheckedResponse
	BlockSplit(*pb.RpcBlockSplitRequest) *pb.RpcBlockSplitResponse
	BlockMerge(*pb.RpcBlockMergeRequest) *pb.RpcBlockMergeResponse
	BlockCopy(*pb.RpcBlockCopyRequest) *pb.RpcBlockCopyResponse
	BlockPaste(*pb.RpcBlockPasteRequest) *pb.RpcBlockPasteResponse
	BlockSetFileName(*pb.RpcBlockSetFileNameRequest) *pb.RpcBlockSetFileNameResponse
	BlockSetImageName(*pb.RpcBlockSetImageNameRequest) *pb.RpcBlockSetImageNameResponse
	BlockSetImageWidth(*pb.RpcBlockSetImageWidthRequest) *pb.RpcBlockSetImageWidthResponse
	BlockSetVideoName(*pb.RpcBlockSetVideoNameRequest) *pb.RpcBlockSetVideoNameResponse
	BlockSetVideoWidth(*pb.RpcBlockSetVideoWidthRequest) *pb.RpcBlockSetVideoWidthResponse
	BlockSetIconName(*pb.RpcBlockSetIconNameRequest) *pb.RpcBlockSetIconNameResponse
	BlockSetLinkTargetBlockId(*pb.RpcBlockSetLinkTargetBlockIdRequest) *pb.RpcBlockSetLinkTargetBlockIdResponse
	Ping(*pb.RpcPingRequest) *pb.RpcPingResponse
	// used only for lib-debug via grpc
	// Streams not supported ### ListenEvents(*pb.Empty)
}

func registerClientCommandsHandler(srv ClientCommandsHandler) {
	clientCommandsHandler = srv
}

func WalletCreate(b []byte) []byte {
	in := new(pb.RpcWalletCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcWalletCreateResponse{Error: &pb.RpcWalletCreateResponseError{Code: pb.RpcWalletCreateResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.WalletCreate(in).Marshal()
	return resp
}

func WalletRecover(b []byte) []byte {
	in := new(pb.RpcWalletRecoverRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcWalletRecoverResponse{Error: &pb.RpcWalletRecoverResponseError{Code: pb.RpcWalletRecoverResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.WalletRecover(in).Marshal()
	return resp
}

func AccountRecover(b []byte) []byte {
	in := new(pb.RpcAccountRecoverRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcAccountRecoverResponse{Error: &pb.RpcAccountRecoverResponseError{Code: pb.RpcAccountRecoverResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.AccountRecover(in).Marshal()
	return resp
}

func AccountCreate(b []byte) []byte {
	in := new(pb.RpcAccountCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcAccountCreateResponse{Error: &pb.RpcAccountCreateResponseError{Code: pb.RpcAccountCreateResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.AccountCreate(in).Marshal()
	return resp
}

func AccountSelect(b []byte) []byte {
	in := new(pb.RpcAccountSelectRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcAccountSelectResponse{Error: &pb.RpcAccountSelectResponseError{Code: pb.RpcAccountSelectResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.AccountSelect(in).Marshal()
	return resp
}

func ImageGetBlob(b []byte) []byte {
	in := new(pb.RpcIpfsImageGetBlobRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcIpfsImageGetBlobResponse{Error: &pb.RpcIpfsImageGetBlobResponseError{Code: pb.RpcIpfsImageGetBlobResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.ImageGetBlob(in).Marshal()
	return resp
}

func VersionGet(b []byte) []byte {
	in := new(pb.RpcVersionGetRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcVersionGetResponse{Error: &pb.RpcVersionGetResponseError{Code: pb.RpcVersionGetResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.VersionGet(in).Marshal()
	return resp
}

func LogSend(b []byte) []byte {
	in := new(pb.RpcLogSendRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcLogSendResponse{Error: &pb.RpcLogSendResponseError{Code: pb.RpcLogSendResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.LogSend(in).Marshal()
	return resp
}

func ConfigGet(b []byte) []byte {
	in := new(pb.RpcConfigGetRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcConfigGetResponse{Error: &pb.RpcConfigGetResponseError{Code: pb.RpcConfigGetResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.ConfigGet(in).Marshal()
	return resp
}

func ExternalDropFiles(b []byte) []byte {
	in := new(pb.RpcExternalDropFilesRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcExternalDropFilesResponse{Error: &pb.RpcExternalDropFilesResponseError{Code: pb.RpcExternalDropFilesResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.ExternalDropFiles(in).Marshal()
	return resp
}

func ExternalDropContent(b []byte) []byte {
	in := new(pb.RpcExternalDropContentRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcExternalDropContentResponse{Error: &pb.RpcExternalDropContentResponseError{Code: pb.RpcExternalDropContentResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.ExternalDropContent(in).Marshal()
	return resp
}

func BlockUpload(b []byte) []byte {
	in := new(pb.RpcBlockUploadRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockUploadResponse{Error: &pb.RpcBlockUploadResponseError{Code: pb.RpcBlockUploadResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockUpload(in).Marshal()
	return resp
}

func BlockReplace(b []byte) []byte {
	in := new(pb.RpcBlockReplaceRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockReplaceResponse{Error: &pb.RpcBlockReplaceResponseError{Code: pb.RpcBlockReplaceResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockReplace(in).Marshal()
	return resp
}

func BlockOpen(b []byte) []byte {
	in := new(pb.RpcBlockOpenRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockOpenResponse{Error: &pb.RpcBlockOpenResponseError{Code: pb.RpcBlockOpenResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockOpen(in).Marshal()
	return resp
}

func BlockCreate(b []byte) []byte {
	in := new(pb.RpcBlockCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockCreateResponse{Error: &pb.RpcBlockCreateResponseError{Code: pb.RpcBlockCreateResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockCreate(in).Marshal()
	return resp
}

func BlockCreatePage(b []byte) []byte {
	in := new(pb.RpcBlockCreatePageRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockCreatePageResponse{Error: &pb.RpcBlockCreatePageResponseError{Code: pb.RpcBlockCreatePageResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockCreatePage(in).Marshal()
	return resp
}

func BlockUnlink(b []byte) []byte {
	in := new(pb.RpcBlockUnlinkRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockUnlinkResponse{Error: &pb.RpcBlockUnlinkResponseError{Code: pb.RpcBlockUnlinkResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockUnlink(in).Marshal()
	return resp
}

func BlockClose(b []byte) []byte {
	in := new(pb.RpcBlockCloseRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockCloseResponse{Error: &pb.RpcBlockCloseResponseError{Code: pb.RpcBlockCloseResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockClose(in).Marshal()
	return resp
}

func BlockDownload(b []byte) []byte {
	in := new(pb.RpcBlockDownloadRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockDownloadResponse{Error: &pb.RpcBlockDownloadResponseError{Code: pb.RpcBlockDownloadResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockDownload(in).Marshal()
	return resp
}

func BlockGetMarks(b []byte) []byte {
	in := new(pb.RpcBlockGetMarksRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockGetMarksResponse{Error: &pb.RpcBlockGetMarksResponseError{Code: pb.RpcBlockGetMarksResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockGetMarks(in).Marshal()
	return resp
}

func BlockHistoryMove(b []byte) []byte {
	in := new(pb.RpcBlockHistoryMoveRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockHistoryMoveResponse{Error: &pb.RpcBlockHistoryMoveResponseError{Code: pb.RpcBlockHistoryMoveResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockHistoryMove(in).Marshal()
	return resp
}

func BlockSetFields(b []byte) []byte {
	in := new(pb.RpcBlockSetFieldsRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetFieldsResponse{Error: &pb.RpcBlockSetFieldsResponseError{Code: pb.RpcBlockSetFieldsResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetFields(in).Marshal()
	return resp
}

func BlockSetRestrictions(b []byte) []byte {
	in := new(pb.RpcBlockSetRestrictionsRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetRestrictionsResponse{Error: &pb.RpcBlockSetRestrictionsResponseError{Code: pb.RpcBlockSetRestrictionsResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetRestrictions(in).Marshal()
	return resp
}

func BlockSetIsArchived(b []byte) []byte {
	in := new(pb.RpcBlockSetIsArchivedRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetIsArchivedResponse{Error: &pb.RpcBlockSetIsArchivedResponseError{Code: pb.RpcBlockSetIsArchivedResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetIsArchived(in).Marshal()
	return resp
}

func BlockListMove(b []byte) []byte {
	in := new(pb.RpcBlockListMoveRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockListMoveResponse{Error: &pb.RpcBlockListMoveResponseError{Code: pb.RpcBlockListMoveResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockListMove(in).Marshal()
	return resp
}

func BlockListSetFields(b []byte) []byte {
	in := new(pb.RpcBlockListSetFieldsRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockListSetFieldsResponse{Error: &pb.RpcBlockListSetFieldsResponseError{Code: pb.RpcBlockListSetFieldsResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockListSetFields(in).Marshal()
	return resp
}

func BlockListSetTextStyle(b []byte) []byte {
	in := new(pb.RpcBlockListSetTextStyleRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockListSetTextStyleResponse{Error: &pb.RpcBlockListSetTextStyleResponseError{Code: pb.RpcBlockListSetTextStyleResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockListSetTextStyle(in).Marshal()
	return resp
}

func BlockListDuplicate(b []byte) []byte {
	in := new(pb.RpcBlockListDuplicateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockListDuplicateResponse{Error: &pb.RpcBlockListDuplicateResponseError{Code: pb.RpcBlockListDuplicateResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockListDuplicate(in).Marshal()
	return resp
}

func BlockSetTextText(b []byte) []byte {
	in := new(pb.RpcBlockSetTextTextRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetTextTextResponse{Error: &pb.RpcBlockSetTextTextResponseError{Code: pb.RpcBlockSetTextTextResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetTextText(in).Marshal()
	return resp
}

func BlockSetTextColor(b []byte) []byte {
	in := new(pb.RpcBlockSetTextColorRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetTextColorResponse{Error: &pb.RpcBlockSetTextColorResponseError{Code: pb.RpcBlockSetTextColorResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetTextColor(in).Marshal()
	return resp
}

func BlockListSetTextColor(b []byte) []byte {
	in := new(pb.RpcBlockListSetTextColorRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockListSetTextColorResponse{Error: &pb.RpcBlockListSetTextColorResponseError{Code: pb.RpcBlockListSetTextColorResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockListSetTextColor(in).Marshal()
	return resp
}

func BlockSetTextBackgroundColor(b []byte) []byte {
	in := new(pb.RpcBlockSetTextBackgroundColorRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetTextBackgroundColorResponse{Error: &pb.RpcBlockSetTextBackgroundColorResponseError{Code: pb.RpcBlockSetTextBackgroundColorResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetTextBackgroundColor(in).Marshal()
	return resp
}

func BlockListSetTextBackgroundColor(b []byte) []byte {
	in := new(pb.RpcBlockListSetTextBackgroundColorRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockListSetTextBackgroundColorResponse{Error: &pb.RpcBlockListSetTextBackgroundColorResponseError{Code: pb.RpcBlockListSetTextBackgroundColorResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockListSetTextBackgroundColor(in).Marshal()
	return resp
}

func BlockSetTextStyle(b []byte) []byte {
	in := new(pb.RpcBlockSetTextStyleRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetTextStyleResponse{Error: &pb.RpcBlockSetTextStyleResponseError{Code: pb.RpcBlockSetTextStyleResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetTextStyle(in).Marshal()
	return resp
}

func BlockSetTextChecked(b []byte) []byte {
	in := new(pb.RpcBlockSetTextCheckedRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetTextCheckedResponse{Error: &pb.RpcBlockSetTextCheckedResponseError{Code: pb.RpcBlockSetTextCheckedResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetTextChecked(in).Marshal()
	return resp
}

func BlockSplit(b []byte) []byte {
	in := new(pb.RpcBlockSplitRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSplitResponse{Error: &pb.RpcBlockSplitResponseError{Code: pb.RpcBlockSplitResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSplit(in).Marshal()
	return resp
}

func BlockMerge(b []byte) []byte {
	in := new(pb.RpcBlockMergeRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockMergeResponse{Error: &pb.RpcBlockMergeResponseError{Code: pb.RpcBlockMergeResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockMerge(in).Marshal()
	return resp
}

func BlockCopy(b []byte) []byte {
	in := new(pb.RpcBlockCopyRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockCopyResponse{Error: &pb.RpcBlockCopyResponseError{Code: pb.RpcBlockCopyResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockCopy(in).Marshal()
	return resp
}

func BlockPaste(b []byte) []byte {
	in := new(pb.RpcBlockPasteRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockPasteResponse{Error: &pb.RpcBlockPasteResponseError{Code: pb.RpcBlockPasteResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockPaste(in).Marshal()
	return resp
}

func BlockSetFileName(b []byte) []byte {
	in := new(pb.RpcBlockSetFileNameRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetFileNameResponse{Error: &pb.RpcBlockSetFileNameResponseError{Code: pb.RpcBlockSetFileNameResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetFileName(in).Marshal()
	return resp
}

func BlockSetImageName(b []byte) []byte {
	in := new(pb.RpcBlockSetImageNameRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetImageNameResponse{Error: &pb.RpcBlockSetImageNameResponseError{Code: pb.RpcBlockSetImageNameResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetImageName(in).Marshal()
	return resp
}

func BlockSetImageWidth(b []byte) []byte {
	in := new(pb.RpcBlockSetImageWidthRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetImageWidthResponse{Error: &pb.RpcBlockSetImageWidthResponseError{Code: pb.RpcBlockSetImageWidthResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetImageWidth(in).Marshal()
	return resp
}

func BlockSetVideoName(b []byte) []byte {
	in := new(pb.RpcBlockSetVideoNameRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetVideoNameResponse{Error: &pb.RpcBlockSetVideoNameResponseError{Code: pb.RpcBlockSetVideoNameResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetVideoName(in).Marshal()
	return resp
}

func BlockSetVideoWidth(b []byte) []byte {
	in := new(pb.RpcBlockSetVideoWidthRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetVideoWidthResponse{Error: &pb.RpcBlockSetVideoWidthResponseError{Code: pb.RpcBlockSetVideoWidthResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetVideoWidth(in).Marshal()
	return resp
}

func BlockSetIconName(b []byte) []byte {
	in := new(pb.RpcBlockSetIconNameRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetIconNameResponse{Error: &pb.RpcBlockSetIconNameResponseError{Code: pb.RpcBlockSetIconNameResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetIconName(in).Marshal()
	return resp
}

func BlockSetLinkTargetBlockId(b []byte) []byte {
	in := new(pb.RpcBlockSetLinkTargetBlockIdRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetLinkTargetBlockIdResponse{Error: &pb.RpcBlockSetLinkTargetBlockIdResponseError{Code: pb.RpcBlockSetLinkTargetBlockIdResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetLinkTargetBlockId(in).Marshal()
	return resp
}

func Ping(b []byte) []byte {
	in := new(pb.RpcPingRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcPingResponse{Error: &pb.RpcPingResponseError{Code: pb.RpcPingResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.Ping(in).Marshal()
	return resp
}

func CommandAsync(cmd string, data []byte, callback func(data []byte)) {
	go func() {
		var cd []byte
		switch cmd {
		case "WalletCreate":
			cd = WalletCreate(data)
		case "WalletRecover":
			cd = WalletRecover(data)
		case "AccountRecover":
			cd = AccountRecover(data)
		case "AccountCreate":
			cd = AccountCreate(data)
		case "AccountSelect":
			cd = AccountSelect(data)
		case "ImageGetBlob":
			cd = ImageGetBlob(data)
		case "VersionGet":
			cd = VersionGet(data)
		case "LogSend":
			cd = LogSend(data)
		case "ConfigGet":
			cd = ConfigGet(data)
		case "ExternalDropFiles":
			cd = ExternalDropFiles(data)
		case "ExternalDropContent":
			cd = ExternalDropContent(data)
		case "BlockUpload":
			cd = BlockUpload(data)
		case "BlockReplace":
			cd = BlockReplace(data)
		case "BlockOpen":
			cd = BlockOpen(data)
		case "BlockCreate":
			cd = BlockCreate(data)
		case "BlockCreatePage":
			cd = BlockCreatePage(data)
		case "BlockUnlink":
			cd = BlockUnlink(data)
		case "BlockClose":
			cd = BlockClose(data)
		case "BlockDownload":
			cd = BlockDownload(data)
		case "BlockGetMarks":
			cd = BlockGetMarks(data)
		case "BlockHistoryMove":
			cd = BlockHistoryMove(data)
		case "BlockSetFields":
			cd = BlockSetFields(data)
		case "BlockSetRestrictions":
			cd = BlockSetRestrictions(data)
		case "BlockSetIsArchived":
			cd = BlockSetIsArchived(data)
		case "BlockListMove":
			cd = BlockListMove(data)
		case "BlockListSetFields":
			cd = BlockListSetFields(data)
		case "BlockListSetTextStyle":
			cd = BlockListSetTextStyle(data)
		case "BlockListDuplicate":
			cd = BlockListDuplicate(data)
		case "BlockSetTextText":
			cd = BlockSetTextText(data)
		case "BlockSetTextColor":
			cd = BlockSetTextColor(data)
		case "BlockListSetTextColor":
			cd = BlockListSetTextColor(data)
		case "BlockSetTextBackgroundColor":
			cd = BlockSetTextBackgroundColor(data)
		case "BlockListSetTextBackgroundColor":
			cd = BlockListSetTextBackgroundColor(data)
		case "BlockSetTextStyle":
			cd = BlockSetTextStyle(data)
		case "BlockSetTextChecked":
			cd = BlockSetTextChecked(data)
		case "BlockSplit":
			cd = BlockSplit(data)
		case "BlockMerge":
			cd = BlockMerge(data)
		case "BlockCopy":
			cd = BlockCopy(data)
		case "BlockPaste":
			cd = BlockPaste(data)
		case "BlockSetFileName":
			cd = BlockSetFileName(data)
		case "BlockSetImageName":
			cd = BlockSetImageName(data)
		case "BlockSetImageWidth":
			cd = BlockSetImageWidth(data)
		case "BlockSetVideoName":
			cd = BlockSetVideoName(data)
		case "BlockSetVideoWidth":
			cd = BlockSetVideoWidth(data)
		case "BlockSetIconName":
			cd = BlockSetIconName(data)
		case "BlockSetLinkTargetBlockId":
			cd = BlockSetLinkTargetBlockId(data)
		case "Ping":
			cd = Ping(data)
		default:
			log.Errorf("unknown command type: %s\n", cmd)
		}
		if callback != nil {
			callback(cd)
		}
	}()
}

type MessageHandler interface {
	Handle(b []byte)
}

func CommandMobile(cmd string, data []byte, callback MessageHandler) {
	CommandAsync(cmd, data, callback.Handle)
}
