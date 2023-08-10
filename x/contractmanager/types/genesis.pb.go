// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/contractmanager/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
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

// Failure message contains information about ACK failures and can be used to
// replay ACK in case of requirement.
// Note that Failure means that sudo handler to cosmwasm contract failed for some reason
type Failure struct {
	// ChannelId
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Address of the failed contract
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Id of the failure under specific address
	Id uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// Packet sequence id to restore
	SequenceId uint64 `protobuf:"varint,4,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	// Acknowledgement type
	AckType string `protobuf:"bytes,5,opt,name=ack_type,json=ackType,proto3" json:"ack_type,omitempty"`
	// IBC Packet
	Packet *types.Packet `protobuf:"bytes,6,opt,name=packet,proto3" json:"packet,omitempty"`
	// Acknowledgement
	Ack *types.Acknowledgement `protobuf:"bytes,7,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}
func (*Failure) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf4a1534315a7490, []int{0}
}
func (m *Failure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Failure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Failure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Failure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failure.Merge(m, src)
}
func (m *Failure) XXX_Size() int {
	return m.Size()
}
func (m *Failure) XXX_DiscardUnknown() {
	xxx_messageInfo_Failure.DiscardUnknown(m)
}

var xxx_messageInfo_Failure proto.InternalMessageInfo

func (m *Failure) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *Failure) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Failure) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Failure) GetSequenceId() uint64 {
	if m != nil {
		return m.SequenceId
	}
	return 0
}

func (m *Failure) GetAckType() string {
	if m != nil {
		return m.AckType
	}
	return ""
}

func (m *Failure) GetPacket() *types.Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *Failure) GetAck() *types.Acknowledgement {
	if m != nil {
		return m.Ack
	}
	return nil
}

