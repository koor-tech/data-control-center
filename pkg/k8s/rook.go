package k8s

import (
	"context"

	statsv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/stats/v1"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
)

func (k *K8s) GetCephResources(ctx context.Context, namespace string) ([]*statsv1.ResourceInfo, error) {
	res := []*statsv1.ResourceInfo{}

	// CephClusters
	clusters, err := k.ListCephClusters(ctx, namespace)
	if err != nil {
		return nil, err
	}

	for _, obj := range clusters {
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
	blockPools, err := k.ListCephStorageBlockPools(ctx, namespace)
	if err != nil {
		return nil, err
	}

	for _, obj := range blockPools {
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
	filesystems, err := k.ListCephFilesystems(ctx, namespace)
	if err != nil {
		return nil, err
	}

	for _, obj := range filesystems {
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
	objectStores, err := k.ListCephObjectStores(ctx, namespace)
	if err != nil {
		return nil, err
	}

	for _, obj := range objectStores {
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
