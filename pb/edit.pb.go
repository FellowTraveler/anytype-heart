// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: edit.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

//*
//  Middleware to front end event message, that will be sent in this scenario:
// Precondition: user A opened a block
// 1. User B opens the same block
// 2. User A receives a message about p.1
type UserBlockJoin struct {
	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *UserBlockJoin) Reset()         { *m = UserBlockJoin{} }
func (m *UserBlockJoin) String() string { return proto.CompactTextString(m) }
func (*UserBlockJoin) ProtoMessage()    {}
func (*UserBlockJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5013c7d48f38d97, []int{0}
}
func (m *UserBlockJoin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBlockJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBlockJoin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBlockJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBlockJoin.Merge(m, src)
}
func (m *UserBlockJoin) XXX_Size() int {
	return m.Size()
}
func (m *UserBlockJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBlockJoin.DiscardUnknown(m)
}

var xxx_messageInfo_UserBlockJoin proto.InternalMessageInfo

func (m *UserBlockJoin) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

//*
//  Middleware to front end event message, that will be sent in this scenario:
// Precondition: user A and user B opened the same block
// 1. User B closes the block
// 2. User A receives a message about p.1
type UserBlockLeft struct {
	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *UserBlockLeft) Reset()         { *m = UserBlockLeft{} }
func (m *UserBlockLeft) String() string { return proto.CompactTextString(m) }
func (*UserBlockLeft) ProtoMessage()    {}
func (*UserBlockLeft) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5013c7d48f38d97, []int{1}
}
func (m *UserBlockLeft) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBlockLeft) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBlockLeft.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBlockLeft) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBlockLeft.Merge(m, src)
}
func (m *UserBlockLeft) XXX_Size() int {
	return m.Size()
}
func (m *UserBlockLeft) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBlockLeft.DiscardUnknown(m)
}

var xxx_messageInfo_UserBlockLeft proto.InternalMessageInfo

func (m *UserBlockLeft) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

