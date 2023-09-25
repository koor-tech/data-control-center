// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/services/stats/v1/stats.proto

package statsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/koor-tech/data-control-center/gen/go/api/services/stats/v1"
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
	// StatsServiceName is the fully-qualified name of the StatsService service.
	StatsServiceName = "api.services.stats.v1.StatsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StatsServiceGetClusterStatsProcedure is the fully-qualified name of the StatsService's
	// GetClusterStats RPC.
	StatsServiceGetClusterStatsProcedure = "/api.services.stats.v1.StatsService/GetClusterStats"
	// StatsServiceGetClusterResourcesProcedure is the fully-qualified name of the StatsService's
	// GetClusterResources RPC.
	StatsServiceGetClusterResourcesProcedure = "/api.services.stats.v1.StatsService/GetClusterResources"
	// StatsServiceGetClusterNodesProcedure is the fully-qualified name of the StatsService's
	// GetClusterNodes RPC.
	StatsServiceGetClusterNodesProcedure = "/api.services.stats.v1.StatsService/GetClusterNodes"
	// StatsServiceGetClusterRadarProcedure is the fully-qualified name of the StatsService's
	// GetClusterRadar RPC.
	StatsServiceGetClusterRadarProcedure = "/api.services.stats.v1.StatsService/GetClusterRadar"
	// StatsServiceGetKoorClusterProcedure is the fully-qualified name of the StatsService's
	// GetKoorCluster RPC.
	StatsServiceGetKoorClusterProcedure = "/api.services.stats.v1.StatsService/GetKoorCluster"
)

// StatsServiceClient is a client for the api.services.stats.v1.StatsService service.
type StatsServiceClient interface {
	GetClusterStats(context.Context, *connect.Request[v1.GetClusterStatsRequest]) (*connect.Response[v1.GetClusterStatsResponse], error)
	GetClusterResources(context.Context, *connect.Request[v1.GetClusterResourcesRequest]) (*connect.Response[v1.GetClusterResourcesResponse], error)
	GetClusterNodes(context.Context, *connect.Request[v1.GetClusterNodesRequest]) (*connect.Response[v1.GetClusterNodesResponse], error)
	GetClusterRadar(context.Context, *connect.Request[v1.GetClusterRadarRequest]) (*connect.Response[v1.GetClusterRadarResponse], error)
	GetKoorCluster(context.Context, *connect.Request[v1.GetKoorClusterRequest]) (*connect.Response[v1.GetKoorClusterResponse], error)
}

// NewStatsServiceClient constructs a client for the api.services.stats.v1.StatsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStatsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) StatsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &statsServiceClient{
		getClusterStats: connect.NewClient[v1.GetClusterStatsRequest, v1.GetClusterStatsResponse](
			httpClient,
			baseURL+StatsServiceGetClusterStatsProcedure,
			opts...,
		),
		getClusterResources: connect.NewClient[v1.GetClusterResourcesRequest, v1.GetClusterResourcesResponse](
			httpClient,
			baseURL+StatsServiceGetClusterResourcesProcedure,
			opts...,
		),
		getClusterNodes: connect.NewClient[v1.GetClusterNodesRequest, v1.GetClusterNodesResponse](
			httpClient,
			baseURL+StatsServiceGetClusterNodesProcedure,
			opts...,
		),
		getClusterRadar: connect.NewClient[v1.GetClusterRadarRequest, v1.GetClusterRadarResponse](
			httpClient,
			baseURL+StatsServiceGetClusterRadarProcedure,
			opts...,
		),
		getKoorCluster: connect.NewClient[v1.GetKoorClusterRequest, v1.GetKoorClusterResponse](
			httpClient,
			baseURL+StatsServiceGetKoorClusterProcedure,
			opts...,
		),
	}
}

// statsServiceClient implements StatsServiceClient.
type statsServiceClient struct {
	getClusterStats     *connect.Client[v1.GetClusterStatsRequest, v1.GetClusterStatsResponse]
	getClusterResources *connect.Client[v1.GetClusterResourcesRequest, v1.GetClusterResourcesResponse]
	getClusterNodes     *connect.Client[v1.GetClusterNodesRequest, v1.GetClusterNodesResponse]
	getClusterRadar     *connect.Client[v1.GetClusterRadarRequest, v1.GetClusterRadarResponse]
	getKoorCluster      *connect.Client[v1.GetKoorClusterRequest, v1.GetKoorClusterResponse]
}

// GetClusterStats calls api.services.stats.v1.StatsService.GetClusterStats.
func (c *statsServiceClient) GetClusterStats(ctx context.Context, req *connect.Request[v1.GetClusterStatsRequest]) (*connect.Response[v1.GetClusterStatsResponse], error) {
	return c.getClusterStats.CallUnary(ctx, req)
}

// GetClusterResources calls api.services.stats.v1.StatsService.GetClusterResources.
func (c *statsServiceClient) GetClusterResources(ctx context.Context, req *connect.Request[v1.GetClusterResourcesRequest]) (*connect.Response[v1.GetClusterResourcesResponse], error) {
	return c.getClusterResources.CallUnary(ctx, req)
}

