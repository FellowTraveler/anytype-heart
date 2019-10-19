// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package lib

import (
	fmt "fmt"
	go_anytype_middleware "github.com/anytypeio/go-anytype-middleware"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0x4a, 0xf3, 0x30,
	0x14, 0xc7, 0x37, 0xf6, 0xf1, 0x09, 0x61, 0xdb, 0x45, 0xf4, 0xaa, 0xda, 0xf8, 0x06, 0x45, 0xf4,
	0x09, 0x5c, 0x91, 0x21, 0x1b, 0x82, 0x13, 0x14, 0xbc, 0x4b, 0xe3, 0xb1, 0x04, 0xd2, 0xa4, 0x36,
	0xd9, 0x60, 0x6f, 0xe1, 0x63, 0x79, 0xb9, 0x4b, 0x2f, 0xa5, 0x05, 0x9f, 0x43, 0x30, 0xc9, 0xda,
	0xae, 0xf6, 0x2e, 0xfc, 0x7f, 0xe7, 0xff, 0xa3, 0x87, 0x53, 0x34, 0xd1, 0x50, 0x6c, 0x38, 0x83,
	0x28, 0x2f, 0x94, 0x51, 0xf8, 0x88, 0xca, 0xad, 0xd9, 0xe6, 0x10, 0x4c, 0x28, 0x63, 0x6a, 0x2d,
	0x8d, 0xcd, 0x03, 0xf4, 0xca, 0x05, 0xf8, 0x77, 0xc6, 0x35, 0xb3, 0xef, 0xcb, 0xef, 0x7f, 0x68,
	0x1a, 0x0b, 0x0e, 0xd2, 0xc4, 0x2a, 0xcb, 0xa8, 0x7c, 0xd1, 0x78, 0x81, 0xc6, 0x4f, 0x54, 0x08,
	0x30, 0x71, 0x01, 0xd4, 0x00, 0x3e, 0x8b, 0x9c, 0x33, 0x6a, 0xc6, 0x2b, 0x78, 0x5b, 0x83, 0x36,
	0x41, 0xd8, 0x43, 0x75, 0xae, 0xa4, 0x06, 0x7c, 0x87, 0x26, 0x36, 0x5f, 0x01, 0x53, 0x1b, 0x28,
	0xf0, 0xe1, 0xbc, 0xcb, 0xbd, 0x8e, 0xf4, 0x61, 0xe7, 0xbb, 0x47, 0xd3, 0x6b, 0xbb, 0x98, 0x17,
	0xd6, 0x8d, 0x36, 0xf0, 0xc6, 0xf3, 0x5e, 0x5e, 0x7f, 0xa2, 0x23, 0x6e, 0xe1, 0xf0, 0xb0, 0xd1,
	0xde, 0x98, 0xf4, 0xe1, 0x8e, 0xef, 0x01, 0x04, 0x30, 0xd3, 0xf5, 0xd9, 0xbc, 0xd7, 0xe7, 0xb1,
	0xf3, 0x2d, 0xd0, 0xf8, 0x36, 0xa3, 0x29, 0xcc, 0xc1, 0xcc, 0x84, 0x4a, 0x1a, 0xf7, 0x68, 0xc6,
	0xdd, 0x7b, 0xb4, 0xa9, 0x93, 0xdd, 0x20, 0x34, 0x07, 0xf3, 0x08, 0x85, 0xe6, 0x4a, 0xe2, 0x60,
	0x3f, 0x5c, 0x87, 0x5e, 0x74, 0xfa, 0x27, 0x73, 0x9a, 0x0b, 0x34, 0x5a, 0xaa, 0x14, 0x1f, 0xef,
	0x67, 0x96, 0x2a, 0xf5, 0xc5, 0x93, 0x76, 0x68, 0x1b, 0xb3, 0xf0, 0xa3, 0x24, 0xc3, 0x5d, 0x49,
	0x86, 0x5f, 0x25, 0x19, 0xbe, 0x57, 0x64, 0xb0, 0xab, 0xc8, 0xe0, 0xb3, 0x22, 0x83, 0xe7, 0x91,
	0xe0, 0x49, 0xf2, 0xff, 0xf7, 0x77, 0xbc, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x7a, 0x13,
	0x2b, 0xcf, 0x02, 0x00, 0x00,
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the gomobile package it is being compiled against.

// ClientCommandsServer is the server API for ClientCommands service.
var handler ClientCommandsServer

type ClientCommandsServer interface {
	WalletCreate(*go_anytype_middleware.WalletCreateRequest) *go_anytype_middleware.WalletCreateResponse
	WalletRecover(*go_anytype_middleware.WalletRecoverRequest) *go_anytype_middleware.WalletRecoverResponse
	AccountRecover(*go_anytype_middleware.AccountRecoverRequest) *go_anytype_middleware.AccountRecoverResponse
	AccountCreate(*go_anytype_middleware.AccountCreateRequest) *go_anytype_middleware.AccountCreateResponse
	AccountSelect(*go_anytype_middleware.AccountSelectRequest) *go_anytype_middleware.AccountSelectResponse
	ImageGetBlob(*go_anytype_middleware.ImageGetBlobRequest) *go_anytype_middleware.ImageGetBlobResponse
	GetVersion(*go_anytype_middleware.GetVersionRequest) *go_anytype_middleware.GetVersionResponse
	Log(*go_anytype_middleware.LogRequest) *go_anytype_middleware.LogResponse
}

func RegisterClientCommandsServer(srv ClientCommandsServer) {
	handler = srv
}

func WalletCreate(b []byte) []byte {
	in := new(go_anytype_middleware.WalletCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&go_anytype_middleware.WalletCreateResponse{Error: &go_anytype_middleware.WalletCreateResponse_Error{Code: go_anytype_middleware.WalletCreateResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := handler.WalletCreate(in).Marshal()
	return resp
}

func WalletRecover(b []byte) []byte {
	in := new(go_anytype_middleware.WalletRecoverRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&go_anytype_middleware.WalletRecoverResponse{Error: &go_anytype_middleware.WalletRecoverResponse_Error{Code: go_anytype_middleware.WalletRecoverResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := handler.WalletRecover(in).Marshal()
	return resp
}

func AccountRecover(b []byte) []byte {
	in := new(go_anytype_middleware.AccountRecoverRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&go_anytype_middleware.AccountRecoverResponse{Error: &go_anytype_middleware.AccountRecoverResponse_Error{Code: go_anytype_middleware.AccountRecoverResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := handler.AccountRecover(in).Marshal()
	return resp
}

func AccountCreate(b []byte) []byte {
	in := new(go_anytype_middleware.AccountCreateRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&go_anytype_middleware.AccountCreateResponse{Error: &go_anytype_middleware.AccountCreateResponse_Error{Code: go_anytype_middleware.AccountCreateResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := handler.AccountCreate(in).Marshal()
	return resp
}

func AccountSelect(b []byte) []byte {
	in := new(go_anytype_middleware.AccountSelectRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&go_anytype_middleware.AccountSelectResponse{Error: &go_anytype_middleware.AccountSelectResponse_Error{Code: go_anytype_middleware.AccountSelectResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := handler.AccountSelect(in).Marshal()
	return resp
}

func ImageGetBlob(b []byte) []byte {
	in := new(go_anytype_middleware.ImageGetBlobRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&go_anytype_middleware.ImageGetBlobResponse{Error: &go_anytype_middleware.ImageGetBlobResponse_Error{Code: go_anytype_middleware.ImageGetBlobResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := handler.ImageGetBlob(in).Marshal()
	return resp
}

func GetVersion(b []byte) []byte {
	in := new(go_anytype_middleware.GetVersionRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&go_anytype_middleware.GetVersionResponse{Error: &go_anytype_middleware.GetVersionResponse_Error{Code: go_anytype_middleware.GetVersionResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := handler.GetVersion(in).Marshal()
	return resp
}

func Log(b []byte) []byte {
	in := new(go_anytype_middleware.LogRequest)
	if err := in.Unmarshal(b); err != nil {
		resp, _ := (&go_anytype_middleware.LogResponse{Error: &go_anytype_middleware.LogResponse_Error{Code: go_anytype_middleware.LogResponse_Error_BAD_INPUT, Description: err.Error()}}).Marshal()
		return resp
	}
	resp, _ := handler.Log(in).Marshal()
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
		case "GetVersion":
			cd = GetVersion(data)
		case "Log":
			cd = Log(data)
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
