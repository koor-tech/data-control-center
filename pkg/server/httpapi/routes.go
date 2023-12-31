package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/version"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("httpapi_routes", fx.Provide(
	New,
))

type Routes struct {
	logger *zap.Logger

	clientCfg *ClientConfig
}

func New(logger *zap.Logger, cfg *config.Config) *Routes {
	providers := make([]*ProviderConfig, len(cfg.OAuth2.Providers))

	for k, p := range cfg.OAuth2.Providers {
		providers[k] = &ProviderConfig{
			Name:  p.Name,
			Label: p.Label,
		}
	}

	return &Routes{
		logger: logger,

		clientCfg: &ClientConfig{
			Version: version.Version,
			Login: LoginConfig{
				Providers: providers,
			},
			ReadOnly:        cfg.ReadOnly,
			UpdateAvailable: nil,
		},
	}
}

func (r *Routes) Register(e *gin.Engine) {
	e.GET("/readiness", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	// API Base
	g := e.Group("/api")
	{
		g.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, PingResponse)
		})

		g.POST("/config", func(c *gin.Context) {
			c.JSON(http.StatusOK, r.clientCfg)
		})

		g.GET("/clear-site-data", func(c *gin.Context) {
			c.Header("Clear-Site-Data", "\"cache\", \"cookies\", \"storage\"")
			c.String(http.StatusOK, "Your local site data should be cleared now, please go back to the data-control-center homepage yourself.")
		})
	}
}

func (r *Routes) SetUpdateAvailable(update string) {
	r.clientCfg.UpdateAvailable = &update
}
