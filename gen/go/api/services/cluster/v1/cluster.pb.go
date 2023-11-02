// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/services/cluster/v1/cluster.proto

package clusterv1

import (
	v1 "github.com/koor-tech/data-control-center/gen/go/api/resources/koor/v1"
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

type GetKoorClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetKoorClusterRequest) Reset() {
	*x = GetKoorClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKoorClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKoorClusterRequest) ProtoMessage() {}

func (x *GetKoorClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKoorClusterRequest.ProtoReflect.Descriptor instead.
func (*GetKoorClusterRequest) Descriptor() ([]byte, []int) {
	return file_api_services_cluster_v1_cluster_proto_rawDescGZIP(), []int{0}
}

type GetKoorClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KoorCluster *v1.KoorCluster `protobuf:"bytes,1,opt,name=koor_cluster,json=koorCluster,proto3" json:"koor_cluster,omitempty"`
}

func (x *GetKoorClusterResponse) Reset() {
	*x = GetKoorClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKoorClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKoorClusterResponse) ProtoMessage() {}

func (x *GetKoorClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKoorClusterResponse.ProtoReflect.Descriptor instead.
func (*GetKoorClusterResponse) Descriptor() ([]byte, []int) {
	return file_api_services_cluster_v1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *GetKoorClusterResponse) GetKoorCluster() *v1.KoorCluster {
	if x != nil {
		return x.KoorCluster
	}
	return nil
}

type GetTroubleshootReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTroubleshootReportRequest) Reset() {
	*x = GetTroubleshootReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTroubleshootReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTroubleshootReportRequest) ProtoMessage() {}

