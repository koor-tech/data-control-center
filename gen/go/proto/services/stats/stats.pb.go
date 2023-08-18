// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/services/stats/stats.proto

package stats

import (
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

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_stats_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_stats_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_stats_stats_proto_rawDescGZIP(), []int{0}
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*StreamResponse_Todo
	Message isStreamResponse_Message `protobuf_oneof:"message"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_stats_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_stats_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_stats_stats_proto_rawDescGZIP(), []int{1}
}

func (m *StreamResponse) GetMessage() isStreamResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *StreamResponse) GetTodo() bool {
	if x, ok := x.GetMessage().(*StreamResponse_Todo); ok {
		return x.Todo
	}
	return false
}

type isStreamResponse_Message interface {
	isStreamResponse_Message()
}

type StreamResponse_Todo struct {
	Todo bool `protobuf:"varint,1,opt,name=todo,proto3,oneof"`
}

func (*StreamResponse_Todo) isStreamResponse_Message() {}

var File_proto_services_stats_stats_proto protoreflect.FileDescriptor

var file_proto_services_stats_stats_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x42, 0x09, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x56, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x4d, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0xc9, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6f, 0x72, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0xa2, 0x02, 0x03,
	0x53, 0x43, 0x58, 0xaa, 0x02, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0xca, 0x02, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0xe2, 0x02, 0x1c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3a, 0x3a, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_services_stats_stats_proto_rawDescOnce sync.Once
	file_proto_services_stats_stats_proto_rawDescData = file_proto_services_stats_stats_proto_rawDesc
)

func file_proto_services_stats_stats_proto_rawDescGZIP() []byte {
	file_proto_services_stats_stats_proto_rawDescOnce.Do(func() {
		file_proto_services_stats_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_services_stats_stats_proto_rawDescData)
	})
	return file_proto_services_stats_stats_proto_rawDescData
}

var file_proto_services_stats_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_services_stats_stats_proto_goTypes = []interface{}{
	(*StreamRequest)(nil),  // 0: services.cluster.StreamRequest
	(*StreamResponse)(nil), // 1: services.cluster.StreamResponse
}
var file_proto_services_stats_stats_proto_depIdxs = []int32{
	0, // 0: services.cluster.Stats.Stream:input_type -> services.cluster.StreamRequest
	1, // 1: services.cluster.Stats.Stream:output_type -> services.cluster.StreamResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_services_stats_stats_proto_init() }
func file_proto_services_stats_stats_proto_init() {
	if File_proto_services_stats_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_services_stats_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file_proto_services_stats_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse); i {
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
	file_proto_services_stats_stats_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StreamResponse_Todo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_services_stats_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_services_stats_stats_proto_goTypes,
		DependencyIndexes: file_proto_services_stats_stats_proto_depIdxs,
		MessageInfos:      file_proto_services_stats_stats_proto_msgTypes,
	}.Build()
	File_proto_services_stats_stats_proto = out.File
	file_proto_services_stats_stats_proto_rawDesc = nil
	file_proto_services_stats_stats_proto_goTypes = nil
	file_proto_services_stats_stats_proto_depIdxs = nil
}
