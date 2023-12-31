// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/services/cluster/v1/cluster.proto

package clusterv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/koor-tech/data-control-center/gen/go/api/services/cluster/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// ClusterServiceName is the fully-qualified name of the ClusterService service.
	ClusterServiceName = "api.services.cluster.v1.ClusterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ClusterServiceGetKoorClusterProcedure is the fully-qualified name of the ClusterService's
	// GetKoorCluster RPC.
	ClusterServiceGetKoorClusterProcedure = "/api.services.cluster.v1.ClusterService/GetKoorCluster"
	// ClusterServiceGetTroubleshootReportProcedure is the fully-qualified name of the ClusterService's
	// GetTroubleshootReport RPC.
	ClusterServiceGetTroubleshootReportProcedure = "/api.services.cluster.v1.ClusterService/GetTroubleshootReport"
	// ClusterServiceGetNetworkTestStatusProcedure is the fully-qualified name of the ClusterService's
	// GetNetworkTestStatus RPC.
	ClusterServiceGetNetworkTestStatusProcedure = "/api.services.cluster.v1.ClusterService/GetNetworkTestStatus"
	// ClusterServiceStartNetworkTestProcedure is the fully-qualified name of the ClusterService's
	// StartNetworkTest RPC.
	ClusterServiceStartNetworkTestProcedure = "/api.services.cluster.v1.ClusterService/StartNetworkTest"
	// ClusterServiceCancelNetworkTestProcedure is the fully-qualified name of the ClusterService's
	// CancelNetworkTest RPC.
	ClusterServiceCancelNetworkTestProcedure = "/api.services.cluster.v1.ClusterService/CancelNetworkTest"
	// ClusterServiceGetNetworkTestResultsProcedure is the fully-qualified name of the ClusterService's
	// GetNetworkTestResults RPC.
	ClusterServiceGetNetworkTestResultsProcedure = "/api.services.cluster.v1.ClusterService/GetNetworkTestResults"
	// ClusterServiceSetScrubbingScheduleProcedure is the fully-qualified name of the ClusterService's
	// SetScrubbingSchedule RPC.
	ClusterServiceSetScrubbingScheduleProcedure = "/api.services.cluster.v1.ClusterService/SetScrubbingSchedule"
	// ClusterServiceGetResourcesProcedure is the fully-qualified name of the ClusterService's
	// GetResources RPC.
	ClusterServiceGetResourcesProcedure = "/api.services.cluster.v1.ClusterService/GetResources"
	// ClusterServiceSaveResourcesProcedure is the fully-qualified name of the ClusterService's
	// SaveResources RPC.
	ClusterServiceSaveResourcesProcedure = "/api.services.cluster.v1.ClusterService/SaveResources"
)

// ClusterServiceClient is a client for the api.services.cluster.v1.ClusterService service.
type ClusterServiceClient interface {
	GetKoorCluster(context.Context, *connect.Request[v1.GetKoorClusterRequest]) (*connect.Response[v1.GetKoorClusterResponse], error)
	GetTroubleshootReport(context.Context, *connect.Request[v1.GetTroubleshootReportRequest]) (*connect.Response[v1.GetTroubleshootReportResponse], error)
	GetNetworkTestStatus(context.Context, *connect.Request[v1.GetNetworkTestStatusRequest]) (*connect.Response[v1.GetNetworkTestStatusResponse], error)
	StartNetworkTest(context.Context, *connect.Request[v1.StartNetworkTestRequest]) (*connect.Response[v1.StartNetworkTestResponse], error)
	CancelNetworkTest(context.Context, *connect.Request[v1.CancelNetworkTestRequest]) (*connect.Response[v1.CancelNetworkTestResponse], error)
	GetNetworkTestResults(context.Context, *connect.Request[v1.GetNetworkTestResultsRequest]) (*connect.Response[v1.GetNetworkTestResultsResponse], error)
	SetScrubbingSchedule(context.Context, *connect.Request[v1.SetScrubbingScheduleRequest]) (*connect.Response[v1.SetScrubbingScheduleResponse], error)
	GetResources(context.Context, *connect.Request[v1.GetResourcesRequest]) (*connect.Response[v1.GetResourcesResponse], error)
	SaveResources(context.Context, *connect.Request[v1.SaveResourcesRequest]) (*connect.Response[v1.SaveResourcesResponse], error)
}

