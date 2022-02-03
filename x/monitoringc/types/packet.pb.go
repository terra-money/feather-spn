// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: monitoringc/packet.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/tendermint/spn/pkg/types"
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

type MonitoringcPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*MonitoringcPacketData_MonitoringPacket
	Packet isMonitoringcPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *MonitoringcPacketData) Reset()         { *m = MonitoringcPacketData{} }
func (m *MonitoringcPacketData) String() string { return proto.CompactTextString(m) }
func (*MonitoringcPacketData) ProtoMessage()    {}
func (*MonitoringcPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e34e8c279cb01f10, []int{0}
}
func (m *MonitoringcPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MonitoringcPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MonitoringcPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MonitoringcPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoringcPacketData.Merge(m, src)
}
func (m *MonitoringcPacketData) XXX_Size() int {
	return m.Size()
}
func (m *MonitoringcPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoringcPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoringcPacketData proto.InternalMessageInfo

type isMonitoringcPacketData_Packet interface {
	isMonitoringcPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type MonitoringcPacketData_MonitoringPacket struct {
	MonitoringPacket *types.MonitoringPacket `protobuf:"bytes,1,opt,name=monitoringPacket,proto3,oneof" json:"monitoringPacket,omitempty"`
}

func (*MonitoringcPacketData_MonitoringPacket) isMonitoringcPacketData_Packet() {}

func (m *MonitoringcPacketData) GetPacket() isMonitoringcPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *MonitoringcPacketData) GetMonitoringPacket() *types.MonitoringPacket {
	if x, ok := m.GetPacket().(*MonitoringcPacketData_MonitoringPacket); ok {
		return x.MonitoringPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MonitoringcPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MonitoringcPacketData_MonitoringPacket)(nil),
	}
}

// MonitoringPacketAck defines a struct for the packet acknowledgment
type MonitoringPacketAck struct {
}

func (m *MonitoringPacketAck) Reset()         { *m = MonitoringPacketAck{} }
func (m *MonitoringPacketAck) String() string { return proto.CompactTextString(m) }
func (*MonitoringPacketAck) ProtoMessage()    {}
func (*MonitoringPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_e34e8c279cb01f10, []int{1}
}
func (m *MonitoringPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MonitoringPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MonitoringPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MonitoringPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoringPacketAck.Merge(m, src)
}
func (m *MonitoringPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *MonitoringPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoringPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoringPacketAck proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MonitoringcPacketData)(nil), "tendermint.spn.monitoringc.MonitoringcPacketData")
	proto.RegisterType((*MonitoringPacketAck)(nil), "tendermint.spn.monitoringc.MonitoringPacketAck")
}

func init() { proto.RegisterFile("monitoringc/packet.proto", fileDescriptor_e34e8c279cb01f10) }

var fileDescriptor_e34e8c279cb01f10 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xcd, 0xcf, 0xcb,
	0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0x4f, 0xd6, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2a, 0x49, 0xcd, 0x4b, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b,
	0xd1, 0x2b, 0x2e, 0xc8, 0xd3, 0x43, 0x52, 0x28, 0x25, 0x56, 0x52, 0x59, 0x90, 0x5a, 0xac, 0x8f,
	0x10, 0x82, 0xe8, 0x51, 0x2a, 0xe7, 0x12, 0xf5, 0x45, 0x28, 0x0b, 0x00, 0x1b, 0xe7, 0x92, 0x58,
	0x92, 0x28, 0x14, 0xc2, 0x25, 0x80, 0x50, 0x0c, 0x11, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36,
	0x52, 0xd3, 0x43, 0xb3, 0x07, 0x6c, 0xb4, 0x9e, 0x2f, 0x9a, 0x6a, 0x0f, 0x86, 0x20, 0x0c, 0x13,
	0x9c, 0x38, 0xb8, 0xd8, 0x20, 0x4e, 0x56, 0x12, 0xe5, 0x12, 0x46, 0xd7, 0xe1, 0x98, 0x9c, 0xed,
	0xe4, 0x7e, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xba, 0xe9, 0x99, 0x25,
	0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x08, 0x07, 0xe8, 0x17, 0x17, 0xe4, 0xe9, 0x57,
	0xe8, 0x23, 0x87, 0x09, 0xd8, 0x39, 0x49, 0x6c, 0x60, 0xff, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0x57, 0x63, 0x4e, 0x2f, 0x01, 0x00, 0x00,
}

func (m *MonitoringcPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MonitoringcPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MonitoringcPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *MonitoringcPacketData_MonitoringPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MonitoringcPacketData_MonitoringPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MonitoringPacket != nil {
		{
			size, err := m.MonitoringPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *MonitoringPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MonitoringPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MonitoringPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MonitoringcPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *MonitoringcPacketData_MonitoringPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MonitoringPacket != nil {
		l = m.MonitoringPacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *MonitoringPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MonitoringcPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: MonitoringcPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonitoringcPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonitoringPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.MonitoringPacket{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &MonitoringcPacketData_MonitoringPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *MonitoringPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: MonitoringPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonitoringPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)