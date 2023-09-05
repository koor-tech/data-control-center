package k8s

import (
	"os"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Retrieve Rook Ceph Pods
// Show other objects as well

type K8s struct {
	client *kubernetes.Clientset
}

func New() (*K8s, error) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return &K8s{
		client: clientset,
	}, nil
}

func (k *K8s) GetPods() ([]*corev1.Pod, error) {

	// TODO get Rook Ceph daemon deployments

	return nil, nil
}

func (k *K8s) GetCephResources() error {

	// TODO Get each Ceph resources per namespace

	return nil
}
