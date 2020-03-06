// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/storage/protos/block.proto

package storage

import (
	fmt "fmt"
	model "github.com/anytypeio/go-anytype-library/pb/model"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type SmartBlockWithMeta struct {
	Blocks     []*model.Block       `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Details    *types.Struct        `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	KeysByHash map[string]*FileKeys `protobuf:"bytes,3,rep,name=keysByHash,proto3" json:"keysByHash,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *SmartBlockWithMeta) Reset()         { *m = SmartBlockWithMeta{} }
func (m *SmartBlockWithMeta) String() string { return proto.CompactTextString(m) }
func (*SmartBlockWithMeta) ProtoMessage()    {}
func (*SmartBlockWithMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_922fa7b2983033eb, []int{0}
}
func (m *SmartBlockWithMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmartBlockWithMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SmartBlockWithMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SmartBlockWithMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartBlockWithMeta.Merge(m, src)
}
func (m *SmartBlockWithMeta) XXX_Size() int {
	return m.Size()
}
func (m *SmartBlockWithMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartBlockWithMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SmartBlockWithMeta proto.InternalMessageInfo

func (m *SmartBlockWithMeta) GetBlocks() []*model.Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *SmartBlockWithMeta) GetDetails() *types.Struct {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *SmartBlockWithMeta) GetKeysByHash() map[string]*FileKeys {
	if m != nil {
		return m.KeysByHash
	}
	return nil
}

type BlockMetaOnly struct {
	Details *types.Struct `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *BlockMetaOnly) Reset()         { *m = BlockMetaOnly{} }
func (m *BlockMetaOnly) String() string { return proto.CompactTextString(m) }
func (*BlockMetaOnly) ProtoMessage()    {}
func (*BlockMetaOnly) Descriptor() ([]byte, []int) {
	return fileDescriptor_922fa7b2983033eb, []int{1}
}
func (m *BlockMetaOnly) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockMetaOnly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockMetaOnly.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockMetaOnly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMetaOnly.Merge(m, src)
}
func (m *BlockMetaOnly) XXX_Size() int {
	return m.Size()
}
func (m *BlockMetaOnly) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMetaOnly.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMetaOnly proto.InternalMessageInfo

func (m *BlockMetaOnly) GetDetails() *types.Struct {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*SmartBlockWithMeta)(nil), "anytype.storage.SmartBlockWithMeta")
	proto.RegisterMapType((map[string]*FileKeys)(nil), "anytype.storage.SmartBlockWithMeta.KeysByHashEntry")
	proto.RegisterType((*BlockMetaOnly)(nil), "anytype.storage.BlockMetaOnly")
}

func init() { proto.RegisterFile("pb/storage/protos/block.proto", fileDescriptor_922fa7b2983033eb) }