// GenesisState defines the contractmanager module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// List of the contract failures
	FailuresList []Failure `protobuf:"bytes,2,rep,name=failures_list,json=failuresList,proto3" json:"failures_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf4a1534315a7490, []int{1}
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

func (m *GenesisState) GetFailuresList() []Failure {
	if m != nil {
		return m.FailuresList
	}
	return nil
}

func init() {
	proto.RegisterType((*Failure)(nil), "neutron.contractmanager.Failure")
	proto.RegisterType((*GenesisState)(nil), "neutron.contractmanager.GenesisState")
}

func init() {
	proto.RegisterFile("neutron/contractmanager/genesis.proto", fileDescriptor_cf4a1534315a7490)
}

var fileDescriptor_cf4a1534315a7490 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0xe3, 0x24, 0x24, 0xd4, 0x29, 0x0c, 0x16, 0x12, 0x47, 0x10, 0x97, 0xa3, 0x2a, 0x52,
	0x16, 0x6c, 0x35, 0x95, 0xba, 0x31, 0xd0, 0x01, 0x54, 0xc1, 0x10, 0x1d, 0x4c, 0x2c, 0x91, 0x63,
	0x3f, 0xae, 0xd6, 0x25, 0xf6, 0x61, 0x3b, 0x85, 0x7e, 0x0b, 0x66, 0x3e, 0x51, 0xc7, 0x8e, 0x4c,
	0x08, 0x25, 0x1f, 0x83, 0x05, 0x9d, 0xcf, 0xb7, 0x40, 0xb3, 0xbd, 0x7b, 0xef, 0xf7, 0xff, 0xdf,
	0xbd, 0xff, 0x3b, 0xfc, 0x42, 0xc3, 0xc6, 0x5b, 0xa3, 0x99, 0x30, 0xda, 0x5b, 0x2e, 0xfc, 0x9a,
	0x6b, 0x5e, 0x80, 0x65, 0x05, 0x68, 0x70, 0xca, 0xd1, 0xca, 0x1a, 0x6f, 0xc8, 0xe3, 0x88, 0xd1,
	0x7f, 0xb0, 0xf1, 0xa3, 0xc2, 0x14, 0x26, 0x30, 0xac, 0xae, 0x1a, 0x7c, 0x7c, 0xbc, 0xcf, 0xb5,
	0xe2, 0x96, 0xaf, 0xa3, 0xe9, 0xf8, 0xb9, 0x5a, 0x0a, 0x26, 0x8c, 0x05, 0x26, 0x2e, 0xb9, 0xd6,
	0xb0, 0x62, 0x57, 0x27, 0x6d, 0xd9, 0x20, 0x47, 0x7f, 0x10, 0x1e, 0xbe, 0xe1, 0x6a, 0xb5, 0xb1,
	0x40, 0x9e, 0x61, 0x1c, 0x87, 0x0b, 0x25, 0x13, 0x94, 0xa1, 0xe9, 0x41, 0x7e, 0x10, 0x3b, 0x17,
	0x92, 0x24, 0x78, 0xc8, 0xa5, 0xb4, 0xe0, 0x5c, 0xd2, 0x0d, 0xb3, 0xf6, 0x91, 0x3c, 0xc4, 0x5d,
	0x25, 0x93, 0x5e, 0x86, 0xa6, 0xfd, 0xbc, 0xab, 0x24, 0x99, 0xe0, 0x91, 0x83, 0x2f, 0x1b, 0xd0,
	0x02, 0x6a, 0xa7, 0x7e, 0x18, 0xe0, 0xb6, 0x75, 0x21, 0xc9, 0x13, 0x7c, 0x9f, 0x8b, 0x72, 0xe1,
	0xaf, 0x2b, 0x48, 0xee, 0x45, 0x2f, 0x51, 0x7e, 0xbc, 0xae, 0x80, 0x9c, 0xe2, 0x41, 0xc5, 0x45,
	0x09, 0x3e, 0x19, 0x64, 0x68, 0x3a, 0x9a, 0x3d, 0xa5, 0x6a, 0x29, 0x68, 0xbd, 0x04, 0x6d, 0xbf,
	0xfc, 0xea, 0x84, 0xce, 0x03, 0x92, 0x47, 0x94, 0x9c, 0xe1, 0x1e, 0x17, 0x65, 0x32, 0x0c, 0x8a,
	0xe3, 0x3b, 0x15, 0xaf, 0x45, 0xa9, 0xcd, 0xd7, 0x15, 0xc8, 0x02, 0xd6, 0xa0, 0x7d, 0x5e, 0x0b,
	0x8e, 0x7e, 0x20, 0x7c, 0xf8, 0xb6, 0xb9, 0xc3, 0x07, 0xcf, 0x3d, 0x90, 0x57, 0xf5, 0xdb, 0xeb,
	0x04, 0xc3, 0xfa, 0xa3, 0xd9, 0x84, 0xee, 0xb9, 0x0b, 0x9d, 0x07, 0xec, 0xbc, 0x7f, 0xf3, 0x6b,
	0xd2, 0xc9, 0xa3, 0x88, 0xbc, 0xc3, 0x0f, 0x3e, 0x37, 0x61, 0xba, 0xc5, 0x4a, 0x39, 0x9f, 0x74,
	0xb3, 0xde, 0x74, 0x34, 0xcb, 0xf6, 0xba, 0xc4, 0xe8, 0xa3, 0xcd, 0x61, 0x2b, 0x7e, 0xaf, 0x9c,
	0x3f, 0x9f, 0xdf, 0x6c, 0x53, 0x74, 0xbb, 0x4d, 0xd1, 0xef, 0x6d, 0x8a, 0xbe, 0xef, 0xd2, 0xce,
	0xed, 0x2e, 0xed, 0xfc, 0xdc, 0xa5, 0x9d, 0x4f, 0x67, 0x85, 0xf2, 0x97, 0x9b, 0x25, 0x15, 0x66,
	0xcd, 0xa2, 0xf3, 0x4b, 0x63, 0x8b, 0xb6, 0x66, 0xdf, 0xfe, 0xfb, 0x2d, 0xea, 0xa4, 0xdd, 0x72,
	0x10, 0x6e, 0x7e, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x4c, 0x8c, 0x41, 0x94, 0x02, 0x00,
	0x00,
}

func (m *Failure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Failure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Failure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ack != nil {
		{
			size, err := m.Ack.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Packet != nil {
		{
			size, err := m.Packet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.AckType) > 0 {
		i -= len(m.AckType)
		copy(dAtA[i:], m.AckType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AckType)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SequenceId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SequenceId))
		i--
		dAtA[i] = 0x20
	}
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if len(m.FailuresList) > 0 {
		for iNdEx := len(m.FailuresList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FailuresList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *Failure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	if m.SequenceId != 0 {
		n += 1 + sovGenesis(uint64(m.SequenceId))
	}
	l = len(m.AckType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Packet != nil {
		l = m.Packet.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Ack != nil {
		l = m.Ack.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FailuresList) > 0 {
		for _, e := range m.FailuresList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Failure) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Failure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Failure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceId", wireType)
			}
			m.SequenceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Packet", wireType)
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
			if m.Packet == nil {
				m.Packet = &types.Packet{}
			}
			if err := m.Packet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ack", wireType)
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
			if m.Ack == nil {
				m.Ack = &types.Acknowledgement{}
			}
			if err := m.Ack.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
				return fmt.Errorf("proto: wrong wireType = %d for field FailuresList", wireType)
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
			m.FailuresList = append(m.FailuresList, Failure{})
			if err := m.FailuresList[len(m.FailuresList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
