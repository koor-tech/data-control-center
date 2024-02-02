package recommender

import (
	cephcache "github.com/koor-tech/data-control-center/pkg/ceph/cache"
	k8scache "github.com/koor-tech/data-control-center/pkg/k8s/cache"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	K8S  *k8scache.Cache
	Ceph *cephcache.Cache
}

type Recommender struct {
	k8s  *k8scache.Cache
	ceph *cephcache.Cache
}

func New(p Params) *Recommender {
	return &Recommender{
		k8s:  p.K8S,
		ceph: p.Ceph,
	}
}
