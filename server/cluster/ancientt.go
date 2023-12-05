package cluster

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	clusterpb "github.com/koor-tech/data-control-center/gen/go/api/services/cluster/v1"
	"github.com/koor-tech/data-control-center/pkg/ancientt"
	grpcerrors "github.com/koor-tech/data-control-center/pkg/grpc/errors"
)

func (s *Server) GetNetworkTestStatus(ctx context.Context, req *connect.Request[clusterpb.GetNetworkTestStatusRequest]) (*connect.Response[clusterpb.GetNetworkTestStatusResponse], error) {
	running := s.ancientt.IsRunning()

	var logs string
	if running {
		var err error
		logs, err = s.ancientt.GetLogs()
		if err != nil {
			logs += fmt.Sprintf("\nError retrieving ancientt command run logs: %s", err)
		}
	}

	return connect.NewResponse(&clusterpb.GetNetworkTestStatusResponse{
		Running: running,
		Logs:    logs,
	}), nil
}

func (s *Server) StartNetworkTest(ctx context.Context, req *connect.Request[clusterpb.StartNetworkTestRequest]) (*connect.Response[clusterpb.StartNetworkTestResponse], error) {
	if s.readOnly {
		return nil, grpcerrors.ErrReadOnly
	}

	if s.ancientt.IsRunning() {
		return connect.NewResponse(&clusterpb.StartNetworkTestResponse{
			AlreadyRunning: true,
		}), nil
	}

	outputFormat := ancientt.CSVOutputFormat
	switch req.Msg.OutputFormat {
	case clusterpb.NetworkTestOutputFormat_NETWORK_TEST_OUTPUT_FORMAT_EXCELIZE:
		outputFormat = ancientt.ExcelizeOutputFormat
	default:
	}

	if err := s.ancientt.Start(&ancientt.RunParams{
		HostNetwork:  req.Msg.HostNetwork,
		OutputFormat: outputFormat,
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&clusterpb.StartNetworkTestResponse{
		AlreadyRunning: false,
	}), nil
}

func (s *Server) CancelNetworkTest(ctx context.Context, req *connect.Request[clusterpb.CancelNetworkTestRequest]) (*connect.Response[clusterpb.CancelNetworkTestResponse], error) {
	if s.readOnly {
		return nil, grpcerrors.ErrReadOnly
	}

	if err := s.ancientt.Cancel(); err != nil {
		return nil, err
	}

	return connect.NewResponse(&clusterpb.CancelNetworkTestResponse{
		NotRunning: false,
	}), nil
}

func (s *Server) GetNetworkTestResults(ctx context.Context, req *connect.Request[clusterpb.GetNetworkTestResultsRequest]) (*connect.Response[clusterpb.GetNetworkTestResultsResponse], error) {
	if s.readOnly {
		return nil, grpcerrors.ErrReadOnly
	}

	name, fType, contents, err := s.ancientt.GetResultFile()
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&clusterpb.GetNetworkTestResultsResponse{
		FileName:     name,
		FileType:     fType,
		FileContents: contents,
	}), nil
}
