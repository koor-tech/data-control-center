package recommendermodules

import (
	"context"
	"slices"

	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	corev1 "k8s.io/api/core/v1"
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
		// Collect the set global config options so we can check if a scrubbing schedule is set
		resOpts := []string{}
		for opt := range item.Spec.Resources {
			resOpts = append(resOpts, opt)
		}

		found := true
		for _, opt := range k8sResourcesClusterResMustHave {
			if !slices.Contains(resOpts, opt) {
				found = false
				break
			}
		}

		if found {
			break
		}

		recs = append(recs, &cephv1.ClusterRecommendation{
			Title:       "Resources: No requests/limits set on cluster components",
			Description: "It is recommended to set sensible resources requests and limits on the most important cluster components (MON, MGR, OSD). See https://rook.io/docs/rook/latest-release/CRDs/Cluster/ceph-cluster-crd/#resource-requirementslimits",
			Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_HIGH,
			Type:        cephv1.RecommendationType_RECOMMENDATION_TYPE_CLUSTER,
		})
	}

	fses, err := p.K8S.GetRookClient().CephV1().CephFilesystems(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range fses.Items {
		if !m.checkIfHasResourcesSet(item.Spec.MetadataServer.Resources) {
			recs = append(recs, &cephv1.ClusterRecommendation{
				Title:       "Resources: No requests/limits set on CephFS Metadata Server ",
				Description: "It is recommended to set sensible resources requests and limits on the CephFS Metadata server component. See https://rook.io/docs/rook/latest-release/CRDs/Cluster/ceph-cluster-crd/#resource-requirementslimits",
				Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_MEDIUM,
				Type:        cephv1.RecommendationType_RECOMMENDATION_TYPE_MISC,
			})
		}
	}

	rgws, err := p.K8S.GetRookClient().CephV1().CephObjectStores(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range rgws.Items {
		if !m.checkIfHasResourcesSet(item.Spec.Gateway.Resources) {
			recs = append(recs, &cephv1.ClusterRecommendation{
				Title:       "Resources: No requests/limits set on Ceph RGW ",
				Description: "It is recommended to set sensible resources requests and limits on the Ceph RGW component. See https://rook.io/docs/rook/latest-release/CRDs/Cluster/ceph-cluster-crd/#resource-requirementslimits",
				Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_MEDIUM,
				Type:        cephv1.RecommendationType_RECOMMENDATION_TYPE_MISC,
			})
		}
	}

	nfses, err := p.K8S.GetRookClient().CephV1().CephNFSes(p.Namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range nfses.Items {
		if !m.checkIfHasResourcesSet(item.Spec.Server.Resources) {
			recs = append(recs, &cephv1.ClusterRecommendation{
				Title:       "Resources: No requests/limits set on Ceph NFS server ",
				Description: "It is recommended to set sensible resources requests and limits on the Ceph NFS server component. See https://rook.io/docs/rook/latest-release/CRDs/Cluster/ceph-cluster-crd/#resource-requirementslimits",
				Level:       cephv1.RecommendationLevel_RECOMMENDATION_LEVEL_MEDIUM,
				Type:        cephv1.RecommendationType_RECOMMENDATION_TYPE_MISC,
			})
		}
	}

	return recs, nil
}

func (m *K8SResources) checkIfHasResourcesSet(res corev1.ResourceRequirements) bool {
	if res.Requests.Cpu() == nil {
		return false
	}

	if res.Requests.Memory() == nil {
		return false
	}

	if res.Limits.Cpu() == nil {
		return false
	}

	if res.Limits.Memory() == nil {
		return false
	}

	return true
}
