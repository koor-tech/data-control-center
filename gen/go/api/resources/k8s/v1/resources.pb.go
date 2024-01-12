// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/resources/k8s/v1/resources.proto

package k8sv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReliabilityScore int32

const (
	ReliabilityScore_RELIABILITY_SCORE_UNSPECIFIED ReliabilityScore = 0
	ReliabilityScore_RELIABILITY_SCORE_UNKNOWN     ReliabilityScore = 1
	ReliabilityScore_RELIABILITY_SCORE_NONE        ReliabilityScore = 2
	ReliabilityScore_RELIABILITY_SCORE_DEGRADED    ReliabilityScore = 3
	ReliabilityScore_RELIABILITY_SCORE_OK          ReliabilityScore = 4
)

// Enum value maps for ReliabilityScore.
var (
	ReliabilityScore_name = map[int32]string{
		0: "RELIABILITY_SCORE_UNSPECIFIED",
		1: "RELIABILITY_SCORE_UNKNOWN",
		2: "RELIABILITY_SCORE_NONE",
		3: "RELIABILITY_SCORE_DEGRADED",
		4: "RELIABILITY_SCORE_OK",
	}
	ReliabilityScore_value = map[string]int32{
		"RELIABILITY_SCORE_UNSPECIFIED": 0,
		"RELIABILITY_SCORE_UNKNOWN":     1,
		"RELIABILITY_SCORE_NONE":        2,
		"RELIABILITY_SCORE_DEGRADED":    3,
		"RELIABILITY_SCORE_OK":          4,
	}
)

func (x ReliabilityScore) Enum() *ReliabilityScore {
	p := new(ReliabilityScore)
	*p = x
	return p
}

func (x ReliabilityScore) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReliabilityScore) Descriptor() protoreflect.EnumDescriptor {
	return file_api_resources_k8s_v1_resources_proto_enumTypes[0].Descriptor()
}

func (ReliabilityScore) Type() protoreflect.EnumType {
	return &file_api_resources_k8s_v1_resources_proto_enumTypes[0]
}

func (x ReliabilityScore) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReliabilityScore.Descriptor instead.
func (ReliabilityScore) EnumDescriptor() ([]byte, []int) {
	return file_api_resources_k8s_v1_resources_proto_rawDescGZIP(), []int{0}
}

type ResourceStatus int32

const (
	ResourceStatus_RESOURCE_STATUS_UNSPECIFIED ResourceStatus = 0
	ResourceStatus_RESOURCE_STATUS_UNKNOWN     ResourceStatus = 1
	ResourceStatus_RESOURCE_STATUS_READY       ResourceStatus = 2
	ResourceStatus_RESOURCE_STATUS_PROGRESSING ResourceStatus = 3
	ResourceStatus_RESOURCE_STATUS_NOT_READY   ResourceStatus = 4
)

// Enum value maps for ResourceStatus.
var (
	ResourceStatus_name = map[int32]string{
		0: "RESOURCE_STATUS_UNSPECIFIED",
		1: "RESOURCE_STATUS_UNKNOWN",
		2: "RESOURCE_STATUS_READY",
		3: "RESOURCE_STATUS_PROGRESSING",
		4: "RESOURCE_STATUS_NOT_READY",
	}
	ResourceStatus_value = map[string]int32{
		"RESOURCE_STATUS_UNSPECIFIED": 0,
		"RESOURCE_STATUS_UNKNOWN":     1,
		"RESOURCE_STATUS_READY":       2,
		"RESOURCE_STATUS_PROGRESSING": 3,
		"RESOURCE_STATUS_NOT_READY":   4,
	}
)

func (x ResourceStatus) Enum() *ResourceStatus {
	p := new(ResourceStatus)
	*p = x
	return p
}

func (x ResourceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_resources_k8s_v1_resources_proto_enumTypes[1].Descriptor()
}

func (ResourceStatus) Type() protoreflect.EnumType {
	return &file_api_resources_k8s_v1_resources_proto_enumTypes[1]
}

