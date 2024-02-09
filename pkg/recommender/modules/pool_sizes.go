package recommendermodules

import (
	"context"
	"fmt"
	"strconv"

	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	rookcephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
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
	for _, item := range blockPools.Items {
		name := item.Spec.Name
		if name == "" {
			name = item.GetName()
		}

		recs = append(recs, m.checkPool(name, item.Spec.PoolSpec)...)
	}

	fses, err := p.K8S.GetRookClient().CephV1().CephFilesystems(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range fses.Items {
		name := item.GetName()

		recs = append(recs, m.checkPool(fmt.Sprintf("%s-metadata", name), item.Spec.MetadataPool)...)

		for key, pool := range item.Spec.DataPools {
			recs = append(recs, m.checkPool(fmt.Sprintf("%s-data%d", name, key), pool.PoolSpec)...)
		}
	}

	objs, err := p.K8S.GetRookClient().CephV1().CephObjectStores(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range objs.Items {
		name := item.GetName()

		recs = append(recs, m.checkPool(fmt.Sprintf("%s.rgw.meta", name), item.Spec.MetadataPool)...)
		recs = append(recs, m.checkPool(fmt.Sprintf("%s.rgw.buckets.data", name), item.Spec.DataPool)...)
	}

	return recs, nil
}

func (m *PoolSizes) checkPool(name string, pSpec rookcephv1.PoolSpec) []*cephv1.ClusterRecommendation {
	recs := []*cephv1.ClusterRecommendation{}

	if pSpec.IsErasureCoded() {
		if pSpec.ErasureCoded.CodingChunks < 2 || pSpec.ErasureCoded.DataChunks < 2 {
			recs = append(recs, &cephv1.ClusterRecommendation{
				Title:       fmt.Sprintf("Pool - %s: Erasure coded chunks not recommended", name),
				Description: "Erasure coded chunk coding/data chunks sizes not recommended, please see https://docs.ceph.com/en/latest/rados/operations/erasure-code/#erasure-code-profiles.",
				Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_HIGH,
				Type:        cephv1.RecommendationType_RECOMMENDATION_TYPE_POOL,
				ExtraData: &cephv1.ClusterRecommendation_RecommendedValue{
					RecommendedValue: &cephv1.RecommendedValue{
						Current:  fmt.Sprintf("Coding chunks: %d, Data chunks: %d", pSpec.ErasureCoded.CodingChunks, pSpec.ErasureCoded.DataChunks),
						Expected: "> 2",
					},
				},
			})
		}

	} else if pSpec.IsReplicated() {
		if pSpec.Replicated.Size < 3 {
			recs = append(recs, &cephv1.ClusterRecommendation{
				Title:       fmt.Sprintf("Pool - %s: Replicated size less than 3", name),
				Description: "Pool replicated size should be 3 or higher in production environments.",
				Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_HIGH,
				Type:        cephv1.RecommendationType_RECOMMENDATION_TYPE_POOL,
				ExtraData: &cephv1.ClusterRecommendation_RecommendedValue{
					RecommendedValue: &cephv1.RecommendedValue{
						Current:  strconv.Itoa(int(pSpec.Replicated.Size)),
						Expected: "3",
					},
				},
			})
		}
	}
	// TODO check sizes for hybrid storage pools

	return recs
}
