package stats

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
	statsv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/stats/v1"
	statspb "github.com/koor-tech/data-control-center/gen/go/api/services/stats/v1"
	"github.com/koor-tech/data-control-center/gen/go/api/services/stats/v1/statsv1connect"
	cephcache "github.com/koor-tech/data-control-center/pkg/ceph/cache"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	k8scache "github.com/koor-tech/data-control-center/pkg/k8s/cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Server is used to implement stats services.
type Server struct {
	statsv1connect.StatsServiceHandler

	logger    *zap.Logger
	auth      *auth.GRPCAuth
	ceph      *cephcache.Cache
	k         *k8scache.Cache
	Namespace string
}

type Params struct {
	fx.In

	Logger   *zap.Logger
	GrpcAuth *auth.GRPCAuth
	K8S      *k8scache.Cache
	Ceph     *cephcache.Cache
	Cfg      *config.Config
}

func New(p Params) (*Server, error) {
	return &Server{
		logger:    p.Logger,
		auth:      p.GrpcAuth,
		ceph:      p.Ceph,
		k:         p.K8S,
		Namespace: p.Cfg.Namespace,
	}, nil
}

func (s *Server) RegisterService(g *gin.RouterGroup) {
	path, handler := statsv1connect.NewStatsServiceHandler(s, connect.WithInterceptors(
		s.auth.NewAuthInterceptor(),
	))
	g.Any(path+"/*path", gin.WrapH(handler))
}

