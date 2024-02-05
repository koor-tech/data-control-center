package recommendermodules

import (
	"context"

	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
)

func init() {
	moduleFactories["pool_sizes"] = NewPoolSizes
}

func NewPoolSizes() (Module, error) {
	return &PoolSizes{}, nil
}

type PoolSizes struct {
	Module
}

func (m *PoolSizes) GetName() string {
	return "pool_sizes"
}

func (m *PoolSizes) Run(ctx context.Context, p *Params) ([]*cephv1.ClusterRecommendation, error) {
	p.K8S.GetCephResources(p.Namespace)

	// TODO iterate over every Ceph* object and make sure pool sizes are safe
	return nil, nil
}
