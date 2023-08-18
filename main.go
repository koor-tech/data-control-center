package main

import (
	"fmt"
	"net/http"

	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	"github.com/koor-tech/data-control-center/pkg/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		LoggerModule,
		config.Module,
		server.HTTPServerModule,
		auth.AuthModule,
		auth.TokenMgrModule,

		fx.Invoke(func(*http.Server) {}),
	).Run()
}

var LoggerModule = fx.Module("logger",
	fx.Provide(
		NewLogger,
	),
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	// Logger Setup
	loggerConfig := zap.NewProductionConfig()
	level, err := zapcore.ParseLevel(cfg.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level from config. %w", err)
	}
	loggerConfig.Level.SetLevel(level)

	logger, err := loggerConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to configure logger. %w", err)
	}

	return logger, nil
}