//*
// Middleware to front end event message, that will be sent in this scenario:
// Precondition: user A and user B opened the same block
// 1. User B sets cursor or selects a text region into a text block
// 2. User A receives a message about p.1
type UserBlockTextRange struct {
	Account *Account     `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	BlockId string       `protobuf:"bytes,2,opt,name=blockId,proto3" json:"blockId,omitempty"`
	Range   *Model_Range `protobuf:"bytes,3,opt,name=range,proto3" json:"range,omitempty"`
}

func (m *UserBlockTextRange) Reset()         { *m = UserBlockTextRange{} }
func (m *UserBlockTextRange) String() string { return proto.CompactTextString(m) }
func (*UserBlockTextRange) ProtoMessage()    {}
func (*UserBlockTextRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5013c7d48f38d97, []int{2}
}
func (m *UserBlockTextRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBlockTextRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBlockTextRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBlockTextRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBlockTextRange.Merge(m, src)
}
func (m *UserBlockTextRange) XXX_Size() int {
	return m.Size()
}
func (m *UserBlockTextRange) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBlockTextRange.DiscardUnknown(m)
}

var xxx_messageInfo_UserBlockTextRange proto.InternalMessageInfo

func (m *UserBlockTextRange) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *UserBlockTextRange) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

func (m *UserBlockTextRange) GetRange() *Model_Range {
	if m != nil {
		return m.Range
	}
	return nil
}

//*
// Middleware to front end event message, that will be sent in this scenario:
// Precondition: user A and user B opened the same block
// 1. User B selects some inner blocks
// 2. User A receives a message about p.1
type UserBlockSelectRange struct {
	Account       *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	BlockIdsArray []string `protobuf:"bytes,2,rep,name=blockIdsArray,proto3" json:"blockIdsArray,omitempty"`
}

func (m *UserBlockSelectRange) Reset()         { *m = UserBlockSelectRange{} }
func (m *UserBlockSelectRange) String() string { return proto.CompactTextString(m) }
func (*UserBlockSelectRange) ProtoMessage()    {}
func (*UserBlockSelectRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5013c7d48f38d97, []int{3}
}
func (m *UserBlockSelectRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBlockSelectRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBlockSelectRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBlockSelectRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBlockSelectRange.Merge(m, src)
}
func (m *UserBlockSelectRange) XXX_Size() int {
	return m.Size()
}
func (m *UserBlockSelectRange) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBlockSelectRange.DiscardUnknown(m)
}

var xxx_messageInfo_UserBlockSelectRange proto.InternalMessageInfo

func (m *UserBlockSelectRange) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *UserBlockSelectRange) GetBlockIdsArray() []string {
	if m != nil {
		return m.BlockIdsArray
	}
	return nil
}

//*
// Middleware to front end event message, that will be sent on one of this scenarios:
// Precondition: user A opened a block
// 1. User A drops a set of files/pictures/videos
// 2. User A creates a MediaBlock and drops a single media, that corresponds to its type.
type FilesUpload struct {
	FilePath []string `protobuf:"bytes,1,rep,name=filePath,proto3" json:"filePath,omitempty"`
	BlockId  string   `protobuf:"bytes,2,opt,name=blockId,proto3" json:"blockId,omitempty"`
}

func (m *FilesUpload) Reset()         { *m = FilesUpload{} }
func (m *FilesUpload) String() string { return proto.CompactTextString(m) }
func (*FilesUpload) ProtoMessage()    {}
func (*FilesUpload) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5013c7d48f38d97, []int{4}
}
func (m *FilesUpload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilesUpload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilesUpload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilesUpload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilesUpload.Merge(m, src)
}
func (m *FilesUpload) XXX_Size() int {
	return m.Size()
}
func (m *FilesUpload) XXX_DiscardUnknown() {
	xxx_messageInfo_FilesUpload.DiscardUnknown(m)
}

var xxx_messageInfo_FilesUpload proto.InternalMessageInfo

func (m *FilesUpload) GetFilePath() []string {
	if m != nil {
		return m.FilePath
	}
	return nil
}

func (m *FilesUpload) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

func init() {
	proto.RegisterType((*UserBlockJoin)(nil), "anytype.UserBlockJoin")
	proto.RegisterType((*UserBlockLeft)(nil), "anytype.UserBlockLeft")
	proto.RegisterType((*UserBlockTextRange)(nil), "anytype.UserBlockTextRange")
	proto.RegisterType((*UserBlockSelectRange)(nil), "anytype.UserBlockSelectRange")
	proto.RegisterType((*FilesUpload)(nil), "anytype.FilesUpload")
}

func init() { proto.RegisterFile("edit.proto", fileDescriptor_f5013c7d48f38d97) }

var fileDescriptor_f5013c7d48f38d97 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xc9, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xcc, 0xab, 0x2c, 0xa9, 0x2c, 0x48, 0x95,
	0xe2, 0xc9, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0x86, 0x08, 0x4b, 0xf1, 0x26, 0x26, 0x27, 0xe7, 0x97,
	0xe6, 0x41, 0x55, 0x29, 0x59, 0x73, 0xf1, 0x86, 0x16, 0xa7, 0x16, 0x39, 0xe5, 0xe4, 0x27, 0x67,
	0x7b, 0xe5, 0x67, 0xe6, 0x09, 0x69, 0x71, 0xb1, 0x43, 0x55, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x1b, 0x09, 0xe8, 0x41, 0x0d, 0xd2, 0x73, 0x84, 0x88, 0x07, 0xc1, 0x14, 0xa0, 0x68, 0xf6, 0x49,
	0x4d, 0x2b, 0x21, 0x49, 0x73, 0x1b, 0x23, 0x97, 0x10, 0x5c, 0x77, 0x48, 0x6a, 0x45, 0x49, 0x50,
	0x62, 0x5e, 0x7a, 0x2a, 0x29, 0x46, 0x08, 0x49, 0x70, 0xb1, 0x27, 0x81, 0x74, 0x7b, 0xa6, 0x48,
	0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x5a, 0x5c, 0xac, 0x45, 0x20, 0xe3, 0x24,
	0x98, 0xc1, 0x66, 0x88, 0xc0, 0xcd, 0xf0, 0x05, 0x85, 0x85, 0x1e, 0xd8, 0xaa, 0x20, 0x88, 0x12,
	0xa5, 0x0c, 0x2e, 0x11, 0xb8, 0x3b, 0x82, 0x53, 0x73, 0x52, 0x93, 0xc9, 0x70, 0x89, 0x0a, 0x17,
	0x2f, 0xd4, 0xea, 0x62, 0xc7, 0xa2, 0xa2, 0xc4, 0x4a, 0x09, 0x26, 0x05, 0x66, 0x0d, 0xce, 0x20,
	0x54, 0x41, 0x25, 0x67, 0x2e, 0x6e, 0xb7, 0xcc, 0x9c, 0xd4, 0xe2, 0xd0, 0x82, 0x9c, 0xfc, 0xc4,
	0x14, 0x21, 0x29, 0x2e, 0x8e, 0xb4, 0xcc, 0x9c, 0xd4, 0x80, 0xc4, 0x92, 0x0c, 0x09, 0x46, 0xb0,
	0x7a, 0x38, 0x1f, 0xb7, 0xd7, 0x9c, 0x64, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e,
	0x21, 0x8a, 0xa9, 0x20, 0x29, 0x89, 0x0d, 0x1c, 0xad, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4f, 0xb6, 0x5b, 0x3d, 0x0a, 0x02, 0x00, 0x00,
}

func (m *UserBlockJoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBlockJoin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBlockJoin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEdit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserBlockLeft) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBlockLeft) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBlockLeft) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEdit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserBlockTextRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBlockTextRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBlockTextRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Range != nil {
		{
			size, err := m.Range.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEdit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BlockId) > 0 {
		i -= len(m.BlockId)
		copy(dAtA[i:], m.BlockId)
		i = encodeVarintEdit(dAtA, i, uint64(len(m.BlockId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEdit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserBlockSelectRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBlockSelectRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBlockSelectRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockIdsArray) > 0 {
		for iNdEx := len(m.BlockIdsArray) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BlockIdsArray[iNdEx])
			copy(dAtA[i:], m.BlockIdsArray[iNdEx])
			i = encodeVarintEdit(dAtA, i, uint64(len(m.BlockIdsArray[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEdit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FilesUpload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilesUpload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilesUpload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockId) > 0 {
		i -= len(m.BlockId)
		copy(dAtA[i:], m.BlockId)
		i = encodeVarintEdit(dAtA, i, uint64(len(m.BlockId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FilePath) > 0 {
		for iNdEx := len(m.FilePath) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FilePath[iNdEx])
			copy(dAtA[i:], m.FilePath[iNdEx])
			i = encodeVarintEdit(dAtA, i, uint64(len(m.FilePath[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintEdit(dAtA []byte, offset int, v uint64) int {
	offset -= sovEdit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserBlockJoin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovEdit(uint64(l))
	}
	return n
}

func (m *UserBlockLeft) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovEdit(uint64(l))
	}
	return n
}

func (m *UserBlockTextRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovEdit(uint64(l))
	}
	l = len(m.BlockId)
	if l > 0 {
		n += 1 + l + sovEdit(uint64(l))
	}
	if m.Range != nil {
		l = m.Range.Size()
		n += 1 + l + sovEdit(uint64(l))
	}
	return n
}

func (m *UserBlockSelectRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovEdit(uint64(l))
	}
	if len(m.BlockIdsArray) > 0 {
		for _, s := range m.BlockIdsArray {
			l = len(s)
			n += 1 + l + sovEdit(uint64(l))
		}
	}
	return n
}

func (m *FilesUpload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FilePath) > 0 {
		for _, s := range m.FilePath {
			l = len(s)
			n += 1 + l + sovEdit(uint64(l))
		}
	}
	l = len(m.BlockId)
	if l > 0 {
		n += 1 + l + sovEdit(uint64(l))
	}
	return n
}

func sovEdit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEdit(x uint64) (n int) {
	return sovEdit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserBlockJoin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEdit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserBlockJoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBlockJoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEdit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEdit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Account == nil {
				m.Account = &Account{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEdit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEdit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEdit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserBlockLeft) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEdit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserBlockLeft: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBlockLeft: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEdit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEdit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Account == nil {
				m.Account = &Account{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEdit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEdit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEdit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserBlockTextRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEdit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserBlockTextRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBlockTextRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEdit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEdit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Account == nil {
				m.Account = &Account{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEdit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEdit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEdit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEdit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Range == nil {
				m.Range = &Model_Range{}
			}
			if err := m.Range.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEdit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEdit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEdit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserBlockSelectRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEdit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserBlockSelectRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBlockSelectRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEdit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEdit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Account == nil {
				m.Account = &Account{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockIdsArray", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEdit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEdit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockIdsArray = append(m.BlockIdsArray, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEdit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEdit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEdit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FilesUpload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEdit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FilesUpload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilesUpload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEdit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEdit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilePath = append(m.FilePath, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEdit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEdit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEdit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEdit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEdit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEdit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEdit
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEdit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEdit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEdit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEdit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEdit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEdit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEdit = fmt.Errorf("proto: unexpected end of group")
)
