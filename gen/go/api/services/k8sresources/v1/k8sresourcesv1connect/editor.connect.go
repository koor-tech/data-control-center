// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/services/k8sresources/v1/editor.proto

package k8sresourcesv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/koor-tech/data-control-center/gen/go/api/services/k8sresources/v1"
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
	// K8sResourcesServiceName is the fully-qualified name of the K8sResourcesService service.
	K8sResourcesServiceName = "api.services.k8sresources.v1.K8sResourcesService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// K8SResourcesServiceGetResourcesProcedure is the fully-qualified name of the K8sResourcesService's
	// GetResources RPC.
	K8SResourcesServiceGetResourcesProcedure = "/api.services.k8sresources.v1.K8sResourcesService/GetResources"
	// K8SResourcesServiceSaveResourceProcedure is the fully-qualified name of the K8sResourcesService's
	// SaveResource RPC.
	K8SResourcesServiceSaveResourceProcedure = "/api.services.k8sresources.v1.K8sResourcesService/SaveResource"
)

// K8SResourcesServiceClient is a client for the api.services.k8sresources.v1.K8sResourcesService
// service.
type K8SResourcesServiceClient interface {
	GetResources(context.Context, *connect.Request[v1.GetResourcesRequest]) (*connect.Response[v1.GetResourcesResponse], error)
	SaveResource(context.Context, *connect.Request[v1.SaveResourceRequest]) (*connect.Response[v1.SaveResourceResponse], error)
}

// NewK8SResourcesServiceClient constructs a client for the
// api.services.k8sresources.v1.K8sResourcesService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewK8SResourcesServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) K8SResourcesServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &k8SResourcesServiceClient{
		getResources: connect.NewClient[v1.GetResourcesRequest, v1.GetResourcesResponse](
			httpClient,
			baseURL+K8SResourcesServiceGetResourcesProcedure,
			opts...,
		),
		saveResource: connect.NewClient[v1.SaveResourceRequest, v1.SaveResourceResponse](
			httpClient,
			baseURL+K8SResourcesServiceSaveResourceProcedure,
			opts...,
		),
	}
}

// k8SResourcesServiceClient implements K8SResourcesServiceClient.
type k8SResourcesServiceClient struct {
	getResources *connect.Client[v1.GetResourcesRequest, v1.GetResourcesResponse]
	saveResource *connect.Client[v1.SaveResourceRequest, v1.SaveResourceResponse]
}

// GetResources calls api.services.k8sresources.v1.K8sResourcesService.GetResources.
func (c *k8SResourcesServiceClient) GetResources(ctx context.Context, req *connect.Request[v1.GetResourcesRequest]) (*connect.Response[v1.GetResourcesResponse], error) {
	return c.getResources.CallUnary(ctx, req)
}

// SaveResource calls api.services.k8sresources.v1.K8sResourcesService.SaveResource.
func (c *k8SResourcesServiceClient) SaveResource(ctx context.Context, req *connect.Request[v1.SaveResourceRequest]) (*connect.Response[v1.SaveResourceResponse], error) {
	return c.saveResource.CallUnary(ctx, req)
}

// K8SResourcesServiceHandler is an implementation of the
// api.services.k8sresources.v1.K8sResourcesService service.
type K8SResourcesServiceHandler interface {
	GetResources(context.Context, *connect.Request[v1.GetResourcesRequest]) (*connect.Response[v1.GetResourcesResponse], error)
	SaveResource(context.Context, *connect.Request[v1.SaveResourceRequest]) (*connect.Response[v1.SaveResourceResponse], error)
}

// NewK8SResourcesServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewK8SResourcesServiceHandler(svc K8SResourcesServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	k8SResourcesServiceGetResourcesHandler := connect.NewUnaryHandler(
		K8SResourcesServiceGetResourcesProcedure,
		svc.GetResources,
		opts...,
	)
	k8SResourcesServiceSaveResourceHandler := connect.NewUnaryHandler(
		K8SResourcesServiceSaveResourceProcedure,
		svc.SaveResource,
		opts...,
	)
	return "/api.services.k8sresources.v1.K8sResourcesService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case K8SResourcesServiceGetResourcesProcedure:
			k8SResourcesServiceGetResourcesHandler.ServeHTTP(w, r)
		case K8SResourcesServiceSaveResourceProcedure:
			k8SResourcesServiceSaveResourceHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedK8SResourcesServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedK8SResourcesServiceHandler struct{}

func (UnimplementedK8SResourcesServiceHandler) GetResources(context.Context, *connect.Request[v1.GetResourcesRequest]) (*connect.Response[v1.GetResourcesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.k8sresources.v1.K8sResourcesService.GetResources is not implemented"))
}

func (UnimplementedK8SResourcesServiceHandler) SaveResource(context.Context, *connect.Request[v1.SaveResourceRequest]) (*connect.Response[v1.SaveResourceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.k8sresources.v1.K8sResourcesService.SaveResource is not implemented"))
}
