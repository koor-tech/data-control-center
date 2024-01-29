// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/resources/ceph/v1/controls.proto

package cephv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Osd Scrubbing schedule config
type OSDScrubbingSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxScrubOps       *int64   `protobuf:"varint,2,opt,name=max_scrub_ops,json=maxScrubOps,proto3,oneof" json:"max_scrub_ops,omitempty"`
	BeginHour         *int64   `protobuf:"varint,3,opt,name=begin_hour,json=beginHour,proto3,oneof" json:"begin_hour,omitempty"`
	EndHour           *int64   `protobuf:"varint,4,opt,name=end_hour,json=endHour,proto3,oneof" json:"end_hour,omitempty"`
	BeginWeekDay      *int64   `protobuf:"varint,5,opt,name=begin_week_day,json=beginWeekDay,proto3,oneof" json:"begin_week_day,omitempty"`
	EndWeekDay        *int64   `protobuf:"varint,6,opt,name=end_week_day,json=endWeekDay,proto3,oneof" json:"end_week_day,omitempty"`
	MinScrubInterval  *int64   `protobuf:"varint,7,opt,name=min_scrub_interval,json=minScrubInterval,proto3,oneof" json:"min_scrub_interval,omitempty"`
	MaxScrubInterval  *int64   `protobuf:"varint,8,opt,name=max_scrub_interval,json=maxScrubInterval,proto3,oneof" json:"max_scrub_interval,omitempty"`
	DeepScrubInterval *int64   `protobuf:"varint,9,opt,name=deep_scrub_interval,json=deepScrubInterval,proto3,oneof" json:"deep_scrub_interval,omitempty"`
	ScrubSleepSeconds *float32 `protobuf:"fixed32,10,opt,name=scrub_sleep_seconds,json=scrubSleepSeconds,proto3,oneof" json:"scrub_sleep_seconds,omitempty"`
}

func (x *OSDScrubbingSchedule) Reset() {
	*x = OSDScrubbingSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_resources_ceph_v1_controls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSDScrubbingSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSDScrubbingSchedule) ProtoMessage() {}

