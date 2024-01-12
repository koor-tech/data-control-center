package k8s

import (
	"context"
	"strings"

	k8sv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/k8s/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectldescribe "k8s.io/kubectl/pkg/describe"
)

func (k *K8s) transformNodeIntoNodeInfo(node *corev1.Node) *k8sv1.NodeInfo {
	status := k8sv1.ResourceStatus_RESOURCE_STATUS_UNKNOWN
	for _, cond := range node.Status.Conditions {
		if cond.Type == corev1.NodeReady {
			if cond.Status == "True" {
				status = k8sv1.ResourceStatus_RESOURCE_STATUS_READY
			} else {
				status = k8sv1.ResourceStatus_RESOURCE_STATUS_NOT_READY
			}
			break
		}
	}

	var internalIP *string
	var externalIP *string
	for _, address := range node.Status.Addresses {
		switch address.Type {
		case corev1.NodeExternalIP:
			fallthrough
		case corev1.NodeExternalDNS:
			address := address.Address
			externalIP = &address

		case corev1.NodeInternalIP:
			fallthrough
		case corev1.NodeInternalDNS:
			address := address.Address
			internalIP = &address
		}
	}

	roles := []string{}
	for key := range node.ObjectMeta.Labels {
		if strings.HasPrefix(key, kubectldescribe.LabelNodeRolePrefix) {
			role, found := strings.CutPrefix(key, kubectldescribe.LabelNodeRolePrefix)
			if found {
				roles = append(roles, role)
			}
		}
	}

	return &k8sv1.NodeInfo{
		Name:       node.Name,
		Roles:      roles,
		InternalIp: internalIP,
		ExternalIp: externalIP,
		Status:     status,
	}
}

func (k *K8s) GetStorageNodes(ctx context.Context, namespace string) ([]*k8sv1.NodeInfo, error) {
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

	storageNodes := []*k8sv1.NodeInfo{}
	for nodeName := range nodes {
		node, err := k.client.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}

		storageNodes = append(storageNodes, k.transformNodeIntoNodeInfo(node))
	}

	return storageNodes, nil
}
