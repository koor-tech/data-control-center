package recommender

import (
	"context"
	"sync"
	"time"

	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	cephcache "github.com/koor-tech/data-control-center/pkg/ceph/cache"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/k8s"
	recommendermodules "github.com/koor-tech/data-control-center/pkg/recommender/modules"
	"go.uber.org/fx"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

var Module = fx.Module("recommender",
	fx.Provide(
		New,
	),
)

type Params struct {
	fx.In

	LC     fx.Lifecycle
	Logger *zap.Logger

	Ceph *cephcache.Cache
	K8S  *k8s.K8s
	Cfg  *config.Config
}

type Recommender struct {
	ctx    context.Context
	logger *zap.Logger

	ceph      *cephcache.Cache
	k8s       *k8s.K8s
	namespace string

	mutex           sync.Mutex
	recommendations []*cephv1.ClusterRecommendation
}

func New(p Params) *Recommender {
	ctx, cancel := context.WithCancel(context.Background())

	r := &Recommender{
		ctx:    ctx,
		logger: p.Logger,

		ceph:      p.Ceph,
		k8s:       p.K8S,
		namespace: p.Cfg.Namespace,
	}

	p.LC.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				for {
					if err := r.Gather(); err != nil {
						r.logger.Error("failed to gather cluster recommendations", zap.Error(err))
					}

					select {
					case <-ctx.Done():
						return
					case <-time.After(60 * time.Second):
					}
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			cancel()
			return nil
		},
	})

	return r
}

func (r *Recommender) Gather() error {
	modules, err := recommendermodules.GetAllModules()
	if err != nil {
		return err
	}

	p := &recommendermodules.Params{
		Logger:    r.logger,
		Namespace: r.namespace,
		K8S:       r.k8s,
		Ceph:      r.ceph,
	}

	errs := multierr.Combine()
	recommendations := []*cephv1.ClusterRecommendation{}

	for name, module := range modules {
		func() {
			ctx, cancel := context.WithTimeout(r.ctx, 20*time.Second)
			defer cancel()

			r.logger.Debug("running module", zap.String("module_name", name))
			rs, err := module.Run(ctx, p)
			if err != nil {
				errs = multierr.Append(errs, err)
			}

			r.logger.Debug("module run completed", zap.String("module_name", name), zap.Int("recommend_length", len(recommendations)))
			recommendations = append(recommendations, rs...)
		}()
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.recommendations = recommendations

	return errs
}

func (r *Recommender) GetRecommendations() []*cephv1.ClusterRecommendation {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.recommendations
}
