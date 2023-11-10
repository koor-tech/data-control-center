// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/services/ceph/v1/ceph.proto

package cephv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CephService_GetCephUsers_FullMethodName    = "/api.services.ceph.v1.CephService/GetCephUsers"
	CephService_CreateCephUsers_FullMethodName = "/api.services.ceph.v1.CephService/CreateCephUsers"
	CephService_DeleteCephUser_FullMethodName  = "/api.services.ceph.v1.CephService/DeleteCephUser"
)

// CephServiceClient is the client API for CephService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CephServiceClient interface {
	GetCephUsers(ctx context.Context, in *GetCephUsersRequest, opts ...grpc.CallOption) (*GetCephUsersResponse, error)
	CreateCephUsers(ctx context.Context, in *CreateCephUsersRequest, opts ...grpc.CallOption) (*CreateCephUsersResponse, error)
	DeleteCephUser(ctx context.Context, in *DeleteCephUserRequest, opts ...grpc.CallOption) (*DeleteCephUserResponse, error)
}

type cephServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCephServiceClient(cc grpc.ClientConnInterface) CephServiceClient {
	return &cephServiceClient{cc}
}

func (c *cephServiceClient) GetCephUsers(ctx context.Context, in *GetCephUsersRequest, opts ...grpc.CallOption) (*GetCephUsersResponse, error) {
	out := new(GetCephUsersResponse)
	err := c.cc.Invoke(ctx, CephService_GetCephUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cephServiceClient) CreateCephUsers(ctx context.Context, in *CreateCephUsersRequest, opts ...grpc.CallOption) (*CreateCephUsersResponse, error) {
	out := new(CreateCephUsersResponse)
	err := c.cc.Invoke(ctx, CephService_CreateCephUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cephServiceClient) DeleteCephUser(ctx context.Context, in *DeleteCephUserRequest, opts ...grpc.CallOption) (*DeleteCephUserResponse, error) {
	out := new(DeleteCephUserResponse)
	err := c.cc.Invoke(ctx, CephService_DeleteCephUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CephServiceServer is the server API for CephService service.
// All implementations should embed UnimplementedCephServiceServer
// for forward compatibility
type CephServiceServer interface {
	GetCephUsers(context.Context, *GetCephUsersRequest) (*GetCephUsersResponse, error)
	CreateCephUsers(context.Context, *CreateCephUsersRequest) (*CreateCephUsersResponse, error)
	DeleteCephUser(context.Context, *DeleteCephUserRequest) (*DeleteCephUserResponse, error)
}

// UnimplementedCephServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCephServiceServer struct {
}

func (UnimplementedCephServiceServer) GetCephUsers(context.Context, *GetCephUsersRequest) (*GetCephUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCephUsers not implemented")
}
func (UnimplementedCephServiceServer) CreateCephUsers(context.Context, *CreateCephUsersRequest) (*CreateCephUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCephUsers not implemented")
}
func (UnimplementedCephServiceServer) DeleteCephUser(context.Context, *DeleteCephUserRequest) (*DeleteCephUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCephUser not implemented")
}

// UnsafeCephServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CephServiceServer will
// result in compilation errors.
type UnsafeCephServiceServer interface {
	mustEmbedUnimplementedCephServiceServer()
}

func RegisterCephServiceServer(s grpc.ServiceRegistrar, srv CephServiceServer) {
	s.RegisterService(&CephService_ServiceDesc, srv)
}

func _CephService_GetCephUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCephUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CephServiceServer).GetCephUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CephService_GetCephUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CephServiceServer).GetCephUsers(ctx, req.(*GetCephUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CephService_CreateCephUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCephUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CephServiceServer).CreateCephUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CephService_CreateCephUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CephServiceServer).CreateCephUsers(ctx, req.(*CreateCephUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CephService_DeleteCephUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCephUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CephServiceServer).DeleteCephUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CephService_DeleteCephUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CephServiceServer).DeleteCephUser(ctx, req.(*DeleteCephUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CephService_ServiceDesc is the grpc.ServiceDesc for CephService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CephService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.services.ceph.v1.CephService",
	HandlerType: (*CephServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCephUsers",
			Handler:    _CephService_GetCephUsers_Handler,
		},
		{
			MethodName: "CreateCephUsers",
			Handler:    _CephService_CreateCephUsers_Handler,
		},
		{
			MethodName: "DeleteCephUser",
			Handler:    _CephService_DeleteCephUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/services/ceph/v1/ceph.proto",
}
