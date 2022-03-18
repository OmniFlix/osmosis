// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/lockup/lock.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type LockQueryType int32

const (
	ByDuration LockQueryType = 0
	ByTime     LockQueryType = 1
)

var LockQueryType_name = map[int32]string{
	0: "ByDuration",
	1: "ByTime",
}

var LockQueryType_value = map[string]int32{
	"ByDuration": 0,
	"ByTime":     1,
}

func (x LockQueryType) String() string {
	return proto.EnumName(LockQueryType_name, int32(x))
}

func (LockQueryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7e9d7527a237b489, []int{0}
}

// PeriodLock is a single unit of lock by period. It's a record of locked coin
// at a specific time. It stores owner, duration, unlock time and the amount of
// coins locked.
type PeriodLock struct {
	ID       uint64                                   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Owner    string                                   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	Duration time.Duration                            `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	EndTime  time.Time                                `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	Coins    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *PeriodLock) Reset()         { *m = PeriodLock{} }
func (m *PeriodLock) String() string { return proto.CompactTextString(m) }
func (*PeriodLock) ProtoMessage()    {}
func (*PeriodLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9d7527a237b489, []int{0}
}
func (m *PeriodLock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeriodLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeriodLock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeriodLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeriodLock.Merge(m, src)
}
func (m *PeriodLock) XXX_Size() int {
	return m.Size()
}
func (m *PeriodLock) XXX_DiscardUnknown() {
	xxx_messageInfo_PeriodLock.DiscardUnknown(m)
}

var xxx_messageInfo_PeriodLock proto.InternalMessageInfo

func (m *PeriodLock) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PeriodLock) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *PeriodLock) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *PeriodLock) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *PeriodLock) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

type QueryCondition struct {
	// type of lock query, ByLockDuration | ByLockTime
	LockQueryType LockQueryType `protobuf:"varint,1,opt,name=lock_query_type,json=lockQueryType,proto3,enum=osmosis.lockup.LockQueryType" json:"lock_query_type,omitempty"`
	// What token denomination are we looking for lockups of
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	// valid when query condition is ByDuration
	Duration time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration" yaml:"duration"`
	// valid when query condition is ByTime
	Timestamp time.Time `protobuf:"bytes,4,opt,name=timestamp,proto3,stdtime" json:"timestamp" yaml:"timestamp"`
}

func (m *QueryCondition) Reset()         { *m = QueryCondition{} }
func (m *QueryCondition) String() string { return proto.CompactTextString(m) }
func (*QueryCondition) ProtoMessage()    {}
func (*QueryCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9d7527a237b489, []int{1}
}
func (m *QueryCondition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCondition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCondition.Merge(m, src)
}
func (m *QueryCondition) XXX_Size() int {
	return m.Size()
}
func (m *QueryCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCondition.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCondition proto.InternalMessageInfo

func (m *QueryCondition) GetLockQueryType() LockQueryType {
	if m != nil {
		return m.LockQueryType
	}
	return ByDuration
}

func (m *QueryCondition) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *QueryCondition) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *QueryCondition) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