// NewClusterServiceClient constructs a client for the api.services.cluster.v1.ClusterService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewClusterServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ClusterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &clusterServiceClient{
		getKoorCluster: connect.NewClient[v1.GetKoorClusterRequest, v1.GetKoorClusterResponse](
			httpClient,
			baseURL+ClusterServiceGetKoorClusterProcedure,
			opts...,
		),
		getTroubleshootReport: connect.NewClient[v1.GetTroubleshootReportRequest, v1.GetTroubleshootReportResponse](
			httpClient,
			baseURL+ClusterServiceGetTroubleshootReportProcedure,
			opts...,
		),
		getNetworkTestStatus: connect.NewClient[v1.GetNetworkTestStatusRequest, v1.GetNetworkTestStatusResponse](
			httpClient,
			baseURL+ClusterServiceGetNetworkTestStatusProcedure,
			opts...,
		),
		startNetworkTest: connect.NewClient[v1.StartNetworkTestRequest, v1.StartNetworkTestResponse](
			httpClient,
			baseURL+ClusterServiceStartNetworkTestProcedure,
			opts...,
		),
		cancelNetworkTest: connect.NewClient[v1.CancelNetworkTestRequest, v1.CancelNetworkTestResponse](
			httpClient,
			baseURL+ClusterServiceCancelNetworkTestProcedure,
			opts...,
		),
		getNetworkTestResults: connect.NewClient[v1.GetNetworkTestResultsRequest, v1.GetNetworkTestResultsResponse](
			httpClient,
			baseURL+ClusterServiceGetNetworkTestResultsProcedure,
			opts...,
		),
		setScrubbingSchedule: connect.NewClient[v1.SetScrubbingScheduleRequest, v1.SetScrubbingScheduleResponse](
			httpClient,
			baseURL+ClusterServiceSetScrubbingScheduleProcedure,
			opts...,
		),
		getResources: connect.NewClient[v1.GetResourcesRequest, v1.GetResourcesResponse](
			httpClient,
			baseURL+ClusterServiceGetResourcesProcedure,
			opts...,
		),
		saveResources: connect.NewClient[v1.SaveResourcesRequest, v1.SaveResourcesResponse](
			httpClient,
			baseURL+ClusterServiceSaveResourcesProcedure,
			opts...,
		),
	}
}

// clusterServiceClient implements ClusterServiceClient.
type clusterServiceClient struct {
	getKoorCluster        *connect.Client[v1.GetKoorClusterRequest, v1.GetKoorClusterResponse]
	getTroubleshootReport *connect.Client[v1.GetTroubleshootReportRequest, v1.GetTroubleshootReportResponse]
	getNetworkTestStatus  *connect.Client[v1.GetNetworkTestStatusRequest, v1.GetNetworkTestStatusResponse]
	startNetworkTest      *connect.Client[v1.StartNetworkTestRequest, v1.StartNetworkTestResponse]
	cancelNetworkTest     *connect.Client[v1.CancelNetworkTestRequest, v1.CancelNetworkTestResponse]
	getNetworkTestResults *connect.Client[v1.GetNetworkTestResultsRequest, v1.GetNetworkTestResultsResponse]
	setScrubbingSchedule  *connect.Client[v1.SetScrubbingScheduleRequest, v1.SetScrubbingScheduleResponse]
	getResources          *connect.Client[v1.GetResourcesRequest, v1.GetResourcesResponse]
	saveResources         *connect.Client[v1.SaveResourcesRequest, v1.SaveResourcesResponse]
}

// GetKoorCluster calls api.services.cluster.v1.ClusterService.GetKoorCluster.
func (c *clusterServiceClient) GetKoorCluster(ctx context.Context, req *connect.Request[v1.GetKoorClusterRequest]) (*connect.Response[v1.GetKoorClusterResponse], error) {
	return c.getKoorCluster.CallUnary(ctx, req)
}

// GetTroubleshootReport calls api.services.cluster.v1.ClusterService.GetTroubleshootReport.
func (c *clusterServiceClient) GetTroubleshootReport(ctx context.Context, req *connect.Request[v1.GetTroubleshootReportRequest]) (*connect.Response[v1.GetTroubleshootReportResponse], error) {
	return c.getTroubleshootReport.CallUnary(ctx, req)
}

// GetNetworkTestStatus calls api.services.cluster.v1.ClusterService.GetNetworkTestStatus.
func (c *clusterServiceClient) GetNetworkTestStatus(ctx context.Context, req *connect.Request[v1.GetNetworkTestStatusRequest]) (*connect.Response[v1.GetNetworkTestStatusResponse], error) {
	return c.getNetworkTestStatus.CallUnary(ctx, req)
}

// StartNetworkTest calls api.services.cluster.v1.ClusterService.StartNetworkTest.
func (c *clusterServiceClient) StartNetworkTest(ctx context.Context, req *connect.Request[v1.StartNetworkTestRequest]) (*connect.Response[v1.StartNetworkTestResponse], error) {
	return c.startNetworkTest.CallUnary(ctx, req)
}

