// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: project/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Params defines the set of params for the project module.
type Params struct {
	TotalSupplyRange   TotalSupplyRange                         `protobuf:"bytes,1,opt,name=totalSupplyRange,proto3" json:"totalSupplyRange"`
	ProjectCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=projectCreationFee,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"projectCreationFee"`
	MaxMetadataLength  uint64                                   `protobuf:"varint,3,opt,name=maxMetadataLength,proto3" json:"maxMetadataLength,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_15316631d2950b4e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetTotalSupplyRange() TotalSupplyRange {
	if m != nil {
		return m.TotalSupplyRange
	}
	return TotalSupplyRange{}
}

func (m *Params) GetProjectCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ProjectCreationFee
	}
	return nil
}

func (m *Params) GetMaxMetadataLength() uint64 {
	if m != nil {
		return m.MaxMetadataLength
	}
	return 0
}

// TotalSupplyRange defines the range of allowed values for total supply
type TotalSupplyRange struct {
	MinTotalSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=minTotalSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"minTotalSupply"`
	MaxTotalSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=maxTotalSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"maxTotalSupply"`
}

func (m *TotalSupplyRange) Reset()         { *m = TotalSupplyRange{} }
func (m *TotalSupplyRange) String() string { return proto.CompactTextString(m) }
func (*TotalSupplyRange) ProtoMessage()    {}
func (*TotalSupplyRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_15316631d2950b4e, []int{1}
}
func (m *TotalSupplyRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TotalSupplyRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TotalSupplyRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TotalSupplyRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalSupplyRange.Merge(m, src)
}
func (m *TotalSupplyRange) XXX_Size() int {
	return m.Size()
}
func (m *TotalSupplyRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalSupplyRange.DiscardUnknown(m)
}

var xxx_messageInfo_TotalSupplyRange proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "tendermint.spn.project.Params")
	proto.RegisterType((*TotalSupplyRange)(nil), "tendermint.spn.project.TotalSupplyRange")
}

func init() { proto.RegisterFile("project/params.proto", fileDescriptor_15316631d2950b4e) }

var fileDescriptor_15316631d2950b4e = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0xcf,
	0x4a, 0x4d, 0x2e, 0xd1, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x2b, 0x49, 0xcd, 0x4b, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2b, 0x2e, 0xc8,
	0xd3, 0x83, 0x2a, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd1, 0x07, 0xb1, 0x20, 0xaa,
	0xa5, 0xe4, 0x92, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xf5, 0x93, 0x12, 0x8b, 0x53, 0xf5, 0xcb, 0x0c,
	0x93, 0x52, 0x4b, 0x12, 0x0d, 0xf5, 0x93, 0xf3, 0x33, 0xf3, 0xa0, 0xf2, 0x92, 0x10, 0xf9, 0x78,
	0x88, 0x46, 0x08, 0x07, 0x22, 0xa5, 0x74, 0x90, 0x89, 0x8b, 0x2d, 0x00, 0x6c, 0xb3, 0x50, 0x14,
	0x97, 0x40, 0x49, 0x7e, 0x49, 0x62, 0x4e, 0x70, 0x69, 0x41, 0x41, 0x4e, 0x65, 0x50, 0x62, 0x5e,
	0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x86, 0x1e, 0x76, 0xe7, 0xe8, 0x85, 0xa0,
	0xa9, 0x77, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0xc3, 0x1c, 0xa1, 0x65, 0x8c, 0x5c, 0x42,
	0x50, 0x4d, 0xce, 0x45, 0xa9, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x6e, 0xa9, 0xa9, 0x12, 0x4c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x92, 0x7a, 0x50, 0x27, 0x81, 0xdc, 0xaf, 0x07, 0x75, 0xbf, 0x9e, 0x73,
	0x7e, 0x66, 0x9e, 0x53, 0x34, 0xc8, 0xbc, 0x5f, 0xf7, 0xe4, 0xd5, 0xd3, 0x33, 0x4b, 0x32, 0x4a,
	0x93, 0xf4, 0x92, 0xf3, 0x73, 0xa1, 0xee, 0x87, 0x52, 0xba, 0xc5, 0x29, 0xd9, 0xfa, 0x25, 0x95,
	0x05, 0xa9, 0xc5, 0x60, 0x0d, 0xab, 0xee, 0xcb, 0x6b, 0x10, 0xa9, 0xb4, 0x38, 0x08, 0x8b, 0x8b,
	0x84, 0x74, 0xb8, 0x04, 0x73, 0x13, 0x2b, 0x7c, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x7d,
	0x52, 0xf3, 0xd2, 0x4b, 0x32, 0x24, 0x98, 0x15, 0x18, 0x35, 0x58, 0x82, 0x30, 0x25, 0xac, 0x58,
	0x66, 0x2c, 0x90, 0x67, 0x50, 0xba, 0xc7, 0xc8, 0x25, 0x80, 0x1e, 0x12, 0x42, 0x29, 0x5c, 0x7c,
	0xb9, 0x99, 0x79, 0x48, 0xc2, 0xe0, 0xb0, 0xe4, 0x74, 0xb2, 0x01, 0xf9, 0xe8, 0xd6, 0x3d, 0x79,
	0x35, 0x22, 0x9c, 0xe9, 0x99, 0x57, 0x72, 0x69, 0x8b, 0x2e, 0x17, 0x34, 0x74, 0x3c, 0xf3, 0x4a,
	0x82, 0xd0, 0xcc, 0x04, 0xdb, 0x92, 0x58, 0x81, 0x6c, 0x0b, 0x13, 0x55, 0x6c, 0x41, 0x31, 0xd3,
	0xc9, 0xf9, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0,
	0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x34, 0x91, 0xcc, 0x47,
	0xa4, 0x11, 0xfd, 0xe2, 0x82, 0x3c, 0xfd, 0x0a, 0x7d, 0x58, 0xca, 0x06, 0x5b, 0x93, 0xc4, 0x06,
	0x4e, 0x70, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0xfe, 0x58, 0x40, 0xf1, 0x02, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxMetadataLength != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxMetadataLength))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ProjectCreationFee) > 0 {
		for iNdEx := len(m.ProjectCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProjectCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.TotalSupplyRange.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TotalSupplyRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TotalSupplyRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TotalSupplyRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxTotalSupply.Size()
		i -= size
		if _, err := m.MaxTotalSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinTotalSupply.Size()
		i -= size
		if _, err := m.MinTotalSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TotalSupplyRange.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.ProjectCreationFee) > 0 {
		for _, e := range m.ProjectCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.MaxMetadataLength != 0 {
		n += 1 + sovParams(uint64(m.MaxMetadataLength))
	}
	return n
}

func (m *TotalSupplyRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinTotalSupply.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxTotalSupply.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupplyRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalSupplyRange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectCreationFee = append(m.ProjectCreationFee, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.ProjectCreationFee[len(m.ProjectCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMetadataLength", wireType)
			}
			m.MaxMetadataLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMetadataLength |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *TotalSupplyRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: TotalSupplyRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TotalSupplyRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTotalSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinTotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTotalSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxTotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
