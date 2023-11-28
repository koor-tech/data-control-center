package cluster

import (
	"context"

	"connectrpc.com/connect"
	clusterpb "github.com/koor-tech/data-control-center/gen/go/api/services/cluster/v1"
)

// TODO implement the ancientt run methods here

func (s *Server) StartNetworkTest(context.Context, *connect.Request[clusterpb.StartNetworkTestRequest]) (*connect.Response[clusterpb.StartNetworkTestResponse], error) {

	// TODO

	return nil, nil
}

func (s *Server) CancelNetworkTest(context.Context, *connect.Request[clusterpb.CancelNetworkTestRequest]) (*connect.Response[clusterpb.CancelNetworkTestResponse], error) {

	// TODO

	return nil, nil
}

func (s *Server) GetNetworkTestResults(context.Context, *connect.Request[clusterpb.GetNetworkTestResultsRequest]) (*connect.Response[clusterpb.GetNetworkTestResultsResponse], error) {

	// TODO

	return nil, nil
}
