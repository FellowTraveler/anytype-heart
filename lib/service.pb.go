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
	// 1242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x98, 0xcd, 0x73, 0x1b, 0x35,
	0x18, 0xc6, 0xeb, 0x81, 0xa1, 0x83, 0x1a, 0x02, 0x15, 0xd0, 0xa1, 0x61, 0xea, 0x7e, 0x43, 0x92,
	0xc2, 0xa6, 0xb4, 0xcc, 0xc0, 0x85, 0x43, 0xed, 0x7c, 0x90, 0x99, 0xa4, 0x04, 0x3b, 0x6d, 0x66,
	0x3a, 0x30, 0xb0, 0xd9, 0x7d, 0xe3, 0x2c, 0x5e, 0x4b, 0xdb, 0x5d, 0xd9, 0xad, 0xaf, 0x9c, 0x38,
	0x72, 0xe0, 0x3f, 0xe2, 0xc2, 0xb1, 0x47, 0x8e, 0x4c, 0xf2, 0x8f, 0x30, 0xd2, 0xbe, 0xfb, 0x2d,
	0xc9, 0x9b, 0x43, 0xdb, 0xe9, 0x3e, 0xbf, 0xf7, 0x79, 0x24, 0xeb, 0x95, 0xb4, 0x36, 0xb9, 0x19,
	0x1d, 0x6f, 0x44, 0x31, 0x17, 0x3c, 0xd9, 0x48, 0x20, 0x9e, 0x05, 0x1e, 0x64, 0xff, 0x3a, 0xea,
	0x31, 0xbd, 0xec, 0xb2, 0xb9, 0x98, 0x47, 0xb0, 0xf2, 0x49, 0x41, 0x7a, 0x7c, 0x32, 0x71, 0x99,
	0x9f, 0xa4, 0xc8, 0xca, 0xb5, 0x42, 0x81, 0x19, 0x30, 0x81, 0xcf, 0x1f, 0xfd, 0xbd, 0x46, 0x96,
	0xfb, 0x61, 0x00, 0x4c, 0xf4, 0xb1, 0x80, 0x1e, 0x91, 0xa5, 0x23, 0x37, 0x0c, 0x41, 0xf4, 0x63,
	0x70, 0x05, 0xd0, 0x3b, 0x0e, 0xda, 0x3b, 0x83, 0xc8, 0x73, 0x52, 0xc9, 0x49, 0x35, 0x67, 0x00,
	0x2f, 0xa7, 0x90, 0x88, 0x95, 0xbb, 0x56, 0x26, 0x89, 0x38, 0x4b, 0x80, 0xbe, 0x20, 0xef, 0xa5,
	0xca, 0x00, 0x3c, 0x3e, 0x83, 0x98, 0x6a, 0xab, 0x50, 0xcc, 0xad, 0xef, 0xd9, 0x21, 0xf4, 0xfe,
	0x99, 0x2c, 0x3f, 0xf1, 0x3c, 0x3e, 0x65, 0xb9, 0x79, 0xb5, 0x0e, 0xc5, 0x86, 0xfb, 0xfd, 0x05,
	0x54, 0x31, 0x74, 0xd4, 0xf0, 0x43, 0xb9, 0xab, 0xad, 0xab, 0x7d, 0x2a, 0xf7, 0xec, 0x50, 0xc3,
	0x7b, 0x08, 0x21, 0x78, 0xc2, 0xe0, 0x9d, 0x8a, 0x0b, 0xbc, 0x73, 0x08, 0xbd, 0x0f, 0xc9, 0x95,
	0xcc, 0x5b, 0xf0, 0x88, 0xde, 0xd6, 0x17, 0x09, 0x1e, 0xe5, 0xbe, 0x77, 0x6c, 0x08, 0xba, 0x7a,
	0x64, 0x69, 0x77, 0xe2, 0x8e, 0x60, 0x07, 0x44, 0x2f, 0xe4, 0xc7, 0x74, 0xb5, 0x52, 0xb3, 0x1b,
	0x9d, 0x24, 0x8e, 0xd2, 0x9d, 0x1d, 0x10, 0x8e, 0x24, 0x72, 0xf7, 0xb5, 0x16, 0x24, 0x86, 0xfc,
	0x48, 0xc8, 0x73, 0x88, 0x93, 0x80, 0xb3, 0x1d, 0x10, 0xf4, 0x56, 0xa5, 0x10, 0x05, 0x55, 0x95,
	0x59, 0xdf, 0xb6, 0x10, 0x68, 0xf9, 0x3d, 0xb9, 0xbc, 0xc7, 0x47, 0x43, 0x60, 0x3e, 0xbd, 0x51,
	0xa1, 0xf7, 0xf8, 0xc8, 0x91, 0x8f, 0x73, 0xb3, 0xae, 0x49, 0x46, 0xa7, 0xa7, 0xe4, 0xdd, 0x3e,
	0x67, 0x27, 0xc1, 0x48, 0x8e, 0xed, 0x66, 0x05, 0x4e, 0x9f, 0x57, 0x86, 0x76, 0xcb, 0x0c, 0xa0,
	0xdf, 0x09, 0xb9, 0xba, 0xf5, 0x5a, 0x40, 0xcc, 0xdc, 0x70, 0x33, 0xe6, 0xd1, 0x76, 0x10, 0x42,
	0x42, 0x3f, 0xaf, 0x94, 0x95, 0x75, 0x47, 0x01, 0xb9, 0xff, 0xea, 0x62, 0x10, 0x73, 0x42, 0xf2,
	0x61, 0x59, 0xee, 0x73, 0x26, 0x80, 0x09, 0xba, 0x66, 0x36, 0x40, 0x24, 0xcf, 0x5a, 0x6f, 0x83,
	0x62, 0xda, 0x80, 0x5c, 0xd9, 0x0b, 0xd8, 0xf8, 0x20, 0x86, 0x59, 0x00, 0xaf, 0x6a, 0x6b, 0x58,
	0x52, 0x0c, 0x6b, 0x58, 0x25, 0x8a, 0x8e, 0xee, 0x85, 0xdc, 0x1b, 0x3f, 0x8b, 0x42, 0xee, 0xfa,
	0xb5, 0x8e, 0x56, 0x8a, 0x93, 0x4a, 0x86, 0x8e, 0xae, 0x21, 0xe8, 0x7a, 0x44, 0x96, 0x94, 0x30,
	0x80, 0x28, 0x74, 0xbd, 0xfa, 0x99, 0x97, 0xd6, 0xa0, 0x66, 0x38, 0xf3, 0xea, 0x4c, 0xd1, 0x28,
	0x4a, 0xf9, 0x21, 0x02, 0x56, 0x6b, 0x94, 0xb4, 0x42, 0x0a, 0x86, 0x46, 0xa9, 0x00, 0xe8, 0xc7,
	0xc9, 0x47, 0xb9, 0x5f, 0x2f, 0x06, 0xd7, 0xf7, 0xe2, 0xe9, 0xe4, 0x38, 0xa1, 0xeb, 0x86, 0xca,
	0x12, 0x93, 0xa7, 0x3c, 0x68, 0xc5, 0x16, 0x1d, 0xa3, 0x88, 0xfe, 0x54, 0x94, 0xf3, 0xd6, 0x34,
	0x1e, 0x55, 0xc4, 0xd0, 0x31, 0x06, 0xb4, 0xb6, 0xba, 0x78, 0xca, 0xea, 0x56, 0xb7, 0x76, 0xc6,
	0xde, 0xb1, 0x21, 0xe8, 0xfa, 0x2b, 0x79, 0xbf, 0xe4, 0x7a, 0xe0, 0x8e, 0x80, 0xde, 0x37, 0x96,
	0x49, 0x39, 0x77, 0xff, 0x6c, 0x11, 0x56, 0xef, 0x4a, 0x16, 0x06, 0x6c, 0xac, 0xef, 0x4a, 0x25,
	0xd9, 0xbb, 0x32, 0x43, 0x8a, 0x23, 0x30, 0x1d, 0x77, 0xc8, 0x13, 0xa0, 0xba, 0xe6, 0x50, 0x8a,
	0x61, 0xfb, 0x54, 0x89, 0xe2, 0xb2, 0x51, 0xcf, 0x37, 0xf9, 0x2b, 0xa6, 0x36, 0x90, 0xae, 0x8b,
	0x33, 0xd1, 0x70, 0xd9, 0x34, 0x20, 0xf4, 0xfe, 0x09, 0xbd, 0x77, 0x40, 0xec, 0xbb, 0xf1, 0x38,
	0xa1, 0xba, 0x32, 0x79, 0xec, 0x29, 0xd5, 0x70, 0x05, 0x37, 0xa9, 0xda, 0x4e, 0x7a, 0xc6, 0x7c,
	0xae, 0xdd, 0x49, 0x52, 0xb0, 0xee, 0x24, 0x04, 0x6a, 0x7e, 0x03, 0x30, 0xf8, 0x49, 0xc1, 0xea,
	0x87, 0x00, 0xfa, 0xfd, 0x42, 0x96, 0xd5, 0xe3, 0x21, 0x88, 0xed, 0x00, 0x42, 0x3f, 0xd1, 0xf6,
	0xd8, 0x10, 0x84, 0x93, 0xca, 0xd6, 0x1e, 0xab, 0x60, 0x18, 0xf0, 0x12, 0xb7, 0xfe, 0x50, 0xbe,
	0x40, 0x25, 0x22, 0x0e, 0x3c, 0x11, 0x70, 0x96, 0xd0, 0x07, 0x86, 0xfa, 0x32, 0x94, 0x87, 0x7d,
	0xd1, 0x0e, 0xc6, 0xc8, 0x39, 0xb9, 0x96, 0x45, 0xca, 0x7e, 0xdf, 0x4d, 0x9e, 0xc4, 0xde, 0x69,
	0x30, 0x03, 0x9f, 0x3a, 0x06, 0x1f, 0xb5, 0x2d, 0x0a, 0x2e, 0xcf, 0xdd, 0x68, 0xcd, 0xd7, 0x1a,
	0x75, 0x2f, 0x48, 0xc4, 0x3e, 0x9f, 0x81, 0xae, 0x51, 0xa5, 0xe6, 0x48, 0xd1, 0xd6, 0xa8, 0x65,
	0x08, 0xbd, 0xc7, 0x84, 0xe6, 0x52, 0xb1, 0x5c, 0x6b, 0x86, 0x5a, 0xcd, 0x92, 0xad, 0xb7, 0x41,
	0x31, 0x6c, 0x46, 0x3e, 0x2e, 0x87, 0x1d, 0xc2, 0x6b, 0x31, 0x14, 0xf3, 0x10, 0xe8, 0x97, 0x16,
	0x13, 0x49, 0x39, 0x0a, 0xcb, 0x33, 0x9d, 0xb6, 0x38, 0xe6, 0x06, 0xa5, 0x49, 0x6e, 0x4e, 0xa3,
	0x30, 0xf0, 0xe4, 0x89, 0xba, 0x6a, 0x70, 0xc9, 0x09, 0xc3, 0xab, 0x9a, 0x9e, 0xc4, 0xa8, 0x11,
	0xf9, 0x20, 0x6b, 0x13, 0x39, 0x12, 0xf9, 0x47, 0x17, 0x54, 0x0c, 0x55, 0xfd, 0x65, 0x09, 0x6a,
	0x90, 0x18, 0xf4, 0x1b, 0xb9, 0x5a, 0x0e, 0xea, 0xf3, 0x90, 0xc7, 0xd4, 0x5a, 0xaf, 0x10, 0xeb,
	0x55, 0xd4, 0x40, 0xcd, 0xeb, 0x96, 0xe6, 0x2d, 0x5c, 0xb7, 0x6a, 0xa6, 0xd3, 0x16, 0xc7, 0xdc,
	0x3f, 0x3a, 0xe4, 0xd3, 0xf2, 0x24, 0x7b, 0xae, 0x37, 0x1e, 0xc5, 0x7c, 0xca, 0xfc, 0x34, 0xfe,
	0xb1, 0x6d, 0x0e, 0x35, 0x38, 0x1f, 0xc4, 0xd7, 0x17, 0x2b, 0xc2, 0xa1, 0xfc, 0xd5, 0x21, 0x37,
	0xeb, 0x9f, 0x41, 0x7d, 0x38, 0xdf, 0x2c, 0x9a, 0x9e, 0x69, 0x48, 0xdf, 0x5e, 0xbc, 0x50, 0xdf,
	0x05, 0xe9, 0x6e, 0xb2, 0x76, 0x41, 0x75, 0x27, 0xad, 0xb7, 0x41, 0x31, 0x2b, 0xc2, 0xd7, 0x9f,
	0xac, 0x03, 0x4e, 0xc1, 0x1b, 0x83, 0x6f, 0x3c, 0x73, 0xd3, 0x05, 0x4d, 0xa1, 0x85, 0x67, 0x6e,
	0x0d, 0xae, 0x5d, 0xfa, 0xc3, 0x28, 0x0c, 0x84, 0xf6, 0xd2, 0x57, 0x8a, 0xf5, 0xd2, 0xcf, 0x88,
	0x9a, 0xe5, 0x3e, 0xc4, 0x23, 0xfd, 0x7b, 0x84, 0x52, 0xac, 0x96, 0x19, 0x51, 0xbb, 0x3d, 0xfb,
	0x3c, 0x9a, 0x6b, 0x6f, 0x4f, 0x29, 0x58, 0x6f, 0x4f, 0x04, 0x6a, 0x43, 0x3c, 0x70, 0x13, 0xa1,
	0x1f, 0xa2, 0x52, 0xac, 0x43, 0xcc, 0x88, 0xe6, 0xa9, 0x24, 0xbf, 0x05, 0x3d, 0x75, 0x27, 0x60,
	0x3c, 0x95, 0x24, 0xe0, 0x48, 0x62, 0xe1, 0xa9, 0x54, 0x26, 0x9b, 0xfd, 0xa8, 0xbe, 0xcc, 0xaa,
	0x24, 0x53, 0x7d, 0xfa, 0x75, 0xb7, 0x12, 0xb5, 0xde, 0x06, 0xc5, 0xac, 0x09, 0x9e, 0xea, 0x59,
	0xd6, 0x51, 0xe0, 0x8b, 0x53, 0x6a, 0x77, 0x50, 0x8c, 0xf5, 0xed, 0xbf, 0xc9, 0x36, 0xa7, 0xf6,
	0x3c, 0xf0, 0x81, 0x5b, 0xa7, 0xa6, 0x88, 0x76, 0x53, 0xab, 0xa0, 0xcd, 0xa9, 0x29, 0xd9, 0x3e,
	0xb5, 0xd4, 0xa1, 0xdd, 0xd4, 0xaa, 0x6c, 0xb3, 0x3d, 0x76, 0x3d, 0xce, 0xac, 0xed, 0x21, 0x81,
	0x76, 0xed, 0x51, 0x26, 0x31, 0xe8, 0xf7, 0x0e, 0xb9, 0x9e, 0x25, 0xc9, 0xaf, 0xb4, 0x87, 0x6e,
	0x3c, 0x52, 0x3f, 0x9c, 0x78, 0xe3, 0x5d, 0x9f, 0x7e, 0x65, 0x30, 0x92, 0xa4, 0x53, 0x41, 0xf3,
	0xec, 0x47, 0x17, 0x29, 0xa9, 0xbd, 0xf2, 0xf4, 0x38, 0x1f, 0x4f, 0xdc, 0x78, 0xbc, 0x0d, 0xc2,
	0x3b, 0xd5, 0xae, 0x64, 0x46, 0x38, 0x0a, 0xb1, 0xae, 0x64, 0x03, 0xc5, 0xb0, 0xef, 0xc8, 0xdb,
	0x07, 0x01, 0x1b, 0xd1, 0xeb, 0x95, 0x1a, 0xf9, 0x28, 0xb7, 0x5b, 0xd1, 0x49, 0x58, 0xfe, 0x90,
	0x2c, 0xc9, 0x9b, 0x00, 0xd8, 0x96, 0xfa, 0xa5, 0x92, 0x2e, 0xe7, 0xec, 0xd6, 0x24, 0x12, 0xf3,
	0x95, 0xd2, 0xff, 0x25, 0xf0, 0xb0, 0xd3, 0xbb, 0xf1, 0xcf, 0x59, 0xb7, 0xf3, 0xe6, 0xac, 0xdb,
	0xf9, 0xef, 0xac, 0xdb, 0xf9, 0xf3, 0xbc, 0x7b, 0xe9, 0xcd, 0x79, 0xf7, 0xd2, 0xbf, 0xe7, 0xdd,
	0x4b, 0x2f, 0xde, 0x0a, 0x83, 0xe3, 0xe3, 0x77, 0xd4, 0x6f, 0x9d, 0x8f, 0xff, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x5a, 0x80, 0x83, 0xba, 0x49, 0x15, 0x00, 0x00,
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
	AccountStop(*pb.RpcAccountStopRequest) *pb.RpcAccountStopResponse
	ImageGetBlob(*pb.RpcIpfsImageGetBlobRequest) *pb.RpcIpfsImageGetBlobResponse
	VersionGet(*pb.RpcVersionGetRequest) *pb.RpcVersionGetResponse
	LogSend(*pb.RpcLogSendRequest) *pb.RpcLogSendResponse
	ConfigGet(*pb.RpcConfigGetRequest) *pb.RpcConfigGetResponse
	ExternalDropFiles(*pb.RpcExternalDropFilesRequest) *pb.RpcExternalDropFilesResponse
	ExternalDropContent(*pb.RpcExternalDropContentRequest) *pb.RpcExternalDropContentResponse
	LinkPreview(*pb.RpcLinkPreviewRequest) *pb.RpcLinkPreviewResponse
	BlockUpload(*pb.RpcBlockUploadRequest) *pb.RpcBlockUploadResponse
	BlockReplace(*pb.RpcBlockReplaceRequest) *pb.RpcBlockReplaceResponse
	BlockOpen(*pb.RpcBlockOpenRequest) *pb.RpcBlockOpenResponse
	BlockOpenBreadcrumbs(*pb.RpcBlockOpenBreadcrumbsRequest) *pb.RpcBlockOpenBreadcrumbsResponse
	BlockCutBreadcrumbs(*pb.RpcBlockCutBreadcrumbsRequest) *pb.RpcBlockCutBreadcrumbsResponse
	BlockCreate(*pb.RpcBlockCreateRequest) *pb.RpcBlockCreateResponse
	BlockCreatePage(*pb.RpcBlockCreatePageRequest) *pb.RpcBlockCreatePageResponse
	BlockUnlink(*pb.RpcBlockUnlinkRequest) *pb.RpcBlockUnlinkResponse
	BlockClose(*pb.RpcBlockCloseRequest) *pb.RpcBlockCloseResponse
	BlockDownload(*pb.RpcBlockDownloadRequest) *pb.RpcBlockDownloadResponse
	BlockGetMarks(*pb.RpcBlockGetMarksRequest) *pb.RpcBlockGetMarksResponse
	BlockUndo(*pb.RpcBlockUndoRequest) *pb.RpcBlockUndoResponse
	BlockRedo(*pb.RpcBlockRedoRequest) *pb.RpcBlockRedoResponse
	BlockSetFields(*pb.RpcBlockSetFieldsRequest) *pb.RpcBlockSetFieldsResponse
	BlockSetRestrictions(*pb.RpcBlockSetRestrictionsRequest) *pb.RpcBlockSetRestrictionsResponse
	BlockSetPageIsArchived(*pb.RpcBlockSetPageIsArchivedRequest) *pb.RpcBlockSetPageIsArchivedResponse
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
	BlockBookmarkFetch(*pb.RpcBlockBookmarkFetchRequest) *pb.RpcBlockBookmarkFetchResponse
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