func (x *GetTroubleshootReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTroubleshootReportRequest.ProtoReflect.Descriptor instead.
func (*GetTroubleshootReportRequest) Descriptor() ([]byte, []int) {
	return file_api_services_cluster_v1_cluster_proto_rawDescGZIP(), []int{2}
}

type GetTroubleshootReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report string `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *GetTroubleshootReportResponse) Reset() {
	*x = GetTroubleshootReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTroubleshootReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTroubleshootReportResponse) ProtoMessage() {}

func (x *GetTroubleshootReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTroubleshootReportResponse.ProtoReflect.Descriptor instead.
func (*GetTroubleshootReportResponse) Descriptor() ([]byte, []int) {
	return file_api_services_cluster_v1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *GetTroubleshootReportResponse) GetReport() string {
	if x != nil {
		return x.Report
	}
	return ""
}

type SetScrubbingScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OsdScrubbingSchedule *v1.OSDScrubbingSchedule `protobuf:"bytes,1,opt,name=osd_scrubbing_schedule,json=osdScrubbingSchedule,proto3" json:"osd_scrubbing_schedule,omitempty"`
}

func (x *SetScrubbingScheduleRequest) Reset() {
	*x = SetScrubbingScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScrubbingScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScrubbingScheduleRequest) ProtoMessage() {}

func (x *SetScrubbingScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScrubbingScheduleRequest.ProtoReflect.Descriptor instead.
func (*SetScrubbingScheduleRequest) Descriptor() ([]byte, []int) {
	return file_api_services_cluster_v1_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *SetScrubbingScheduleRequest) GetOsdScrubbingSchedule() *v1.OSDScrubbingSchedule {
	if x != nil {
		return x.OsdScrubbingSchedule
	}
	return nil
}

type SetScrubbingScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OsdScrubbingSchedule *v1.OSDScrubbingSchedule `protobuf:"bytes,1,opt,name=osd_scrubbing_schedule,json=osdScrubbingSchedule,proto3" json:"osd_scrubbing_schedule,omitempty"`
}

func (x *SetScrubbingScheduleResponse) Reset() {
	*x = SetScrubbingScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScrubbingScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScrubbingScheduleResponse) ProtoMessage() {}

func (x *SetScrubbingScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_cluster_v1_cluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScrubbingScheduleResponse.ProtoReflect.Descriptor instead.
func (*SetScrubbingScheduleResponse) Descriptor() ([]byte, []int) {
	return file_api_services_cluster_v1_cluster_proto_rawDescGZIP(), []int{5}
}

func (x *SetScrubbingScheduleResponse) GetOsdScrubbingSchedule() *v1.OSDScrubbingSchedule {
	if x != nil {
		return x.OsdScrubbingSchedule
	}
	return nil
}

var File_api_services_cluster_v1_cluster_proto protoreflect.FileDescriptor

var file_api_services_cluster_v1_cluster_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x6b, 0x6f, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x6f, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4b, 0x6f, 0x6f, 0x72, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4b, 0x6f, 0x6f, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x6b, 0x6f, 0x6f, 0x72, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x6f, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x6f, 0x6f, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x0b, 0x6b, 0x6f, 0x6f, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x1d,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x53, 0x63, 0x72,
	0x75, 0x62, 0x62, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x61, 0x0a, 0x16, 0x6f, 0x73, 0x64, 0x5f, 0x73, 0x63, 0x72,
	0x75, 0x62, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x6f, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x53,
	0x44, 0x53, 0x63, 0x72, 0x75, 0x62, 0x62, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x14, 0x6f, 0x73, 0x64, 0x53, 0x63, 0x72, 0x75, 0x62, 0x62, 0x69, 0x6e, 0x67,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x1c, 0x53, 0x65, 0x74,
	0x53, 0x63, 0x72, 0x75, 0x62, 0x62, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x16, 0x6f, 0x73, 0x64,
	0x5f, 0x73, 0x63, 0x72, 0x75, 0x62, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x6f, 0x6f, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x53, 0x44, 0x53, 0x63, 0x72, 0x75, 0x62, 0x62, 0x69, 0x6e, 0x67, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x14, 0x6f, 0x73, 0x64, 0x53, 0x63, 0x72, 0x75, 0x62,
	0x62, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x32, 0x98, 0x03, 0x0a,
	0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x73, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4b, 0x6f, 0x6f, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b,
	0x6f, 0x6f, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b,
	0x6f, 0x6f, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x88, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x72, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x35,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x85, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x53, 0x63, 0x72, 0x75, 0x62, 0x62, 0x69, 0x6e, 0x67,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x63, 0x72, 0x75, 0x62, 0x62, 0x69, 0x6e, 0x67, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x63, 0x72, 0x75,
	0x62, 0x62, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xfd, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6f, 0x72, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53, 0x43,
	0xaa, 0x02, 0x17, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x41, 0x70, 0x69,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_services_cluster_v1_cluster_proto_rawDescOnce sync.Once
	file_api_services_cluster_v1_cluster_proto_rawDescData = file_api_services_cluster_v1_cluster_proto_rawDesc
)

func file_api_services_cluster_v1_cluster_proto_rawDescGZIP() []byte {
	file_api_services_cluster_v1_cluster_proto_rawDescOnce.Do(func() {
		file_api_services_cluster_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_cluster_v1_cluster_proto_rawDescData)
	})
	return file_api_services_cluster_v1_cluster_proto_rawDescData
}

var file_api_services_cluster_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_services_cluster_v1_cluster_proto_goTypes = []interface{}{
	(*GetKoorClusterRequest)(nil),         // 0: api.services.cluster.v1.GetKoorClusterRequest
	(*GetKoorClusterResponse)(nil),        // 1: api.services.cluster.v1.GetKoorClusterResponse
	(*GetTroubleshootReportRequest)(nil),  // 2: api.services.cluster.v1.GetTroubleshootReportRequest
	(*GetTroubleshootReportResponse)(nil), // 3: api.services.cluster.v1.GetTroubleshootReportResponse
	(*SetScrubbingScheduleRequest)(nil),   // 4: api.services.cluster.v1.SetScrubbingScheduleRequest
	(*SetScrubbingScheduleResponse)(nil),  // 5: api.services.cluster.v1.SetScrubbingScheduleResponse
	(*v1.KoorCluster)(nil),                // 6: api.resources.koor.v1.KoorCluster
	(*v1.OSDScrubbingSchedule)(nil),       // 7: api.resources.koor.v1.OSDScrubbingSchedule
}
var file_api_services_cluster_v1_cluster_proto_depIdxs = []int32{
	6, // 0: api.services.cluster.v1.GetKoorClusterResponse.koor_cluster:type_name -> api.resources.koor.v1.KoorCluster
	7, // 1: api.services.cluster.v1.SetScrubbingScheduleRequest.osd_scrubbing_schedule:type_name -> api.resources.koor.v1.OSDScrubbingSchedule
	7, // 2: api.services.cluster.v1.SetScrubbingScheduleResponse.osd_scrubbing_schedule:type_name -> api.resources.koor.v1.OSDScrubbingSchedule
	0, // 3: api.services.cluster.v1.ClusterService.GetKoorCluster:input_type -> api.services.cluster.v1.GetKoorClusterRequest
	2, // 4: api.services.cluster.v1.ClusterService.GetTroubleshootReport:input_type -> api.services.cluster.v1.GetTroubleshootReportRequest
	4, // 5: api.services.cluster.v1.ClusterService.SetScrubbingSchedule:input_type -> api.services.cluster.v1.SetScrubbingScheduleRequest
	1, // 6: api.services.cluster.v1.ClusterService.GetKoorCluster:output_type -> api.services.cluster.v1.GetKoorClusterResponse
	3, // 7: api.services.cluster.v1.ClusterService.GetTroubleshootReport:output_type -> api.services.cluster.v1.GetTroubleshootReportResponse
	5, // 8: api.services.cluster.v1.ClusterService.SetScrubbingSchedule:output_type -> api.services.cluster.v1.SetScrubbingScheduleResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_services_cluster_v1_cluster_proto_init() }
func file_api_services_cluster_v1_cluster_proto_init() {
	if File_api_services_cluster_v1_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_cluster_v1_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKoorClusterRequest); i {
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
		file_api_services_cluster_v1_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKoorClusterResponse); i {
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
		file_api_services_cluster_v1_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTroubleshootReportRequest); i {
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
		file_api_services_cluster_v1_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTroubleshootReportResponse); i {
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
		file_api_services_cluster_v1_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScrubbingScheduleRequest); i {
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
		file_api_services_cluster_v1_cluster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScrubbingScheduleResponse); i {
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
			RawDescriptor: file_api_services_cluster_v1_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_services_cluster_v1_cluster_proto_goTypes,
		DependencyIndexes: file_api_services_cluster_v1_cluster_proto_depIdxs,
		MessageInfos:      file_api_services_cluster_v1_cluster_proto_msgTypes,
	}.Build()
	File_api_services_cluster_v1_cluster_proto = out.File
	file_api_services_cluster_v1_cluster_proto_rawDesc = nil
	file_api_services_cluster_v1_cluster_proto_goTypes = nil
	file_api_services_cluster_v1_cluster_proto_depIdxs = nil
}
