package recommendermodules

import (
	"context"
	"fmt"
	"strconv"

	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	moduleFactories["pool_sizes"] = NewPoolSizes
}

func NewPoolSizes() (Module, error) {
	return &PoolSizes{}, nil
}

type PoolSizes struct {
	Module
}

func (m *PoolSizes) GetName() string {
	return "pool_sizes"
}

func (m *PoolSizes) Run(ctx context.Context, p *Params) ([]*cephv1.ClusterRecommendation, error) {
	recs := []*cephv1.ClusterRecommendation{}

	blockPools, err := p.K8S.GetRookClient().CephV1().CephBlockPools(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, bp := range blockPools.Items {
		poolName := bp.Spec.Name
		if poolName == "" {
			poolName = bp.GetName()
		}

		if bp.Spec.IsErasureCoded() {
			if bp.Spec.ErasureCoded.CodingChunks < 2 || bp.Spec.ErasureCoded.DataChunks < 2 {
				recs = append(recs, &cephv1.ClusterRecommendation{
					Title:       fmt.Sprintf("Pool - %s: Erasure coded chunks not recommended", poolName),
					Description: "Erasure coded chunk coding/data chunks sizes not recommended, please see https://docs.ceph.com/en/latest/rados/operations/erasure-code/#erasure-code-profiles.",
					Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_HIGH,
					Type:        cephv1.RecommendationType_RECOMMENDATION_TYPE_POOL,
					ExtraData: &cephv1.ClusterRecommendation_RecommendedValue{
						RecommendedValue: &cephv1.RecommendedValue{
							Current:  fmt.Sprintf("Coding chunks: %d, Data chunks: %d", bp.Spec.ErasureCoded.CodingChunks, bp.Spec.ErasureCoded.DataChunks),
							Expected: "> 2",
						},
					},
				})
			}

		} else if bp.Spec.IsReplicated() {
			if bp.Spec.Replicated.Size < 3 {
				recs = append(recs, &cephv1.ClusterRecommendation{
					Title:       fmt.Sprintf("Pool - %s: Replicated size less than 3", poolName),
					Description: "Pool replicated size should be 3 or higher in production environments.",
					Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_HIGH,
					Type:        cephv1.RecommendationType_RECOMMENDATION_TYPE_POOL,
					ExtraData: &cephv1.ClusterRecommendation_RecommendedValue{
						RecommendedValue: &cephv1.RecommendedValue{
							Current:  strconv.Itoa(int(bp.Spec.Replicated.Size)),
							Expected: "3",
						},
					},
				})
			}
		} else if bp.Spec.IsHybridStoragePool() {
			// TODO what to check for a hybrid storage pool
		}
	}

	// TODO iterate over CephFilesystesm, CephObjectStorageGateways, CephNFSes to check the pool sizes

	return recs, nil
}
