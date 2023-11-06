package k8s

import (
	"github.com/koor-tech/data-control-center/pkg/config"
	rookclient "github.com/rook/rook/pkg/client/clientset/versioned"
	"go.uber.org/fx"
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

func (k *K8s) GetClient() *kubernetes.Clientset {
	return k.client
}

func (k *K8s) GetRookClient() *rookclient.Clientset {
	return k.rookclient
}
