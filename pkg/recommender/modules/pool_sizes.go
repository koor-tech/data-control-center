package recommendermodules

import (
	"context"
	"fmt"

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
		if bp.Spec.IsErasureCoded() {
			// TODO
		} else if bp.Spec.IsReplicated() {
			// TODO
			if bp.Spec.Replicated.Size < 3 {
				poolName := bp.Spec.Name
				if poolName == "" {
					poolName = bp.GetName()
				}

				recs = append(recs, &cephv1.ClusterRecommendation{
					Title:       fmt.Sprintf("Pool - %s: Replicated size less than 3", poolName),
					Description: "Pool replicated size Should be 3 or higher in production environments.",
					Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_HIGH,
				})
			}
		} else if bp.Spec.IsHybridStoragePool() {
			// TODO
		}
	}

	// TODO iterate over CephFilesystesm, CephObjectStorageGateways, CephNFSes to check the pool sizes

	return recs, nil
}
