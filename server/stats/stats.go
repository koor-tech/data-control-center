package stats

import (
	"context"
	"fmt"
	"github.com/koor-tech/data-control-center/gen/go/api/services/response"
	pb "github.com/koor-tech/data-control-center/gen/go/api/services/response/responseconnect"
	"github.com/koor-tech/data-control-center/gen/go/api/services/stats"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	st, err := s.cephApiService.GetClusterHealth(context.TODO())
	if err != nil {
		return serverResponse.NewFailureResponse(http.StatusInternalServerError, fmt.Errorf("error caused by %w", err)), nil
	}

	var crashes []*stats.Crash

	for _, check := range st.Health.Checks {
		crashes = append(crashes, &stats.Crash{Description: check.Summary.Message})
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

	for _, filesystem := range st.FsMap.Filesystems {
		itemsUp, err := convertMdsItems(filesystem.MdsMap.Up)
		if err != nil {
			return serverResponse.NewFailureResponse(http.StatusInternalServerError, fmt.Errorf("error caused by %w", err)), nil

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

	poolCount := len(st.Pools)
	activeAndCleanPGs := 0
	for _, pool := range st.Pools {
		activeAndCleanPGs += pool.PgStatus.ActiveClean
	}

	stored := 0
	for _, pool := range st.DF.Pools {
		stored += pool.DFPoolStats.Stored
	}

	clientReadBytes := formatBytes(int64(st.ClientPerf.ReadBytesSec))
	readOps := st.ClientPerf.ReadOpPerSec
	writeOps := st.ClientPerf.WriteOpPerSec

	resp := response.Response_ClusterStatusResponse{
		ClusterStatusResponse: &stats.ClusterStatusResponse{
			Id:      st.MonStatus.Monmap.Fsid,
			Status:  st.Health.Status,
			Crashes: crashes,
			Services: &stats.Services{
				Mon: &stats.MonService{
					DaemonCount: int32(daemonCount),
					Quorum:      strings.Join(monNames, ","),
					CreatedAt:   timestamppb.New(st.MonStatus.Monmap.Created),
					UpdatedAt:   timestamppb.New(st.MonStatus.Monmap.Modified),
				},
				Mgr: &stats.MgrService{
					Active:    st.MgrMap.ActiveName,
					Standbys:  standBys,
					UpdatedAt: timestamppb.New(st.MgrMap.ActiveChange.Time),
				},
				Mds: &stats.MdsService{
					DaemonsUp:       int32(daemonsUp),
					HotStandbyCount: int32(standbyCountWanted),
				},
				Osd: &stats.OsdService{
					OsdCount:       int32(osdCount),
					OsdUp:          int32(osdUp),
					OsdIn:          int32(osdIn),
					OsdInUpdatedAt: timestamppb.New(st.OsdMap.LastInChange.Time),
					OsdUpUpdatedAt: timestamppb.New(st.OsdMap.LastUpChange.Time),
				},
				Rgw: &stats.RgwService{
					ActiveDaemon: int32(st.Rgw),
					HostCount:    1,
					ZoneCount:    1,
				},
			},
			Data: &stats.Data{
				Volumes: int32(1), // TODO still figuring out
				Pools: &stats.Pools{
					Pools: int32(poolCount),
					Pgs: &stats.Pgs{
						ActiveClean: int32(activeAndCleanPGs),
					},
				},
				Objects: &stats.Objects{
					ObjectCount: int32(st.PGInfo.ObjectStats.NumObjects),
					Size:        formatBytes(int64(stored)),
				},
				Usage: &stats.Usage{
					Used:      st.DF.Stats.TotalUsedBytes,
					Available: st.DF.Stats.TotalAvailBytes,
					Total:     st.DF.Stats.TotalBytes,
				},
			},
			Io: &stats.Io{
				ClientRead:     fmt.Sprintf("%s/s rd", clientReadBytes),
				ClientReadOps:  fmt.Sprintf("%d ops/s rd", readOps),
				ClientWriteOps: fmt.Sprintf("%d ops/s wr", writeOps),
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
		return nil, fmt.Errorf("unsupported type for MdsItems: %T", src)
	}

	return items, nil
}

func formatBytes(n int64) string {
	units := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB"}

	var i int
	size := float64(n)

	for i = 0; size >= 1024 && i < len(units)-1; i++ {
		size /= 1024
	}

	return fmt.Sprintf("%.2f %s", size, units[i])
}
