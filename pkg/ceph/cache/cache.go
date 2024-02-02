package cephcache

import (
	"context"
	"sync"
	"time"

	"github.com/koor-tech/data-control-center/pkg/cache"
	"github.com/koor-tech/data-control-center/pkg/ceph"
	"go.uber.org/fx"
	"go.uber.org/multierr"
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

	ceph *ceph.MgrService

	healthStatus cache.CacheEntry[*ceph.HealthStatus]
	rbdImages    cache.CacheEntry[[]ceph.BlockImage]
}

type Params struct {
	fx.In

	LC     fx.Lifecycle
	Logger *zap.Logger
	Ceph   *ceph.MgrService
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

		healthStatus, err := c.ceph.GetHealthFull(ctx)
		if err != nil {
			c.logger.Error("failed to update cluster deployments cache", zap.Error(err))
			errs = multierr.Append(errs, err)
			return
		}

		c.healthStatus.Set(healthStatus)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		blockImages, err := c.ceph.GetBlockImage(ctx)
		if err != nil {
			c.logger.Error("failed to update block image cache", zap.Error(err))
			errs = multierr.Append(errs, err)
			return
		}

		c.rbdImages.Set(blockImages)
	}()

	wg.Wait()

	return errs
}

func (c *Cache) GetHealthFull(_ context.Context) (*ceph.HealthStatus, bool) {
	return c.healthStatus.Get()
}

func (c *Cache) GetBlockImages(_ context.Context) ([]ceph.BlockImage, bool) {
	return c.rbdImages.Get()
}