func (x ResourceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceStatus.Descriptor instead.
func (ResourceStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_resources_k8s_v1_resources_proto_rawDescGZIP(), []int{1}
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Kind      string `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	Object    string `protobuf:"bytes,5,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_resources_k8s_v1_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_k8s_v1_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_api_resources_k8s_v1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Resource) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Resource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Resource) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

type Resources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*Resource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *Resources) Reset() {
	*x = Resources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_resources_k8s_v1_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resources) ProtoMessage() {}

func (x *Resources) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_k8s_v1_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resources.ProtoReflect.Descriptor instead.
func (*Resources) Descriptor() ([]byte, []int) {
	return file_api_resources_k8s_v1_resources_proto_rawDescGZIP(), []int{1}
}

func (x *Resources) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type ResourceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apiversion  string           `protobuf:"bytes,1,opt,name=apiversion,proto3" json:"apiversion,omitempty"`
	Kind        string           `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Namespace   string           `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name        string           `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Status      ResourceStatus   `protobuf:"varint,5,opt,name=status,proto3,enum=api.resources.k8s.v1.ResourceStatus" json:"status,omitempty"`
	Replicas    int32            `protobuf:"varint,6,opt,name=replicas,proto3" json:"replicas,omitempty"`
	Reliability ReliabilityScore `protobuf:"varint,7,opt,name=reliability,proto3,enum=api.resources.k8s.v1.ReliabilityScore" json:"reliability,omitempty"`
	Version     *string          `protobuf:"bytes,8,opt,name=version,proto3,oneof" json:"version,omitempty"`
}

func (x *ResourceInfo) Reset() {
	*x = ResourceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_resources_k8s_v1_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInfo) ProtoMessage() {}

func (x *ResourceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_k8s_v1_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInfo.ProtoReflect.Descriptor instead.
func (*ResourceInfo) Descriptor() ([]byte, []int) {
	return file_api_resources_k8s_v1_resources_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceInfo) GetApiversion() string {
	if x != nil {
		return x.Apiversion
	}
	return ""
}

func (x *ResourceInfo) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ResourceInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ResourceInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceInfo) GetStatus() ResourceStatus {
	if x != nil {
		return x.Status
	}
	return ResourceStatus_RESOURCE_STATUS_UNSPECIFIED
}

func (x *ResourceInfo) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *ResourceInfo) GetReliability() ReliabilityScore {
	if x != nil {
		return x.Reliability
	}
	return ReliabilityScore_RELIABILITY_SCORE_UNSPECIFIED
}

func (x *ResourceInfo) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status     ResourceStatus         `protobuf:"varint,2,opt,name=status,proto3,enum=api.resources.k8s.v1.ResourceStatus" json:"status,omitempty"`
	Roles      []string               `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	InternalIp *string                `protobuf:"bytes,4,opt,name=internal_ip,json=internalIp,proto3,oneof" json:"internal_ip,omitempty"`
	ExternalIp *string                `protobuf:"bytes,5,opt,name=external_ip,json=externalIp,proto3,oneof" json:"external_ip,omitempty"`
	Age        *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=age,proto3,oneof" json:"age,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_resources_k8s_v1_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_resources_k8s_v1_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_api_resources_k8s_v1_resources_proto_rawDescGZIP(), []int{3}
}

func (x *NodeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeInfo) GetStatus() ResourceStatus {
	if x != nil {
		return x.Status
	}
	return ResourceStatus_RESOURCE_STATUS_UNSPECIFIED
}

func (x *NodeInfo) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *NodeInfo) GetInternalIp() string {
	if x != nil && x.InternalIp != nil {
		return *x.InternalIp
	}
	return ""
}

func (x *NodeInfo) GetExternalIp() string {
	if x != nil && x.ExternalIp != nil {
		return *x.ExternalIp
	}
	return ""
}

func (x *NodeInfo) GetAge() *timestamppb.Timestamp {
	if x != nil {
		return x.Age
	}
	return nil
}

var File_api_resources_k8s_v1_resources_proto protoreflect.FileDescriptor