// CancelNetworkTest calls api.services.cluster.v1.ClusterService.CancelNetworkTest.
func (c *clusterServiceClient) CancelNetworkTest(ctx context.Context, req *connect.Request[v1.CancelNetworkTestRequest]) (*connect.Response[v1.CancelNetworkTestResponse], error) {
	return c.cancelNetworkTest.CallUnary(ctx, req)
}

// GetNetworkTestResults calls api.services.cluster.v1.ClusterService.GetNetworkTestResults.
func (c *clusterServiceClient) GetNetworkTestResults(ctx context.Context, req *connect.Request[v1.GetNetworkTestResultsRequest]) (*connect.Response[v1.GetNetworkTestResultsResponse], error) {
	return c.getNetworkTestResults.CallUnary(ctx, req)
}

// SetScrubbingSchedule calls api.services.cluster.v1.ClusterService.SetScrubbingSchedule.
func (c *clusterServiceClient) SetScrubbingSchedule(ctx context.Context, req *connect.Request[v1.SetScrubbingScheduleRequest]) (*connect.Response[v1.SetScrubbingScheduleResponse], error) {
	return c.setScrubbingSchedule.CallUnary(ctx, req)
}

// GetResources calls api.services.cluster.v1.ClusterService.GetResources.
func (c *clusterServiceClient) GetResources(ctx context.Context, req *connect.Request[v1.GetResourcesRequest]) (*connect.Response[v1.GetResourcesResponse], error) {
	return c.getResources.CallUnary(ctx, req)
}

// SaveResources calls api.services.cluster.v1.ClusterService.SaveResources.
func (c *clusterServiceClient) SaveResources(ctx context.Context, req *connect.Request[v1.SaveResourcesRequest]) (*connect.Response[v1.SaveResourcesResponse], error) {
	return c.saveResources.CallUnary(ctx, req)
}

// ClusterServiceHandler is an implementation of the api.services.cluster.v1.ClusterService service.
type ClusterServiceHandler interface {
	GetKoorCluster(context.Context, *connect.Request[v1.GetKoorClusterRequest]) (*connect.Response[v1.GetKoorClusterResponse], error)
	GetTroubleshootReport(context.Context, *connect.Request[v1.GetTroubleshootReportRequest]) (*connect.Response[v1.GetTroubleshootReportResponse], error)
	GetNetworkTestStatus(context.Context, *connect.Request[v1.GetNetworkTestStatusRequest]) (*connect.Response[v1.GetNetworkTestStatusResponse], error)
	StartNetworkTest(context.Context, *connect.Request[v1.StartNetworkTestRequest]) (*connect.Response[v1.StartNetworkTestResponse], error)
	CancelNetworkTest(context.Context, *connect.Request[v1.CancelNetworkTestRequest]) (*connect.Response[v1.CancelNetworkTestResponse], error)
	GetNetworkTestResults(context.Context, *connect.Request[v1.GetNetworkTestResultsRequest]) (*connect.Response[v1.GetNetworkTestResultsResponse], error)
	SetScrubbingSchedule(context.Context, *connect.Request[v1.SetScrubbingScheduleRequest]) (*connect.Response[v1.SetScrubbingScheduleResponse], error)
	GetResources(context.Context, *connect.Request[v1.GetResourcesRequest]) (*connect.Response[v1.GetResourcesResponse], error)
	SaveResources(context.Context, *connect.Request[v1.SaveResourcesRequest]) (*connect.Response[v1.SaveResourcesResponse], error)
}

// NewClusterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewClusterServiceHandler(svc ClusterServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	clusterServiceGetKoorClusterHandler := connect.NewUnaryHandler(
		ClusterServiceGetKoorClusterProcedure,
		svc.GetKoorCluster,
		opts...,
	)
	clusterServiceGetTroubleshootReportHandler := connect.NewUnaryHandler(
		ClusterServiceGetTroubleshootReportProcedure,
		svc.GetTroubleshootReport,
		opts...,
	)
	clusterServiceGetNetworkTestStatusHandler := connect.NewUnaryHandler(
		ClusterServiceGetNetworkTestStatusProcedure,
		svc.GetNetworkTestStatus,
		opts...,
	)
	clusterServiceStartNetworkTestHandler := connect.NewUnaryHandler(
		ClusterServiceStartNetworkTestProcedure,
		svc.StartNetworkTest,
		opts...,
	)
	clusterServiceCancelNetworkTestHandler := connect.NewUnaryHandler(
		ClusterServiceCancelNetworkTestProcedure,
		svc.CancelNetworkTest,
		opts...,
	)
	clusterServiceGetNetworkTestResultsHandler := connect.NewUnaryHandler(
		ClusterServiceGetNetworkTestResultsProcedure,
		svc.GetNetworkTestResults,
		opts...,
	)
	clusterServiceSetScrubbingScheduleHandler := connect.NewUnaryHandler(
		ClusterServiceSetScrubbingScheduleProcedure,
		svc.SetScrubbingSchedule,
		opts...,
	)
	clusterServiceGetResourcesHandler := connect.NewUnaryHandler(
		ClusterServiceGetResourcesProcedure,
		svc.GetResources,
		opts...,
	)
	clusterServiceSaveResourcesHandler := connect.NewUnaryHandler(
		ClusterServiceSaveResourcesProcedure,
		svc.SaveResources,
		opts...,
	)
	return "/api.services.cluster.v1.ClusterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ClusterServiceGetKoorClusterProcedure:
			clusterServiceGetKoorClusterHandler.ServeHTTP(w, r)
		case ClusterServiceGetTroubleshootReportProcedure:
			clusterServiceGetTroubleshootReportHandler.ServeHTTP(w, r)
		case ClusterServiceGetNetworkTestStatusProcedure:
			clusterServiceGetNetworkTestStatusHandler.ServeHTTP(w, r)
		case ClusterServiceStartNetworkTestProcedure:
			clusterServiceStartNetworkTestHandler.ServeHTTP(w, r)
		case ClusterServiceCancelNetworkTestProcedure:
			clusterServiceCancelNetworkTestHandler.ServeHTTP(w, r)
		case ClusterServiceGetNetworkTestResultsProcedure:
			clusterServiceGetNetworkTestResultsHandler.ServeHTTP(w, r)
		case ClusterServiceSetScrubbingScheduleProcedure:
			clusterServiceSetScrubbingScheduleHandler.ServeHTTP(w, r)
		case ClusterServiceGetResourcesProcedure:
			clusterServiceGetResourcesHandler.ServeHTTP(w, r)
		case ClusterServiceSaveResourcesProcedure:
			clusterServiceSaveResourcesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedClusterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedClusterServiceHandler struct{}

func (UnimplementedClusterServiceHandler) GetKoorCluster(context.Context, *connect.Request[v1.GetKoorClusterRequest]) (*connect.Response[v1.GetKoorClusterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.cluster.v1.ClusterService.GetKoorCluster is not implemented"))
}

func (UnimplementedClusterServiceHandler) GetTroubleshootReport(context.Context, *connect.Request[v1.GetTroubleshootReportRequest]) (*connect.Response[v1.GetTroubleshootReportResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.cluster.v1.ClusterService.GetTroubleshootReport is not implemented"))
}

func (UnimplementedClusterServiceHandler) GetNetworkTestStatus(context.Context, *connect.Request[v1.GetNetworkTestStatusRequest]) (*connect.Response[v1.GetNetworkTestStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.cluster.v1.ClusterService.GetNetworkTestStatus is not implemented"))
}

func (UnimplementedClusterServiceHandler) StartNetworkTest(context.Context, *connect.Request[v1.StartNetworkTestRequest]) (*connect.Response[v1.StartNetworkTestResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.cluster.v1.ClusterService.StartNetworkTest is not implemented"))
}

func (UnimplementedClusterServiceHandler) CancelNetworkTest(context.Context, *connect.Request[v1.CancelNetworkTestRequest]) (*connect.Response[v1.CancelNetworkTestResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.cluster.v1.ClusterService.CancelNetworkTest is not implemented"))
}

func (UnimplementedClusterServiceHandler) GetNetworkTestResults(context.Context, *connect.Request[v1.GetNetworkTestResultsRequest]) (*connect.Response[v1.GetNetworkTestResultsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.cluster.v1.ClusterService.GetNetworkTestResults is not implemented"))
}

func (UnimplementedClusterServiceHandler) SetScrubbingSchedule(context.Context, *connect.Request[v1.SetScrubbingScheduleRequest]) (*connect.Response[v1.SetScrubbingScheduleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.cluster.v1.ClusterService.SetScrubbingSchedule is not implemented"))
}

func (UnimplementedClusterServiceHandler) GetResources(context.Context, *connect.Request[v1.GetResourcesRequest]) (*connect.Response[v1.GetResourcesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.cluster.v1.ClusterService.GetResources is not implemented"))
}

func (UnimplementedClusterServiceHandler) SaveResources(context.Context, *connect.Request[v1.SaveResourcesRequest]) (*connect.Response[v1.SaveResourcesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.cluster.v1.ClusterService.SaveResources is not implemented"))
}
