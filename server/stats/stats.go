package stats

import (
	"context"
	"fmt"
	"math"

	"connectrpc.com/connect"
	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	k8sv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/k8s/v1"
	statspb "github.com/koor-tech/data-control-center/gen/go/api/services/stats/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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

	blockImages, _ := s.ceph.GetBlockImages(ctx)
	volumes := len(blockImages)

	resp := &statspb.GetClusterStatsResponse{
		Stats: &cephv1.ClusterStats{
			Id:      st.MonStatus.Monmap.Fsid,
			Status:  st.ClusterHealthStatus(),
			Crashes: st.Crashes(),
			Services: &cephv1.Services{
				Mon: &cephv1.MonService{
					DaemonCount:  int32(len(st.MonStatus.Monmap.Mons)),
					Quorum:       monNames,
					CreatedSince: timestamppb.New(st.MonStatus.Monmap.Created),
					UpdatedSince: timestamppb.New(st.MonStatus.Monmap.Modified),
				},
				Mgr: &cephv1.MgrService{
					Active:       st.MgrMap.ActiveName,
					Standbys:     standBys,
					UpdatedSince: timestamppb.New(st.MgrMap.ActiveChange.Time),
				},
				Mds: &cephv1.MdsService{
					DaemonsUp:       int32(daemonsUp),
					HotStandbyCount: int32(standbyCountWanted),
				},
				Osd: &cephv1.OsdService{
					OsdCount:          int32(osdCount),
					OsdUp:             int32(osdUp),
					OsdIn:             int32(osdIn),
					OsdInUpdatedSince: timestamppb.New(st.OsdMap.LastInChange.Time),
					OsdUpUpdatedSince: timestamppb.New(st.OsdMap.LastUpChange.Time),
				},
				Rgw: &cephv1.RgwService{
					ActiveDaemon: int32(st.Rgw),
					HostCount:    int32(st.Hosts),
					ZoneCount:    1, // TODO still figuring out https://linear.app/koorinc/issue/KSD-290/
				},
			},
			Data: &cephv1.Data{
				Volumes: int32(volumes),
				Pools: &cephv1.Pools{
					Pools: int32(poolCount),
					Pgs: &cephv1.PGs{
						ActiveClean: int32(activeAndCleanPGs),
					},
				},
				Objects: &cephv1.Objects{
					ObjectCount: int32(st.PGInfo.ObjectStats.NumObjects),
					Size:        st.ObjectSize(),
				},
				Usage: &cephv1.Usage{
					Used:      st.DF.Stats.TotalUsedBytes,
					Available: st.DF.Stats.TotalAvailBytes,
					Total:     st.DF.Stats.TotalBytes,
				},
			},
			Iops: &cephv1.IOPS{
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
	radar := &cephv1.ClusterRadar{}

	// Cluster Health
	cephStatus, _ := s.ceph.GetHealthFull(ctx)

	radar.ClusterHealth = cephStatus.ClusterHealth()

	// Nodes Health
	nodes, _ := s.k.GetStorageNodes(s.Namespace)
	totalNodes := len(nodes)
	healthyNodes := 0
	for _, node := range nodes {
		if node.Status == k8sv1.ResourceStatus_RESOURCE_STATUS_READY {
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