// GetClusterNodes calls api.services.stats.v1.StatsService.GetClusterNodes.
func (c *statsServiceClient) GetClusterNodes(ctx context.Context, req *connect.Request[v1.GetClusterNodesRequest]) (*connect.Response[v1.GetClusterNodesResponse], error) {
	return c.getClusterNodes.CallUnary(ctx, req)
}

// GetClusterRadar calls api.services.stats.v1.StatsService.GetClusterRadar.
func (c *statsServiceClient) GetClusterRadar(ctx context.Context, req *connect.Request[v1.GetClusterRadarRequest]) (*connect.Response[v1.GetClusterRadarResponse], error) {
	return c.getClusterRadar.CallUnary(ctx, req)
}

// GetKoorCluster calls api.services.stats.v1.StatsService.GetKoorCluster.
func (c *statsServiceClient) GetKoorCluster(ctx context.Context, req *connect.Request[v1.GetKoorClusterRequest]) (*connect.Response[v1.GetKoorClusterResponse], error) {
	return c.getKoorCluster.CallUnary(ctx, req)
}

// StatsServiceHandler is an implementation of the api.services.stats.v1.StatsService service.
type StatsServiceHandler interface {
	GetClusterStats(context.Context, *connect.Request[v1.GetClusterStatsRequest]) (*connect.Response[v1.GetClusterStatsResponse], error)
	GetClusterResources(context.Context, *connect.Request[v1.GetClusterResourcesRequest]) (*connect.Response[v1.GetClusterResourcesResponse], error)
	GetClusterNodes(context.Context, *connect.Request[v1.GetClusterNodesRequest]) (*connect.Response[v1.GetClusterNodesResponse], error)
	GetClusterRadar(context.Context, *connect.Request[v1.GetClusterRadarRequest]) (*connect.Response[v1.GetClusterRadarResponse], error)
	GetKoorCluster(context.Context, *connect.Request[v1.GetKoorClusterRequest]) (*connect.Response[v1.GetKoorClusterResponse], error)
}

// NewStatsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStatsServiceHandler(svc StatsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	statsServiceGetClusterStatsHandler := connect.NewUnaryHandler(
		StatsServiceGetClusterStatsProcedure,
		svc.GetClusterStats,
		opts...,
	)
	statsServiceGetClusterResourcesHandler := connect.NewUnaryHandler(
		StatsServiceGetClusterResourcesProcedure,
		svc.GetClusterResources,
		opts...,
	)
	statsServiceGetClusterNodesHandler := connect.NewUnaryHandler(
		StatsServiceGetClusterNodesProcedure,
		svc.GetClusterNodes,
		opts...,
	)
	statsServiceGetClusterRadarHandler := connect.NewUnaryHandler(
		StatsServiceGetClusterRadarProcedure,
		svc.GetClusterRadar,
		opts...,
	)
	statsServiceGetKoorClusterHandler := connect.NewUnaryHandler(
		StatsServiceGetKoorClusterProcedure,
		svc.GetKoorCluster,
		opts...,
	)
	return "/api.services.stats.v1.StatsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StatsServiceGetClusterStatsProcedure:
			statsServiceGetClusterStatsHandler.ServeHTTP(w, r)
		case StatsServiceGetClusterResourcesProcedure:
			statsServiceGetClusterResourcesHandler.ServeHTTP(w, r)
		case StatsServiceGetClusterNodesProcedure:
			statsServiceGetClusterNodesHandler.ServeHTTP(w, r)
		case StatsServiceGetClusterRadarProcedure:
			statsServiceGetClusterRadarHandler.ServeHTTP(w, r)
		case StatsServiceGetKoorClusterProcedure:
			statsServiceGetKoorClusterHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStatsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStatsServiceHandler struct{}

func (UnimplementedStatsServiceHandler) GetClusterStats(context.Context, *connect.Request[v1.GetClusterStatsRequest]) (*connect.Response[v1.GetClusterStatsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.stats.v1.StatsService.GetClusterStats is not implemented"))
}

func (UnimplementedStatsServiceHandler) GetClusterResources(context.Context, *connect.Request[v1.GetClusterResourcesRequest]) (*connect.Response[v1.GetClusterResourcesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.stats.v1.StatsService.GetClusterResources is not implemented"))
}

func (UnimplementedStatsServiceHandler) GetClusterNodes(context.Context, *connect.Request[v1.GetClusterNodesRequest]) (*connect.Response[v1.GetClusterNodesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.stats.v1.StatsService.GetClusterNodes is not implemented"))
}

func (UnimplementedStatsServiceHandler) GetClusterRadar(context.Context, *connect.Request[v1.GetClusterRadarRequest]) (*connect.Response[v1.GetClusterRadarResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.stats.v1.StatsService.GetClusterRadar is not implemented"))
}

func (UnimplementedStatsServiceHandler) GetKoorCluster(context.Context, *connect.Request[v1.GetKoorClusterRequest]) (*connect.Response[v1.GetKoorClusterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.stats.v1.StatsService.GetKoorCluster is not implemented"))
}
