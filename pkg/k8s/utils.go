package k8s

import "context"

func (k *K8s) HasKSD(ctx context.Context) (bool, error) {

	// TODO check if koor-operator is being used

	return false, nil
}
