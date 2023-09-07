// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/services/stats/stats.proto

package statsconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	common "github.com/koor-tech/data-control-center/gen/go/api/resources/common"
	stats "github.com/koor-tech/data-control-center/gen/go/api/resources/stats"
	stats1 "github.com/koor-tech/data-control-center/gen/go/api/services/stats"
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
	StatsServiceName = "stats.StatsService"
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
	StatsServiceGetClusterStatsProcedure = "/stats.StatsService/GetClusterStats"
	// StatsServiceGetClusterResourcesProcedure is the fully-qualified name of the StatsService's
	// GetClusterResources RPC.
	StatsServiceGetClusterResourcesProcedure = "/stats.StatsService/GetClusterResources"
	// StatsServiceGetClusterNodesProcedure is the fully-qualified name of the StatsService's
	// GetClusterNodes RPC.
	StatsServiceGetClusterNodesProcedure = "/stats.StatsService/GetClusterNodes"
)

// StatsServiceClient is a client for the stats.StatsService service.
type StatsServiceClient interface {
	GetClusterStats(context.Context, *connect.Request[common.EmptyRequest]) (*connect.Response[stats.ClusterStats], error)
	GetClusterResources(context.Context, *connect.Request[common.EmptyRequest]) (*connect.Response[stats1.ClusterResourcesResponse], error)
	GetClusterNodes(context.Context, *connect.Request[common.EmptyRequest]) (*connect.Response[stats1.ClusterNodesResponse], error)
}

// NewStatsServiceClient constructs a client for the stats.StatsService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStatsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) StatsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &statsServiceClient{
		getClusterStats: connect.NewClient[common.EmptyRequest, stats.ClusterStats](
			httpClient,
			baseURL+StatsServiceGetClusterStatsProcedure,
			opts...,
		),
		getClusterResources: connect.NewClient[common.EmptyRequest, stats1.ClusterResourcesResponse](
			httpClient,
			baseURL+StatsServiceGetClusterResourcesProcedure,
			opts...,
		),
		getClusterNodes: connect.NewClient[common.EmptyRequest, stats1.ClusterNodesResponse](
			httpClient,
			baseURL+StatsServiceGetClusterNodesProcedure,
			opts...,
		),
	}
}

// statsServiceClient implements StatsServiceClient.
type statsServiceClient struct {
	getClusterStats     *connect.Client[common.EmptyRequest, stats.ClusterStats]
	getClusterResources *connect.Client[common.EmptyRequest, stats1.ClusterResourcesResponse]
	getClusterNodes     *connect.Client[common.EmptyRequest, stats1.ClusterNodesResponse]
}

// GetClusterStats calls stats.StatsService.GetClusterStats.
func (c *statsServiceClient) GetClusterStats(ctx context.Context, req *connect.Request[common.EmptyRequest]) (*connect.Response[stats.ClusterStats], error) {
	return c.getClusterStats.CallUnary(ctx, req)
}

// GetClusterResources calls stats.StatsService.GetClusterResources.
func (c *statsServiceClient) GetClusterResources(ctx context.Context, req *connect.Request[common.EmptyRequest]) (*connect.Response[stats1.ClusterResourcesResponse], error) {
	return c.getClusterResources.CallUnary(ctx, req)
}

// GetClusterNodes calls stats.StatsService.GetClusterNodes.
func (c *statsServiceClient) GetClusterNodes(ctx context.Context, req *connect.Request[common.EmptyRequest]) (*connect.Response[stats1.ClusterNodesResponse], error) {
	return c.getClusterNodes.CallUnary(ctx, req)
}

// StatsServiceHandler is an implementation of the stats.StatsService service.
type StatsServiceHandler interface {
	GetClusterStats(context.Context, *connect.Request[common.EmptyRequest]) (*connect.Response[stats.ClusterStats], error)
	GetClusterResources(context.Context, *connect.Request[common.EmptyRequest]) (*connect.Response[stats1.ClusterResourcesResponse], error)
	GetClusterNodes(context.Context, *connect.Request[common.EmptyRequest]) (*connect.Response[stats1.ClusterNodesResponse], error)
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
	return "/stats.StatsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StatsServiceGetClusterStatsProcedure:
			statsServiceGetClusterStatsHandler.ServeHTTP(w, r)
		case StatsServiceGetClusterResourcesProcedure:
			statsServiceGetClusterResourcesHandler.ServeHTTP(w, r)
		case StatsServiceGetClusterNodesProcedure:
			statsServiceGetClusterNodesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStatsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStatsServiceHandler struct{}

func (UnimplementedStatsServiceHandler) GetClusterStats(context.Context, *connect.Request[common.EmptyRequest]) (*connect.Response[stats.ClusterStats], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("stats.StatsService.GetClusterStats is not implemented"))
}

func (UnimplementedStatsServiceHandler) GetClusterResources(context.Context, *connect.Request[common.EmptyRequest]) (*connect.Response[stats1.ClusterResourcesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("stats.StatsService.GetClusterResources is not implemented"))
}

func (UnimplementedStatsServiceHandler) GetClusterNodes(context.Context, *connect.Request[common.EmptyRequest]) (*connect.Response[stats1.ClusterNodesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("stats.StatsService.GetClusterNodes is not implemented"))
}
