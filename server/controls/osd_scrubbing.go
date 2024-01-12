package controls

import (
	"context"
	"strconv"

	"connectrpc.com/connect"
	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	pb "github.com/koor-tech/data-control-center/gen/go/api/services/controls/v1"
	rookcephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) GetScrubbingSchedule(ctx context.Context, req *connect.Request[pb.GetScrubbingScheduleRequest]) (*connect.Response[pb.GetScrubbingScheduleResponse], error) {

	list, err := s.k.GetRookClient().CephV1().CephClusters(s.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var cluster rookcephv1.CephCluster
	for _, c := range list.Items {
		cluster = c
		break
	}

	cephConfig := cluster.Spec.CephConfig

	schedule := &cephv1.OSDScrubbingSchedule{}
	schedule.Default()

	globalCfg, ok := cephConfig["global"]
	if !ok {
		return connect.NewResponse(&pb.GetScrubbingScheduleResponse{
			OsdScrubbingSchedule: schedule,
		}), nil
	}

	// osd_max_scrubs
	if val, ok := globalCfg["osd_max_scrubs"]; ok {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		value := int64(parsed)
		schedule.MaxScrubOps = &value
	}

	// osd_scrub_begin_hour
	if val, ok := globalCfg["osd_scrub_begin_hour"]; ok {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		value := int64(parsed)
		schedule.BeginHour = &value
	}

	// osd_scrub_end_hour
	if val, ok := globalCfg["osd_scrub_end_hour"]; ok {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		value := int64(parsed)
		schedule.EndHour = &value
	}

	// osd_scrub_begin_week_day
	if val, ok := globalCfg["osd_scrub_begin_week_day"]; ok {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		value := int64(parsed)
		schedule.BeginWeekDay = &value
	}

	// osd_scrub_end_week_day
	if val, ok := globalCfg["osd_scrub_end_week_day"]; ok {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		value := int64(parsed)
		schedule.EndWeekDay = &value
	}

	// osd_scrub_end_week_day
	if val, ok := globalCfg["osd_scrub_end_week_day"]; ok {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		value := int64(parsed)
		schedule.EndWeekDay = &value
	}

	// osd_scrub_min_interval
	if val, ok := globalCfg["osd_scrub_min_interval"]; ok {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		value := int64(parsed)
		schedule.MinScrubInterval = &value
	}

	// osd_scrub_max_interval
	if val, ok := globalCfg["osd_scrub_max_interval"]; ok {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		value := int64(parsed)
		schedule.MaxScrubInterval = &value
	}

	// osd_deep_scrub_interval
	if val, ok := globalCfg["osd_deep_scrub_interval"]; ok {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		value := int64(parsed)
		schedule.DeepScrubInterval = &value
	}

	// osd_scrub_sleep
	if val, ok := globalCfg["osd_scrub_sleep"]; ok {
		parsed, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return nil, err
		}

		value := float32(parsed)
		schedule.ScrubSleepSeconds = &value
	}

	return connect.NewResponse(&pb.GetScrubbingScheduleResponse{
		OsdScrubbingSchedule: schedule,
	}), nil
}

func (s *Server) SetScrubbingSchedule(ctx context.Context, req *connect.Request[pb.SetScrubbingScheduleRequest]) (*connect.Response[pb.SetScrubbingScheduleResponse], error) {
	schedule := req.Msg.OsdScrubbingSchedule
	schedule.Default()

	// TODO

	return connect.NewResponse(&pb.SetScrubbingScheduleResponse{}), nil
}
