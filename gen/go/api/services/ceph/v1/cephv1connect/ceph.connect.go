// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/services/ceph/v1/ceph.proto

package cephv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/koor-tech/data-control-center/gen/go/api/services/ceph/v1"
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
	// CephServiceName is the fully-qualified name of the CephService service.
	CephServiceName = "api.services.ceph.v1.CephService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CephServiceGetCephUsersProcedure is the fully-qualified name of the CephService's GetCephUsers
	// RPC.
	CephServiceGetCephUsersProcedure = "/api.services.ceph.v1.CephService/GetCephUsers"
	// CephServiceCreateCephUsersProcedure is the fully-qualified name of the CephService's
	// CreateCephUsers RPC.
	CephServiceCreateCephUsersProcedure = "/api.services.ceph.v1.CephService/CreateCephUsers"
)

// CephServiceClient is a client for the api.services.ceph.v1.CephService service.
type CephServiceClient interface {
	GetCephUsers(context.Context, *connect.Request[v1.GetCephUsersRequest]) (*connect.Response[v1.GetCephUsersResponse], error)
	CreateCephUsers(context.Context, *connect.Request[v1.CreateCephUsersRequest]) (*connect.Response[v1.CreateCephUsersResponse], error)
}

// NewCephServiceClient constructs a client for the api.services.ceph.v1.CephService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCephServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CephServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &cephServiceClient{
		getCephUsers: connect.NewClient[v1.GetCephUsersRequest, v1.GetCephUsersResponse](
			httpClient,
			baseURL+CephServiceGetCephUsersProcedure,
			opts...,
		),
		createCephUsers: connect.NewClient[v1.CreateCephUsersRequest, v1.CreateCephUsersResponse](
			httpClient,
			baseURL+CephServiceCreateCephUsersProcedure,
			opts...,
		),
	}
}

// cephServiceClient implements CephServiceClient.
type cephServiceClient struct {
	getCephUsers    *connect.Client[v1.GetCephUsersRequest, v1.GetCephUsersResponse]
	createCephUsers *connect.Client[v1.CreateCephUsersRequest, v1.CreateCephUsersResponse]
}

// GetCephUsers calls api.services.ceph.v1.CephService.GetCephUsers.
func (c *cephServiceClient) GetCephUsers(ctx context.Context, req *connect.Request[v1.GetCephUsersRequest]) (*connect.Response[v1.GetCephUsersResponse], error) {
	return c.getCephUsers.CallUnary(ctx, req)
}

// CreateCephUsers calls api.services.ceph.v1.CephService.CreateCephUsers.
func (c *cephServiceClient) CreateCephUsers(ctx context.Context, req *connect.Request[v1.CreateCephUsersRequest]) (*connect.Response[v1.CreateCephUsersResponse], error) {
	return c.createCephUsers.CallUnary(ctx, req)
}

// CephServiceHandler is an implementation of the api.services.ceph.v1.CephService service.
type CephServiceHandler interface {
	GetCephUsers(context.Context, *connect.Request[v1.GetCephUsersRequest]) (*connect.Response[v1.GetCephUsersResponse], error)
	CreateCephUsers(context.Context, *connect.Request[v1.CreateCephUsersRequest]) (*connect.Response[v1.CreateCephUsersResponse], error)
}

// NewCephServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCephServiceHandler(svc CephServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	cephServiceGetCephUsersHandler := connect.NewUnaryHandler(
		CephServiceGetCephUsersProcedure,
		svc.GetCephUsers,
		opts...,
	)
	cephServiceCreateCephUsersHandler := connect.NewUnaryHandler(
		CephServiceCreateCephUsersProcedure,
		svc.CreateCephUsers,
		opts...,
	)
	return "/api.services.ceph.v1.CephService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CephServiceGetCephUsersProcedure:
			cephServiceGetCephUsersHandler.ServeHTTP(w, r)
		case CephServiceCreateCephUsersProcedure:
			cephServiceCreateCephUsersHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCephServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCephServiceHandler struct{}

func (UnimplementedCephServiceHandler) GetCephUsers(context.Context, *connect.Request[v1.GetCephUsersRequest]) (*connect.Response[v1.GetCephUsersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.ceph.v1.CephService.GetCephUsers is not implemented"))
}

func (UnimplementedCephServiceHandler) CreateCephUsers(context.Context, *connect.Request[v1.CreateCephUsersRequest]) (*connect.Response[v1.CreateCephUsersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.ceph.v1.CephService.CreateCephUsers is not implemented"))
}