// SyntheticLock is a single unit of synthetic lockup
// TODO: Change this to have
// * underlying_lock_id
// * synthetic_coin
// * end_time
// * duration
// * owner
// We then index synthetic locks by the denom, just like we do with normal
// locks. Ideally we even get an interface, so we can re-use that same logic.
// I currently have no idea how reward distribution is supposed to be working...
// EVENTUALLY
// we make a "constrained_coin" field, which is what the current "coins" field
// is. Constrained coin field can be a #post-v7 feature, since we aren't
// allowing partial unlocks of synthetic lockups.
type SyntheticLock struct {
	// underlying native lockup id for this synthetic lockup
	UnderlyingLockID uint64 `protobuf:"varint,1,opt,name=UnderlyingLockID,proto3" json:"UnderlyingLockID,omitempty"`
	SynthDenom       string `protobuf:"bytes,2,opt,name=synthDenom,proto3" json:"synthDenom,omitempty"`
	// used for unbonding synthetic lockups, for active synthetic lockups, this
	// value is set to uninitialized value
	EndTime  time.Time     `protobuf:"bytes,3,opt,name=endTime,proto3,stdtime" json:"endTime" yaml:"end_time"`
	Duration time.Duration `protobuf:"bytes,4,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
}

func (m *SyntheticLock) Reset()         { *m = SyntheticLock{} }
func (m *SyntheticLock) String() string { return proto.CompactTextString(m) }
func (*SyntheticLock) ProtoMessage()    {}
func (*SyntheticLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9d7527a237b489, []int{2}
}
func (m *SyntheticLock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyntheticLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyntheticLock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SyntheticLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyntheticLock.Merge(m, src)
}
func (m *SyntheticLock) XXX_Size() int {
	return m.Size()
}
func (m *SyntheticLock) XXX_DiscardUnknown() {
	xxx_messageInfo_SyntheticLock.DiscardUnknown(m)
}

var xxx_messageInfo_SyntheticLock proto.InternalMessageInfo

func (m *SyntheticLock) GetUnderlyingLockID() uint64 {
	if m != nil {
		return m.UnderlyingLockID
	}
	return 0
}

func (m *SyntheticLock) GetSynthDenom() string {
	if m != nil {
		return m.SynthDenom
	}
	return ""
}

func (m *SyntheticLock) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *SyntheticLock) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterEnum("osmosis.lockup.LockQueryType", LockQueryType_name, LockQueryType_value)
	proto.RegisterType((*PeriodLock)(nil), "osmosis.lockup.PeriodLock")
	proto.RegisterType((*QueryCondition)(nil), "osmosis.lockup.QueryCondition")
	proto.RegisterType((*SyntheticLock)(nil), "osmosis.lockup.SyntheticLock")
}

func init() { proto.RegisterFile("osmosis/lockup/lock.proto", fileDescriptor_7e9d7527a237b489) }

var fileDescriptor_7e9d7527a237b489 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xb6, 0xdd, 0xa4, 0xb4, 0x0f, 0x92, 0x46, 0xa7, 0x0e, 0x69, 0x00, 0x3b, 0xf2, 0x80, 0xa2,
	0x8a, 0xda, 0xa4, 0x0c, 0x48, 0x8c, 0x6e, 0x18, 0x2a, 0x18, 0xa8, 0x29, 0x0c, 0x2c, 0x91, 0x7f,
	0x1c, 0x8e, 0x15, 0xdb, 0x67, 0x6c, 0xa7, 0xe0, 0xff, 0x80, 0xb1, 0x63, 0x99, 0xd9, 0xf8, 0x4b,
	0x3a, 0x76, 0x64, 0x4a, 0x51, 0x22, 0x16, 0xc6, 0xfe, 0x05, 0xe8, 0xee, 0x7c, 0x69, 0xda, 0x0a,
	0xa9, 0x12, 0x62, 0x3a, 0xbf, 0xf7, 0xee, 0xfb, 0xde, 0xbb, 0xef, 0x7d, 0x32, 0x6c, 0x91, 0x3c,
	0x26, 0x79, 0x98, 0x9b, 0x11, 0xf1, 0xc6, 0x93, 0x94, 0x1d, 0x46, 0x9a, 0x91, 0x82, 0xa0, 0x66,
	0x55, 0x32, 0x78, 0xa9, 0xb3, 0x19, 0x90, 0x80, 0xb0, 0x92, 0x49, 0xbf, 0xf8, 0xad, 0x8e, 0x1a,
	0x10, 0x12, 0x44, 0xd8, 0x64, 0x91, 0x3b, 0xf9, 0x60, 0xfa, 0x93, 0xcc, 0x29, 0x42, 0x92, 0x54,
	0x75, 0xed, 0x7a, 0xbd, 0x08, 0x63, 0x9c, 0x17, 0x4e, 0x9c, 0x0a, 0x02, 0x8f, 0xf5, 0x31, 0x5d,
	0x27, 0xc7, 0xe6, 0x51, 0xdf, 0xc5, 0x85, 0xd3, 0x37, 0x3d, 0x12, 0x56, 0x04, 0xfa, 0x2f, 0x05,
	0xe0, 0x35, 0xce, 0x42, 0xe2, 0xbf, 0x22, 0xde, 0x18, 0x35, 0x41, 0xd9, 0x1f, 0xb4, 0xe5, 0xae,
	0xdc, 0xab, 0xd9, 0xca, 0xfe, 0x00, 0x3d, 0x82, 0x3a, 0xf9, 0x94, 0xe0, 0xac, 0xad, 0x74, 0xe5,
	0xde, 0xba, 0xd5, 0xba, 0x98, 0x6a, 0xf7, 0x4a, 0x27, 0x8e, 0x9e, 0xeb, 0x2c, 0xad, 0xdb, 0xbc,
	0x8c, 0x46, 0xb0, 0x26, 0x26, 0x6b, 0xaf, 0x74, 0xe5, 0xde, 0xdd, 0xdd, 0x2d, 0x83, 0x8f, 0x66,
	0x88, 0xd1, 0x8c, 0x41, 0x75, 0xc1, 0xea, 0x9f, 0x4e, 0x35, 0xe9, 0xf7, 0x54, 0x43, 0x02, 0xf2,
	0x98, 0xc4, 0x61, 0x81, 0xe3, 0xb4, 0x28, 0x2f, 0xa6, 0xda, 0x06, 0xe7, 0x17, 0x35, 0xfd, 0xe4,
	0x5c, 0x93, 0xed, 0x05, 0x3b, 0xb2, 0x61, 0x0d, 0x27, 0xfe, 0x90, 0xbe, 0xb3, 0x5d, 0x63, 0x9d,
	0x3a, 0x37, 0x3a, 0x1d, 0x0a, 0x11, 0xac, 0xfb, 0xb4, 0xd5, 0x25, 0xa9, 0x40, 0xea, 0xc7, 0x94,
	0xf4, 0x0e, 0x4e, 0x7c, 0x7a, 0x15, 0x39, 0x50, 0xa7, 0x92, 0xe4, 0xed, 0x7a, 0x77, 0x85, 0x8d,
	0xce, 0x45, 0x33, 0xa8, 0x68, 0x46, 0x25, 0x9a, 0xb1, 0x47, 0xc2, 0xc4, 0x7a, 0x42, 0xf9, 0xbe,
	0x9f, 0x6b, 0xbd, 0x20, 0x2c, 0x46, 0x13, 0xd7, 0xf0, 0x48, 0x6c, 0x56, 0x0a, 0xf3, 0x63, 0x27,
	0xf7, 0xc7, 0x66, 0x51, 0xa6, 0x38, 0x67, 0x80, 0xdc, 0xe6, 0xcc, 0xfa, 0x57, 0x05, 0x9a, 0x07,
	0x13, 0x9c, 0x95, 0x7b, 0x24, 0xf1, 0x43, 0xf6, 0x92, 0x17, 0xb0, 0x41, 0x77, 0x3f, 0xfc, 0x48,
	0xd3, 0x43, 0x8a, 0x61, 0xc2, 0x37, 0x77, 0x1f, 0x1a, 0x57, 0xbd, 0x61, 0xd0, 0xd5, 0x30, 0xf0,
	0x61, 0x99, 0x62, 0xbb, 0x11, 0x2d, 0x87, 0x68, 0x13, 0xea, 0x3e, 0x4e, 0x48, 0xcc, 0x57, 0x64,
	0xf3, 0x80, 0xca, 0x74, 0xfb, 0x85, 0x5c, 0x53, 0xe9, 0x6f, 0xd2, 0xbf, 0x83, 0xf5, 0x85, 0xbd,
	0x6e, 0xa1, 0xfd, 0x83, 0x8a, 0xb5, 0xc5, 0x59, 0x17, 0x50, 0x2e, 0xfe, 0x25, 0x95, 0x7e, 0xa2,
	0x40, 0xe3, 0x4d, 0x99, 0x14, 0x23, 0x5c, 0x84, 0x1e, 0xb3, 0xe1, 0x36, 0xb4, 0xde, 0x26, 0x3e,
	0xce, 0xa2, 0x32, 0x4c, 0x02, 0x9a, 0x59, 0x98, 0xf2, 0x46, 0x1e, 0xa9, 0x00, 0x39, 0x05, 0x0f,
	0x96, 0x44, 0x58, 0xca, 0xa0, 0x03, 0x10, 0x7b, 0xae, 0x84, 0xf8, 0x77, 0xbf, 0x2c, 0xbb, 0xbd,
	0xf6, 0x3f, 0xdd, 0xbe, 0xdd, 0x87, 0xc6, 0x95, 0xe5, 0xa3, 0x26, 0x80, 0x55, 0x0a, 0xee, 0x96,
	0x84, 0x00, 0x56, 0xad, 0x92, 0x0e, 0xd5, 0x92, 0x3b, 0xb5, 0x2f, 0xdf, 0x54, 0xc9, 0x7a, 0x79,
	0x3a, 0x53, 0xe5, 0xb3, 0x99, 0x2a, 0xff, 0x9c, 0xa9, 0xf2, 0xf1, 0x5c, 0x95, 0xce, 0xe6, 0xaa,
	0xf4, 0x63, 0xae, 0x4a, 0xef, 0xfb, 0x4b, 0xa6, 0xad, 0x1c, 0xb6, 0x13, 0x39, 0x6e, 0x2e, 0x02,
	0xf3, 0xe8, 0x99, 0xf9, 0x59, 0xfc, 0xaa, 0x98, 0x87, 0xdd, 0x55, 0xf6, 0x9e, 0xa7, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0x73, 0xd1, 0xc8, 0xc9, 0x04, 0x00, 0x00,
}

func (m *PeriodLock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeriodLock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeriodLock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLock(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintLock(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLock(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCondition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCondition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCondition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintLock(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	n4, err4 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintLock(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x1a
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintLock(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.LockQueryType != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.LockQueryType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SyntheticLock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyntheticLock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SyntheticLock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n5, err5 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintLock(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x22
	n6, err6 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVarintLock(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0x1a
	if len(m.SynthDenom) > 0 {
		i -= len(m.SynthDenom)
		copy(dAtA[i:], m.SynthDenom)
		i = encodeVarintLock(dAtA, i, uint64(len(m.SynthDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.UnderlyingLockID != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.UnderlyingLockID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLock(dAtA []byte, offset int, v uint64) int {
	offset -= sovLock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PeriodLock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLock(uint64(m.ID))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovLock(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovLock(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovLock(uint64(l))
		}
	}
	return n
}

func (m *QueryCondition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LockQueryType != 0 {
		n += 1 + sovLock(uint64(m.LockQueryType))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovLock(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovLock(uint64(l))
	return n
}

func (m *SyntheticLock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UnderlyingLockID != 0 {
		n += 1 + sovLock(uint64(m.UnderlyingLockID))
	}
	l = len(m.SynthDenom)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovLock(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovLock(uint64(l))
	return n
}

func sovLock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLock(x uint64) (n int) {
	return sovLock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PeriodLock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
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
			return fmt.Errorf("proto: PeriodLock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeriodLock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
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
func (m *QueryCondition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
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
			return fmt.Errorf("proto: QueryCondition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCondition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockQueryType", wireType)
			}
			m.LockQueryType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockQueryType |= LockQueryType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
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
func (m *SyntheticLock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
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
			return fmt.Errorf("proto: SyntheticLock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyntheticLock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnderlyingLockID", wireType)
			}
			m.UnderlyingLockID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnderlyingLockID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SynthDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SynthDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
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
func skipLock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLock
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
					return 0, ErrIntOverflowLock
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
					return 0, ErrIntOverflowLock
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
				return 0, ErrInvalidLengthLock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLock = fmt.Errorf("proto: unexpected end of group")
)
