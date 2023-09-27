package k8s

import (
	"context"
	"strings"

	statsv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/stats/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectldescribe "k8s.io/kubectl/pkg/describe"
)

func (k *K8s) GetK8SClusterNodes(ctx context.Context) ([]*statsv1.NodeInfo, error) {
	nodes, err := k.client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	res := []*statsv1.NodeInfo{}
	for _, node := range nodes.Items {
		res = append(res, k.TransformNodeIntoNodeInfo(&node))
	}

	return res, nil
}

func (k *K8s) TransformNodeIntoNodeInfo(node *v1.Node) *statsv1.NodeInfo {
	status := statsv1.ResourceStatus_RESOURCE_STATUS_UNKNOWN
	for _, cond := range node.Status.Conditions {
		if cond.Type == v1.NodeReady {
			if cond.Status == "True" {
				status = statsv1.ResourceStatus_RESOURCE_STATUS_READY
			} else {
				status = statsv1.ResourceStatus_RESOURCE_STATUS_NOT_READY
			}
			break
		}
	}

	var internalIP *string
	var externalIP *string
	for _, address := range node.Status.Addresses {
		switch address.Type {
		case v1.NodeExternalIP:
			fallthrough
		case v1.NodeExternalDNS:
			externalIP = &address.Address

		case v1.NodeInternalIP:
			fallthrough
		case v1.NodeInternalDNS:
			internalIP = &address.Address
		}
	}

	roles := []string{}
	for _, label := range node.Labels {
		if strings.HasPrefix(label, kubectldescribe.LabelNodeRolePrefix) {
			role, found := strings.CutPrefix(label, kubectldescribe.LabelNodeRolePrefix)
			if found {
				roles = append(roles, role)
			}
		}
	}

	return &statsv1.NodeInfo{
		Name:       node.Name,
		Roles:      roles,
		InternalIp: internalIP,
		ExternalIp: externalIP,
		Status:     status,
	}
}

func (k *K8s) GetStorageNodes(ctx context.Context, namespace string) ([]*statsv1.NodeInfo, error) {
	pods, err := k.client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: RookCephPodsLabel,
	})
	if err != nil {
		return nil, err
	}

	nodes := map[string]interface{}{}
	for _, pod := range pods.Items {
		if pod.Spec.NodeName != "" {
			if _, ok := nodes[pod.Spec.NodeName]; !ok {
				nodes[pod.Spec.NodeName] = nil
			}
		}
	}

	storageNodes := []*statsv1.NodeInfo{}
	for nodeName := range nodes {
		node, err := k.client.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}

		storageNodes = append(storageNodes, k.TransformNodeIntoNodeInfo(node))
	}

	return storageNodes, nil
}
