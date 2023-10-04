package k8s

import (
	"context"

	statsv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/stats/v1"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8s) GetCephResources(ctx context.Context, namespace string) ([]*statsv1.ResourceInfo, error) {
	res := []*statsv1.ResourceInfo{}

	// CephClusters
	clusters, err := k.rookclient.CephV1().CephClusters(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, obj := range clusters.Items {
		status := statsv1.ResourceStatus_RESOURCE_STATUS_UNKNOWN
		if obj.Status.CephStatus != nil {
			if obj.Status.State == cephv1.ClusterStateCreated || obj.Status.State == cephv1.ClusterStateConnected {
				status = statsv1.ResourceStatus_RESOURCE_STATUS_READY
			} else if obj.Status.State == cephv1.ClusterStateConnecting || obj.Status.State == cephv1.ClusterStateCreating || obj.Status.State == cephv1.ClusterStateUpdating {
				status = statsv1.ResourceStatus_RESOURCE_STATUS_NOT_READY
			}
		}

		res = append(res, &statsv1.ResourceInfo{
			Apiversion: obj.APIVersion,
			Kind:       obj.Kind,
			Namespace:  obj.Namespace,
			Name:       obj.Name,
			Status:     status,
			Replicas:   int32(obj.Spec.Mon.Count),
		})
	}

	// CephStorageBlockPools
	blockPools, err := k.rookclient.CephV1().CephBlockPools(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, obj := range blockPools.Items {
		status := statsv1.ResourceStatus_RESOURCE_STATUS_UNKNOWN
		if obj.Status != nil {
			if obj.Status.Phase == cephv1.ConditionReady || obj.Status.Phase == cephv1.ConditionProgressing {
				status = statsv1.ResourceStatus_RESOURCE_STATUS_READY
			} else {
				status = statsv1.ResourceStatus_RESOURCE_STATUS_NOT_READY
			}
		}

		res = append(res, &statsv1.ResourceInfo{
			Apiversion: obj.APIVersion,
			Kind:       obj.Kind,
			Namespace:  obj.Namespace,
			Name:       obj.Name,
			Status:     status,
			Replicas:   int32(obj.Spec.Replicated.Size),
		})
	}

	// CephFilesystems
	filesystems, err := k.rookclient.CephV1().CephFilesystems(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, obj := range filesystems.Items {
		status := statsv1.ResourceStatus_RESOURCE_STATUS_UNKNOWN
		if obj.Status != nil {
			if obj.Status.Phase == cephv1.ConditionReady || obj.Status.Phase == cephv1.ConditionProgressing {
				status = statsv1.ResourceStatus_RESOURCE_STATUS_READY
			} else {
				status = statsv1.ResourceStatus_RESOURCE_STATUS_NOT_READY
			}
		}

		res = append(res, &statsv1.ResourceInfo{
			Apiversion: obj.APIVersion,
			Kind:       obj.Kind,
			Namespace:  obj.Namespace,
			Name:       obj.Name,
			Status:     status,
			Replicas:   obj.Spec.MetadataServer.ActiveCount * 2,
		})
	}

	// CephObjectStores
	objectStores, err := k.rookclient.CephV1().CephObjectStores(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, obj := range objectStores.Items {
		status := statsv1.ResourceStatus_RESOURCE_STATUS_UNKNOWN
		if obj.Status != nil {
			if obj.Status.Phase == cephv1.ConditionReady || obj.Status.Phase == cephv1.ConditionProgressing {
				status = statsv1.ResourceStatus_RESOURCE_STATUS_READY
			} else {
				status = statsv1.ResourceStatus_RESOURCE_STATUS_NOT_READY
			}
		}

		res = append(res, &statsv1.ResourceInfo{
			Apiversion: obj.APIVersion,
			Kind:       obj.Kind,
			Namespace:  obj.Namespace,
			Name:       obj.Name,
			Status:     status,
			Replicas:   obj.Spec.Gateway.Instances,
		})
	}

	return res, nil
}
