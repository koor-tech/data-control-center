package cluster

import (
	"context"

	"connectrpc.com/connect"
	clusterpb "github.com/koor-tech/data-control-center/gen/go/api/services/cluster/v1"
	"github.com/koor-tech/data-control-center/pkg/ancientt"
)

func (s *Server) StartNetworkTest(context.Context, *connect.Request[clusterpb.StartNetworkTestRequest]) (*connect.Response[clusterpb.StartNetworkTestResponse], error) {
	if s.ancientt.IsRunning() {
		return connect.NewResponse(&clusterpb.StartNetworkTestResponse{
			AlreadyRunning: true,
		}), nil
	}

	if err := s.ancientt.Start(ancientt.CSVOutputFormat); err != nil {
		return nil, err
	}

	return connect.NewResponse(&clusterpb.StartNetworkTestResponse{
		AlreadyRunning: false,
	}), nil
}

func (s *Server) CancelNetworkTest(context.Context, *connect.Request[clusterpb.CancelNetworkTestRequest]) (*connect.Response[clusterpb.CancelNetworkTestResponse], error) {
	if err := s.ancientt.Cancel(); err != nil {
		return nil, err
	}

	return connect.NewResponse(&clusterpb.CancelNetworkTestResponse{
		NotRunning: false,
	}), nil
}

func (s *Server) GetNetworkTestResults(context.Context, *connect.Request[clusterpb.GetNetworkTestResultsRequest]) (*connect.Response[clusterpb.GetNetworkTestResultsResponse], error) {

	// TODO s.ancientt.GetResultFile()

	return nil, nil
}
