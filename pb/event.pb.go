// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: event.proto

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	// Types that are valid to be assigned to Message:
	//	*Event_AccountShow
	//	*Event_UnitsShow
	//	*Event_UnitsCreate
	Message isEvent_Message `protobuf_oneof:"message"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_Message interface {
	isEvent_Message()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Event_AccountShow struct {
	AccountShow *AccountShow `protobuf:"bytes,1,opt,name=accountShow,proto3,oneof"`
}
type Event_UnitsShow struct {
	UnitsShow *Units `protobuf:"bytes,102,opt,name=unitsShow,proto3,oneof"`
}
type Event_UnitsCreate struct {
	UnitsCreate *Units `protobuf:"bytes,103,opt,name=unitsCreate,proto3,oneof"`
}

func (*Event_AccountShow) isEvent_Message() {}
func (*Event_UnitsShow) isEvent_Message()   {}
func (*Event_UnitsCreate) isEvent_Message() {}

func (m *Event) GetMessage() isEvent_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Event) GetAccountShow() *AccountShow {
	if x, ok := m.GetMessage().(*Event_AccountShow); ok {
		return x.AccountShow
	}
	return nil
}

func (m *Event) GetUnitsShow() *Units {
	if x, ok := m.GetMessage().(*Event_UnitsShow); ok {
		return x.UnitsShow
	}
	return nil
}

func (m *Event) GetUnitsCreate() *Units {
	if x, ok := m.GetMessage().(*Event_UnitsCreate); ok {
		return x.UnitsCreate
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_AccountShow)(nil),
		(*Event_UnitsShow)(nil),
		(*Event_UnitsCreate)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// message
	switch x := m.Message.(type) {
	case *Event_AccountShow:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AccountShow); err != nil {
			return err
		}
	case *Event_UnitsShow:
		_ = b.EncodeVarint(102<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UnitsShow); err != nil {
			return err
		}
	case *Event_UnitsCreate:
		_ = b.EncodeVarint(103<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UnitsCreate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Message has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 1: // message.accountShow
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AccountShow)
		err := b.DecodeMessage(msg)
		m.Message = &Event_AccountShow{msg}
		return true, err
	case 102: // message.unitsShow
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Units)
		err := b.DecodeMessage(msg)
		m.Message = &Event_UnitsShow{msg}
		return true, err
	case 103: // message.unitsCreate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Units)
		err := b.DecodeMessage(msg)
		m.Message = &Event_UnitsCreate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// message
	switch x := m.Message.(type) {
	case *Event_AccountShow:
		s := proto.Size(x.AccountShow)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_UnitsShow:
		s := proto.Size(x.UnitsShow)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_UnitsCreate:
		s := proto.Size(x.UnitsCreate)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AccountShow struct {
	Index   int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Account *Account `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *AccountShow) Reset()         { *m = AccountShow{} }
func (m *AccountShow) String() string { return proto.CompactTextString(m) }
func (*AccountShow) ProtoMessage()    {}
func (*AccountShow) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}
func (m *AccountShow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountShow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountShow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountShow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountShow.Merge(m, src)
}
func (m *AccountShow) XXX_Size() int {
	return m.Size()
}
func (m *AccountShow) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountShow.DiscardUnknown(m)
}

var xxx_messageInfo_AccountShow proto.InternalMessageInfo

func (m *AccountShow) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AccountShow) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "anytype.Event")
	proto.RegisterType((*AccountShow)(nil), "anytype.AccountShow")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x4b, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xcc, 0xab, 0x2c, 0xa9, 0x2c, 0x48,
	0x95, 0xe2, 0x4d, 0x4c, 0x4e, 0xce, 0x2f, 0x85, 0x89, 0x4b, 0x71, 0x95, 0xe6, 0x65, 0x42, 0xd9,
	0x4a, 0x1b, 0x18, 0xb9, 0x58, 0x5d, 0x41, 0x7a, 0x84, 0x2c, 0xb8, 0xb8, 0xa1, 0xca, 0x82, 0x33,
	0xf2, 0xcb, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x44, 0xf4, 0xa0, 0x66, 0xe8, 0x39, 0x22,
	0xe4, 0x3c, 0x18, 0x82, 0x90, 0x95, 0x0a, 0xe9, 0x71, 0x71, 0x82, 0x4c, 0x2c, 0x06, 0xeb, 0x4b,
	0x03, 0xeb, 0xe3, 0x83, 0xeb, 0x0b, 0x05, 0xc9, 0x78, 0x30, 0x04, 0x21, 0x94, 0x08, 0x19, 0x71,
	0x71, 0x83, 0x39, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x12, 0xe9, 0x38, 0x74, 0x20, 0x2b, 0x72,
	0xe2, 0xe4, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x55, 0xf2, 0xe7, 0xe2, 0x46, 0x72,
	0x8c, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x5e, 0x4a, 0x6a, 0x05, 0xd8, 0xc5, 0xcc, 0x41, 0x10, 0x8e,
	0x90, 0x16, 0x17, 0x3b, 0xd4, 0x89, 0x12, 0x4c, 0x60, 0xf3, 0x05, 0xd0, 0x7d, 0x12, 0x04, 0x53,
	0xe0, 0x24, 0x73, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x4c, 0x05, 0x49,
	0x49, 0x6c, 0xe0, 0x80, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x98, 0x48, 0x41, 0x5b,
	0x01, 0x00, 0x00,
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Message != nil {
		nn1, err1 := m.Message.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += nn1
	}
	return i, nil
}

func (m *Event_AccountShow) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.AccountShow != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.AccountShow.Size()))
		n2, err2 := m.AccountShow.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	return i, nil
}
func (m *Event_UnitsShow) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.UnitsShow != nil {
		dAtA[i] = 0xb2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.UnitsShow.Size()))
		n3, err3 := m.UnitsShow.MarshalTo(dAtA[i:])
		if err3 != nil {
			return 0, err3
		}
		i += n3
	}
	return i, nil
}
func (m *Event_UnitsCreate) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.UnitsCreate != nil {
		dAtA[i] = 0xba
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.UnitsCreate.Size()))
		n4, err4 := m.UnitsCreate.MarshalTo(dAtA[i:])
		if err4 != nil {
			return 0, err4
		}
		i += n4
	}
	return i, nil
}
func (m *AccountShow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountShow) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.Index))
	}
	if m.Account != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.Account.Size()))
		n5, err5 := m.Account.MarshalTo(dAtA[i:])
		if err5 != nil {
			return 0, err5
		}
		i += n5
	}
	return i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Message != nil {
		n += m.Message.Size()
	}
	return n
}

func (m *Event_AccountShow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccountShow != nil {
		l = m.AccountShow.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}
func (m *Event_UnitsShow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UnitsShow != nil {
		l = m.UnitsShow.Size()
		n += 2 + l + sovEvent(uint64(l))
	}
	return n
}
func (m *Event_UnitsCreate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UnitsCreate != nil {
		l = m.UnitsCreate.Size()
		n += 2 + l + sovEvent(uint64(l))
	}
	return n
}
func (m *AccountShow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovEvent(uint64(m.Index))
	}
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountShow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AccountShow{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Message = &Event_AccountShow{v}
			iNdEx = postIndex
		case 102:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnitsShow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Units{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Message = &Event_UnitsShow{v}
			iNdEx = postIndex
		case 103:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnitsCreate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Units{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Message = &Event_UnitsCreate{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *AccountShow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: AccountShow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountShow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
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
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEvent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEvent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthEvent
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent   = fmt.Errorf("proto: integer overflow")
)
