package k8s

import "github.com/koor-tech/data-control-center/gen/go/api/resources/stats"

func (k *K8s) GetClusterRadar() (*stats.ClusterRadar, error) {
	radar := &stats.ClusterRadar{}

	// TODO calculate the numbers

	return radar, nil
}
