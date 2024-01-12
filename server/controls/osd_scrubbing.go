package controls

import (
	"context"

	"connectrpc.com/connect"
	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	pb "github.com/koor-tech/data-control-center/gen/go/api/services/controls/v1"
)

func (s *Server) GetScrubbingSchedule(ctx context.Context, req *connect.Request[pb.GetScrubbingScheduleRequest]) (*connect.Response[pb.GetScrubbingScheduleResponse], error) {
	schedule := &cephv1.OSDScrubbingSchedule{}

	// TODO

	return connect.NewResponse(&pb.GetScrubbingScheduleResponse{
		OsdScrubbingSchedule: schedule,
	}), nil
}

func (s *Server) SetScrubbingSchedule(ctx context.Context, req *connect.Request[pb.SetScrubbingScheduleRequest]) (*connect.Response[pb.SetScrubbingScheduleResponse], error) {
	// TODO

	return connect.NewResponse(&pb.SetScrubbingScheduleResponse{}), nil
}
