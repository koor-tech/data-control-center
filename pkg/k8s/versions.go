package k8s

import (
	"context"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8s) GetOperatorVersion(ctx context.Context, namespace string) (string, error) {
	deployment, err := k.client.AppsV1().Deployments(namespace).Get(ctx, "rook-ceph-operator", metav1.GetOptions{})
	if err != nil {
		return "UNKNOWN", err
	}
	if deployment == nil {
		return "UNKNOWN", errors.New("unable to get rook-ceph-operator deployment")
	}

	resInfo := k.ConvertDeploymentToResourceInfo(*deployment)

	version := "UNKNOWN"
	if resInfo.Version != nil {
		version = *resInfo.Version
	}

	return version, nil
}
