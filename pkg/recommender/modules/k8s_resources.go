package recommendermodules

import (
	"context"

	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var k8sResourcesClusterResMustHave = []string{
	"mon",
	"mgr",
	"osd",
	"mds",
	"rgw",
}

func init() {
	moduleFactories["k8s_resources"] = NewK8SResources
}

func NewK8SResources() (Module, error) {
	return &K8SResources{}, nil
}

type K8SResources struct {
	Module
}

func (m *K8SResources) GetName() string {
	return "k8s_resources"
}

func (m *K8SResources) Run(ctx context.Context, p *Params) ([]*cephv1.ClusterRecommendation, error) {
	recs := []*cephv1.ClusterRecommendation{}

	clusters, err := p.K8S.GetRookClient().CephV1().CephClusters(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range clusters.Items {
		// TODO check against k8sResourcesClusterResMustHave
		_ = item
	}

	fses, err := p.K8S.GetRookClient().CephV1().CephFilesystems(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range fses.Items {
		_ = item
	}

	rgws, err := p.K8S.GetRookClient().CephV1().CephObjectStores(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range rgws.Items {
		_ = item
	}

	nfses, err := p.K8S.GetRookClient().CephV1().CephNFSes(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range nfses.Items {
		_ = item
	}

	// TODO check CephFilesystems, CephObjectStores, and CephNFSes

	return recs, nil
}
