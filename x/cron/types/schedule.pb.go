// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/cron/schedule.proto

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

type Schedule struct {
	// Name of schedule
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Period in blocks
	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// Msgs that will be executed every period amount of time
	Msgs []MsgExecuteContract `protobuf:"bytes,3,rep,name=msgs,proto3" json:"msgs"`
	// Last execution's block height
	LastExecuteHeight uint64 `protobuf:"varint,4,opt,name=last_execute_height,json=lastExecuteHeight,proto3" json:"last_execute_height,omitempty"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ace1b59de613ef, []int{0}
}
func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(m, src)
}
func (m *Schedule) XXX_Size() int {
	return m.Size()
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

func (m *Schedule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Schedule) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *Schedule) GetMsgs() []MsgExecuteContract {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func (m *Schedule) GetLastExecuteHeight() uint64 {
	if m != nil {
		return m.LastExecuteHeight
	}
	return 0
}

type MsgExecuteContract struct {
	// Contract is the address of the smart contract
	Contract string `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	// Msg is json encoded message to be passed to the contract
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *MsgExecuteContract) Reset()         { *m = MsgExecuteContract{} }
func (m *MsgExecuteContract) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteContract) ProtoMessage()    {}
func (*MsgExecuteContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ace1b59de613ef, []int{1}
}
func (m *MsgExecuteContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteContract.Merge(m, src)
}
func (m *MsgExecuteContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteContract proto.InternalMessageInfo

func (m *MsgExecuteContract) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *MsgExecuteContract) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ScheduleCount struct {
	// Count is the number of current schedules
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *ScheduleCount) Reset()         { *m = ScheduleCount{} }
func (m *ScheduleCount) String() string { return proto.CompactTextString(m) }
func (*ScheduleCount) ProtoMessage()    {}
func (*ScheduleCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ace1b59de613ef, []int{2}
}
func (m *ScheduleCount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScheduleCount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduleCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleCount.Merge(m, src)
}
func (m *ScheduleCount) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleCount) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleCount.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleCount proto.InternalMessageInfo

func (m *ScheduleCount) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Schedule)(nil), "neutron.cron.Schedule")
	proto.RegisterType((*MsgExecuteContract)(nil), "neutron.cron.MsgExecuteContract")
	proto.RegisterType((*ScheduleCount)(nil), "neutron.cron.ScheduleCount")
}

func init() { proto.RegisterFile("neutron/cron/schedule.proto", fileDescriptor_49ace1b59de613ef) }

