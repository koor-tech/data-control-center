package auth

import (
	"github.com/koor-tech/data-control-center/pkg/config"
	"go.uber.org/fx"
)

var TokenMgrModule = fx.Module("tokenMgr",
	fx.Provide(
		NewTokenMgrFromConfig,
	),
)

func NewTokenMgrFromConfig(cfg *config.Config) *TokenMgr {
	return NewTokenMgr(cfg.JWT.Secret)
}

var AuthModule = fx.Module("auth",
	fx.Provide(
		NewGRPCAuth,
	),
)
