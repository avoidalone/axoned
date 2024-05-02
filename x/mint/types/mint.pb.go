// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mint/v1beta1/mint.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Minter holds the state of minting within the blockchain.
type Minter struct {
	// Current annual inflation rate.
	Inflation cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation"`
	// Current annual anticipated provisions.
	AnnualProvisions cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"annual_provisions"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_06339c129491fd39, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

// Params defines the parameters for the mint module.
type Params struct {
	// Denomination of the coin to be minted.
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// Annual inflation coefficient, influencing the inflation rate based on the bonded ratio.
	// Values range from 0 to 1, with higher values indicating higher inflation.
	InflationCoef cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=inflation_coef,json=inflationCoef,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_coef"`
	// Estimated number of blocks per year.
	BlocksPerYear uint64 `protobuf:"varint,3,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty"`
	// Maximum annual inflation rate.
	InflationMax *cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=inflation_max,json=inflationMax,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_max,omitempty"`
	// Minimum annual inflation rate.
	InflationMin *cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=inflation_min,json=inflationMin,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation_min,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_06339c129491fd39, []int{1}
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

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "mint.v1beta1.Minter")
	proto.RegisterType((*Params)(nil), "mint.v1beta1.Params")
}

func init() { proto.RegisterFile("mint/v1beta1/mint.proto", fileDescriptor_06339c129491fd39) }

var fileDescriptor_06339c129491fd39 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0x87, 0x3b, 0x5c, 0x2e, 0x09, 0x13, 0xb8, 0x7f, 0x1a, 0x13, 0x2b, 0xc6, 0x42, 0x88, 0x31,
	0x6c, 0x6c, 0x43, 0xdc, 0xb9, 0x70, 0x81, 0x2c, 0x25, 0x12, 0x16, 0x46, 0x5d, 0xd8, 0x4c, 0xcb,
	0x50, 0x26, 0xd0, 0x39, 0xa4, 0x33, 0x90, 0xf2, 0x16, 0x2e, 0x5d, 0xfa, 0x10, 0x6e, 0x7c, 0x03,
	0x96, 0xc4, 0x95, 0x71, 0x81, 0x06, 0x5e, 0xc4, 0xb4, 0x43, 0x4a, 0xe2, 0x4e, 0xe2, 0xa6, 0x39,
	0xe7, 0x77, 0x3a, 0xdf, 0x7c, 0x99, 0x1c, 0xbc, 0x1b, 0x30, 0x2e, 0xed, 0x49, 0xdd, 0xa5, 0x92,
	0xd4, 0xed, 0xb8, 0xb1, 0x46, 0x21, 0x48, 0xd0, 0x0b, 0x49, 0xbd, 0x1e, 0x94, 0xf6, 0x3c, 0x10,
	0x01, 0x08, 0x27, 0x99, 0xd9, 0xaa, 0x51, 0x3f, 0x96, 0x76, 0x7c, 0xf0, 0x41, 0xe5, 0x71, 0xa5,
	0xd2, 0xea, 0x33, 0xc2, 0xb9, 0x16, 0xe3, 0x92, 0x86, 0xfa, 0x25, 0xce, 0x33, 0xde, 0x1b, 0x12,
	0xc9, 0x80, 0x1b, 0xa8, 0x82, 0x6a, 0xf9, 0x46, 0x7d, 0xb6, 0x28, 0x6b, 0x6f, 0x8b, 0xf2, 0xbe,
	0x22, 0x89, 0xee, 0xc0, 0x62, 0x60, 0x07, 0x44, 0xf6, 0xad, 0x0b, 0xea, 0x13, 0x6f, 0xda, 0xa4,
	0xde, 0xcb, 0xd3, 0x31, 0x5e, 0x5f, 0xd4, 0xa4, 0x5e, 0x67, 0xc3, 0xd0, 0xef, 0xf0, 0x7f, 0xc2,
	0xf9, 0x98, 0x0c, 0x63, 0x9d, 0x09, 0x13, 0x0c, 0xb8, 0x30, 0x32, 0xdb, 0x82, 0xff, 0x29, 0x56,
	0x3b, 0x45, 0x55, 0xdf, 0x33, 0x38, 0xd7, 0x26, 0x21, 0x09, 0x84, 0x7e, 0x80, 0x71, 0xfc, 0x0e,
	0x4e, 0x97, 0x72, 0x08, 0x94, 0x7c, 0x27, 0x1f, 0x27, 0xcd, 0x38, 0xd0, 0xaf, 0xf1, 0x9f, 0x54,
	0xcb, 0xf1, 0x80, 0xf6, 0xb6, 0xd7, 0x28, 0xa6, 0xa0, 0x73, 0xa0, 0x3d, 0xfd, 0x08, 0xff, 0x75,
	0x87, 0xe0, 0x0d, 0x84, 0x33, 0xa2, 0xa1, 0x33, 0xa5, 0x24, 0x34, 0x7e, 0x55, 0x50, 0x2d, 0xdb,
	0x29, 0xaa, 0xb8, 0x4d, 0xc3, 0x1b, 0x4a, 0x42, 0xfd, 0x0a, 0x6f, 0x0e, 0x3a, 0x01, 0x89, 0x8c,
	0x6c, 0x2a, 0x80, 0xbe, 0x27, 0x50, 0x48, 0x39, 0x2d, 0x12, 0x7d, 0xe1, 0x32, 0x6e, 0xfc, 0xfe,
	0x01, 0x2e, 0xe3, 0xa7, 0xd9, 0x87, 0xc7, 0xb2, 0xd6, 0x38, 0x9b, 0x2d, 0x4d, 0x34, 0x5f, 0x9a,
	0xe8, 0x63, 0x69, 0xa2, 0xfb, 0x95, 0xa9, 0xcd, 0x57, 0xa6, 0xf6, 0xba, 0x32, 0xb5, 0xdb, 0x43,
	0x9f, 0xc9, 0xfe, 0xd8, 0xb5, 0x3c, 0x08, 0x6c, 0x12, 0x01, 0xa7, 0xea, 0xdb, 0xb5, 0xa3, 0x64,
	0x39, 0x6d, 0x39, 0x1d, 0x51, 0xe1, 0xe6, 0x92, 0x25, 0x3b, 0xf9, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0xaf, 0x0f, 0xfe, 0x22, 0xbe, 0x02, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	if m.InflationMin != nil {
		{
			size := m.InflationMin.Size()
			i -= size
			if _, err := m.InflationMin.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintMint(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.InflationMax != nil {
		{
			size := m.InflationMax.Size()
			i -= size
			if _, err := m.InflationMax.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintMint(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.BlocksPerYear != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.InflationCoef.Size()
		i -= size
		if _, err := m.InflationCoef.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.InflationCoef.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovMint(uint64(m.BlocksPerYear))
	}
	if m.InflationMax != nil {
		l = m.InflationMax.Size()
		n += 1 + l + sovMint(uint64(l))
	}
	if m.InflationMin != nil {
		l = m.InflationMin.Size()
		n += 1 + l + sovMint(uint64(l))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationCoef", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationCoef.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.LegacyDec
			m.InflationMax = &v
			if err := m.InflationMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.LegacyDec
			m.InflationMin = &v
			if err := m.InflationMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
