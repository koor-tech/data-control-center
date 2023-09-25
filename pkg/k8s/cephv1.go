package k8s

import (
	"context"

	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetCephV1Resource[T any](ctx context.Context, client *kubernetes.Clientset, resource string, namespace string, name string, opts metav1.GetOptions) (*T, error) {
	return GetCustomResource[T](ctx, client, cephv1.SchemeGroupVersion.WithResource(resource), namespace, name, opts)
}

func ListCephV1Resources[T any](ctx context.Context, client *kubernetes.Clientset, resource string, namespace string, opts metav1.ListOptions) ([]*T, error) {
	return ListCustomResource[T](ctx, client, cephv1.SchemeGroupVersion.WithResource(resource), namespace, opts)
}

func (k *K8s) ListCephClusters(ctx context.Context, namespace string) ([]*cephv1.CephCluster, error) {
	list, err := ListCephV1Resources[cephv1.CephCluster](ctx, k.client, "cephclusters", namespace, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return list, err
}

func (k *K8s) ListCephStorageBlockPools(ctx context.Context, namespace string) ([]*cephv1.CephBlockPool, error) {
	list, err := ListCephV1Resources[cephv1.CephBlockPool](ctx, k.client, "cephblockpools", namespace, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return list, err
}

func (k *K8s) ListCephFilesystems(ctx context.Context, namespace string) ([]*cephv1.CephFilesystem, error) {
	list, err := ListCephV1Resources[cephv1.CephFilesystem](ctx, k.client, "cephfilesystems", namespace, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return list, err
}

func (k *K8s) ListCephObjectStores(ctx context.Context, namespace string) ([]*cephv1.CephObjectStore, error) {
	list, err := ListCephV1Resources[cephv1.CephObjectStore](ctx, k.client, "cephobjectstores", namespace, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return list, err
}
