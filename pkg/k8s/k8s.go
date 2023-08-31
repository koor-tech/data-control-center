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

func GetPods() ([]*corev1.Pod, error) {

	return nil, nil
}

// TODO how should we list the daemons? List the Ceph* resources?
