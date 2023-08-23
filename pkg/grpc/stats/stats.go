package stats

import (
	"context"
	pb "github.com/koor-tech/data-control-center/gen/go/api/services/stats"
)

// Server is used to implement stats services.
type Server struct {
	pb.UnimplementedStatsServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetClusterStats(ctx context.Context, _ *pb.EmptyRequest) (*pb.ClusterStatusResponse, error) {

	daemonCrashes := []*pb.DaemonCrash{
		{
			Description: "1 daemons have recently crashed",
		},
		{
			Description: "1 mgr modules have recently crashed",
		},
	}

	svc := pb.Services{
		Mon: &pb.MonService{
			DaemonCount: 3,
			Quorum:      "a,b,c",
			Age:         "13d",
		},
		Mgr: &pb.MgrService{
			Active:   "a",
			Standbys: []string{"b"},
			Since:    "6h",
		},
		Mds: &pb.MdsService{
			DaemonsUp:       1, // needs change 1/1 daemons up, 1 hot standby
			HotStandbyCount: 1,
		},
		Osd: &pb.OsdService{
			OsdCount: 5,
			OsdUp:    4,    //(since 3d),
			OsdIn:    4,    // 4 in (since 3d) ?
			Since:    "3d", //todo remove this and add in -up and -in
		},
		Rgw: &pb.RgwService{
			ActiveDaemon: 1,
			HostCount:    1,
			ZoneCount:    1,
		},
	}

	data := pb.Data{
		Volumes: &pb.VolumeStatus{
			Description: " 1/1 healthy",
		},
		Pools: &pb.PoolStatus{
			PoolCount: 13,
			Pgs:       217,
		},
		Objects: &pb.ObjectStatus{
			ObjectCount: "15.67k", //bytes
			Size:        "239 MB", //bytes
		},
		Usage: &pb.UsageStatus{
			Used:      "3.3 GiB", //used
			Available: "117 GiB", // avail
			Total:     "120 GiB", //total  in bytes
		},
		Pgs: &pb.PGs{
			ActiveClean: 217, //active+clean
		},
	}

	io := pb.IoStatus{
		ClientRead:     "852 B/s", // 852 B/s rd
		ClientReadOps:  "1 op/s",  // 1 op/s rd
		ClientWriteOps: "0 op/s",  // 0 op/s wr
	}

	return &pb.ClusterStatusResponse{
		Id:            "60d54b9e-87c1-41ec-bc14-9e9925338e59",
		Health:        pb.HealthStatus_HEALTH_WARN,
		DaemonCrashes: daemonCrashes,
		Services:      &svc,
		Data:          &data,
		Io:            &io,
	}, nil
}
