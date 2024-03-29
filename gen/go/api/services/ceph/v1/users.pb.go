// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api/services/ceph/v1/users.proto

package cephv1

import (
	v1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
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

type ListCephUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCephUsersRequest) Reset() {
	*x = ListCephUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_ceph_v1_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCephUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCephUsersRequest) ProtoMessage() {}

func (x *ListCephUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_ceph_v1_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCephUsersRequest.ProtoReflect.Descriptor instead.
func (*ListCephUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_services_ceph_v1_users_proto_rawDescGZIP(), []int{0}
}

type ListCephUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CephUsers []*v1.User `protobuf:"bytes,1,rep,name=ceph_users,json=cephUsers,proto3" json:"ceph_users,omitempty"`
}

func (x *ListCephUsersResponse) Reset() {
	*x = ListCephUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_ceph_v1_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCephUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCephUsersResponse) ProtoMessage() {}

func (x *ListCephUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_ceph_v1_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCephUsersResponse.ProtoReflect.Descriptor instead.
func (*ListCephUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_services_ceph_v1_users_proto_rawDescGZIP(), []int{1}
}

func (x *ListCephUsersResponse) GetCephUsers() []*v1.User {
	if x != nil {
		return x.CephUsers
	}
	return nil
}

type CreateCephUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CephUser *v1.User `protobuf:"bytes,1,opt,name=ceph_user,json=cephUser,proto3" json:"ceph_user,omitempty"`
}

func (x *CreateCephUserRequest) Reset() {
	*x = CreateCephUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_ceph_v1_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCephUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCephUserRequest) ProtoMessage() {}

func (x *CreateCephUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_ceph_v1_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCephUserRequest.ProtoReflect.Descriptor instead.
func (*CreateCephUserRequest) Descriptor() ([]byte, []int) {
	return file_api_services_ceph_v1_users_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCephUserRequest) GetCephUser() *v1.User {
	if x != nil {
		return x.CephUser
	}
	return nil
}

type CreateCephUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CephUser *v1.User `protobuf:"bytes,1,opt,name=ceph_user,json=cephUser,proto3" json:"ceph_user,omitempty"`
}

func (x *CreateCephUserResponse) Reset() {
	*x = CreateCephUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_ceph_v1_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCephUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCephUserResponse) ProtoMessage() {}

func (x *CreateCephUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_ceph_v1_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCephUserResponse.ProtoReflect.Descriptor instead.
func (*CreateCephUserResponse) Descriptor() ([]byte, []int) {
	return file_api_services_ceph_v1_users_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCephUserResponse) GetCephUser() *v1.User {
	if x != nil {
		return x.CephUser
	}
	return nil
}

type DeleteCephUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *DeleteCephUserRequest) Reset() {
	*x = DeleteCephUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_ceph_v1_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCephUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCephUserRequest) ProtoMessage() {}

func (x *DeleteCephUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_ceph_v1_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCephUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteCephUserRequest) Descriptor() ([]byte, []int) {
	return file_api_services_ceph_v1_users_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCephUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type DeleteCephUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteCephUserResponse) Reset() {
	*x = DeleteCephUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_ceph_v1_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCephUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCephUserResponse) ProtoMessage() {}

func (x *DeleteCephUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_ceph_v1_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCephUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteCephUserResponse) Descriptor() ([]byte, []int) {
	return file_api_services_ceph_v1_users_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCephUserResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_api_services_ceph_v1_users_proto protoreflect.FileDescriptor

var file_api_services_ceph_v1_users_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x65, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x65, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x70, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a,
	0x63, 0x65, 0x70, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x63,
	0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x51, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x65, 0x70, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x08, 0x63, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x22, 0x52, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x65, 0x70, 0x68, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x63, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x33, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65,
	0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xdc, 0x02, 0x0a, 0x10, 0x43, 0x65, 0x70, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x65, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xe6, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x65, 0x70, 0x68, 0x2e,
	0x76, 0x31, 0x42, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6f,
	0x72, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x65, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x65, 0x70, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x41, 0x53, 0x43, 0xaa, 0x02, 0x14, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x65, 0x70, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x41, 0x70, 0x69,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x43, 0x65, 0x70, 0x68, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x20, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x5c, 0x43, 0x65, 0x70, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x43, 0x65, 0x70, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_services_ceph_v1_users_proto_rawDescOnce sync.Once
	file_api_services_ceph_v1_users_proto_rawDescData = file_api_services_ceph_v1_users_proto_rawDesc
)

