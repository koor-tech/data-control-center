package k8s

import (
	"context"
	"sync"
	"time"

	koorv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/koor/v1"
	statsv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/stats/v1"
	"github.com/koor-tech/data-control-center/pkg/cache"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/k8s"
	"go.uber.org/fx"
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

	clusterDeployments cache.CacheEntry[[]*statsv1.ResourceInfo]
	cephResources      cache.CacheEntry[[]*statsv1.ResourceInfo]
	storageNodes       cache.CacheEntry[[]*statsv1.NodeInfo]
	koorCluster        cache.CacheEntry[*koorv1.KoorCluster]
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
		OnStart: func(context.Context) error {
			c.run()

			c.wg.Add(1)
			go func() {
				defer c.wg.Done()
				for {
					select {
					case <-time.After(3 * time.Second):
						c.run()

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

func (c *Cache) run() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		clusterDeployments, err := c.k8s.GetClusterDeployments(c.ctx, c.namespace)
		if err != nil {
			c.logger.Error("failed to update cluster deployments cache", zap.Error(err))
			return
		}
		c.clusterDeployments.Set(clusterDeployments)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		nodes, err := c.k8s.GetStorageNodes(c.ctx, c.namespace)
		if err != nil {
			c.logger.Error("failed to update storage nodes cache", zap.Error(err))
			return
		}
		c.storageNodes.Set(nodes)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resources, err := c.k8s.GetCephResources(c.ctx, c.namespace)
		if err != nil {
			c.logger.Error("failed to update ceph resources cache", zap.Error(err))
			return
		}
		c.clusterDeployments.Set(resources)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cluster, err := c.k8s.GetKoorCluster(c.ctx, c.namespace)
		if err != nil {
			c.logger.Error("failed to update koor cluster cache", zap.Error(err))
			return
		}
		c.koorCluster.Set(cluster)
	}()

	wg.Wait()
}

func (c *Cache) GetClusterDeployments(namespace string) ([]*statsv1.ResourceInfo, bool) {
	return c.clusterDeployments.Get()
}

func (c *Cache) GetStorageNodes(namespace string) ([]*statsv1.NodeInfo, bool) {
	return c.storageNodes.Get()
}

func (c *Cache) GetCephResources(namespace string) ([]*statsv1.ResourceInfo, bool) {
	return c.cephResources.Get()
}

func (c *Cache) GetKoorCluster(namespace string) (*koorv1.KoorCluster, bool) {
	return c.koorCluster.Get()
}