var fileDescriptor_49ace1b59de613ef = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x9b, 0xff, 0xba, 0xb1, 0xe5, 0xaf, 0xa0, 0x71, 0x48, 0x99, 0x10, 0x4b, 0x41, 0xe8,
	0xc5, 0x14, 0xd4, 0x93, 0xc7, 0x0e, 0x61, 0x17, 0x2f, 0xf5, 0xe6, 0x65, 0x74, 0x59, 0x48, 0x0b,
	0x6b, 0x53, 0x9a, 0x54, 0xe6, 0xb7, 0xf0, 0x33, 0xf8, 0x69, 0x76, 0xdc, 0xd1, 0x93, 0x48, 0xfb,
	0x45, 0x24, 0x69, 0x26, 0x82, 0x97, 0xf0, 0x7b, 0x78, 0xde, 0x27, 0xef, 0xfb, 0x26, 0xf0, 0xa2,
	0x64, 0x8d, 0xaa, 0x45, 0x19, 0x51, 0x7d, 0x48, 0x9a, 0xb1, 0x75, 0xb3, 0x61, 0xa4, 0xaa, 0x85,
	0x12, 0xe8, 0xc8, 0x9a, 0x44, 0x9b, 0xb3, 0x29, 0x17, 0x5c, 0x18, 0x23, 0xd2, 0xd4, 0xd7, 0x04,
	0xef, 0x00, 0x8e, 0x9f, 0x6c, 0x0c, 0x21, 0xe8, 0x96, 0x69, 0xc1, 0x3c, 0xe0, 0x83, 0x70, 0x92,
	0x18, 0x46, 0xe7, 0x70, 0x54, 0xb1, 0x3a, 0x17, 0x6b, 0xef, 0x9f, 0x0f, 0x42, 0x37, 0xb1, 0x0a,
	0xdd, 0x43, 0xb7, 0x90, 0x5c, 0x7a, 0x03, 0x7f, 0x10, 0xfe, 0xbf, 0xf1, 0xc9, 0xef, 0x5e, 0xe4,
	0x51, 0xf2, 0x87, 0x2d, 0xa3, 0x8d, 0x62, 0x73, 0x51, 0xaa, 0x3a, 0xa5, 0x2a, 0x76, 0x77, 0x9f,
	0x97, 0x4e, 0x62, 0x32, 0x88, 0xc0, 0xb3, 0x4d, 0x2a, 0xd5, 0x92, 0xf5, 0x35, 0xcb, 0x8c, 0xe5,
	0x3c, 0x53, 0x9e, 0x6b, 0x1a, 0x9c, 0x6a, 0xcb, 0xa6, 0x17, 0xc6, 0x08, 0x62, 0x88, 0xfe, 0xde,
	0x88, 0x66, 0x70, 0x4c, 0x2d, 0xdb, 0x89, 0x7f, 0x34, 0x3a, 0x81, 0x83, 0x42, 0x72, 0x33, 0xf2,
	0x24, 0xd1, 0x18, 0x5c, 0xc1, 0xe3, 0xc3, 0x9e, 0x73, 0xd1, 0x94, 0x0a, 0x4d, 0xe1, 0x90, 0x6a,
	0x30, 0xd9, 0x61, 0xd2, 0x8b, 0x78, 0xb1, 0x6b, 0x31, 0xd8, 0xb7, 0x18, 0x7c, 0xb5, 0x18, 0xbc,
	0x75, 0xd8, 0xd9, 0x77, 0xd8, 0xf9, 0xe8, 0xb0, 0xf3, 0x4c, 0x78, 0xae, 0xb2, 0x66, 0x45, 0xa8,
	0x28, 0x22, 0xbb, 0xec, 0xb5, 0xa8, 0xf9, 0x81, 0xa3, 0x97, 0xbb, 0x68, 0xdb, 0x7f, 0x83, 0x7a,
	0xad, 0x98, 0x5c, 0x8d, 0xcc, 0x03, 0xdf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x19, 0x38,
	0x8d, 0xa3, 0x01, 0x00, 0x00,
}

func (m *Schedule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schedule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schedule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastExecuteHeight != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.LastExecuteHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSchedule(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Period != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecuteContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ScheduleCount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleCount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleCount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintSchedule(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchedule(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchedule(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Schedule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	if m.Period != 0 {
		n += 1 + sovSchedule(uint64(m.Period))
	}
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovSchedule(uint64(l))
		}
	}
	if m.LastExecuteHeight != 0 {
		n += 1 + sovSchedule(uint64(m.LastExecuteHeight))
	}
	return n
}

func (m *MsgExecuteContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func (m *ScheduleCount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovSchedule(uint64(m.Count))
	}
	return n
}

func sovSchedule(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSchedule(x uint64) (n int) {
	return sovSchedule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Schedule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: Schedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, MsgExecuteContract{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastExecuteHeight", wireType)
			}
			m.LastExecuteHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastExecuteHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
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
func (m *MsgExecuteContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: MsgExecuteContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
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
func (m *ScheduleCount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: ScheduleCount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleCount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedule
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
func skipSchedule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchedule
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
					return 0, ErrIntOverflowSchedule
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
					return 0, ErrIntOverflowSchedule
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
				return 0, ErrInvalidLengthSchedule
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchedule
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchedule
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchedule        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchedule          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchedule = fmt.Errorf("proto: unexpected end of group")
)
