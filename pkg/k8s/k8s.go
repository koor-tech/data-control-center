package k8s

import (
	"context"
	"encoding/json"
	"os"

	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
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

func GetCephResource[T any](client *kubernetes.Clientset, resource string, name string) (*T, error) {
	cliSet := dynamic.New(client.RESTClient())
	obj, err := cliSet.Resource(cephv1.SchemeGroupVersion.WithResource(resource)).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	bytes, err := obj.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var crdObj *T
	if err := json.Unmarshal(bytes, &crdObj); err != nil {
		return nil, err
	}

	return crdObj, nil
}
