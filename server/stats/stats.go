package stats

import (
	"context"
	"fmt"
	"github.com/koor-tech/data-control-center/gen/go/api/services/response"
	pb "github.com/koor-tech/data-control-center/gen/go/api/services/response/responseconnect"
	"github.com/koor-tech/data-control-center/gen/go/api/services/stats"
	"go.uber.org/zap"
	"log"
	"strings"

	"github.com/koor-tech/data-control-center/internal/ceph"
	serverResponse "github.com/koor-tech/data-control-center/server/response"

	"github.com/koor-tech/data-control-center/pkg/config"
	"net/http"

	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
)

// Server is used to implement stats services.
type Server struct {
	auth           *auth.GRPCAuth
	cephApiService *ceph.Service
	logger         *zap.Logger
}

func New(logger *zap.Logger, grpcAuth *auth.GRPCAuth, cfg *config.Config) (*Server, error) {

	return &Server{
		logger:         logger,
		auth:           grpcAuth,
		cephApiService: ceph.NewService(logger, &cfg.Ceph),
	}, nil
}

func (s *Server) RegisterService(g *gin.RouterGroup) {
	path, handler := pb.NewStatsServiceHandler(s, connect.WithInterceptors(
		s.auth.NewAuthInterceptor(),
	))
	g.Any(path+"/*path", gin.WrapH(handler))
}

func (s *Server) GetClusterStats(ctx context.Context, req *connect.Request[response.EmptyRequest]) (*connect.Response[response.Response], error) {

	//daemonCrashes := []*pb.DaemonCrash{
	//	{
	//		Description: "1 daemons have recently crashed",
	//	},
	//	{
	//		Description: "1 mgr modules have recently crashed",
	//	},
	//}
	//
	//svc := pb.Services{
	//	Mon: &pb.MonService{
	//		DaemonCount: 3,
	//		Quorum:      "a,b,c",
	//		Age:         "13d",
	//	},
	//	Mgr: &pb.MgrService{
	//		Active:   "a",
	//		Standbys: []string{"b"},
	//		Since:    "6h",
	//	},
	//	Mds: &pb.MdsService{
	//		DaemonsUp:       1, // needs change 1/1 daemons up, 1 hot standby
	//		HotStandbyCount: 1,
	//	},
	//	Osd: &pb.OsdService{
	//		OsdCount: 5,
	//		OsdUp:    4,    //(since 3d),
	//		OsdIn:    4,    // 4 in (since 3d) ?
	//		Since:    "3d", //todo remove this and add in -up and -in
	//	},
	//	Rgw: &pb.RgwService{
	//		ActiveDaemon: 1,
	//		HostCount:    1,
	//		ZoneCount:    1,
	//	},
	//}
	//
	//data := pb.Data{
	//	Volumes: &pb.VolumeStatus{
	//		Description: " 1/1 healthy",
	//	},
	//	Pools: &pb.PoolStatus{
	//		PoolCount: 13,
	//		Pgs:       217,
	//	},
	//	Objects: &pb.ObjectStatus{
	//		ObjectCount: "15.67k", //bytes
	//		Size:        "239 MB", //bytes
	//	},
	//	Usage: &pb.UsageStatus{
	//		Used:      3.3,   //used
	//		Available: 176.3, // avail
	//		Total:     120,   //total  in bytes
	//	},
	//	Pgs: &pb.PGs{
	//		ActiveClean: 217, //active+clean
	//	},
	//}
	//
	//io := pb.IoStatus{
	//	ClientRead:     "852 B/s", // 852 B/s rd
	//	ClientReadOps:  "1 op/s",  // 1 op/s rd
	//	ClientWriteOps: "0 op/s",  // 0 op/s wr
	//}

	//Health := pb.HealthStatus_HEALTH_OK
	//if hs.Health.Status == "HEALTH_WARN" {
	//	Health = pb.HealthStatus_HEALTH_WARN
	//}

	//
	//return &connect.Response[pb.ClusterStatusResponse]{
	//	Msg: &pb.ClusterStatusResponse{
	//		Id: "123",
	//	},
	//}, nil
	// *connect.Response[response.Response]
	//return &connect.Response[connect.Response{}], nil

	st, err := s.cephApiService.GetClusterHealth(context.TODO())
	if err != nil {
		return serverResponse.NewFailureResponse(http.StatusInternalServerError, fmt.Errorf("error caused by %w", err)), nil
	}

	daemonCount := len(st.MonStatus.Monmap.Mons)
	var monNames []string
	for _, mon := range st.MonStatus.Monmap.Mons {
		monNames = append(monNames, mon.Name)
	}

	var standBys []string
	for _, standBy := range st.MgrMap.Standbys {
		standBys = append(standBys, standBy.Name)
	}

	var (
		daemonsUp          int
		standbyCountWanted int
	)

	// TODO check the input to test how works in different scenarios
	for _, filesystem := range st.FsMap.Filesystems {
		itemsUp, err := convertMdsItems(filesystem.MdsMap.Up)
		if err != nil {
			log.Fatal(err)
		}
		daemonsUp = len(itemsUp)
		standbyCountWanted = filesystem.MdsMap.StandbyCountWanted
	}

	var (
		osdCount int
		osdUp    int
		osdIn    int
	)

	for _, osd := range st.OsdMap.Osds {
		osdUp += osd.Up
		osdIn += osd.In
		osdCount++
	}
	/*
		//	Osd: &pb.OsdService{
		//		OsdCount: 5,
		//		OsdUp:    4,    //(since 3d),
		//		OsdIn:    4,    // 4 in (since 3d) ?
		//		Since:    "3d", //todo remove this and add in -up and -in
		//	},
		//	Rgw: &pb.RgwService{
		//		ActiveDaemon: 1,
		//		HostCount:    1,
		//		ZoneCount:    1,
		//	},
	*/

	resp := response.Response_ClusterStatusResponse{
		ClusterStatusResponse: &stats.ClusterStatusResponse{
			Id:     st.MonStatus.Monmap.Fsid,
			Status: st.Health.Status,
			Services: &stats.Services{
				Mon: &stats.MonService{
					DaemonCount: int32(daemonCount),
					Quorum:      strings.Join(monNames, ","),
					Age:         "13 d", // TODO how to calculate it
				},
				Mgr: &stats.MgrService{
					Active:   st.MgrMap.ActiveName,
					Standbys: standBys,
					Since:    "6h", // TODO how to calculate it
				},
				Mds: &stats.MdsService{
					DaemonsUp:       int32(daemonsUp),
					HotStandbyCount: int32(standbyCountWanted),
				},
				Osd: &stats.OsdService{
					OsdCount: int32(osdCount),
					OsdUp:    int32(osdUp),
					OsdIn:    int32(osdIn),
					Since:    "3d", //TODO how to calculate it
				},
				Rgw: &stats.RgwService{
					ActiveDaemon: 1,
					HostCount:    1,
					ZoneCount:    1,
				},
			},
		},
	}

	return serverResponse.NewSuccessResponse(http.StatusOK, &resp), nil
}

type MdsItems []map[string]interface{}

func convertMdsItems(src interface{}) (MdsItems, error) {
	var items MdsItems

	switch src := src.(type) {
	case []interface{}:
		for _, item := range src {
			if m, ok := item.(map[string]interface{}); ok {
				items = append(items, m)
			}
		}
	case map[string]interface{}:
		items = append(items, src)
	default:
		return nil, fmt.Errorf("Unsupported type for MdsItems: %T", src)
	}

	return items, nil
}
