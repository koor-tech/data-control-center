// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/services/cluster/v1/cluster.proto

package clusterv1

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
	ClusterService_GetKoorCluster_FullMethodName        = "/api.services.cluster.v1.ClusterService/GetKoorCluster"
	ClusterService_GetTroubleshootReport_FullMethodName = "/api.services.cluster.v1.ClusterService/GetTroubleshootReport"
	ClusterService_GetNetworkTestStatus_FullMethodName  = "/api.services.cluster.v1.ClusterService/GetNetworkTestStatus"
	ClusterService_StartNetworkTest_FullMethodName      = "/api.services.cluster.v1.ClusterService/StartNetworkTest"
	ClusterService_CancelNetworkTest_FullMethodName     = "/api.services.cluster.v1.ClusterService/CancelNetworkTest"
	ClusterService_GetNetworkTestResults_FullMethodName = "/api.services.cluster.v1.ClusterService/GetNetworkTestResults"
	ClusterService_SetScrubbingSchedule_FullMethodName  = "/api.services.cluster.v1.ClusterService/SetScrubbingSchedule"
	ClusterService_GetResources_FullMethodName          = "/api.services.cluster.v1.ClusterService/GetResources"
	ClusterService_SaveResources_FullMethodName         = "/api.services.cluster.v1.ClusterService/SaveResources"
)

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	GetKoorCluster(ctx context.Context, in *GetKoorClusterRequest, opts ...grpc.CallOption) (*GetKoorClusterResponse, error)
	GetTroubleshootReport(ctx context.Context, in *GetTroubleshootReportRequest, opts ...grpc.CallOption) (*GetTroubleshootReportResponse, error)
	GetNetworkTestStatus(ctx context.Context, in *GetNetworkTestStatusRequest, opts ...grpc.CallOption) (*GetNetworkTestStatusResponse, error)
	StartNetworkTest(ctx context.Context, in *StartNetworkTestRequest, opts ...grpc.CallOption) (*StartNetworkTestResponse, error)
	CancelNetworkTest(ctx context.Context, in *CancelNetworkTestRequest, opts ...grpc.CallOption) (*CancelNetworkTestResponse, error)
	GetNetworkTestResults(ctx context.Context, in *GetNetworkTestResultsRequest, opts ...grpc.CallOption) (*GetNetworkTestResultsResponse, error)
	SetScrubbingSchedule(ctx context.Context, in *SetScrubbingScheduleRequest, opts ...grpc.CallOption) (*SetScrubbingScheduleResponse, error)
	GetResources(ctx context.Context, in *GetResourcesRequest, opts ...grpc.CallOption) (*GetResourcesResponse, error)
	SaveResources(ctx context.Context, in *SaveResourcesRequest, opts ...grpc.CallOption) (*SaveResourcesResponse, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) GetKoorCluster(ctx context.Context, in *GetKoorClusterRequest, opts ...grpc.CallOption) (*GetKoorClusterResponse, error) {
	out := new(GetKoorClusterResponse)
	err := c.cc.Invoke(ctx, ClusterService_GetKoorCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetTroubleshootReport(ctx context.Context, in *GetTroubleshootReportRequest, opts ...grpc.CallOption) (*GetTroubleshootReportResponse, error) {
	out := new(GetTroubleshootReportResponse)
	err := c.cc.Invoke(ctx, ClusterService_GetTroubleshootReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetNetworkTestStatus(ctx context.Context, in *GetNetworkTestStatusRequest, opts ...grpc.CallOption) (*GetNetworkTestStatusResponse, error) {
	out := new(GetNetworkTestStatusResponse)
	err := c.cc.Invoke(ctx, ClusterService_GetNetworkTestStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) StartNetworkTest(ctx context.Context, in *StartNetworkTestRequest, opts ...grpc.CallOption) (*StartNetworkTestResponse, error) {
	out := new(StartNetworkTestResponse)
	err := c.cc.Invoke(ctx, ClusterService_StartNetworkTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) CancelNetworkTest(ctx context.Context, in *CancelNetworkTestRequest, opts ...grpc.CallOption) (*CancelNetworkTestResponse, error) {
	out := new(CancelNetworkTestResponse)
	err := c.cc.Invoke(ctx, ClusterService_CancelNetworkTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetNetworkTestResults(ctx context.Context, in *GetNetworkTestResultsRequest, opts ...grpc.CallOption) (*GetNetworkTestResultsResponse, error) {
	out := new(GetNetworkTestResultsResponse)
	err := c.cc.Invoke(ctx, ClusterService_GetNetworkTestResults_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) SetScrubbingSchedule(ctx context.Context, in *SetScrubbingScheduleRequest, opts ...grpc.CallOption) (*SetScrubbingScheduleResponse, error) {
	out := new(SetScrubbingScheduleResponse)
	err := c.cc.Invoke(ctx, ClusterService_SetScrubbingSchedule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetResources(ctx context.Context, in *GetResourcesRequest, opts ...grpc.CallOption) (*GetResourcesResponse, error) {
	out := new(GetResourcesResponse)
	err := c.cc.Invoke(ctx, ClusterService_GetResources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) SaveResources(ctx context.Context, in *SaveResourcesRequest, opts ...grpc.CallOption) (*SaveResourcesResponse, error) {
	out := new(SaveResourcesResponse)
	err := c.cc.Invoke(ctx, ClusterService_SaveResources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations should embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	GetKoorCluster(context.Context, *GetKoorClusterRequest) (*GetKoorClusterResponse, error)
	GetTroubleshootReport(context.Context, *GetTroubleshootReportRequest) (*GetTroubleshootReportResponse, error)
	GetNetworkTestStatus(context.Context, *GetNetworkTestStatusRequest) (*GetNetworkTestStatusResponse, error)
	StartNetworkTest(context.Context, *StartNetworkTestRequest) (*StartNetworkTestResponse, error)
	CancelNetworkTest(context.Context, *CancelNetworkTestRequest) (*CancelNetworkTestResponse, error)
	GetNetworkTestResults(context.Context, *GetNetworkTestResultsRequest) (*GetNetworkTestResultsResponse, error)
	SetScrubbingSchedule(context.Context, *SetScrubbingScheduleRequest) (*SetScrubbingScheduleResponse, error)
	GetResources(context.Context, *GetResourcesRequest) (*GetResourcesResponse, error)
	SaveResources(context.Context, *SaveResourcesRequest) (*SaveResourcesResponse, error)
}

// UnimplementedClusterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (UnimplementedClusterServiceServer) GetKoorCluster(context.Context, *GetKoorClusterRequest) (*GetKoorClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKoorCluster not implemented")
}
func (UnimplementedClusterServiceServer) GetTroubleshootReport(context.Context, *GetTroubleshootReportRequest) (*GetTroubleshootReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTroubleshootReport not implemented")
}
func (UnimplementedClusterServiceServer) GetNetworkTestStatus(context.Context, *GetNetworkTestStatusRequest) (*GetNetworkTestStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkTestStatus not implemented")
}
func (UnimplementedClusterServiceServer) StartNetworkTest(context.Context, *StartNetworkTestRequest) (*StartNetworkTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNetworkTest not implemented")
}
func (UnimplementedClusterServiceServer) CancelNetworkTest(context.Context, *CancelNetworkTestRequest) (*CancelNetworkTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelNetworkTest not implemented")
}
func (UnimplementedClusterServiceServer) GetNetworkTestResults(context.Context, *GetNetworkTestResultsRequest) (*GetNetworkTestResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkTestResults not implemented")
}
func (UnimplementedClusterServiceServer) SetScrubbingSchedule(context.Context, *SetScrubbingScheduleRequest) (*SetScrubbingScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetScrubbingSchedule not implemented")
}
func (UnimplementedClusterServiceServer) GetResources(context.Context, *GetResourcesRequest) (*GetResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResources not implemented")
}
func (UnimplementedClusterServiceServer) SaveResources(context.Context, *SaveResourcesRequest) (*SaveResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveResources not implemented")
}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
	s.RegisterService(&ClusterService_ServiceDesc, srv)
}

func _ClusterService_GetKoorCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKoorClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetKoorCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_GetKoorCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetKoorCluster(ctx, req.(*GetKoorClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetTroubleshootReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTroubleshootReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetTroubleshootReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_GetTroubleshootReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetTroubleshootReport(ctx, req.(*GetTroubleshootReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetNetworkTestStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkTestStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetNetworkTestStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_GetNetworkTestStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetNetworkTestStatus(ctx, req.(*GetNetworkTestStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_StartNetworkTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartNetworkTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).StartNetworkTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_StartNetworkTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).StartNetworkTest(ctx, req.(*StartNetworkTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_CancelNetworkTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelNetworkTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).CancelNetworkTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_CancelNetworkTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).CancelNetworkTest(ctx, req.(*CancelNetworkTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetNetworkTestResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkTestResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetNetworkTestResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_GetNetworkTestResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetNetworkTestResults(ctx, req.(*GetNetworkTestResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_SetScrubbingSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetScrubbingScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).SetScrubbingSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_SetScrubbingSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).SetScrubbingSchedule(ctx, req.(*SetScrubbingScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_GetResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetResources(ctx, req.(*GetResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_SaveResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).SaveResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_SaveResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).SaveResources(ctx, req.(*SaveResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.services.cluster.v1.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKoorCluster",
			Handler:    _ClusterService_GetKoorCluster_Handler,
		},
		{
			MethodName: "GetTroubleshootReport",
			Handler:    _ClusterService_GetTroubleshootReport_Handler,
		},
		{
			MethodName: "GetNetworkTestStatus",
			Handler:    _ClusterService_GetNetworkTestStatus_Handler,
		},
		{
			MethodName: "StartNetworkTest",
			Handler:    _ClusterService_StartNetworkTest_Handler,
		},
		{
			MethodName: "CancelNetworkTest",
			Handler:    _ClusterService_CancelNetworkTest_Handler,
		},
		{
			MethodName: "GetNetworkTestResults",
			Handler:    _ClusterService_GetNetworkTestResults_Handler,
		},
		{
			MethodName: "SetScrubbingSchedule",
			Handler:    _ClusterService_SetScrubbingSchedule_Handler,
		},
		{
			MethodName: "GetResources",
			Handler:    _ClusterService_GetResources_Handler,
		},
		{
			MethodName: "SaveResources",
			Handler:    _ClusterService_SaveResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/services/cluster/v1/cluster.proto",
}