func AccountStop(b []byte) []byte {
	in := new(pb.RpcAccountStopRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcAccountStopResponse{Error: &pb.RpcAccountStopResponseError{Code: pb.RpcAccountStopResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.AccountStop(in).Marshal()
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

func LinkPreview(b []byte) []byte {
	in := new(pb.RpcLinkPreviewRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcLinkPreviewResponse{Error: &pb.RpcLinkPreviewResponseError{Code: pb.RpcLinkPreviewResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.LinkPreview(in).Marshal()
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

func BlockOpenBreadcrumbs(b []byte) []byte {
	in := new(pb.RpcBlockOpenBreadcrumbsRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockOpenBreadcrumbsResponse{Error: &pb.RpcBlockOpenBreadcrumbsResponseError{Code: pb.RpcBlockOpenBreadcrumbsResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockOpenBreadcrumbs(in).Marshal()
	return resp
}

func BlockCutBreadcrumbs(b []byte) []byte {
	in := new(pb.RpcBlockCutBreadcrumbsRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockCutBreadcrumbsResponse{Error: &pb.RpcBlockCutBreadcrumbsResponseError{Code: pb.RpcBlockCutBreadcrumbsResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockCutBreadcrumbs(in).Marshal()
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

func BlockUndo(b []byte) []byte {
	in := new(pb.RpcBlockUndoRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockUndoResponse{Error: &pb.RpcBlockUndoResponseError{Code: pb.RpcBlockUndoResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockUndo(in).Marshal()
	return resp
}

func BlockRedo(b []byte) []byte {
	in := new(pb.RpcBlockRedoRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockRedoResponse{Error: &pb.RpcBlockRedoResponseError{Code: pb.RpcBlockRedoResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockRedo(in).Marshal()
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

func BlockSetPageIsArchived(b []byte) []byte {
	in := new(pb.RpcBlockSetPageIsArchivedRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockSetPageIsArchivedResponse{Error: &pb.RpcBlockSetPageIsArchivedResponseError{Code: pb.RpcBlockSetPageIsArchivedResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockSetPageIsArchived(in).Marshal()
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

func BlockBookmarkFetch(b []byte) []byte {
	in := new(pb.RpcBlockBookmarkFetchRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&pb.RpcBlockBookmarkFetchResponse{Error: &pb.RpcBlockBookmarkFetchResponseError{Code: pb.RpcBlockBookmarkFetchResponseError_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := clientCommandsHandler.BlockBookmarkFetch(in).Marshal()
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
		case "AccountStop":
			cd = AccountStop(data)
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
		case "LinkPreview":
			cd = LinkPreview(data)
		case "BlockUpload":
			cd = BlockUpload(data)
		case "BlockReplace":
			cd = BlockReplace(data)
		case "BlockOpen":
			cd = BlockOpen(data)
		case "BlockOpenBreadcrumbs":
			cd = BlockOpenBreadcrumbs(data)
		case "BlockCutBreadcrumbs":
			cd = BlockCutBreadcrumbs(data)
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
		case "BlockUndo":
			cd = BlockUndo(data)
		case "BlockRedo":
			cd = BlockRedo(data)
		case "BlockSetFields":
			cd = BlockSetFields(data)
		case "BlockSetRestrictions":
			cd = BlockSetRestrictions(data)
		case "BlockSetPageIsArchived":
			cd = BlockSetPageIsArchived(data)
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
		case "BlockBookmarkFetch":
			cd = BlockBookmarkFetch(data)
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