var fileDescriptor_922fa7b2983033eb = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4a, 0x2b, 0x31,
	0x14, 0x86, 0x9b, 0x96, 0xdb, 0xcb, 0x4d, 0xb9, 0xf4, 0x12, 0x2e, 0x38, 0x16, 0x1d, 0x4a, 0x57,
	0x5d, 0xd8, 0x04, 0xdb, 0x8d, 0xb8, 0x1c, 0x50, 0x04, 0x15, 0x61, 0xba, 0x50, 0xdc, 0x25, 0x6d,
	0x3a, 0x0d, 0x4d, 0x9b, 0x21, 0xc9, 0x08, 0x79, 0x0b, 0x1f, 0xc2, 0x87, 0x71, 0xd9, 0xa5, 0x4b,
	0x69, 0x5f, 0x44, 0x26, 0x33, 0xa3, 0xd2, 0xee, 0xdc, 0xcd, 0xe1, 0xff, 0xe7, 0x7c, 0xdf, 0x21,
	0xf0, 0x38, 0x65, 0xc4, 0x58, 0xa5, 0x69, 0xc2, 0x49, 0xaa, 0x95, 0x55, 0x86, 0x30, 0xa9, 0x26,
	0x0b, 0xec, 0x07, 0xd4, 0xa6, 0x2b, 0x67, 0x5d, 0xca, 0x71, 0xd9, 0xe9, 0x1c, 0xa5, 0x8c, 0x2c,
	0xd5, 0x94, 0xcb, 0xaa, 0xed, 0x07, 0x53, 0xd4, 0x7d, 0xba, 0xb3, 0x6d, 0x26, 0x24, 0xaf, 0xd2,
	0x44, 0xa9, 0x44, 0x96, 0x09, 0xcb, 0x66, 0xc4, 0x58, 0x9d, 0x4d, 0x6c, 0x91, 0xf6, 0x5e, 0xea,
	0x10, 0x8d, 0x97, 0x54, 0xdb, 0x28, 0xe7, 0xdf, 0x0b, 0x3b, 0xbf, 0xe5, 0x96, 0xa2, 0x13, 0xd8,
	0xf4, 0x42, 0x26, 0x00, 0xdd, 0x46, 0xbf, 0x35, 0xfc, 0x8f, 0x2b, 0x25, 0x4f, 0xc6, 0xbe, 0x1d,
	0x97, 0x1d, 0x74, 0x0a, 0x7f, 0x4f, 0xb9, 0xa5, 0x42, 0x9a, 0xa0, 0xde, 0x05, 0xfd, 0xd6, 0xf0,
	0x00, 0x17, 0x50, 0x5c, 0x41, 0xf1, 0xd8, 0x43, 0xe3, 0xaa, 0x87, 0xc6, 0x10, 0x2e, 0xb8, 0x33,
	0x91, 0xbb, 0xa2, 0x66, 0x1e, 0x34, 0x3c, 0x64, 0x84, 0x77, 0xee, 0xc6, 0xfb, 0x66, 0xf8, 0xfa,
	0xf3, 0xaf, 0x8b, 0x95, 0xd5, 0x2e, 0xfe, 0xb6, 0xa6, 0xf3, 0x00, 0xdb, 0x3b, 0x31, 0xfa, 0x07,
	0x1b, 0x0b, 0xee, 0x02, 0xd0, 0x05, 0xfd, 0x3f, 0x71, 0xfe, 0x89, 0x08, 0xfc, 0xf5, 0x44, 0x65,
	0xc6, 0x4b, 0xd5, 0xc3, 0x3d, 0xe8, 0xa5, 0x90, 0x3c, 0x5f, 0x13, 0x17, 0xbd, 0xf3, 0xfa, 0x19,
	0xe8, 0x45, 0xf0, 0xaf, 0xd7, 0xc8, 0x15, 0xee, 0x56, 0xd2, 0xfd, 0xe0, 0xe4, 0xe8, 0xe6, 0x75,
	0x13, 0x82, 0xf5, 0x26, 0x04, 0xef, 0x9b, 0x10, 0x3c, 0x6f, 0xc3, 0xda, 0x7a, 0x1b, 0xd6, 0xde,
	0xb6, 0x61, 0xed, 0x71, 0x98, 0x08, 0x3b, 0xcf, 0x18, 0x9e, 0xa8, 0x25, 0x29, 0x6d, 0x84, 0x22,
	0x89, 0x1a, 0x94, 0xc3, 0x40, 0x0a, 0xa6, 0xa9, 0x76, 0xe4, 0xeb, 0xa1, 0x59, 0xd3, 0x73, 0x46,
	0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0xc2, 0x2c, 0xb0, 0x4b, 0x02, 0x00, 0x00,
}

func (m *SmartBlockWithMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmartBlockWithMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SmartBlockWithMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeysByHash) > 0 {
		for k := range m.KeysByHash {
			v := m.KeysByHash[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintBlock(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintBlock(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintBlock(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Details != nil {
		{
			size, err := m.Details.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Blocks) > 0 {
		for iNdEx := len(m.Blocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Blocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BlockMetaOnly) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockMetaOnly) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockMetaOnly) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Details != nil {
		{
			size, err := m.Details.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SmartBlockWithMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Blocks) > 0 {
		for _, e := range m.Blocks {
			l = e.Size()
			n += 1 + l + sovBlock(uint64(l))
		}
	}
	if m.Details != nil {
		l = m.Details.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if len(m.KeysByHash) > 0 {
		for k, v := range m.KeysByHash {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovBlock(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovBlock(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovBlock(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *BlockMetaOnly) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Details != nil {
		l = m.Details.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SmartBlockWithMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: SmartBlockWithMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmartBlockWithMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocks = append(m.Blocks, &model.Block{})
			if err := m.Blocks[len(m.Blocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Details == nil {
				m.Details = &types.Struct{}
			}
			if err := m.Details.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeysByHash", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.KeysByHash == nil {
				m.KeysByHash = make(map[string]*FileKeys)
			}
			var mapkey string
			var mapvalue *FileKeys
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBlock
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBlock
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthBlock
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthBlock
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBlock
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthBlock
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthBlock
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &FileKeys{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipBlock(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthBlock
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.KeysByHash[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *BlockMetaOnly) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: BlockMetaOnly: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockMetaOnly: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Details == nil {
				m.Details = &types.Struct{}
			}
			if err := m.Details.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)
