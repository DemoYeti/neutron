// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the dex module's genesis state.
type GenesisState struct {
	Params                        Params                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TickLiquidityList             []*TickLiquidity         `protobuf:"bytes,2,rep,name=tickLiquidityList,proto3" json:"tickLiquidityList,omitempty"`
	InactiveLimitOrderTrancheList []*LimitOrderTranche     `protobuf:"bytes,3,rep,name=inactiveLimitOrderTrancheList,proto3" json:"inactiveLimitOrderTrancheList,omitempty"`
	LimitOrderTrancheUserList     []*LimitOrderTrancheUser `protobuf:"bytes,4,rep,name=limitOrderTrancheUserList,proto3" json:"limitOrderTrancheUserList,omitempty"`
	PoolMetadataList              []PoolMetadata           `protobuf:"bytes,5,rep,name=poolMetadataList,proto3" json:"poolMetadataList"`
	PoolCount                     uint64                   `protobuf:"varint,6,opt,name=poolCount,proto3" json:"poolCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ccf50bb6779a67a, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTickLiquidityList() []*TickLiquidity {
	if m != nil {
		return m.TickLiquidityList
	}
	return nil
}

func (m *GenesisState) GetInactiveLimitOrderTrancheList() []*LimitOrderTranche {
	if m != nil {
		return m.InactiveLimitOrderTrancheList
	}
	return nil
}

func (m *GenesisState) GetLimitOrderTrancheUserList() []*LimitOrderTrancheUser {
	if m != nil {
		return m.LimitOrderTrancheUserList
	}
	return nil
}

func (m *GenesisState) GetPoolMetadataList() []PoolMetadata {
	if m != nil {
		return m.PoolMetadataList
	}
	return nil
}

func (m *GenesisState) GetPoolCount() uint64 {
	if m != nil {
		return m.PoolCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "duality.dex.GenesisState")
}

func init() { proto.RegisterFile("duality/dex/genesis.proto", fileDescriptor_6ccf50bb6779a67a) }

var fileDescriptor_6ccf50bb6779a67a = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0xdb, 0x0b, 0x97, 0xe4, 0x0e, 0x77, 0x71, 0x6f, 0x75, 0x51, 0x1a, 0x2d, 0x0d, 0x89,
	0x09, 0xd1, 0xd8, 0x46, 0x7c, 0x03, 0x8c, 0x71, 0x21, 0xfe, 0x09, 0xe2, 0xc6, 0x4d, 0x33, 0xb4,
	0x63, 0x19, 0x6d, 0x3b, 0xd8, 0x99, 0x1a, 0x78, 0x0b, 0x1f, 0x8b, 0x25, 0x4b, 0x57, 0xc6, 0xc0,
	0x8b, 0x98, 0xf9, 0x43, 0x6c, 0x25, 0xe2, 0xae, 0x3d, 0xe7, 0x77, 0xbe, 0xef, 0x7c, 0x27, 0x03,
	0x1a, 0x61, 0x0e, 0x63, 0xcc, 0xa6, 0x5e, 0x88, 0x26, 0x5e, 0x84, 0x52, 0x44, 0x31, 0x75, 0xc7,
	0x19, 0x61, 0xc4, 0xa8, 0xab, 0x96, 0x1b, 0xa2, 0x89, 0xb5, 0x1d, 0x91, 0x88, 0x88, 0xba, 0xc7,
	0xbf, 0x24, 0x62, 0x99, 0xc5, 0xe9, 0x31, 0xcc, 0x60, 0xa2, 0x86, 0xad, 0xfd, 0x62, 0x27, 0xc6,
	0x09, 0x66, 0x3e, 0xc9, 0x42, 0x94, 0xf9, 0x2c, 0x83, 0x69, 0x30, 0x42, 0x7e, 0x4e, 0x51, 0xa6,
	0xd8, 0xbd, 0x1f, 0x58, 0x85, 0x39, 0x45, 0x8c, 0xe1, 0xe0, 0xd1, 0x8f, 0xf1, 0x53, 0x8e, 0x43,
	0xbe, 0xa2, 0x24, 0x9a, 0xa5, 0x75, 0x08, 0x89, 0xfd, 0x04, 0x31, 0x18, 0x42, 0x06, 0x25, 0xd0,
	0x9a, 0x57, 0xc0, 0xdf, 0x33, 0x19, 0xf2, 0x86, 0x41, 0x86, 0x8c, 0x23, 0x50, 0x93, 0x6b, 0x9b,
	0xba, 0xa3, 0xb7, 0xeb, 0x9d, 0x2d, 0xb7, 0x10, 0xda, 0xbd, 0x16, 0xad, 0x6e, 0x75, 0xf6, 0xd6,
	0xd4, 0xfa, 0x0a, 0x34, 0x2e, 0xc1, 0x7f, 0x6e, 0xde, 0x5b, 0x79, 0xf7, 0x30, 0x65, 0xe6, 0x2f,
	0xa7, 0xd2, 0xae, 0x77, 0xac, 0xd2, 0xf4, 0xa0, 0x48, 0x09, 0x11, 0xbd, 0xbf, 0x3e, 0x6a, 0x3c,
	0x80, 0x5d, 0x9c, 0xc2, 0x80, 0xe1, 0x67, 0xd4, 0xe3, 0xd9, 0xaf, 0x78, 0xf4, 0x81, 0x4c, 0x2e,
	0xb4, 0x2b, 0x42, 0xdb, 0x2e, 0x69, 0xaf, 0x91, 0x4a, 0x7f, 0xb3, 0x94, 0x71, 0x0f, 0x1a, 0xf1,
	0xd7, 0xc6, 0x2d, 0x45, 0x99, 0xf0, 0xa9, 0x0a, 0x9f, 0xd6, 0x66, 0x1f, 0x4e, 0x2b, 0xaf, 0xef,
	0xa5, 0x8c, 0x73, 0xf0, 0x8f, 0x9f, 0xff, 0x42, 0x5d, 0x5f, 0xc8, 0xff, 0x16, 0xf2, 0x8d, 0xf2,
	0x81, 0x0b, 0x90, 0x3a, 0xf3, 0xda, 0xa0, 0xb1, 0x03, 0xfe, 0xf0, 0xda, 0x09, 0xc9, 0x53, 0x66,
	0xd6, 0x1c, 0xbd, 0x5d, 0xed, 0x7f, 0x16, 0xba, 0xa7, 0xb3, 0x85, 0xad, 0xcf, 0x17, 0xb6, 0xfe,
	0xbe, 0xb0, 0xf5, 0x97, 0xa5, 0xad, 0xcd, 0x97, 0xb6, 0xf6, 0xba, 0xb4, 0xb5, 0xbb, 0x83, 0x08,
	0xb3, 0x51, 0x3e, 0x74, 0x03, 0x92, 0x78, 0xca, 0xf4, 0x30, 0x86, 0x43, 0xba, 0xfa, 0xf1, 0x26,
	0xf2, 0x25, 0x4d, 0xc7, 0x88, 0x0e, 0x6b, 0xe2, 0x81, 0x1c, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xb6, 0x99, 0x5f, 0x53, 0x10, 0x03, 0x00, 0x00,
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
	if m.PoolCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PoolCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PoolMetadataList) > 0 {
		for iNdEx := len(m.PoolMetadataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolMetadataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LimitOrderTrancheUserList) > 0 {
		for iNdEx := len(m.LimitOrderTrancheUserList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderTrancheUserList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InactiveLimitOrderTrancheList) > 0 {
		for iNdEx := len(m.InactiveLimitOrderTrancheList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InactiveLimitOrderTrancheList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TickLiquidityList) > 0 {
		for iNdEx := len(m.TickLiquidityList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TickLiquidityList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TickLiquidityList) > 0 {
		for _, e := range m.TickLiquidityList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.InactiveLimitOrderTrancheList) > 0 {
		for _, e := range m.InactiveLimitOrderTrancheList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LimitOrderTrancheUserList) > 0 {
		for _, e := range m.LimitOrderTrancheUserList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PoolMetadataList) > 0 {
		for _, e := range m.PoolMetadataList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PoolCount != 0 {
		n += 1 + sovGenesis(uint64(m.PoolCount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickLiquidityList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TickLiquidityList = append(m.TickLiquidityList, &TickLiquidity{})
			if err := m.TickLiquidityList[len(m.TickLiquidityList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveLimitOrderTrancheList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InactiveLimitOrderTrancheList = append(m.InactiveLimitOrderTrancheList, &LimitOrderTranche{})
			if err := m.InactiveLimitOrderTrancheList[len(m.InactiveLimitOrderTrancheList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderTrancheUserList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LimitOrderTrancheUserList = append(m.LimitOrderTrancheUserList, &LimitOrderTrancheUser{})
			if err := m.LimitOrderTrancheUserList[len(m.LimitOrderTrancheUserList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolMetadataList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolMetadataList = append(m.PoolMetadataList, PoolMetadata{})
			if err := m.PoolMetadataList[len(m.PoolMetadataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCount", wireType)
			}
			m.PoolCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)