// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: event.proto

package benchserder

import (
	"fmt"
	"math"
	math_bits "math/bits"
	"time"

	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Event)(nil), "github.com.narqo.benchserder.Event")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 897 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x63, 0x92, 0x36, 0xf6, 0xd8, 0xb1, 0x9d, 0x69, 0xd3, 0x0c, 0xa1, 0x64, 0x4d, 0x39,
	0xe0, 0x16, 0x75, 0x2d, 0x5a, 0x15, 0x09, 0x6e, 0x76, 0x1b, 0xa1, 0x40, 0xa1, 0xd5, 0x36, 0xf4,
	0x80, 0x90, 0x56, 0xb3, 0x3b, 0x63, 0x77, 0xb4, 0xbb, 0x33, 0xcb, 0xec, 0xd8, 0x52, 0xf8, 0x04,
	0x1c, 0xf9, 0x08, 0xe5, 0xcc, 0x17, 0xe1, 0x53, 0x04, 0x04, 0x5f, 0x80, 0x73, 0x4e, 0x68, 0xde,
	0xec, 0xda, 0x26, 0x91, 0x90, 0x97, 0xdb, 0xee, 0x7b, 0xef, 0xff, 0x9b, 0xf7, 0x66, 0xde, 0xbc,
	0x41, 0x6d, 0xbe, 0xe0, 0xd2, 0xf8, 0xb9, 0x56, 0x46, 0xe1, 0xbb, 0x33, 0x61, 0xde, 0xcc, 0x23,
	0x3f, 0x56, 0x99, 0x2f, 0xa9, 0xfe, 0x41, 0xf9, 0x11, 0x97, 0xf1, 0x9b, 0x82, 0x6b, 0xc6, 0xf5,
	0xd1, 0xc3, 0x95, 0x77, 0x34, 0x53, 0x33, 0x35, 0x02, 0x51, 0x34, 0x9f, 0xc2, 0x1f, 0xfc, 0xc0,
	0x97, 0x83, 0x1d, 0x79, 0x33, 0xa5, 0x66, 0x29, 0x5f, 0x45, 0x19, 0x91, 0xf1, 0xc2, 0xd0, 0x2c,
	0x2f, 0x03, 0x0e, 0x85, 0x34, 0x5c, 0x4b, 0x9a, 0x8e, 0x62, 0x9a, 0xa6, 0x11, 0x8d, 0x93, 0x6b,
	0x0e, 0x39, 0x4f, 0x53, 0x1a, 0xa5, 0xbc, 0x74, 0xdc, 0x5e, 0x3a, 0x32, 0x25, 0xf9, 0xf9, 0x35,
	0xeb, 0x54, 0xd3, 0x39, 0x2b, 0xad, 0x07, 0x4b, 0xab, 0x51, 0x09, 0x97, 0x85, 0x33, 0xdf, 0xfb,
	0xb5, 0x87, 0x6e, 0x9c, 0xd8, 0x92, 0xf1, 0x6b, 0xd4, 0xa2, 0x79, 0x1e, 0x82, 0x97, 0x34, 0x06,
	0x8d, 0x61, 0x6b, 0xf2, 0xd9, 0xe5, 0x85, 0xf7, 0x64, 0xad, 0x4a, 0xd8, 0x83, 0xd1, 0xda, 0x1e,
	0x8c, 0xae, 0x42, 0xc7, 0x79, 0x7e, 0x66, 0xbf, 0x82, 0x26, 0x2d, 0xbf, 0xf0, 0xd7, 0x68, 0xd7,
	0x68, 0x1a, 0x27, 0x5c, 0x93, 0x77, 0x06, 0x8d, 0x61, 0xf7, 0xd1, 0x63, 0xff, 0xbf, 0xb6, 0xd5,
	0xaf, 0x90, 0xfe, 0x99, 0x53, 0x39, 0x5e, 0xc5, 0xc0, 0xa7, 0x68, 0x87, 0x32, 0xc1, 0xc8, 0x36,
	0x64, 0xf8, 0xe4, 0xf2, 0xc2, 0xfb, 0xa4, 0x5e, 0x86, 0x4c, 0xb0, 0x00, 0x10, 0xf8, 0x53, 0x74,
	0x38, 0xa5, 0x31, 0x8f, 0x94, 0x4a, 0x42, 0x6a, 0x8c, 0x16, 0xd1, 0xdc, 0x08, 0x25, 0x43, 0xc1,
	0xc8, 0x8e, 0xa5, 0x07, 0x07, 0x95, 0x7b, 0xbc, 0xf2, 0x9e, 0x32, 0x3c, 0x44, 0xfd, 0x95, 0x4e,
	0x3a, 0xc1, 0x0d, 0x10, 0x74, 0x97, 0x02, 0x09, 0x91, 0xf7, 0x51, 0x1f, 0xf2, 0x16, 0x72, 0x16,
	0x72, 0x69, 0x4f, 0x8e, 0x91, 0x9b, 0x10, 0xd9, 0xab, 0xec, 0x27, 0xce, 0x8c, 0x9f, 0x22, 0x14,
	0xa7, 0x22, 0x4e, 0x42, 0xdb, 0x16, 0x64, 0x77, 0xd0, 0x18, 0xb6, 0x1f, 0x1d, 0xf9, 0xae, 0x67,
	0xfc, 0xaa, 0x67, 0xfc, 0xb3, 0xaa, 0x67, 0x26, 0xcd, 0xdf, 0x2e, 0xbc, 0xad, 0x9f, 0x7f, 0xf7,
	0x1a, 0x41, 0x0b, 0x74, 0xd6, 0x83, 0x03, 0x84, 0xa7, 0x42, 0x17, 0x26, 0x2c, 0x78, 0x51, 0xd8,
	0x52, 0x00, 0xd6, 0xac, 0x01, 0xeb, 0x83, 0xfe, 0x95, 0x93, 0x03, 0xf3, 0x25, 0xda, 0x4f, 0xe9,
	0x55, 0x64, 0xab, 0x06, 0xb2, 0x67, 0xe5, 0xeb, 0xc4, 0xe7, 0x08, 0x4c, 0x21, 0x5c, 0x35, 0xc7,
	0x43, 0x35, 0x78, 0x7b, 0x56, 0x0c, 0x3d, 0xfb, 0xaf, 0xfc, 0xb4, 0xc5, 0xcd, 0xb9, 0xe3, 0xb5,
	0xeb, 0xe6, 0x17, 0x38, 0x35, 0x10, 0xed, 0x51, 0x68, 0x4e, 0x0d, 0x67, 0x21, 0x35, 0xa4, 0x53,
	0xeb, 0x28, 0x9c, 0x6e, 0x6c, 0xf0, 0x09, 0x6a, 0x6b, 0x1e, 0x73, 0xb1, 0x70, 0x94, 0xbd, 0x1a,
	0x14, 0x54, 0x09, 0xc7, 0x06, 0x7f, 0x81, 0x3a, 0x42, 0x16, 0x86, 0xa6, 0xa9, 0x2b, 0xac, 0x5b,
	0x83, 0xd3, 0x2e, 0x95, 0x50, 0xd4, 0xf7, 0xa8, 0xb7, 0x04, 0x95, 0xd7, 0xb1, 0xf7, 0xff, 0xaf,
	0x63, 0xb7, 0x42, 0x97, 0xb7, 0xf2, 0xa3, 0x15, 0x3d, 0x56, 0x73, 0x69, 0xf4, 0x39, 0xe9, 0xbb,
	0x1b, 0x51, 0x9a, 0x9f, 0x3a, 0x2b, 0x4e, 0x10, 0xa9, 0x02, 0x45, 0x96, 0xeb, 0xb2, 0xa7, 0x22,
	0x5a, 0x70, 0x46, 0xf6, 0xa1, 0xb6, 0x8f, 0x37, 0xcc, 0x67, 0xa2, 0x54, 0x3a, 0xd9, 0xb1, 0xc5,
	0x06, 0x77, 0x4a, 0xe4, 0xe9, 0x92, 0x38, 0xb1, 0x40, 0xec, 0x95, 0xe3, 0xbc, 0x1c, 0x6a, 0x18,
	0x32, 0x42, 0x60, 0x5a, 0xce, 0xa6, 0xb2, 0x6d, 0xc8, 0x2d, 0x58, 0xfc, 0xe1, 0x86, 0x8b, 0x8f,
	0x33, 0x5b, 0x4f, 0xb9, 0x7c, 0xc5, 0xb0, 0xeb, 0x31, 0xbe, 0x10, 0x31, 0x0f, 0xcd, 0x79, 0xce,
	0xc9, 0x6d, 0xb7, 0x9e, 0x33, 0x9d, 0x9d, 0xe7, 0x1c, 0x0f, 0x50, 0x9b, 0xcb, 0x85, 0xd0, 0x4a,
	0x66, 0x5c, 0x1a, 0x72, 0x00, 0x01, 0xeb, 0x26, 0xfc, 0x1a, 0x75, 0xed, 0x90, 0x0f, 0x0b, 0x96,
	0x84, 0x29, 0x5f, 0xf0, 0x94, 0xdc, 0x81, 0xc4, 0x1e, 0x6c, 0x98, 0xd8, 0xe9, 0x32, 0xab, 0x8e,
	0xe5, 0xbc, 0x62, 0xc9, 0x73, 0x4b, 0xc1, 0x43, 0xd4, 0xfe, 0x51, 0x49, 0x1e, 0xaa, 0xe9, 0xb4,
	0xe0, 0x86, 0x1c, 0x0e, 0x1a, 0xc3, 0xed, 0xc9, 0xee, 0xe5, 0x85, 0xb7, 0x2d, 0xa4, 0x09, 0x90,
	0xf5, 0xbd, 0x00, 0x17, 0xfe, 0x12, 0x21, 0x78, 0x37, 0xc2, 0x44, 0x48, 0x46, 0x08, 0xf4, 0xc8,
	0xa6, 0x67, 0xf2, 0x95, 0x90, 0x2c, 0x68, 0x81, 0xdc, 0x7e, 0xe2, 0x0f, 0x50, 0x27, 0x17, 0x72,
	0x66, 0xdf, 0xb2, 0x70, 0xae, 0x53, 0xf2, 0xae, 0x2b, 0xb8, 0xb2, 0x7d, 0xab, 0x53, 0xfc, 0x12,
	0xed, 0x55, 0xcf, 0x5d, 0xc8, 0xa8, 0xa1, 0xe4, 0xa8, 0x56, 0x17, 0x3c, 0xa3, 0x86, 0x06, 0x9d,
	0x8a, 0x60, 0xff, 0xf0, 0x3d, 0xb4, 0xe7, 0x86, 0xa0, 0x2a, 0x42, 0x49, 0x33, 0x4e, 0xde, 0x73,
	0xab, 0x82, 0xf1, 0x45, 0xf1, 0x0d, 0xcd, 0x38, 0xfe, 0xb0, 0x8a, 0xa9, 0xba, 0xf5, 0x2e, 0xc4,
	0x74, 0xc0, 0x58, 0xf5, 0xea, 0x03, 0xb4, 0xef, 0x82, 0xd6, 0x0f, 0xf5, 0x7d, 0x37, 0xbe, 0xc1,
	0xf1, 0x6c, 0x75, 0xb2, 0xf7, 0x51, 0xff, 0x5a, 0x3f, 0x1f, 0x0f, 0x1a, 0xc3, 0x66, 0xd0, 0x13,
	0x57, 0xba, 0x72, 0x84, 0x6e, 0x95, 0x40, 0x3b, 0x2c, 0xdc, 0xc3, 0xc2, 0x19, 0xf1, 0x20, 0x1a,
	0x3b, 0x57, 0xb0, 0xe6, 0xf9, 0xbc, 0xf9, 0xd3, 0x5b, 0x6f, 0xeb, 0xef, 0x5f, 0xbc, 0xad, 0xc9,
	0xc1, 0x1f, 0x7f, 0x1e, 0x6f, 0xbd, 0xfd, 0xeb, 0xb8, 0xf1, 0x5d, 0x7b, 0x6d, 0x37, 0xa2, 0x9b,
	0x30, 0x06, 0x1e, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x00, 0x87, 0x5b, 0xf5, 0xbd, 0x08, 0x00,
	0x00,
}

