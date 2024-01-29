// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/services/stats/v1/stats.proto

package statsv1

import (
	v1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	v11 "github.com/koor-tech/data-control-center/gen/go/api/resources/k8s/v1"
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

type GetClusterStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClusterStatsRequest) Reset() {
	*x = GetClusterStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_stats_v1_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterStatsRequest) ProtoMessage() {}

func (x *GetClusterStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_stats_v1_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterStatsRequest.ProtoReflect.Descriptor instead.
func (*GetClusterStatsRequest) Descriptor() ([]byte, []int) {
	return file_api_services_stats_v1_stats_proto_rawDescGZIP(), []int{0}
}

type GetClusterStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats *v1.ClusterStats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetClusterStatsResponse) Reset() {
	*x = GetClusterStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_stats_v1_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterStatsResponse) ProtoMessage() {}

func (x *GetClusterStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_stats_v1_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterStatsResponse.ProtoReflect.Descriptor instead.
func (*GetClusterStatsResponse) Descriptor() ([]byte, []int) {
	return file_api_services_stats_v1_stats_proto_rawDescGZIP(), []int{1}
}

func (x *GetClusterStatsResponse) GetStats() *v1.ClusterStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type GetClusterResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClusterResourcesRequest) Reset() {
	*x = GetClusterResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_stats_v1_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterResourcesRequest) ProtoMessage() {}

