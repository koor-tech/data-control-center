package server

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/koor-tech/data-control-center/gen/go/proto/services/auth/authconnect"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	"github.com/koor-tech/data-control-center/pkg/server/api"
	"github.com/koor-tech/data-control-center/pkg/server/oauth2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"go.uber.org/zap"

	// Services
	serverauth "github.com/koor-tech/data-control-center/server/auth"
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

    // TODO use fx framework
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
	e.Use(sessions.SessionsMany([]string{"datacontrolcenter_oauth2_state", "datacontrolcenter_token"}, sessStore))

	// GZIP
	e.Use(gzip.Gzip(gzip.DefaultCompression))

	// Prometheus Metrics endpoint
	e.GET("/metrics", gin.WrapH(promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer, promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
			// Opt into OpenMetrics e.g. to support exemplars
			EnableOpenMetrics: true,
		}),
	)))

	// Register app routes
	oauth := oauth2.New(p.Logger.Named("oauth"), p.TokenMgr, p.Config.OAuth2.Providers)
	rs := api.New(p.Logger, p.Config)
	rs.Register(e, oauth)

	fs := static.LocalFile(".output/public/", false)
	fileserver := http.FileServer(fs)
	fileserver = http.StripPrefix("/", fileserver)

	e.NoRoute(func(c *gin.Context) {
		requestPath := c.Request.URL.Path
		if strings.HasPrefix(requestPath, "/api") || requestPath == "/" {
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
    authSvc := &serverauth.Server{}
    authPath, authHandler := authconnect.NewAuthHandler(authSvc)
    e.Any(authPath, gin.WrapH(authHandler))

	return e
}

func wrapLogger(log *zap.Logger) *zap.Logger {
	return log.Named("http_server")
}
