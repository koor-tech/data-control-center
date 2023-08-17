package providers

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/koor-tech/data-control-center/pkg/config"
	"golang.org/x/oauth2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type IProvider interface {
	SetOauthConfig(*oauth2.Config)
	SetMapping(*config.OAuth2Mapping)

	GetName() string

	GetRedirect(state string) string
	GetUserInfo(string) (*UserInfo, error)
}

type BaseProvider struct {
	oauthConfig *oauth2.Config
	mapping     *config.OAuth2Mapping

	UserInfoURL   string

	Name string
}

func (b *BaseProvider) SetOauthConfig(cfg *oauth2.Config) {
	b.oauthConfig = cfg
}

func (b *BaseProvider) SetMapping(mapping *config.OAuth2Mapping) {
	b.mapping = mapping
}

func (b *BaseProvider) GetName() string {
	return b.Name
}

func (b *BaseProvider) GetRedirect(state string) string {
	return b.oauthConfig.AuthCodeURL(state)
}

type UserInfo struct {
	ID       int64
	Username string
}
