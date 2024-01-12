package cluster

import (
	"context"

	"connectrpc.com/connect"
	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	clusterpb "github.com/koor-tech/data-control-center/gen/go/api/services/cluster/v1"
)

func (s *Server) GetScrubbingSchedule(ctx context.Context, req *connect.Request[clusterpb.GetScrubbingScheduleRequest]) (*connect.Response[clusterpb.GetScrubbingScheduleResponse], error) {
	schedule := &cephv1.OSDScrubbingSchedule{}

	// TODO

	return connect.NewResponse(&clusterpb.GetScrubbingScheduleResponse{
		OsdScrubbingSchedule: schedule,
	}), nil
}

func (s *Server) SetScrubbingSchedule(ctx context.Context, req *connect.Request[clusterpb.SetScrubbingScheduleRequest]) (*connect.Response[clusterpb.SetScrubbingScheduleResponse], error) {
	// TODO

	return connect.NewResponse(&clusterpb.SetScrubbingScheduleResponse{}), nil
}