func file_api_services_ceph_v1_users_proto_rawDescGZIP() []byte {
	file_api_services_ceph_v1_users_proto_rawDescOnce.Do(func() {
		file_api_services_ceph_v1_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_ceph_v1_users_proto_rawDescData)
	})
	return file_api_services_ceph_v1_users_proto_rawDescData
}

var file_api_services_ceph_v1_users_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_services_ceph_v1_users_proto_goTypes = []interface{}{
	(*ListCephUsersRequest)(nil),   // 0: api.services.ceph.v1.ListCephUsersRequest
	(*ListCephUsersResponse)(nil),  // 1: api.services.ceph.v1.ListCephUsersResponse
	(*CreateCephUserRequest)(nil),  // 2: api.services.ceph.v1.CreateCephUserRequest
	(*CreateCephUserResponse)(nil), // 3: api.services.ceph.v1.CreateCephUserResponse
	(*DeleteCephUserRequest)(nil),  // 4: api.services.ceph.v1.DeleteCephUserRequest
	(*DeleteCephUserResponse)(nil), // 5: api.services.ceph.v1.DeleteCephUserResponse
	(*v1.User)(nil),                // 6: api.resources.ceph.v1.User
}
var file_api_services_ceph_v1_users_proto_depIdxs = []int32{
	6, // 0: api.services.ceph.v1.ListCephUsersResponse.ceph_users:type_name -> api.resources.ceph.v1.User
	6, // 1: api.services.ceph.v1.CreateCephUserRequest.ceph_user:type_name -> api.resources.ceph.v1.User
	6, // 2: api.services.ceph.v1.CreateCephUserResponse.ceph_user:type_name -> api.resources.ceph.v1.User
	0, // 3: api.services.ceph.v1.CephUsersService.ListCephUsers:input_type -> api.services.ceph.v1.ListCephUsersRequest
	2, // 4: api.services.ceph.v1.CephUsersService.CreateCephUser:input_type -> api.services.ceph.v1.CreateCephUserRequest
	4, // 5: api.services.ceph.v1.CephUsersService.DeleteCephUser:input_type -> api.services.ceph.v1.DeleteCephUserRequest
	1, // 6: api.services.ceph.v1.CephUsersService.ListCephUsers:output_type -> api.services.ceph.v1.ListCephUsersResponse
	3, // 7: api.services.ceph.v1.CephUsersService.CreateCephUser:output_type -> api.services.ceph.v1.CreateCephUserResponse
	5, // 8: api.services.ceph.v1.CephUsersService.DeleteCephUser:output_type -> api.services.ceph.v1.DeleteCephUserResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_services_ceph_v1_users_proto_init() }
func file_api_services_ceph_v1_users_proto_init() {
	if File_api_services_ceph_v1_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_ceph_v1_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCephUsersRequest); i {
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
		file_api_services_ceph_v1_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCephUsersResponse); i {
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
		file_api_services_ceph_v1_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCephUserRequest); i {
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
		file_api_services_ceph_v1_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCephUserResponse); i {
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
		file_api_services_ceph_v1_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCephUserRequest); i {
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
		file_api_services_ceph_v1_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCephUserResponse); i {
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
			RawDescriptor: file_api_services_ceph_v1_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_services_ceph_v1_users_proto_goTypes,
		DependencyIndexes: file_api_services_ceph_v1_users_proto_depIdxs,
		MessageInfos:      file_api_services_ceph_v1_users_proto_msgTypes,
	}.Build()
	File_api_services_ceph_v1_users_proto = out.File
	file_api_services_ceph_v1_users_proto_rawDesc = nil
	file_api_services_ceph_v1_users_proto_goTypes = nil
	file_api_services_ceph_v1_users_proto_depIdxs = nil
}
