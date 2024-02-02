package k8scache

import (
	"context"
	"sync"
	"time"

	k8sv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/k8s/v1"
	"github.com/koor-tech/data-control-center/pkg/cache"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/k8s"
	"go.uber.org/fx"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

var Module = fx.Module("k8s_cache",
	fx.Provide(
		New,
	),
	fx.Decorate(wrapLogger),
)

func wrapLogger(log *zap.Logger) *zap.Logger {
	return log.Named("k8s_cache")
}

type CacheEntry struct {
	Data      any
	UpdatedAt time.Time
	State     bool
}

type Cache struct {
	ctx       context.Context
	wg        sync.WaitGroup
	logger    *zap.Logger
	k8s       *k8s.K8s
	namespace string

	clusterDeployments cache.CacheEntry[[]*k8sv1.ResourceInfo]
	cephResources      cache.CacheEntry[[]*k8sv1.ResourceInfo]
	storageNodes       cache.CacheEntry[[]*k8sv1.NodeInfo]
}

type Params struct {
	fx.In

	LC     fx.Lifecycle
	Logger *zap.Logger

	K8S *k8s.K8s
	Cfg *config.Config
}

func New(p Params) *Cache {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Cache{
		ctx:       ctx,
		wg:        sync.WaitGroup{},
		logger:    p.Logger,
		k8s:       p.K8S,
		namespace: p.Cfg.Namespace,
	}

	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := c.run(ctx); err != nil {
				return err
			}

			c.wg.Add(1)
			go func() {
				defer c.wg.Done()
				for {
					select {
					case <-time.After(3 * time.Second):
						func() {
							ctx, cancel := context.WithTimeout(c.ctx, 15*time.Second)
							defer cancel()

							c.run(ctx)
						}()

					case <-c.ctx.Done():
						return
					}

				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			cancel()

			c.wg.Wait()

			return nil
		},
	})

	return c
}

func (c *Cache) run(ctx context.Context) error {
	wg := sync.WaitGroup{}

	errs := multierr.Combine()

	wg.Add(1)
	go func() {
		defer wg.Done()
		clusterDeployments, err := c.k8s.GetClusterDeployments(ctx, c.namespace)
		if err != nil {
			c.logger.Error("failed to update cluster deployments cache", zap.Error(err))
			errs = multierr.Append(errs, err)
			return
		}
		c.clusterDeployments.Set(clusterDeployments)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		nodes, err := c.k8s.GetStorageNodes(ctx, c.namespace)
		if err != nil {
			c.logger.Error("failed to update storage nodes cache", zap.Error(err))
			errs = multierr.Append(errs, err)
			return
		}
		c.storageNodes.Set(nodes)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resources, err := c.k8s.GetCephResources(ctx, c.namespace)
		if err != nil {
			c.logger.Error("failed to update ceph resources cache", zap.Error(err))
			errs = multierr.Append(errs, err)
			return
		}
		c.cephResources.Set(resources)
	}()

	wg.Wait()

	return errs
}

func (c *Cache) GetClusterDeployments(namespace string) ([]*k8sv1.ResourceInfo, bool) {
	return c.clusterDeployments.Get()
}

func (c *Cache) GetStorageNodes(namespace string) ([]*k8sv1.NodeInfo, bool) {
	return c.storageNodes.Get()
}

func (c *Cache) GetCephResources(namespace string) ([]*k8sv1.ResourceInfo, bool) {
	return c.cephResources.Get()
}
