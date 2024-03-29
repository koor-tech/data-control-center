package troubleshooting

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/koor-tech/data-control-center/gen/go/api/services/troubleshooting/v1"
	"github.com/koor-tech/data-control-center/pkg/ancientt"
	grpcerrors "github.com/koor-tech/data-control-center/pkg/grpc/errors"
)

func (s *Server) GetNetworkTestStatus(ctx context.Context, req *connect.Request[pb.GetNetworkTestStatusRequest]) (*connect.Response[pb.GetNetworkTestStatusResponse], error) {
	running := s.ancientt.IsRunning()

	var logs string
	if running {
		var err error
		logs, err = s.ancientt.GetLogs()
		if err != nil {
			logs += fmt.Sprintf("\nError retrieving ancientt command run logs: %s", err)
		}
	}

	completed, err := s.ancientt.IsComplete()
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&pb.GetNetworkTestStatusResponse{
		Running:  running && !completed,
		Complete: completed,
		Logs:     logs,
	}), nil
}

func (s *Server) StartNetworkTest(ctx context.Context, req *connect.Request[pb.StartNetworkTestRequest]) (*connect.Response[pb.StartNetworkTestResponse], error) {
	if s.readOnly {
		return nil, grpcerrors.ErrReadOnly
	}

	if s.ancientt.IsRunning() {
		return connect.NewResponse(&pb.StartNetworkTestResponse{
			AlreadyRunning: true,
		}), nil
	}

	outputFormat := ancientt.CSVOutputFormat
	switch req.Msg.OutputFormat {
	case pb.NetworkTestOutputFormat_NETWORK_TEST_OUTPUT_FORMAT_EXCELIZE:
		outputFormat = ancientt.ExcelizeOutputFormat
	default:
	}

	if err := s.ancientt.Start(&ancientt.RunParams{
		HostNetwork:  req.Msg.HostNetwork,
		OutputFormat: outputFormat,
	}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&pb.StartNetworkTestResponse{
		AlreadyRunning: false,
	}), nil
}

func (s *Server) CancelNetworkTest(ctx context.Context, req *connect.Request[pb.CancelNetworkTestRequest]) (*connect.Response[pb.CancelNetworkTestResponse], error) {
	if s.readOnly {
		return nil, grpcerrors.ErrReadOnly
	}

	if err := s.ancientt.Cancel(); err != nil {
		return nil, err
	}

	return connect.NewResponse(&pb.CancelNetworkTestResponse{
		NotRunning: false,
	}), nil
}

func (s *Server) GetNetworkTestResults(ctx context.Context, req *connect.Request[pb.GetNetworkTestResultsRequest]) (*connect.Response[pb.GetNetworkTestResultsResponse], error) {
	if s.readOnly {
		return nil, grpcerrors.ErrReadOnly
	}

	completed, err := s.ancientt.IsComplete()
	if err != nil {
		return nil, err
	}

	if !completed {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("Network Test is not complete yet!"))
	}

	name, fType, contents, err := s.ancientt.GetResultFile()
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&pb.GetNetworkTestResultsResponse{
		FileName:     name,
		FileType:     fType,
		FileContents: contents,
	}), nil
}
