// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/lib/pb/storage/protos/block.proto

package storage

import (
	fmt "fmt"
	model "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
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

type SmartBlockSnapshot struct {
	State      map[string]uint64    `protobuf:"bytes,1,rep,name=state,proto3" json:"state,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Blocks     []*model.Block       `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Details    *types.Struct        `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	KeysByHash map[string]*FileKeys `protobuf:"bytes,4,rep,name=keysByHash,proto3" json:"keysByHash,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ClientTime int64                `protobuf:"varint,5,opt,name=clientTime,proto3" json:"clientTime,omitempty"`
}

func (m *SmartBlockSnapshot) Reset()         { *m = SmartBlockSnapshot{} }
func (m *SmartBlockSnapshot) String() string { return proto.CompactTextString(m) }
func (*SmartBlockSnapshot) ProtoMessage()    {}
func (*SmartBlockSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dc566bdc18e510, []int{0}
}
func (m *SmartBlockSnapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmartBlockSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SmartBlockSnapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SmartBlockSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartBlockSnapshot.Merge(m, src)
}
func (m *SmartBlockSnapshot) XXX_Size() int {
	return m.Size()
}
func (m *SmartBlockSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartBlockSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_SmartBlockSnapshot proto.InternalMessageInfo

func (m *SmartBlockSnapshot) GetState() map[string]uint64 {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *SmartBlockSnapshot) GetBlocks() []*model.Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *SmartBlockSnapshot) GetDetails() *types.Struct {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *SmartBlockSnapshot) GetKeysByHash() map[string]*FileKeys {
	if m != nil {
		return m.KeysByHash
	}
	return nil
}

func (m *SmartBlockSnapshot) GetClientTime() int64 {
	if m != nil {
		return m.ClientTime
	}
	return 0
}

func init() {
	proto.RegisterType((*SmartBlockSnapshot)(nil), "anytype.storage.SmartBlockSnapshot")
	proto.RegisterMapType((map[string]*FileKeys)(nil), "anytype.storage.SmartBlockSnapshot.KeysByHashEntry")
	proto.RegisterMapType((map[string]uint64)(nil), "anytype.storage.SmartBlockSnapshot.StateEntry")
}

func init() {
	proto.RegisterFile("pkg/lib/pb/storage/protos/block.proto", fileDescriptor_98dc566bdc18e510)
}

var fileDescriptor_98dc566bdc18e510 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x4f, 0xc2, 0x40,
	0x1c, 0xc5, 0x39, 0x0a, 0x18, 0x8f, 0x01, 0x73, 0x21, 0xb1, 0x12, 0xd3, 0x34, 0x46, 0x93, 0x0e,
	0xe4, 0x2e, 0xc2, 0x42, 0x1c, 0x89, 0x1a, 0x13, 0xb7, 0xd6, 0xc1, 0xb8, 0x5d, 0xe1, 0x28, 0x4d,
	0x8f, 0x5e, 0xd3, 0x3b, 0x4c, 0xfa, 0x2d, 0xfc, 0x50, 0x0e, 0x8e, 0x8c, 0x8e, 0x06, 0xbe, 0x88,
	0xe9, 0xb5, 0x60, 0x05, 0x07, 0xb7, 0xfe, 0xfb, 0x7f, 0xbf, 0xf7, 0x7f, 0xaf, 0x29, 0xbc, 0x4a,
	0xa2, 0x80, 0xf0, 0xd0, 0x27, 0x89, 0x4f, 0xa4, 0x12, 0x29, 0x0d, 0x18, 0x49, 0x52, 0xa1, 0x84,
	0x24, 0x3e, 0x17, 0x93, 0x08, 0xeb, 0x01, 0x75, 0x68, 0x9c, 0xa9, 0x2c, 0x61, 0xb8, 0xd4, 0xf4,
	0x2e, 0x2b, 0xdc, 0x42, 0x4c, 0x19, 0xdf, 0x52, 0x7a, 0x90, 0x05, 0xf6, 0x4b, 0xb5, 0xe7, 0x3e,
	0x0b, 0x39, 0x2b, 0x55, 0xe7, 0x81, 0x10, 0x01, 0x2f, 0x37, 0xfe, 0x72, 0x46, 0xa4, 0x4a, 0x97,
	0x13, 0x55, 0x6c, 0x2f, 0xde, 0x0d, 0x88, 0xbc, 0x05, 0x4d, 0xd5, 0x38, 0xcf, 0xe3, 0xc5, 0x34,
	0x91, 0x73, 0xa1, 0xd0, 0x2d, 0x6c, 0x4a, 0x45, 0x15, 0x33, 0x81, 0x6d, 0x38, 0xed, 0x01, 0xc6,
	0x7b, 0x09, 0xf1, 0x21, 0x83, 0xbd, 0x1c, 0xb8, 0x8b, 0x55, 0x9a, 0xb9, 0x05, 0x8c, 0xfa, 0xb0,
	0xa5, 0x6b, 0x4a, 0xb3, 0xae, 0x6d, 0xba, 0x3b, 0x1b, 0xdd, 0x03, 0x6b, 0xde, 0x2d, 0x35, 0xe8,
	0x1a, 0x1e, 0x4d, 0x99, 0xa2, 0x21, 0x97, 0xa6, 0x61, 0x03, 0xa7, 0x3d, 0x38, 0xc5, 0x45, 0x74,
	0xbc, 0x8d, 0x8e, 0x3d, 0x1d, 0xdd, 0xdd, 0xea, 0x90, 0x07, 0x61, 0xc4, 0x32, 0x39, 0xce, 0x1e,
	0xa8, 0x9c, 0x9b, 0x0d, 0x7d, 0x64, 0xf8, 0x9f, 0xac, 0x8f, 0x3b, 0xaa, 0x08, 0x5c, 0xb1, 0x41,
	0x16, 0x84, 0x13, 0x1e, 0xb2, 0x58, 0x3d, 0x85, 0x0b, 0x66, 0x36, 0x6d, 0xe0, 0x18, 0x6e, 0xe5,
	0x4d, 0x6f, 0x04, 0xe1, 0x4f, 0x55, 0x74, 0x02, 0x8d, 0x88, 0x65, 0x26, 0xb0, 0x81, 0x73, 0xec,
	0xe6, 0x8f, 0xa8, 0x0b, 0x9b, 0xaf, 0x94, 0x2f, 0x99, 0x59, 0xb7, 0x81, 0xd3, 0x70, 0x8b, 0xe1,
	0xa6, 0x3e, 0x02, 0xbd, 0x67, 0xd8, 0xd9, 0x3b, 0xfc, 0x07, 0x4e, 0xaa, 0x78, 0x7b, 0x70, 0x76,
	0x50, 0xe7, 0x3e, 0xe4, 0x2c, 0xb7, 0xa9, 0x38, 0x8f, 0xfb, 0x1f, 0x6b, 0x0b, 0xac, 0xd6, 0x16,
	0xf8, 0x5a, 0x5b, 0xe0, 0x6d, 0x63, 0xd5, 0x56, 0x1b, 0xab, 0xf6, 0xb9, 0xb1, 0x6a, 0x2f, 0xe8,
	0xf0, 0x27, 0xf1, 0x5b, 0xfa, 0x83, 0x0e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xf5, 0x39,
	0x64, 0x9f, 0x02, 0x00, 0x00,
}

func (m *SmartBlockSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmartBlockSnapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SmartBlockSnapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClientTime != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.ClientTime))
		i--
		dAtA[i] = 0x28
	}
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
			dAtA[i] = 0x22
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
		dAtA[i] = 0x1a
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
			dAtA[i] = 0x12
		}
	}
	if len(m.State) > 0 {
		for k := range m.State {
			v := m.State[k]
			baseI := i
			i = encodeVarintBlock(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintBlock(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintBlock(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
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
func (m *SmartBlockSnapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.State) > 0 {
		for k, v := range m.State {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovBlock(uint64(len(k))) + 1 + sovBlock(uint64(v))
			n += mapEntrySize + 1 + sovBlock(uint64(mapEntrySize))
		}
	}
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
	if m.ClientTime != 0 {
		n += 1 + sovBlock(uint64(m.ClientTime))
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SmartBlockSnapshot) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SmartBlockSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmartBlockSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
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
			if m.State == nil {
				m.State = make(map[string]uint64)
			}
			var mapkey string
			var mapvalue uint64
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
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBlock
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
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
			m.State[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
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
		case 3:
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
		case 4:
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientTime", wireType)
			}
			m.ClientTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
