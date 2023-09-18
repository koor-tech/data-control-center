package k8s

import (
	"context"
	"strings"

	"github.com/koor-tech/data-control-center/gen/go/api/resources/stats"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectldescribe "k8s.io/kubectl/pkg/describe"
)

func (k *K8s) GetK8SClusterNodes(ctx context.Context) ([]*stats.NodeInfo, error) {
	nodes, err := k.client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	res := []*stats.NodeInfo{}
	for _, node := range nodes.Items {
		res = append(res, k.TransformNodeIntoNodeInfo(&node))
	}

	return res, nil
}

func (k *K8s) TransformNodeIntoNodeInfo(node *v1.Node) *stats.NodeInfo {
	status := stats.ResourceStatus_RESOURCE_UNKNOWN
	if node.Status.Phase >= v1.NodePhase(v1.NodeReady) {
		status = stats.ResourceStatus_RESOURCE_READY
	} else {
		status = stats.ResourceStatus_RESOURCE_NOT_READY
	}

	internalIP := ""
	externalIP := ""
	for _, address := range node.Status.Addresses {
		switch address.Type {
		case v1.NodeExternalIP:
			fallthrough
		case v1.NodeExternalDNS:
			externalIP = address.Address

		case v1.NodeInternalIP:
			fallthrough
		case v1.NodeInternalDNS:
			internalIP = address.Address
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

	return &stats.NodeInfo{
		Name:       node.Name,
		Roles:      roles,
		InternalIp: internalIP,
		ExternalIp: externalIP,
		Status:     status,
	}
}

func (k *K8s) GetStorageNodes(ctx context.Context, namespace string) ([]*stats.NodeInfo, error) {
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

	storageNodes := []*stats.NodeInfo{}
	for nodeName := range nodes {
		node, err := k.client.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}

		storageNodes = append(storageNodes, k.TransformNodeIntoNodeInfo(node))
	}

	return storageNodes, nil
}
