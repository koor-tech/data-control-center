// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/services/k8sresources/v1/editor.proto

package k8sresourcesv1

import (
	v1 "github.com/koor-tech/data-control-center/gen/go/api/resources/k8s/v1"
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

type GetResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources *v1.Resources `protobuf:"bytes,1,opt,name=resources,proto3" json:"resources,omitempty"`
}

func (x *GetResourcesResponse) Reset() {
	*x = GetResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_k8sresources_v1_editor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesResponse) ProtoMessage() {}

func (x *GetResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_k8sresources_v1_editor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesResponse.ProtoReflect.Descriptor instead.
func (*GetResourcesResponse) Descriptor() ([]byte, []int) {
	return file_api_services_k8sresources_v1_editor_proto_rawDescGZIP(), []int{0}
}

func (x *GetResourcesResponse) GetResources() *v1.Resources {
	if x != nil {
		return x.Resources
	}
	return nil
}

type GetResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetResourcesRequest) Reset() {
	*x = GetResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_k8sresources_v1_editor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesRequest) ProtoMessage() {}

func (x *GetResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_k8sresources_v1_editor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesRequest.ProtoReflect.Descriptor instead.
func (*GetResourcesRequest) Descriptor() ([]byte, []int) {
	return file_api_services_k8sresources_v1_editor_proto_rawDescGZIP(), []int{1}
}

type SaveResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *SaveResourceRequest) Reset() {
	*x = SaveResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_k8sresources_v1_editor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResourceRequest) ProtoMessage() {}

func (x *SaveResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_k8sresources_v1_editor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResourceRequest.ProtoReflect.Descriptor instead.
func (*SaveResourceRequest) Descriptor() ([]byte, []int) {
	return file_api_services_k8sresources_v1_editor_proto_rawDescGZIP(), []int{2}
}

func (x *SaveResourceRequest) GetResource() *v1.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type SaveResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *SaveResourceResponse) Reset() {
	*x = SaveResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_k8sresources_v1_editor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResourceResponse) ProtoMessage() {}

func (x *SaveResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_k8sresources_v1_editor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResourceResponse.ProtoReflect.Descriptor instead.
func (*SaveResourceResponse) Descriptor() ([]byte, []int) {
	return file_api_services_k8sresources_v1_editor_proto_rawDescGZIP(), []int{3}
}

func (x *SaveResourceResponse) GetResource() *v1.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

var File_api_services_k8sresources_v1_editor_proto protoreflect.FileDescriptor

var file_api_services_k8sresources_v1_editor_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6b,
	0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x55, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a,
	0x13, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x52, 0x0a, 0x14, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x32, 0x87, 0x02, 0x0a, 0x13, 0x4b, 0x38, 0x73, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6b,
	0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x9f,
	0x02, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x6f, 0x6f, 0x72, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x6b, 0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x6b, 0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x41, 0x53, 0x4b, 0xaa, 0x02, 0x1c, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1c, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5c, 0x4b, 0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x28, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x4b, 0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1f, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a,
	0x4b, 0x38, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_services_k8sresources_v1_editor_proto_rawDescOnce sync.Once
	file_api_services_k8sresources_v1_editor_proto_rawDescData = file_api_services_k8sresources_v1_editor_proto_rawDesc
)

func file_api_services_k8sresources_v1_editor_proto_rawDescGZIP() []byte {
	file_api_services_k8sresources_v1_editor_proto_rawDescOnce.Do(func() {
		file_api_services_k8sresources_v1_editor_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_k8sresources_v1_editor_proto_rawDescData)
	})
	return file_api_services_k8sresources_v1_editor_proto_rawDescData
}

var file_api_services_k8sresources_v1_editor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_services_k8sresources_v1_editor_proto_goTypes = []interface{}{
	(*GetResourcesResponse)(nil), // 0: api.services.k8sresources.v1.GetResourcesResponse
	(*GetResourcesRequest)(nil),  // 1: api.services.k8sresources.v1.GetResourcesRequest
	(*SaveResourceRequest)(nil),  // 2: api.services.k8sresources.v1.SaveResourceRequest
	(*SaveResourceResponse)(nil), // 3: api.services.k8sresources.v1.SaveResourceResponse
	(*v1.Resources)(nil),         // 4: api.resources.k8s.v1.Resources
	(*v1.Resource)(nil),          // 5: api.resources.k8s.v1.Resource
}
var file_api_services_k8sresources_v1_editor_proto_depIdxs = []int32{
	4, // 0: api.services.k8sresources.v1.GetResourcesResponse.resources:type_name -> api.resources.k8s.v1.Resources
	5, // 1: api.services.k8sresources.v1.SaveResourceRequest.resource:type_name -> api.resources.k8s.v1.Resource
	5, // 2: api.services.k8sresources.v1.SaveResourceResponse.resource:type_name -> api.resources.k8s.v1.Resource
	1, // 3: api.services.k8sresources.v1.K8sResourcesService.GetResources:input_type -> api.services.k8sresources.v1.GetResourcesRequest
	2, // 4: api.services.k8sresources.v1.K8sResourcesService.SaveResource:input_type -> api.services.k8sresources.v1.SaveResourceRequest
	0, // 5: api.services.k8sresources.v1.K8sResourcesService.GetResources:output_type -> api.services.k8sresources.v1.GetResourcesResponse
	3, // 6: api.services.k8sresources.v1.K8sResourcesService.SaveResource:output_type -> api.services.k8sresources.v1.SaveResourceResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_services_k8sresources_v1_editor_proto_init() }
func file_api_services_k8sresources_v1_editor_proto_init() {
	if File_api_services_k8sresources_v1_editor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_k8sresources_v1_editor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcesResponse); i {
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
		file_api_services_k8sresources_v1_editor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcesRequest); i {
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
		file_api_services_k8sresources_v1_editor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveResourceRequest); i {
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
		file_api_services_k8sresources_v1_editor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveResourceResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_services_k8sresources_v1_editor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_services_k8sresources_v1_editor_proto_goTypes,
		DependencyIndexes: file_api_services_k8sresources_v1_editor_proto_depIdxs,
		MessageInfos:      file_api_services_k8sresources_v1_editor_proto_msgTypes,
	}.Build()
	File_api_services_k8sresources_v1_editor_proto = out.File
	file_api_services_k8sresources_v1_editor_proto_rawDesc = nil
	file_api_services_k8sresources_v1_editor_proto_goTypes = nil
	file_api_services_k8sresources_v1_editor_proto_depIdxs = nil
}
