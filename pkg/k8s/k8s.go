package k8s

import (
	"context"
	"os"
	"path"

	"github.com/koor-tech/data-control-center/gen/go/api/resources/stats"
	"github.com/koor-tech/data-control-center/pkg/config"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	"go.uber.org/fx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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
	kubeconfig := os.Getenv("KUBECONFIG")
	if cfg.Kubernetes.Kubeconfig != "" {
		kubeconfig = cfg.Kubernetes.Kubeconfig
	}
	if kubeconfig == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		kubeconfig = path.Join(home, ".kube/config")
	}

	// Use the current context from the provided kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
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
			Apiversion: deployment.APIVersion,
			Kind:       deployment.Kind,
			Namespace:  deployment.Namespace,
			Name:       deployment.Name,
			Status:     status,
		})
	}

	return res, nil
}

func (k *K8s) GetCephResources(ctx context.Context) ([]*stats.ResourceInfo, error) {
	list, err := ListCephV1Resources[cephv1.CephCluster](ctx, k.client, "cephcluster", metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	_ = list
	// TODO

	return nil, nil
}
