package k8s

import (
	"context"

	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetCephV1Resource[T any](ctx context.Context, client *kubernetes.Clientset, resource string, name string, opts metav1.GetOptions) (*T, error) {
	return GetCustomResource[T](ctx, client, cephv1.SchemeGroupVersion.WithResource(resource), resource, name, opts)
}

func ListCephV1Resources[T any](ctx context.Context, client *kubernetes.Clientset, resource string, opts metav1.ListOptions) ([]*T, error) {
	return ListCustomResource[T](ctx, client, cephv1.SchemeGroupVersion.WithResource(resource), resource, opts)
}
