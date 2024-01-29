// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/services/controls/v1/osds.proto

package controlsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/koor-tech/data-control-center/gen/go/api/services/controls/v1"
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
	// OSDsServiceName is the fully-qualified name of the OSDsService service.
	OSDsServiceName = "api.services.controls.v1.OSDsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OSDsServiceGetScrubbingScheduleProcedure is the fully-qualified name of the OSDsService's
	// GetScrubbingSchedule RPC.
	OSDsServiceGetScrubbingScheduleProcedure = "/api.services.controls.v1.OSDsService/GetScrubbingSchedule"
	// OSDsServiceSetScrubbingScheduleProcedure is the fully-qualified name of the OSDsService's
	// SetScrubbingSchedule RPC.
	OSDsServiceSetScrubbingScheduleProcedure = "/api.services.controls.v1.OSDsService/SetScrubbingSchedule"
)

// OSDsServiceClient is a client for the api.services.controls.v1.OSDsService service.
type OSDsServiceClient interface {
	GetScrubbingSchedule(context.Context, *connect.Request[v1.GetScrubbingScheduleRequest]) (*connect.Response[v1.GetScrubbingScheduleResponse], error)
	SetScrubbingSchedule(context.Context, *connect.Request[v1.SetScrubbingScheduleRequest]) (*connect.Response[v1.SetScrubbingScheduleResponse], error)
}

// NewOSDsServiceClient constructs a client for the api.services.controls.v1.OSDsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOSDsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OSDsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &oSDsServiceClient{
		getScrubbingSchedule: connect.NewClient[v1.GetScrubbingScheduleRequest, v1.GetScrubbingScheduleResponse](
			httpClient,
			baseURL+OSDsServiceGetScrubbingScheduleProcedure,
			opts...,
		),
		setScrubbingSchedule: connect.NewClient[v1.SetScrubbingScheduleRequest, v1.SetScrubbingScheduleResponse](
			httpClient,
			baseURL+OSDsServiceSetScrubbingScheduleProcedure,
			opts...,
		),
	}
}

// oSDsServiceClient implements OSDsServiceClient.
type oSDsServiceClient struct {
	getScrubbingSchedule *connect.Client[v1.GetScrubbingScheduleRequest, v1.GetScrubbingScheduleResponse]
	setScrubbingSchedule *connect.Client[v1.SetScrubbingScheduleRequest, v1.SetScrubbingScheduleResponse]
}

// GetScrubbingSchedule calls api.services.controls.v1.OSDsService.GetScrubbingSchedule.
func (c *oSDsServiceClient) GetScrubbingSchedule(ctx context.Context, req *connect.Request[v1.GetScrubbingScheduleRequest]) (*connect.Response[v1.GetScrubbingScheduleResponse], error) {
	return c.getScrubbingSchedule.CallUnary(ctx, req)
}

// SetScrubbingSchedule calls api.services.controls.v1.OSDsService.SetScrubbingSchedule.
func (c *oSDsServiceClient) SetScrubbingSchedule(ctx context.Context, req *connect.Request[v1.SetScrubbingScheduleRequest]) (*connect.Response[v1.SetScrubbingScheduleResponse], error) {
	return c.setScrubbingSchedule.CallUnary(ctx, req)
}

// OSDsServiceHandler is an implementation of the api.services.controls.v1.OSDsService service.
type OSDsServiceHandler interface {
	GetScrubbingSchedule(context.Context, *connect.Request[v1.GetScrubbingScheduleRequest]) (*connect.Response[v1.GetScrubbingScheduleResponse], error)
	SetScrubbingSchedule(context.Context, *connect.Request[v1.SetScrubbingScheduleRequest]) (*connect.Response[v1.SetScrubbingScheduleResponse], error)
}

// NewOSDsServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOSDsServiceHandler(svc OSDsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	oSDsServiceGetScrubbingScheduleHandler := connect.NewUnaryHandler(
		OSDsServiceGetScrubbingScheduleProcedure,
		svc.GetScrubbingSchedule,
		opts...,
	)
	oSDsServiceSetScrubbingScheduleHandler := connect.NewUnaryHandler(
		OSDsServiceSetScrubbingScheduleProcedure,
		svc.SetScrubbingSchedule,
		opts...,
	)
	return "/api.services.controls.v1.OSDsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OSDsServiceGetScrubbingScheduleProcedure:
			oSDsServiceGetScrubbingScheduleHandler.ServeHTTP(w, r)
		case OSDsServiceSetScrubbingScheduleProcedure:
			oSDsServiceSetScrubbingScheduleHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOSDsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOSDsServiceHandler struct{}

func (UnimplementedOSDsServiceHandler) GetScrubbingSchedule(context.Context, *connect.Request[v1.GetScrubbingScheduleRequest]) (*connect.Response[v1.GetScrubbingScheduleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.controls.v1.OSDsService.GetScrubbingSchedule is not implemented"))
}

func (UnimplementedOSDsServiceHandler) SetScrubbingSchedule(context.Context, *connect.Request[v1.SetScrubbingScheduleRequest]) (*connect.Response[v1.SetScrubbingScheduleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.services.controls.v1.OSDsService.SetScrubbingSchedule is not implemented"))
}