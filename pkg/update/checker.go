package update

import (
	"context"
	"strings"
	"time"

	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/server/httpapi"
	"github.com/koor-tech/data-control-center/pkg/version"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("update_checker",
	fx.Provide(
		New,
	),
)

type Checker struct {
	ctx    context.Context
	logger *zap.Logger

	routes *httpapi.Routes

	interval time.Duration
}

type Params struct {
	fx.In

	LC fx.Lifecycle

	Logger *zap.Logger
	Cfg    *config.Config

	Routes *httpapi.Routes
}

func New(p Params) *Checker {
	ctx, cancel := context.WithCancel(context.Background())

	c := &Checker{
		ctx:    ctx,
		logger: p.Logger.Named("update_checker"),

		routes: p.Routes,

		interval: p.Cfg.UpdateCheck.Interval,
	}

	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				c.run(c.ctx)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			cancel()
			return nil
		},
	})

	return c
}

func (c *Checker) run(ctx context.Context) error {
	for {
		select {
		case <-c.ctx.Done():
			return nil

		case <-time.After(c.interval):
			if err := c.check(ctx); err != nil {
				c.logger.Error("failed to check github for latest version", zap.Error(err))
				continue
			}
		}
	}
}

func (c *Checker) check(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	latest, err := GetLatestRelease(ctx)
	if err != nil {
		return err
	}

	latest = strings.ReplaceAll(latest, "v", "")

	if latest == version.Version {
		// No new version available
		return nil
	}

	c.logger.Info("new version found", zap.String("current", version.Version), zap.String("new", latest))
	c.routes.SetUpdateAvailable(latest)

	return nil
}
