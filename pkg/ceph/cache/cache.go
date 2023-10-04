package cache

import (
	"context"
	"sync"
	"time"

	"github.com/koor-tech/data-control-center/internal/ceph"
	"github.com/koor-tech/data-control-center/pkg/cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("ceph_cache",
	fx.Provide(
		New,
	),
	fx.Decorate(wrapLogger),
)

func wrapLogger(log *zap.Logger) *zap.Logger {
	return log.Named("ceph_cache")
}

type Cache struct {
	ctx    context.Context
	wg     sync.WaitGroup
	logger *zap.Logger

	ceph *ceph.Service

	healthStatus cache.CacheEntry[*ceph.HealthStatus]
}

type Params struct {
	fx.In

	LC     fx.Lifecycle
	Logger *zap.Logger
	Ceph   *ceph.Service
}

func New(p Params) *Cache {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Cache{
		ctx:    ctx,
		wg:     sync.WaitGroup{},
		logger: p.Logger,

		ceph: p.Ceph,
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
		healthStatus, err := c.ceph.GetHealthFull(c.ctx)
		if err != nil {
			c.logger.Error("failed to update cluster deployments cache", zap.Error(err))
			return
		}
		c.healthStatus.Set(healthStatus)
	}()

	wg.Wait()
}

func (c *Cache) GetHealthFull(ctx context.Context) (*ceph.HealthStatus, bool) {
	return c.healthStatus.Get()
}
