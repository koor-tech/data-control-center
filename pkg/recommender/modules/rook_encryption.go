package recommendermodules

import (
	"context"

	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	moduleFactories["k8s_resources"] = NewRookEncryption
}

func NewRookEncryption() (Module, error) {
	return &RookEncryption{}, nil
}

type RookEncryption struct {
	Module
}

func (m *RookEncryption) GetName() string {
	return "rook_encryption"
}

func (m *RookEncryption) Run(ctx context.Context, p *Params) ([]*cephv1.ClusterRecommendation, error) {
	recs := []*cephv1.ClusterRecommendation{}

	clusters, err := p.K8S.GetRookClient().CephV1().CephClusters(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range clusters.Items {
		if item.Spec.Network.Connections.Encryption == nil || !item.Spec.Network.Connections.Encryption.Enabled {
			continue
		}

		recs = append(recs, &cephv1.ClusterRecommendation{
			Title:       "Ceph network encryption not active",
			Description: "It is recommended to enable Ceph network encryption for clusters that might be running in \"untrusted\" networks.",
			Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_INFORMAL,
			Type:        cephv1.RecommendationType_RECOMMENDATION_TYPE_CLUSTER,
		})
	}

	return recs, nil
}
