package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap/zapcore"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/koor-tech/data-control-center/internal/ceph"
	cephcache "github.com/koor-tech/data-control-center/pkg/ceph/cache"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	"github.com/koor-tech/data-control-center/pkg/k8s"
	k8scache "github.com/koor-tech/data-control-center/pkg/k8s/cache"
	"github.com/koor-tech/data-control-center/pkg/server/httpapi"
	"github.com/koor-tech/data-control-center/pkg/server/oauth2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"go.uber.org/zap"

	// Connect Services - Need to be added here
	serverauth "github.com/koor-tech/data-control-center/server/auth"
	serverceph "github.com/koor-tech/data-control-center/server/ceph"
	servercluster "github.com/koor-tech/data-control-center/server/cluster"
	serverstats "github.com/koor-tech/data-control-center/server/stats"
)

var HTTPServerModule = fx.Module("httpserver",
	fx.Provide(
		NewHTTP,
	),
	fx.Decorate(wrapLogger),
)

type ServerParams struct {
	fx.In

	LC fx.Lifecycle

	Logger   *zap.Logger
	Config   *config.Config
	TokenMgr *auth.TokenMgr

	Services []Service `group:"connectServices"`
}

type ServerResult struct {
	fx.Out

	Server *http.Server
}

// NewHTTP builds an HTTP server that will begin serving requests
// when the Fx application starts.
func NewHTTP(p ServerParams) (ServerResult, error) {
	// Create HTTP Server for graceful shutdown handling
	srv := &http.Server{
		Addr:    p.Config.HTTP.Listen,
		Handler: setupHTTPServer(p).Handler(),
	}

	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			p.Logger.Info("http server listening", zap.String("address", p.Config.HTTP.Listen))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return ServerResult{
		Server: srv,
	}, nil
}

func setupHTTPServer(p ServerParams) *gin.Engine {
	// Gin HTTP Server
	gin.SetMode(p.Config.Mode)
	e := gin.New()

	e.UseH2C = true

	// Add Zap Logger to Gin
	e.Use(ginzap.Ginzap(p.Logger, time.RFC3339, true))
	e.Use(ginzap.RecoveryWithZap(p.Logger, true))

	// Sessions
	sessStore := cookie.NewStore([]byte(p.Config.HTTP.Sessions.CookieSecret))
	sessStore.Options(sessions.Options{
		Domain:   p.Config.HTTP.Sessions.Domain,
		Path:     "/",
		MaxAge:   int((24 * time.Hour).Seconds()),
		HttpOnly: true,
		Secure:   true,
	})
	e.Use(sessions.SessionsMany([]string{"datacontrolcenter_oauth2_state"}, sessStore))

	// Prometheus Metrics endpoint
	e.GET("/metrics", gin.WrapH(promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer, promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
			// Opt into OpenMetrics e.g. to support exemplars
			EnableOpenMetrics: true,
		}),
	)))

	// Register HTTP API routes
	rs := httpapi.New(p.Logger, p.Config)
	rs.Register(e)

	if len(p.Config.OAuth2.Providers) > 0 {
		oauth := oauth2.New(p.Logger.Named("oauth"), p.TokenMgr, p.Config.OAuth2.Providers)
		oauth.Register(e)
	}

	// Setup nuxt generated files serving
	fs := static.LocalFile(".output/public/", false)
	fileserver := http.FileServer(fs)
	fileserver = http.StripPrefix("/", fileserver)

	e.NoRoute(func(c *gin.Context) {
		requestPath := c.Request.URL.Path
		if strings.HasPrefix(requestPath, "/api") {
			return
		}

		if strings.HasSuffix(requestPath, "/") || !strings.Contains(requestPath, ".") {
			c.Request.URL.Path = "/"
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}

		fileserver.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	})
	// Register output dir for assets and other static files
	e.Use(static.Serve("/", fs))

	// Register Connect services
	g := e.Group("/api", func(c *gin.Context) {
		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, "/api")
		c.Next()
	})

	for _, svc := range p.Services {
		svc.RegisterService(g)
	}

	return e
}

func wrapLogger(log *zap.Logger) *zap.Logger {
	return log.Named("http_server")
}

func StartHTTPServer() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),

		LoggerModule,
		config.Module,
		HTTPServerModule,
		auth.AuthModule,
		auth.TokenMgrModule,
		k8s.Module,
		k8scache.Module,
		ceph.Module,
		cephcache.Module,

		// Connect Services - Need to be added here
		fx.Provide(
			AsService(serverauth.New),
			AsService(servercluster.New),
			AsService(serverstats.New),
			AsService(serverceph.New),
		),

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
