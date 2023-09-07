// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/services/stats/stats.proto

package stats

import (
	context "context"
	common "github.com/koor-tech/data-control-center/gen/go/api/resources/common"
	stats "github.com/koor-tech/data-control-center/gen/go/api/resources/stats"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StatsService_GetClusterStats_FullMethodName     = "/stats.StatsService/GetClusterStats"
	StatsService_GetClusterResources_FullMethodName = "/stats.StatsService/GetClusterResources"
	StatsService_GetClusterNodes_FullMethodName     = "/stats.StatsService/GetClusterNodes"
)

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsServiceClient interface {
	GetClusterStats(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*stats.ClusterStats, error)
	GetClusterResources(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*ClusterResourcesResponse, error)
	GetClusterNodes(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*ClusterNodesResponse, error)
}

type statsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsServiceClient(cc grpc.ClientConnInterface) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) GetClusterStats(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*stats.ClusterStats, error) {
	out := new(stats.ClusterStats)
	err := c.cc.Invoke(ctx, StatsService_GetClusterStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetClusterResources(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*ClusterResourcesResponse, error) {
	out := new(ClusterResourcesResponse)
	err := c.cc.Invoke(ctx, StatsService_GetClusterResources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetClusterNodes(ctx context.Context, in *common.EmptyRequest, opts ...grpc.CallOption) (*ClusterNodesResponse, error) {
	out := new(ClusterNodesResponse)
	err := c.cc.Invoke(ctx, StatsService_GetClusterNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServiceServer is the server API for StatsService service.
// All implementations should embed UnimplementedStatsServiceServer
// for forward compatibility
type StatsServiceServer interface {
	GetClusterStats(context.Context, *common.EmptyRequest) (*stats.ClusterStats, error)
	GetClusterResources(context.Context, *common.EmptyRequest) (*ClusterResourcesResponse, error)
	GetClusterNodes(context.Context, *common.EmptyRequest) (*ClusterNodesResponse, error)
}

// UnimplementedStatsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStatsServiceServer struct {
}

func (UnimplementedStatsServiceServer) GetClusterStats(context.Context, *common.EmptyRequest) (*stats.ClusterStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterStats not implemented")
}
func (UnimplementedStatsServiceServer) GetClusterResources(context.Context, *common.EmptyRequest) (*ClusterResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterResources not implemented")
}
func (UnimplementedStatsServiceServer) GetClusterNodes(context.Context, *common.EmptyRequest) (*ClusterNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterNodes not implemented")
}

// UnsafeStatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServiceServer will
// result in compilation errors.
type UnsafeStatsServiceServer interface {
	mustEmbedUnimplementedStatsServiceServer()
}

func RegisterStatsServiceServer(s grpc.ServiceRegistrar, srv StatsServiceServer) {
	s.RegisterService(&StatsService_ServiceDesc, srv)
}

func _StatsService_GetClusterStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetClusterStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatsService_GetClusterStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetClusterStats(ctx, req.(*common.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetClusterResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetClusterResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatsService_GetClusterResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetClusterResources(ctx, req.(*common.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetClusterNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetClusterNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatsService_GetClusterNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetClusterNodes(ctx, req.(*common.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatsService_ServiceDesc is the grpc.ServiceDesc for StatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stats.StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClusterStats",
			Handler:    _StatsService_GetClusterStats_Handler,
		},
		{
			MethodName: "GetClusterResources",
			Handler:    _StatsService_GetClusterResources_Handler,
		},
		{
			MethodName: "GetClusterNodes",
			Handler:    _StatsService_GetClusterNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/services/stats/stats.proto",
}
