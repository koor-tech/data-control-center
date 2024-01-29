// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/services/ceph/v1/users.proto

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
	CephUsersService_ListCephUsers_FullMethodName  = "/api.services.ceph.v1.CephUsersService/ListCephUsers"
	CephUsersService_CreateCephUser_FullMethodName = "/api.services.ceph.v1.CephUsersService/CreateCephUser"
	CephUsersService_DeleteCephUser_FullMethodName = "/api.services.ceph.v1.CephUsersService/DeleteCephUser"
)

// CephUsersServiceClient is the client API for CephUsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CephUsersServiceClient interface {
	ListCephUsers(ctx context.Context, in *ListCephUsersRequest, opts ...grpc.CallOption) (*ListCephUsersResponse, error)
	CreateCephUser(ctx context.Context, in *CreateCephUserRequest, opts ...grpc.CallOption) (*CreateCephUserResponse, error)
	DeleteCephUser(ctx context.Context, in *DeleteCephUserRequest, opts ...grpc.CallOption) (*DeleteCephUserResponse, error)
}

type cephUsersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCephUsersServiceClient(cc grpc.ClientConnInterface) CephUsersServiceClient {
	return &cephUsersServiceClient{cc}
}

func (c *cephUsersServiceClient) ListCephUsers(ctx context.Context, in *ListCephUsersRequest, opts ...grpc.CallOption) (*ListCephUsersResponse, error) {
	out := new(ListCephUsersResponse)
	err := c.cc.Invoke(ctx, CephUsersService_ListCephUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cephUsersServiceClient) CreateCephUser(ctx context.Context, in *CreateCephUserRequest, opts ...grpc.CallOption) (*CreateCephUserResponse, error) {
	out := new(CreateCephUserResponse)
	err := c.cc.Invoke(ctx, CephUsersService_CreateCephUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cephUsersServiceClient) DeleteCephUser(ctx context.Context, in *DeleteCephUserRequest, opts ...grpc.CallOption) (*DeleteCephUserResponse, error) {
	out := new(DeleteCephUserResponse)
	err := c.cc.Invoke(ctx, CephUsersService_DeleteCephUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CephUsersServiceServer is the server API for CephUsersService service.
// All implementations should embed UnimplementedCephUsersServiceServer
// for forward compatibility
type CephUsersServiceServer interface {
	ListCephUsers(context.Context, *ListCephUsersRequest) (*ListCephUsersResponse, error)
	CreateCephUser(context.Context, *CreateCephUserRequest) (*CreateCephUserResponse, error)
	DeleteCephUser(context.Context, *DeleteCephUserRequest) (*DeleteCephUserResponse, error)
}

// UnimplementedCephUsersServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCephUsersServiceServer struct {
}

func (UnimplementedCephUsersServiceServer) ListCephUsers(context.Context, *ListCephUsersRequest) (*ListCephUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCephUsers not implemented")
}
func (UnimplementedCephUsersServiceServer) CreateCephUser(context.Context, *CreateCephUserRequest) (*CreateCephUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCephUser not implemented")
}
func (UnimplementedCephUsersServiceServer) DeleteCephUser(context.Context, *DeleteCephUserRequest) (*DeleteCephUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCephUser not implemented")
}

// UnsafeCephUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CephUsersServiceServer will
// result in compilation errors.
type UnsafeCephUsersServiceServer interface {
	mustEmbedUnimplementedCephUsersServiceServer()
}

func RegisterCephUsersServiceServer(s grpc.ServiceRegistrar, srv CephUsersServiceServer) {
	s.RegisterService(&CephUsersService_ServiceDesc, srv)
}

func _CephUsersService_ListCephUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCephUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CephUsersServiceServer).ListCephUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CephUsersService_ListCephUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CephUsersServiceServer).ListCephUsers(ctx, req.(*ListCephUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CephUsersService_CreateCephUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCephUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CephUsersServiceServer).CreateCephUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CephUsersService_CreateCephUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CephUsersServiceServer).CreateCephUser(ctx, req.(*CreateCephUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CephUsersService_DeleteCephUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCephUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CephUsersServiceServer).DeleteCephUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CephUsersService_DeleteCephUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CephUsersServiceServer).DeleteCephUser(ctx, req.(*DeleteCephUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CephUsersService_ServiceDesc is the grpc.ServiceDesc for CephUsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CephUsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.services.ceph.v1.CephUsersService",
	HandlerType: (*CephUsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCephUsers",
			Handler:    _CephUsersService_ListCephUsers_Handler,
		},
		{
			MethodName: "CreateCephUser",
			Handler:    _CephUsersService_CreateCephUser_Handler,
		},
		{
			MethodName: "DeleteCephUser",
			Handler:    _CephUsersService_DeleteCephUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/services/ceph/v1/users.proto",
}