func (x *GetClusterResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_stats_v1_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterResourcesRequest.ProtoReflect.Descriptor instead.
func (*GetClusterResourcesRequest) Descriptor() ([]byte, []int) {
	return file_api_services_stats_v1_stats_proto_rawDescGZIP(), []int{2}
}

type GetClusterResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources   []*v11.ResourceInfo `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	Deployments []*v11.ResourceInfo `protobuf:"bytes,2,rep,name=deployments,proto3" json:"deployments,omitempty"`
}

func (x *GetClusterResourcesResponse) Reset() {
	*x = GetClusterResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_stats_v1_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterResourcesResponse) ProtoMessage() {}

func (x *GetClusterResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_stats_v1_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterResourcesResponse.ProtoReflect.Descriptor instead.
func (*GetClusterResourcesResponse) Descriptor() ([]byte, []int) {
	return file_api_services_stats_v1_stats_proto_rawDescGZIP(), []int{3}
}

func (x *GetClusterResourcesResponse) GetResources() []*v11.ResourceInfo {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *GetClusterResourcesResponse) GetDeployments() []*v11.ResourceInfo {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type GetClusterNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClusterNodesRequest) Reset() {
	*x = GetClusterNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_stats_v1_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterNodesRequest) ProtoMessage() {}

func (x *GetClusterNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_stats_v1_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterNodesRequest.ProtoReflect.Descriptor instead.
func (*GetClusterNodesRequest) Descriptor() ([]byte, []int) {
	return file_api_services_stats_v1_stats_proto_rawDescGZIP(), []int{4}
}

type GetClusterNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*v11.NodeInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *GetClusterNodesResponse) Reset() {
	*x = GetClusterNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_stats_v1_stats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterNodesResponse) ProtoMessage() {}

func (x *GetClusterNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_stats_v1_stats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterNodesResponse.ProtoReflect.Descriptor instead.
func (*GetClusterNodesResponse) Descriptor() ([]byte, []int) {
	return file_api_services_stats_v1_stats_proto_rawDescGZIP(), []int{5}
}

func (x *GetClusterNodesResponse) GetNodes() []*v11.NodeInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type GetClusterRadarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClusterRadarRequest) Reset() {
	*x = GetClusterRadarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_stats_v1_stats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterRadarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterRadarRequest) ProtoMessage() {}

func (x *GetClusterRadarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_stats_v1_stats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterRadarRequest.ProtoReflect.Descriptor instead.
func (*GetClusterRadarRequest) Descriptor() ([]byte, []int) {
	return file_api_services_stats_v1_stats_proto_rawDescGZIP(), []int{6}
}

type GetClusterRadarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Radar *v1.ClusterRadar `protobuf:"bytes,1,opt,name=radar,proto3" json:"radar,omitempty"`
}

func (x *GetClusterRadarResponse) Reset() {
	*x = GetClusterRadarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_stats_v1_stats_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterRadarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterRadarResponse) ProtoMessage() {}

func (x *GetClusterRadarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_stats_v1_stats_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterRadarResponse.ProtoReflect.Descriptor instead.
func (*GetClusterRadarResponse) Descriptor() ([]byte, []int) {
	return file_api_services_stats_v1_stats_proto_rawDescGZIP(), []int{7}
}

func (x *GetClusterRadarResponse) GetRadar() *v1.ClusterRadar {
	if x != nil {
		return x.Radar
	}
	return nil
}

var File_api_services_stats_v1_stats_proto protoreflect.FileDescriptor

var file_api_services_stats_v1_stats_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x65, 0x70, 0x68, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x54, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xa5, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x61, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x54,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x64, 0x61,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x72, 0x61, 0x64,
	0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x64, 0x61, 0x72, 0x52, 0x05, 0x72,
	0x61, 0x64, 0x61, 0x72, 0x32, 0xea, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7e, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x72, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x72, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x64, 0x61, 0x72,
	0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x61, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x61, 0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xed, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6f, 0x72, 0x2d, 0x74,
	0x65, 0x63, 0x68, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41,
	0x53, 0x53, 0xaa, 0x02, 0x15, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x41, 0x70, 0x69,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x21, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_services_stats_v1_stats_proto_rawDescOnce sync.Once
	file_api_services_stats_v1_stats_proto_rawDescData = file_api_services_stats_v1_stats_proto_rawDesc
)

func file_api_services_stats_v1_stats_proto_rawDescGZIP() []byte {
	file_api_services_stats_v1_stats_proto_rawDescOnce.Do(func() {
		file_api_services_stats_v1_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_stats_v1_stats_proto_rawDescData)
	})
	return file_api_services_stats_v1_stats_proto_rawDescData
}

var file_api_services_stats_v1_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_services_stats_v1_stats_proto_goTypes = []interface{}{
	(*GetClusterStatsRequest)(nil),      // 0: api.services.stats.v1.GetClusterStatsRequest
	(*GetClusterStatsResponse)(nil),     // 1: api.services.stats.v1.GetClusterStatsResponse
	(*GetClusterResourcesRequest)(nil),  // 2: api.services.stats.v1.GetClusterResourcesRequest
	(*GetClusterResourcesResponse)(nil), // 3: api.services.stats.v1.GetClusterResourcesResponse
	(*GetClusterNodesRequest)(nil),      // 4: api.services.stats.v1.GetClusterNodesRequest
	(*GetClusterNodesResponse)(nil),     // 5: api.services.stats.v1.GetClusterNodesResponse
	(*GetClusterRadarRequest)(nil),      // 6: api.services.stats.v1.GetClusterRadarRequest
	(*GetClusterRadarResponse)(nil),     // 7: api.services.stats.v1.GetClusterRadarResponse
	(*v1.ClusterStats)(nil),             // 8: api.resources.ceph.v1.ClusterStats
	(*v11.ResourceInfo)(nil),            // 9: api.resources.k8s.v1.ResourceInfo
	(*v11.NodeInfo)(nil),                // 10: api.resources.k8s.v1.NodeInfo
	(*v1.ClusterRadar)(nil),             // 11: api.resources.ceph.v1.ClusterRadar
}
var file_api_services_stats_v1_stats_proto_depIdxs = []int32{
	8,  // 0: api.services.stats.v1.GetClusterStatsResponse.stats:type_name -> api.resources.ceph.v1.ClusterStats
	9,  // 1: api.services.stats.v1.GetClusterResourcesResponse.resources:type_name -> api.resources.k8s.v1.ResourceInfo
	9,  // 2: api.services.stats.v1.GetClusterResourcesResponse.deployments:type_name -> api.resources.k8s.v1.ResourceInfo
	10, // 3: api.services.stats.v1.GetClusterNodesResponse.nodes:type_name -> api.resources.k8s.v1.NodeInfo
	11, // 4: api.services.stats.v1.GetClusterRadarResponse.radar:type_name -> api.resources.ceph.v1.ClusterRadar
	0,  // 5: api.services.stats.v1.StatsService.GetClusterStats:input_type -> api.services.stats.v1.GetClusterStatsRequest
	2,  // 6: api.services.stats.v1.StatsService.GetClusterResources:input_type -> api.services.stats.v1.GetClusterResourcesRequest
	4,  // 7: api.services.stats.v1.StatsService.GetClusterNodes:input_type -> api.services.stats.v1.GetClusterNodesRequest
	6,  // 8: api.services.stats.v1.StatsService.GetClusterRadar:input_type -> api.services.stats.v1.GetClusterRadarRequest
	1,  // 9: api.services.stats.v1.StatsService.GetClusterStats:output_type -> api.services.stats.v1.GetClusterStatsResponse
	3,  // 10: api.services.stats.v1.StatsService.GetClusterResources:output_type -> api.services.stats.v1.GetClusterResourcesResponse
	5,  // 11: api.services.stats.v1.StatsService.GetClusterNodes:output_type -> api.services.stats.v1.GetClusterNodesResponse
	7,  // 12: api.services.stats.v1.StatsService.GetClusterRadar:output_type -> api.services.stats.v1.GetClusterRadarResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_services_stats_v1_stats_proto_init() }
func file_api_services_stats_v1_stats_proto_init() {
	if File_api_services_stats_v1_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_stats_v1_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterStatsRequest); i {
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
		file_api_services_stats_v1_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterStatsResponse); i {
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
		file_api_services_stats_v1_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterResourcesRequest); i {
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
		file_api_services_stats_v1_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterResourcesResponse); i {
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
		file_api_services_stats_v1_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterNodesRequest); i {
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
		file_api_services_stats_v1_stats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterNodesResponse); i {
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
		file_api_services_stats_v1_stats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterRadarRequest); i {
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
		file_api_services_stats_v1_stats_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterRadarResponse); i {
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
			RawDescriptor: file_api_services_stats_v1_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_services_stats_v1_stats_proto_goTypes,
		DependencyIndexes: file_api_services_stats_v1_stats_proto_depIdxs,
		MessageInfos:      file_api_services_stats_v1_stats_proto_msgTypes,
	}.Build()
	File_api_services_stats_v1_stats_proto = out.File
	file_api_services_stats_v1_stats_proto_rawDesc = nil
	file_api_services_stats_v1_stats_proto_goTypes = nil
	file_api_services_stats_v1_stats_proto_depIdxs = nil
}