var file_api_resources_k8s_v1_resources_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x49, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x3c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0xc3, 0x02,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x70, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x48,
	0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x69, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x6c,
	0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x99, 0x02, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x88, 0x01, 0x01, 0x12, 0x24,
	0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x70, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x61, 0x67, 0x65, 0x2a,
	0xaa, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x4c, 0x49, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x5f, 0x44, 0x45, 0x47, 0x52, 0x41, 0x44, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x5f, 0x4f, 0x4b, 0x10, 0x04, 0x2a, 0xa9, 0x01, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x19, 0x0a,
	0x15, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x47,
	0x52, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x04, 0x42, 0xe9, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6f, 0x72, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x38,
	0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x52, 0x4b, 0xaa, 0x02, 0x14, 0x41, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x38, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x14, 0x41, 0x70, 0x69, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x5c, 0x4b, 0x38, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x20, 0x41, 0x70, 0x69, 0x5c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5c, 0x4b, 0x38, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x4b, 0x38, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_resources_k8s_v1_resources_proto_rawDescOnce sync.Once
	file_api_resources_k8s_v1_resources_proto_rawDescData = file_api_resources_k8s_v1_resources_proto_rawDesc
)

func file_api_resources_k8s_v1_resources_proto_rawDescGZIP() []byte {
	file_api_resources_k8s_v1_resources_proto_rawDescOnce.Do(func() {
		file_api_resources_k8s_v1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_resources_k8s_v1_resources_proto_rawDescData)
	})
	return file_api_resources_k8s_v1_resources_proto_rawDescData
}

var file_api_resources_k8s_v1_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_resources_k8s_v1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_resources_k8s_v1_resources_proto_goTypes = []interface{}{
	(ReliabilityScore)(0),         // 0: api.resources.k8s.v1.ReliabilityScore
	(ResourceStatus)(0),           // 1: api.resources.k8s.v1.ResourceStatus
	(*Resource)(nil),              // 2: api.resources.k8s.v1.Resource
	(*Resources)(nil),             // 3: api.resources.k8s.v1.Resources
	(*ResourceInfo)(nil),          // 4: api.resources.k8s.v1.ResourceInfo
	(*NodeInfo)(nil),              // 5: api.resources.k8s.v1.NodeInfo
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_api_resources_k8s_v1_resources_proto_depIdxs = []int32{
	2, // 0: api.resources.k8s.v1.Resources.resources:type_name -> api.resources.k8s.v1.Resource
	1, // 1: api.resources.k8s.v1.ResourceInfo.status:type_name -> api.resources.k8s.v1.ResourceStatus
	0, // 2: api.resources.k8s.v1.ResourceInfo.reliability:type_name -> api.resources.k8s.v1.ReliabilityScore
	1, // 3: api.resources.k8s.v1.NodeInfo.status:type_name -> api.resources.k8s.v1.ResourceStatus
	6, // 4: api.resources.k8s.v1.NodeInfo.age:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_resources_k8s_v1_resources_proto_init() }
func file_api_resources_k8s_v1_resources_proto_init() {
	if File_api_resources_k8s_v1_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_resources_k8s_v1_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_api_resources_k8s_v1_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resources); i {
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
		file_api_resources_k8s_v1_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceInfo); i {
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
		file_api_resources_k8s_v1_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
	file_api_resources_k8s_v1_resources_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_api_resources_k8s_v1_resources_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_resources_k8s_v1_resources_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_resources_k8s_v1_resources_proto_goTypes,
		DependencyIndexes: file_api_resources_k8s_v1_resources_proto_depIdxs,
		EnumInfos:         file_api_resources_k8s_v1_resources_proto_enumTypes,
		MessageInfos:      file_api_resources_k8s_v1_resources_proto_msgTypes,
	}.Build()
	File_api_resources_k8s_v1_resources_proto = out.File
	file_api_resources_k8s_v1_resources_proto_rawDesc = nil
	file_api_resources_k8s_v1_resources_proto_goTypes = nil
	file_api_resources_k8s_v1_resources_proto_depIdxs = nil
}
