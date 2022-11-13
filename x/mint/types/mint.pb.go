// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params holds parameters for the mint module.
type Params struct {
	// mint_denom defines denomination of coin to be minted
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// mint_pool_address defines the address where inflation will be minted. The default is FeeCollector,
	// but if it is set to FeeCollector, minted inflation could be mixed together with collected tx fees.
	// Therefore, it is recommended to specify a separate address depending on usage.
	MintPoolAddress string `protobuf:"bytes,2,opt,name=mint_pool_address,json=mintPoolAddress,proto3" json:"mint_pool_address,omitempty"`
	// block_time_threshold defines block time threshold to prevent from any inflationary manipulation attacks
	// it is used for maximum block duration when calculating block inflation
	BlockTimeThreshold time.Duration `protobuf:"bytes,3,opt,name=block_time_threshold,json=blockTimeThreshold,proto3,stdduration" json:"block_time_threshold"`
	// inflation_schedules defines a list of inflation schedules
	InflationSchedules []InflationSchedule `protobuf:"bytes,4,rep,name=inflation_schedules,json=inflationSchedules,proto3" json:"inflation_schedules"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe08af702efa1523, []int{0}
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

func (m *Params) GetMintPoolAddress() string {
	if m != nil {
		return m.MintPoolAddress
	}
	return ""
}

func (m *Params) GetBlockTimeThreshold() time.Duration {
	if m != nil {
		return m.BlockTimeThreshold
	}
	return 0
}

func (m *Params) GetInflationSchedules() []InflationSchedule {
	if m != nil {
		return m.InflationSchedules
	}
	return nil
}

// InflationSchedule defines the start and end time of the inflation period, and the amount of inflation during that
// period.
type InflationSchedule struct {
	// start_time defines the start date time for the inflation schedule
	StartTime time.Time `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	// end_time defines the end date time for the inflation schedule
	EndTime time.Time `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	// amount defines the total amount of inflation for the schedule
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *InflationSchedule) Reset()         { *m = InflationSchedule{} }
func (m *InflationSchedule) String() string { return proto.CompactTextString(m) }
func (*InflationSchedule) ProtoMessage()    {}
func (*InflationSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe08af702efa1523, []int{1}
}
func (m *InflationSchedule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationSchedule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationSchedule.Merge(m, src)
}
func (m *InflationSchedule) XXX_Size() int {
	return m.Size()
}
func (m *InflationSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_InflationSchedule proto.InternalMessageInfo

func (m *InflationSchedule) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *InflationSchedule) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Params)(nil), "fury.mint.v1beta1.Params")
	proto.RegisterType((*InflationSchedule)(nil), "fury.mint.v1beta1.InflationSchedule")
}

func init() { proto.RegisterFile("fury/mint/v1beta1/mint.proto", fileDescriptor_fe08af702efa1523) }

var fileDescriptor_fe08af702efa1523 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0xb3, 0x69, 0x89, 0xcd, 0xe4, 0x50, 0x32, 0x56, 0x88, 0x91, 0x6e, 0x42, 0x0e, 0x12,
	0x84, 0xce, 0xd0, 0xd4, 0x93, 0x37, 0x97, 0x22, 0xf4, 0x22, 0x65, 0x8d, 0x20, 0x5e, 0x96, 0xd9,
	0x9d, 0xe9, 0x66, 0xc9, 0xce, 0xbc, 0x61, 0x67, 0xb6, 0xda, 0x8f, 0x20, 0x5e, 0x7a, 0xf4, 0x23,
	0xf5, 0xd8, 0xa3, 0x78, 0x88, 0x92, 0x7c, 0x03, 0x3f, 0x81, 0xcc, 0xec, 0xae, 0x4a, 0x2b, 0x78,
	0x4a, 0xe6, 0x79, 0x9f, 0xe7, 0xf7, 0xfe, 0x61, 0xd1, 0x38, 0x29, 0x84, 0x4e, 0x84, 0x32, 0x54,
	0x66, 0xca, 0xd0, 0xcb, 0xe3, 0x58, 0x18, 0x76, 0xec, 0x1e, 0x64, 0x55, 0x80, 0x01, 0xfc, 0xa8,
	0x71, 0x10, 0x27, 0xd6, 0x8e, 0xe1, 0x41, 0x0a, 0x29, 0x38, 0x07, 0xb5, 0xff, 0x2a, 0xf3, 0xd0,
	0x4f, 0x01, 0xd2, 0x5c, 0x50, 0xf7, 0x8a, 0xcb, 0x0b, 0xca, 0xcb, 0x82, 0x99, 0x0c, 0x54, 0x5d,
	0x1f, 0xdd, 0xad, 0x9b, 0x4c, 0x0a, 0x6d, 0x98, 0x5c, 0x55, 0x86, 0xc9, 0xe7, 0x36, 0xea, 0x9c,
	0xb3, 0x82, 0x49, 0x8d, 0x0f, 0x11, 0xb2, 0x1d, 0x23, 0x2e, 0x14, 0xc8, 0x81, 0x37, 0xf6, 0xa6,
	0xdd, 0xb0, 0x6b, 0x95, 0x53, 0x2b, 0xe0, 0x67, 0xa8, 0xef, 0xca, 0x2b, 0x80, 0x3c, 0x62, 0x9c,
	0x17, 0x42, 0xeb, 0x41, 0xdb, 0xb9, 0xf6, 0x6d, 0xe1, 0x1c, 0x20, 0x7f, 0x59, 0xc9, 0xf8, 0x2d,
	0x3a, 0x88, 0x73, 0x48, 0x96, 0x91, 0x6d, 0x17, 0x99, 0x45, 0x21, 0xf4, 0x02, 0x72, 0x3e, 0xd8,
	0x19, 0x7b, 0xd3, 0xde, 0xec, 0x31, 0xa9, 0xa6, 0x22, 0xcd, 0x54, 0xe4, 0xb4, 0x9e, 0x3a, 0xd8,
	0xbb, 0x59, 0x8f, 0x5a, 0x5f, 0xbe, 0x8f, 0xbc, 0x10, 0x3b, 0xc0, 0x3c, 0x93, 0x62, 0xde, 0xc4,
	0x71, 0x84, 0x1e, 0x66, 0xea, 0x22, 0x77, 0xd6, 0x48, 0x27, 0x0b, 0xc1, 0xcb, 0x5c, 0xe8, 0xc1,
	0xee, 0x78, 0x67, 0xda, 0x9b, 0x4d, 0xc9, 0x3f, 0x0f, 0x47, 0xce, 0x9a, 0xc4, 0x9b, 0x3a, 0x10,
	0xec, 0xda, 0x26, 0x21, 0xce, 0xee, 0x16, 0xf4, 0xe4, 0x53, 0x1b, 0xf5, 0xef, 0xf9, 0xf1, 0x3b,
	0x84, 0xb4, 0x61, 0x85, 0x71, 0xdb, 0xb8, 0xc3, 0xf4, 0x66, 0xc3, 0x7b, 0x3b, 0xcc, 0x9b, 0xcb,
	0x06, 0x87, 0x96, 0xff, 0x73, 0x3d, 0xea, 0x5f, 0x31, 0x99, 0xbf, 0x98, 0xfc, 0xc9, 0x4e, 0xae,
	0xed, 0x66, 0x5d, 0x27, 0x58, 0x3b, 0x0e, 0xd1, 0x9e, 0x50, 0xbc, 0xe2, 0xb6, 0xff, 0xcb, 0x7d,
	0x52, 0x73, 0xf7, 0x2b, 0x6e, 0x93, 0xac, 0xa8, 0x0f, 0x84, 0xe2, 0x8e, 0xf9, 0x0a, 0x75, 0x98,
	0x84, 0x52, 0x19, 0x77, 0xed, 0x6e, 0x40, 0x6c, 0xea, 0xdb, 0x7a, 0xf4, 0x34, 0xcd, 0xcc, 0xa2,
	0x8c, 0x49, 0x02, 0x92, 0x26, 0xa0, 0x25, 0xe8, 0xfa, 0xe7, 0x48, 0xf3, 0x25, 0x35, 0x57, 0x2b,
	0xa1, 0xc9, 0x99, 0x32, 0x61, 0x9d, 0x0e, 0x5e, 0xdf, 0x6c, 0x7c, 0xef, 0x76, 0xe3, 0x7b, 0x3f,
	0x36, 0xbe, 0x77, 0xbd, 0xf5, 0x5b, 0xb7, 0x5b, 0xbf, 0xf5, 0x75, 0xeb, 0xb7, 0xde, 0x3f, 0xff,
	0x9b, 0x54, 0xdf, 0xfc, 0x48, 0x09, 0xf3, 0x01, 0x8a, 0xe5, 0x6f, 0x81, 0x5e, 0x9e, 0xd0, 0x8f,
	0xd5, 0x47, 0xee, 0xd8, 0x71, 0xc7, 0x6d, 0x74, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xde, 0xcd,
	0x22, 0x06, 0x02, 0x03, 0x00, 0x00,
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
	if len(m.InflationSchedules) > 0 {
		for iNdEx := len(m.InflationSchedules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InflationSchedules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.BlockTimeThreshold, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.BlockTimeThreshold):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintMint(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.MintPoolAddress) > 0 {
		i -= len(m.MintPoolAddress)
		copy(dAtA[i:], m.MintPoolAddress)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintPoolAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InflationSchedule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationSchedule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationSchedule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintMint(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintMint(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
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
	l = len(m.MintPoolAddress)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.BlockTimeThreshold)
	n += 1 + l + sovMint(uint64(l))
	if len(m.InflationSchedules) > 0 {
		for _, e := range m.InflationSchedules {
			l = e.Size()
			n += 1 + l + sovMint(uint64(l))
		}
	}
	return n
}

func (m *InflationSchedule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovMint(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovMint(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintPoolAddress", wireType)
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
			m.MintPoolAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTimeThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.BlockTimeThreshold, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationSchedules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InflationSchedules = append(m.InflationSchedules, InflationSchedule{})
			if err := m.InflationSchedules[len(m.InflationSchedules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *InflationSchedule) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: InflationSchedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationSchedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
