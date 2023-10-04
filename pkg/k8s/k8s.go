package k8s

import (
	"context"

	statsv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/stats/v1"
	"github.com/koor-tech/data-control-center/pkg/config"
	rookclient "github.com/rook/rook/pkg/client/clientset/versioned"
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
	client     *kubernetes.Clientset
	rookclient *rookclient.Clientset
}

func New(cfg *config.Config) (*K8s, error) {
	// Copied from the `GetConfig()` method docs of `sigs.k8s.io/controller-runtime`
	// https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client/config#GetConfig
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

	rookClientset, err := rookclient.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &K8s{
		client:     clientset,
		rookclient: rookClientset,
	}, nil
}

var (
	ImportantDeployPrefixes = []string{"rook-ceph-", "csi-", "extended-ceph-exporter"}
)

func (k *K8s) GetClusterDeployments(ctx context.Context, namespace string) ([]*statsv1.ResourceInfo, error) {
	deployments, err := k.client.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	res := []*statsv1.ResourceInfo{}
	for _, deployment := range deployments.Items {
		status := statsv1.ResourceStatus_RESOURCE_STATUS_UNKNOWN
		if deployment.Status.ReadyReplicas >= deployment.Status.Replicas {
			status = statsv1.ResourceStatus_RESOURCE_STATUS_READY
		} else {
			status = statsv1.ResourceStatus_RESOURCE_STATUS_NOT_READY
		}

		res = append(res, &statsv1.ResourceInfo{
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
