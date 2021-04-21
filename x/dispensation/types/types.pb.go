// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/dispensation/v1/types.proto

package types

import (
	fmt "fmt"
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

type DistributionType int32

const (
	DistributionType_DISTRIBUTION_TYPE_UNSPECIFIED DistributionType = 0
	DistributionType_DISTRIBUTION_TYPE_AIRDROP     DistributionType = 1
)

var DistributionType_name = map[int32]string{
	0: "DISTRIBUTION_TYPE_UNSPECIFIED",
	1: "DISTRIBUTION_TYPE_AIRDROP",
}

var DistributionType_value = map[string]int32{
	"DISTRIBUTION_TYPE_UNSPECIFIED": 0,
	"DISTRIBUTION_TYPE_AIRDROP":     1,
}

func (x DistributionType) String() string {
	return proto.EnumName(DistributionType_name, int32(x))
}

func (DistributionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bfdf912039cd8799, []int{0}
}

type GenesisState struct {
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfdf912039cd8799, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("sifnode.dispensation.v1.DistributionType", DistributionType_name, DistributionType_value)
	proto.RegisterType((*GenesisState)(nil), "sifnode.dispensation.v1.GenesisState")
}

func init() {
	proto.RegisterFile("sifnode/dispensation/v1/types.proto", fileDescriptor_bfdf912039cd8799)
}

var fileDescriptor_bfdf912039cd8799 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0xce, 0x4c, 0xcb,
	0xcb, 0x4f, 0x49, 0xd5, 0x4f, 0xc9, 0x2c, 0x2e, 0x48, 0xcd, 0x2b, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf,
	0xd3, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x87, 0x2a, 0xd2, 0x43, 0x56, 0xa4, 0x57, 0x66, 0x28, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f,
	0x56, 0xa3, 0x0f, 0x62, 0x41, 0x94, 0x2b, 0xf1, 0x71, 0xf1, 0xb8, 0xa7, 0xe6, 0xa5, 0x16, 0x67,
	0x16, 0x07, 0x97, 0x24, 0x96, 0xa4, 0x6a, 0x85, 0x70, 0x09, 0xb8, 0x64, 0x16, 0x97, 0x14, 0x65,
	0x26, 0x95, 0x82, 0x34, 0x86, 0x54, 0x16, 0xa4, 0x0a, 0x29, 0x72, 0xc9, 0xba, 0x78, 0x06, 0x87,
	0x04, 0x79, 0x3a, 0x85, 0x86, 0x78, 0xfa, 0xfb, 0xc5, 0x87, 0x44, 0x06, 0xb8, 0xc6, 0x87, 0xfa,
	0x05, 0x07, 0xb8, 0x3a, 0x7b, 0xba, 0x79, 0xba, 0xba, 0x08, 0x30, 0x08, 0xc9, 0x72, 0x49, 0x62,
	0x2a, 0x71, 0xf4, 0x0c, 0x72, 0x09, 0xf2, 0x0f, 0x10, 0x60, 0x74, 0xf2, 0x3a, 0xf1, 0x48, 0x8e,
	0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58,
	0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x83, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4,
	0xfc, 0x5c, 0xfd, 0xe0, 0xcc, 0xb4, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c, 0x7d, 0x98, 0x3f, 0x2b, 0x50,
	0x7d, 0x0a, 0xf6, 0x66, 0x12, 0x1b, 0xd8, 0xe1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8,
	0x73, 0x7c, 0x91, 0x0e, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
