package oauth2

import (
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/server/oauth2/providers"
	"go.uber.org/fx"
	"golang.org/x/oauth2"
)

var ConfigModule = fx.Module("oauth2_config",
	fx.Provide(NewConfig),
)

func NewConfig(cfg *config.Config) map[string]providers.IProvider {
	oauthConfigs := make(map[string]providers.IProvider, len(cfg.OAuth2.Providers))
	for _, p := range cfg.OAuth2.Providers {
		cfg := &oauth2.Config{
			RedirectURL:  p.RedirectURL,
			ClientID:     p.ClientID,
			ClientSecret: p.ClientSecret,
			Scopes:       p.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:   p.Endpoints.AuthURL,
				TokenURL:  p.Endpoints.TokenURL,
				AuthStyle: oauth2.AuthStyleInParams,
			},
		}
		var provider providers.IProvider
		switch p.Type {
		default:
			provider = &providers.Generic{
				BaseProvider: providers.BaseProvider{
					ProviderConfig: p,
					Name:           p.Name,
					UserInfoURL:    p.Endpoints.UserInfoURL,
				},
			}
		}

		provider.SetOauthConfig(cfg)
		provider.SetMapping(p.Mapping)

		oauthConfigs[p.Name] = provider
	}

	return oauthConfigs
}