func (x *OSDScrubbingSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_ceph_v1_controls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSDScrubbingSchedule.ProtoReflect.Descriptor instead.
func (*OSDScrubbingSchedule) Descriptor() ([]byte, []int) {
	return file_api_resources_ceph_v1_controls_proto_rawDescGZIP(), []int{0}
}

func (x *OSDScrubbingSchedule) GetMaxScrubOps() int64 {
	if x != nil && x.MaxScrubOps != nil {
		return *x.MaxScrubOps
	}
	return 0
}

func (x *OSDScrubbingSchedule) GetBeginHour() int64 {
	if x != nil && x.BeginHour != nil {
		return *x.BeginHour
	}
	return 0
}

func (x *OSDScrubbingSchedule) GetEndHour() int64 {
	if x != nil && x.EndHour != nil {
		return *x.EndHour
	}
	return 0
}

func (x *OSDScrubbingSchedule) GetBeginWeekDay() int64 {
	if x != nil && x.BeginWeekDay != nil {
		return *x.BeginWeekDay
	}
	return 0
}

func (x *OSDScrubbingSchedule) GetEndWeekDay() int64 {
	if x != nil && x.EndWeekDay != nil {
		return *x.EndWeekDay
	}
	return 0
}

func (x *OSDScrubbingSchedule) GetMinScrubInterval() int64 {
	if x != nil && x.MinScrubInterval != nil {
		return *x.MinScrubInterval
	}
	return 0
}

func (x *OSDScrubbingSchedule) GetMaxScrubInterval() int64 {
	if x != nil && x.MaxScrubInterval != nil {
		return *x.MaxScrubInterval
	}
	return 0
}

func (x *OSDScrubbingSchedule) GetDeepScrubInterval() int64 {
	if x != nil && x.DeepScrubInterval != nil {
		return *x.DeepScrubInterval
	}
	return 0
}

func (x *OSDScrubbingSchedule) GetScrubSleepSeconds() float32 {
	if x != nil && x.ScrubSleepSeconds != nil {
		return *x.ScrubSleepSeconds
	}
	return 0
}

var File_api_resources_ceph_v1_controls_proto protoreflect.FileDescriptor

var file_api_resources_ceph_v1_controls_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x65, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x05, 0x0a, 0x14, 0x4f, 0x53, 0x44, 0x53, 0x63,
	0x72, 0x75, 0x62, 0x62, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x37, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x5f, 0x6f, 0x70, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0xfa, 0x42, 0x09, 0x22, 0x07, 0x10, 0xff, 0xac,
	0xe2, 0x04, 0x28, 0x00, 0x30, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x53, 0x63, 0x72,
	0x75, 0x62, 0x4f, 0x70, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69,
	0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xfa, 0x42,
	0x06, 0x22, 0x04, 0x10, 0x17, 0x28, 0x00, 0x30, 0x01, 0x48, 0x01, 0x52, 0x09, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xfa, 0x42, 0x06,
	0x22, 0x04, 0x10, 0x17, 0x28, 0x00, 0x30, 0x01, 0x48, 0x02, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x48,
	0x6f, 0x75, 0x72, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x0e, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f,
	0x77, 0x65, 0x65, 0x6b, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b,
	0xfa, 0x42, 0x06, 0x22, 0x04, 0x10, 0x06, 0x28, 0x00, 0x30, 0x01, 0x48, 0x03, 0x52, 0x0c, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x57, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79, 0x88, 0x01, 0x01, 0x12, 0x32,
	0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xfa, 0x42, 0x06, 0x22, 0x04, 0x10, 0x06, 0x28, 0x00, 0x30,
	0x01, 0x48, 0x04, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x57, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79, 0x88,
	0x01, 0x01, 0x12, 0x42, 0x0a, 0x12, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f,
	0xfa, 0x42, 0x0a, 0x22, 0x08, 0x10, 0xff, 0xc7, 0xaf, 0xa0, 0x25, 0x28, 0x00, 0x30, 0x01, 0x48,
	0x05, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x53, 0x63, 0x72, 0x75, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x63,
	0x72, 0x75, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x0f, 0xfa, 0x42, 0x0a, 0x22, 0x08, 0x10, 0xff, 0xc7, 0xaf, 0xa0, 0x25, 0x28,
	0x00, 0x30, 0x01, 0x48, 0x06, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x53, 0x63, 0x72, 0x75, 0x62, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x13, 0x64, 0x65,
	0x65, 0x70, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0xfa, 0x42, 0x0a, 0x22, 0x08, 0x10, 0xff,
	0xc7, 0xaf, 0xa0, 0x25, 0x28, 0x00, 0x30, 0x01, 0x48, 0x07, 0x52, 0x11, 0x64, 0x65, 0x65, 0x70,
	0x53, 0x63, 0x72, 0x75, 0x62, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x44, 0x0a, 0x13, 0x73, 0x63, 0x72, 0x75, 0x62, 0x5f, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0f, 0xfa,
	0x42, 0x0c, 0x0a, 0x0a, 0x15, 0xf9, 0x02, 0x15, 0x50, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x48, 0x08,
	0x52, 0x11, 0x73, 0x63, 0x72, 0x75, 0x62, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x73,
	0x63, 0x72, 0x75, 0x62, 0x5f, 0x6f, 0x70, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x68, 0x6f, 0x75, 0x72, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x77,
	0x65, 0x65, 0x6b, 0x5f, 0x64, 0x61, 0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x77, 0x65, 0x65, 0x6b, 0x5f, 0x64, 0x61, 0x79, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x6d, 0x69, 0x6e,
	0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42,
	0x15, 0x0a, 0x13, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x64, 0x65, 0x65, 0x70, 0x5f,
	0x73, 0x63, 0x72, 0x75, 0x62, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x16,
	0x0a, 0x14, 0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x5f, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0xef, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70,
	0x68, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x6f, 0x6f, 0x72, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x65, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x65, 0x70,
	0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x52, 0x43, 0xaa, 0x02, 0x15, 0x41, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x65, 0x70, 0x68, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x5c, 0x43, 0x65, 0x70, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x41, 0x70, 0x69, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5c, 0x43, 0x65, 0x70, 0x68, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0x3a,
	0x43, 0x65, 0x70, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_resources_ceph_v1_controls_proto_rawDescOnce sync.Once
	file_api_resources_ceph_v1_controls_proto_rawDescData = file_api_resources_ceph_v1_controls_proto_rawDesc
)

func file_api_resources_ceph_v1_controls_proto_rawDescGZIP() []byte {
	file_api_resources_ceph_v1_controls_proto_rawDescOnce.Do(func() {
		file_api_resources_ceph_v1_controls_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_resources_ceph_v1_controls_proto_rawDescData)
	})
	return file_api_resources_ceph_v1_controls_proto_rawDescData
}

var file_api_resources_ceph_v1_controls_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_resources_ceph_v1_controls_proto_goTypes = []interface{}{
	(*OSDScrubbingSchedule)(nil), // 0: api.resources.ceph.v1.OSDScrubbingSchedule
}
var file_api_resources_ceph_v1_controls_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_resources_ceph_v1_controls_proto_init() }
func file_api_resources_ceph_v1_controls_proto_init() {
	if File_api_resources_ceph_v1_controls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_resources_ceph_v1_controls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSDScrubbingSchedule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_resources_ceph_v1_controls_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_resources_ceph_v1_controls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_resources_ceph_v1_controls_proto_goTypes,
		DependencyIndexes: file_api_resources_ceph_v1_controls_proto_depIdxs,
		MessageInfos:      file_api_resources_ceph_v1_controls_proto_msgTypes,
	}.Build()
	File_api_resources_ceph_v1_controls_proto = out.File
	file_api_resources_ceph_v1_controls_proto_rawDesc = nil
	file_api_resources_ceph_v1_controls_proto_goTypes = nil
	file_api_resources_ceph_v1_controls_proto_depIdxs = nil
}