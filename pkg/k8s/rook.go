package k8s

import (
	"context"
	koorv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	statsv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/stats/v1"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
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
				status = statsv1.ResourceStatus_RESOURCE_STATUS_PROGRESSING
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

func (k *K8s) getAsYaml(resource interface{}) (string, error) {
	resourceYaml, err := yaml.Marshal(resource)
	if err != nil {
		return "", err
	}
	return string(resourceYaml), nil
}

var cephObjects = []string{
	"cephclusters",
	"cephblockpoolradosnamespaces",
	"cephblockpools",
	"cephbucketnotifications",
	"cephbuckettopics",
	"cephclients",
	"cephcosidrivers",
	"cephfilesystemmirrors",
	"cephfilesystems",
	"cephfilesystemsubvolumegroups",
	"cephnfses",
	"cephobjectrealms",
	"cephobjectstores",
	"cephobjectstoreusers",
	"cephobjectzonegroups",
	"cephobjectzones",
	"cephrbdmirrors",
}

type Resource struct {
	Name      string
	Namespace string
	Content   string
	Kind      string
	Object    string
}

const (
	cephGroupResource = "ceph.rook.io"
	cephGroupVersion  = "v1"
)

func (k *K8s) GetYamlCephResources(ctx context.Context, namespace string) ([]Resource, error) {

	var resources []Resource
	for _, cephObject := range cephObjects {
		resourceId := schema.GroupVersionResource{
			Group:    cephGroupResource,
			Version:  cephGroupVersion,
			Resource: cephObject,
		}

		cephResources, err := k.dynamicClientSet.Resource(resourceId).Namespace(namespace).List(ctx, v1.ListOptions{})
		if err != nil {
			if k8sErrors.IsNotFound(err) {
				continue

			}
			return nil, err
		}

		for _, cephResource := range cephResources.Items {
			objectAsYaml, err := k.getAsYaml(cephResource.Object)
			if err != nil {
				return nil, err
			}
			resource := Resource{
				Name:      cephResource.GetName(),
				Namespace: cephResource.GetNamespace(),
				Content:   objectAsYaml,
				Kind:      cephResource.GetKind(),
				Object:    cephObject,
			}
			resources = append(resources, resource)
		}
	}

	return resources, nil
}

func (k *K8s) SaveResource(ctx context.Context, resource *koorv1.Resource) error {
	var (
		object *unstructured.Unstructured
		err    error
	)
	err = yaml.Unmarshal([]byte(resource.Content), &object)

	resourceId := schema.GroupVersionResource{
		Group:    cephGroupResource,
		Version:  cephGroupVersion,
		Resource: resource.Object,
	}

	cephResource, err := k.dynamicClientSet.Resource(resourceId).Namespace(resource.Namespace).Get(ctx, resource.Name, v1.GetOptions{})
	if err != nil {
		return err
	}
	object.SetResourceVersion(cephResource.GetResourceVersion())

	_, err = k.dynamicClientSet.Resource(resourceId).Namespace(resource.Namespace).Update(ctx, object, v1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}