func (m *Event) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AppToken)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Tracker != 0 {
		n += 1 + sovEvent(uint64(m.Tracker))
	}
	l = len(m.Adid)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FacebookAttributionId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.FacebookAnonId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.TrackingEnabled)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ClickTime)
	n += 1 + l + sovEvent(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.FirstSessionTime)
	n += 1 + l + sovEvent(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastSessionTime)
	n += 1 + l + sovEvent(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastEventTime)
	n += 1 + l + sovEvent(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastRevenueTime)
	n += 1 + l + sovEvent(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovEvent(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ReceivedAt)
	n += 1 + l + sovEvent(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.InstallTime)
	n += 1 + l + sovEvent(uint64(l))
	if m.InstallTracker != 0 {
		n += 1 + sovEvent(uint64(m.InstallTracker))
	}
	l = len(m.InstallCountry)
	if l > 0 {
		n += 2 + l + sovEvent(uint64(l))
	}
	l = m.InstallImpressionBased.ProtoSize()
	n += 2 + l + sovEvent(uint64(l))
	l = len(m.EventToken)
	if l > 0 {
		n += 2 + l + sovEvent(uint64(l))
	}
	l = m.Revenue.ProtoSize()
	n += 2 + l + sovEvent(uint64(l))
	l = len(m.DeviceType)
	if l > 0 {
		n += 2 + l + sovEvent(uint64(l))
	}
	l = len(m.Environment)
	if l > 0 {
		n += 2 + l + sovEvent(uint64(l))
	}
	l = m.NullSdkLevel.ProtoSize()
	n += 2 + l + sovEvent(uint64(l))
	if m.ZoneOffset != 0 {
		n += 2 + sovEvent(uint64(m.ZoneOffset))
	}
	if m.FraudKind != 0 {
		n += 2 + sovEvent(uint64(m.FraudKind))
	}
	l = len(m.PingbackUrl)
	if l > 0 {
		n += 2 + l + sovEvent(uint64(l))
	}
	if m.CallbackData != nil {
		l = m.CallbackData.ProtoSize()
		n += 2 + l + sovEvent(uint64(l))
	}
	l = len(m.FirstOsName)
	if l > 0 {
		n += 2 + l + sovEvent(uint64(l))
	}
	l = len(m.FirstCountry)
	if l > 0 {
		n += 2 + l + sovEvent(uint64(l))
	}
	l = len(m.FirstDeviceType)
	if l > 0 {
		n += 2 + l + sovEvent(uint64(l))
	}
	if m.ImpressionBased {
		n += 3
	}
	if m.DeviceReattributed {
		n += 3
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