func (s *Server) GetClusterStats(ctx context.Context, req *connect.Request[statspb.GetClusterStatsRequest]) (*connect.Response[statspb.GetClusterStatsResponse], error) {
	st, _ := s.ceph.GetHealthFull(ctx)

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
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error caused by %w", err))

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

	now := time.Now()

	resp := &statspb.GetClusterStatsResponse{
		Stats: &statsv1.ClusterStats{
			Id:      st.MonStatus.Monmap.Fsid,
			Status:  st.ClusterHealthStatus(),
			Crashes: st.Crashes(),
			Services: &statsv1.Services{
				Mon: &statsv1.MonService{
					DaemonCount:  int32(len(st.MonStatus.Monmap.Mons)),
					Quorum:       strings.Join(monNames, ","),
					CreatedSince: int64(now.Sub(st.MonStatus.Monmap.Created)),
					UpdatedSince: int64(now.Sub(st.MonStatus.Monmap.Modified)),
				},
				Mgr: &statsv1.MgrService{
					Active:       st.MgrMap.ActiveName,
					Standbys:     standBys,
					UpdatedSince: int64(now.Sub(st.MgrMap.ActiveChange.Time)),
				},
				Mds: &statsv1.MdsService{
					DaemonsUp:       int32(daemonsUp),
					HotStandbyCount: int32(standbyCountWanted),
				},
				Osd: &statsv1.OsdService{
					OsdCount:          int32(osdCount),
					OsdUp:             int32(osdUp),
					OsdIn:             int32(osdIn),
					OsdInUpdatedSince: int64(now.Sub(st.OsdMap.LastInChange.Time)),
					OsdUpUpdatedSince: int64(now.Sub(st.OsdMap.LastUpChange.Time)),
				},
				Rgw: &statsv1.RgwService{
					ActiveDaemon: int32(st.Rgw),
					HostCount:    int32(st.Hosts),
					ZoneCount:    1, // TODO still figuring out https://linear.app/koorinc/issue/KSD-290/
				},
			},
			Data: &statsv1.Data{
				Volumes: int32(1), // TODO still figuring out https://linear.app/koorinc/issue/KSD-290/
				Pools: &statsv1.Pools{
					Pools: int32(poolCount),
					Pgs: &statsv1.PGs{
						ActiveClean: int32(activeAndCleanPGs),
					},
				},
				Objects: &statsv1.Objects{
					ObjectCount: int32(st.PGInfo.ObjectStats.NumObjects),
					Size:        st.ObjectSize(),
				},
				Usage: &statsv1.Usage{
					Used:      st.DF.Stats.TotalUsedBytes,
					Available: st.DF.Stats.TotalAvailBytes,
					Total:     st.DF.Stats.TotalBytes,
				},
			},
			Iops: &statsv1.IOPS{
				ClientRead:     int64(st.ClientPerf.ReadBytesSec),
				ClientWrite:    int64(st.ClientPerf.WriteBytesSec),
				ClientReadOps:  int64(st.ClientPerf.ReadOpPerSec),
				ClientWriteOps: int64(st.ClientPerf.WriteOpPerSec),
			},
		},
	}

	return &connect.Response[statspb.GetClusterStatsResponse]{
		Msg: resp,
	}, nil
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

func (s *Server) GetClusterResources(ctx context.Context, req *connect.Request[statspb.GetClusterResourcesRequest]) (*connect.Response[statspb.GetClusterResourcesResponse], error) {
	deployments, _ := s.k.GetClusterDeployments(s.Namespace)

	resources, _ := s.k.GetCephResources(s.Namespace)
	return &connect.Response[statspb.GetClusterResourcesResponse]{
		Msg: &statspb.GetClusterResourcesResponse{
			Deployments: deployments,
			Resources:   resources,
		},
	}, nil
}

// GetClusterNodes Iterate over every Ceph cluster Pod (not CSI Pods) to collect the nodes used for the storage cluster
func (s *Server) GetClusterNodes(ctx context.Context, req *connect.Request[statspb.GetClusterNodesRequest]) (*connect.Response[statspb.GetClusterNodesResponse], error) {
	nodes, _ := s.k.GetStorageNodes(s.Namespace)

	return &connect.Response[statspb.GetClusterNodesResponse]{
		Msg: &statspb.GetClusterNodesResponse{
			Nodes: nodes,
		},
	}, nil
}

func (s *Server) GetClusterRadar(ctx context.Context, req *connect.Request[statspb.GetClusterRadarRequest]) (*connect.Response[statspb.GetClusterRadarResponse], error) {
	radar := &statsv1.ClusterRadar{}

	// Cluster Health
	cephStatus, _ := s.ceph.GetHealthFull(ctx)

	radar.ClusterHealth = cephStatus.ClusterHealth()

	// Nodes Health
	nodes, _ := s.k.GetStorageNodes(s.Namespace)
	totalNodes := len(nodes)
	healthyNodes := 0
	for _, node := range nodes {
		if node.Status == statsv1.ResourceStatus_RESOURCE_STATUS_READY {
			healthyNodes++
		}
	}
	radar.NodesHealth = float32(healthyNodes/totalNodes) * 100

	// Storage Capacity
	// TODO is the total used raw ratio the correct indicator to use?
	radar.CapacityAvailable = float32(math.Ceil((1-cephStatus.DF.Stats.TotalUsedRawRatio)*100*10) / 10)

	// Stability
	stabilityReduction := len(cephStatus.Health.Checks) * 10
	radar.Stability = float32(100 - stabilityReduction)
	if radar.Stability < 0 {
		radar.Stability = 0
	}

	// Reliability
	// TODO calculate from `Ceph*` CustomResources `.replicated.size`, `replicas`, etc., for the components
	radar.Reliability = 100

	return &connect.Response[statspb.GetClusterRadarResponse]{
		Msg: &statspb.GetClusterRadarResponse{
			Radar: radar,
		},
	}, nil
}

func (s *Server) GetKoorCluster(ctx context.Context, req *connect.Request[statspb.GetKoorClusterRequest]) (*connect.Response[statspb.GetKoorClusterResponse], error) {
	kc, _ := s.k.GetKoorCluster(s.Namespace)

	res := connect.NewResponse(&statspb.GetKoorClusterResponse{
		KoorCluster: kc,
	})
	return res, nil
}
