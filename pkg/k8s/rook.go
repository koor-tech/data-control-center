package k8s

import (
	"context"

	"github.com/koor-tech/data-control-center/gen/go/api/resources/stats"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
)

func (k *K8s) GetCephResources(ctx context.Context, namespace string) ([]*stats.ResourceInfo, error) {
	res := []*stats.ResourceInfo{}

	// CephClusters
	clusters, err := k.ListCephClusters(ctx, namespace)
	if err != nil {
		return nil, err
	}

	for _, obj := range clusters {
		status := stats.ResourceStatus_RESOURCE_UNKNOWN
		if obj.Status.CephStatus != nil {
			if obj.Status.State == cephv1.ClusterStateCreated || obj.Status.State == cephv1.ClusterStateConnected {
				status = stats.ResourceStatus_RESOURCE_READY
			} else if obj.Status.State == cephv1.ClusterStateConnecting || obj.Status.State == cephv1.ClusterStateCreating || obj.Status.State == cephv1.ClusterStateUpdating {
				status = stats.ResourceStatus_RESOURCE_NOT_READY
			}
		}

		res = append(res, &stats.ResourceInfo{
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
		status := stats.ResourceStatus_RESOURCE_UNKNOWN
		if obj.Status != nil {
			if obj.Status.Phase == cephv1.ConditionReady || obj.Status.Phase == cephv1.ConditionProgressing {
				status = stats.ResourceStatus_RESOURCE_READY
			} else {
				status = stats.ResourceStatus_RESOURCE_NOT_READY
			}
		}

		res = append(res, &stats.ResourceInfo{
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
		status := stats.ResourceStatus_RESOURCE_UNKNOWN
		if obj.Status != nil {
			if obj.Status.Phase == cephv1.ConditionReady || obj.Status.Phase == cephv1.ConditionProgressing {
				status = stats.ResourceStatus_RESOURCE_READY
			} else {
				status = stats.ResourceStatus_RESOURCE_NOT_READY
			}
		}

		res = append(res, &stats.ResourceInfo{
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
		status := stats.ResourceStatus_RESOURCE_UNKNOWN
		if obj.Status != nil {
			if obj.Status.Phase == cephv1.ConditionReady || obj.Status.Phase == cephv1.ConditionProgressing {
				status = stats.ResourceStatus_RESOURCE_READY
			} else {
				status = stats.ResourceStatus_RESOURCE_NOT_READY
			}
		}

		res = append(res, &stats.ResourceInfo{
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
