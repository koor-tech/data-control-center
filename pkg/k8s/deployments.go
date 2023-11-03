package k8s

import (
	"context"

	statsv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/stats/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ImportantDeployPrefixes = []string{"rook-ceph-", "csi-", "extended-ceph-exporter"}
)

func (k *K8s) GetClusterDeployments(ctx context.Context, namespace string) ([]*statsv1.ResourceInfo, error) {
	deployments, err := k.client.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	res := []*statsv1.ResourceInfo{}
	for _, deployment := range deployments.Items {
		status := statsv1.ResourceStatus_RESOURCE_STATUS_UNKNOWN
		if deployment.Status.ReadyReplicas >= deployment.Status.Replicas {
			status = statsv1.ResourceStatus_RESOURCE_STATUS_READY
		} else {
			status = statsv1.ResourceStatus_RESOURCE_STATUS_NOT_READY
		}

		res = append(res, &statsv1.ResourceInfo{
			// Why do we set it manually? https://github.com/kubernetes/client-go/issues/541
			Apiversion: "apps/v1",
			Kind:       "Deployment",
			Namespace:  deployment.Namespace,
			Name:       deployment.Name,
			Status:     status,
		})
	}

	return res, nil
}
