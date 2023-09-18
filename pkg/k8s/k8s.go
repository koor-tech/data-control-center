package k8s

import (
	"context"

	"github.com/koor-tech/data-control-center/gen/go/api/resources/stats"
	"github.com/koor-tech/data-control-center/pkg/config"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	"go.uber.org/fx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	cruntimeconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

var Module = fx.Module("k8s",
	fx.Provide(
		New,
	),
)

// Retrieve Rook Ceph Pods
// Show other objects as well

type K8s struct {
	client *kubernetes.Clientset
}

func New(cfg *config.Config) (*K8s, error) {
	// Copied from the `GetConfig()` method docs
	//
	// > Config precedence:
	// >
	// > * --kubeconfig flag pointing at a file
	// > * KUBECONFIG environment variable pointing at a file
	// > * In-cluster config if running in cluster
	// > * $HOME/.kube/config if exists.
	config, err := cruntimeconfig.GetConfig()
	if err != nil {
		return nil, err
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &K8s{
		client: clientset,
	}, nil
}

var (
	ImportantDeployPrefixes = []string{"rook-ceph-", "csi-", "extended-ceph-exporter"}
)

func (k *K8s) GetClusterDeployments(ctx context.Context, namespace string) ([]*stats.ResourceInfo, error) {
	deployments, err := k.client.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	res := []*stats.ResourceInfo{}
	for _, deployment := range deployments.Items {
		status := stats.ResourceStatus_RESOURCE_UNKNOWN
		if deployment.Status.ReadyReplicas >= deployment.Status.Replicas {
			status = stats.ResourceStatus_RESOURCE_READY
		} else {
			status = stats.ResourceStatus_RESOURCE_NOT_READY
		}

		res = append(res, &stats.ResourceInfo{
			// Why do we set it manually? https://github.com/kubernetes/client-go/issues/541
			Apiversion: "apps/v1",
			Kind:       "Deployment",
			Namespace:  deployment.Namespace,
			Name:       deployment.Name,
			Status:     status,
		})
	}

	return res, nil
}

func (k *K8s) GetCephResources(ctx context.Context, namespace string) ([]*stats.ResourceInfo, error) {
	list, err := ListCephV1Resources[cephv1.CephCluster](ctx, k.client, "cephclusters", namespace, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	res := []*stats.ResourceInfo{}
	for _, obj := range list {
		status := stats.ResourceStatus_RESOURCE_UNKNOWN
		if obj.Status.CephStatus != nil {
			if obj.Status.State == cephv1.ClusterStateCreated || obj.Status.State == cephv1.ClusterStateConnected {
				status = stats.ResourceStatus_RESOURCE_READY
			} else {
				status = stats.ResourceStatus_RESOURCE_NOT_READY
			}
		}

		res = append(res, &stats.ResourceInfo{
			Apiversion: obj.APIVersion,
			Kind:       obj.Kind,
			Namespace:  obj.Namespace,
			Name:       obj.Name,
			Status:     status,
		})
	}

	// TODO iterate over all the important ceph cluster resources (e.g., CephCluster, CephStorageBlockPool, CephFilesystems, CephObjectStores)

	return res, nil
}
