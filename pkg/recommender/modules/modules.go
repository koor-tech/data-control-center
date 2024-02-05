package recommendermodules

import (
	"context"
	"fmt"

	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	cephcache "github.com/koor-tech/data-control-center/pkg/ceph/cache"
	"github.com/koor-tech/data-control-center/pkg/k8s"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var ErrModuleNotFound = fmt.Errorf("module not found")

type Params struct {
	fx.In

	Logger    *zap.Logger
	Namespace string

	Ceph *cephcache.Cache
	K8S  *k8s.K8s
}

type NewModuleFn = func() (Module, error)

var moduleFactories = map[string]NewModuleFn{}

var modules = map[string]Module{}

type Module interface {
	Run(ctx context.Context, p *Params) ([]*cephv1.ClusterRecommendation, error)
}

func GetAllModules() (map[string]Module, error) {
	for name := range moduleFactories {
		if _, err := getModule(name); err != nil {
			return nil, err
		}
	}

	return modules, nil
}

func getModule(name string) (Module, error) {
	if _, ok := moduleFactories[name]; !ok {
		return nil, ErrModuleNotFound
	}

	if _, ok := modules[name]; !ok {
		module, err := moduleFactories[name]()
		if err != nil {
			return nil, err
		}

		modules[name] = module
	}

	return modules[name], nil
}
