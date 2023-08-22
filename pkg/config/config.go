package config

import (
	"fmt"

	"github.com/creasty/defaults"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Module("config",
	fx.Provide(
		Load,
	),
)

type Result struct {
	fx.Out

	Config *Config
}

func Load() (*Config, error) {
	// Viper Config reading setup
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/config")
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	c := &Config{}
	if err := defaults.Set(c); err != nil {
		return nil, fmt.Errorf("failed to set config defaults: %w", err)
	}

	if err := viper.Unmarshal(c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return c, nil
}

func LoadTest() (*Config, error) {
	c := &Config{}
	if err := defaults.Set(c); err != nil {
		return nil, fmt.Errorf("failed to set config defaults: %w", err)
	}

	return c, nil
}

type Config struct {
	LogLevel string `default:"DEBUG" yaml:"logLevel"`
	Mode     string `default:"debug" yaml:"mode"`

	HTTP   HTTP    `yaml:"http"`
	GRPC   GRPC    `yaml:"grpc"`
	JWT    JWT     `yaml:"jwt"`
	OAuth2 OAuth2  `yaml:"oauth2"`
	Users  []*User `yaml:"users"`
}

type HTTP struct {
	Listen   string   `default:":8080" yaml:"listen"`
	Sessions Sessions `yaml:"sessions"`
}

type Sessions struct {
	CookieSecret string `yaml:"cookieSecret"`
	Domain       string `default:"localhost" yaml:"domain"`
}

type GRPC struct {
	Listen string `default:":9090" yaml:"listen"`
}

type JWT struct {
	Secret string `yaml:"secret"`
}

type OAuth2 struct {
	Providers []*OAuth2Provider
}

type OAuth2ProviderType string

const (
	OAuth2ProviderGeneric OAuth2ProviderType = "generic"
	OAuth2ProviderDiscord OAuth2ProviderType = "discord"
)

type OAuth2Provider struct {
	Name         string             `yaml:"name"`
	Label        string             `yaml:"label"`
	Homepage     string             `yaml:"homepage"`
	Type         OAuth2ProviderType `yaml:"type"`
	RedirectURL  string             `yaml:"redirectURL"`
	ClientID     string             `yaml:"clientID"`
	ClientSecret string             `yaml:"clientSecret"`
	Scopes       []string           `yaml:"scopes"`
	Endpoints    OAuth2Endpoints    `yaml:"endpoints"`
	Mapping      *OAuth2Mapping     `yaml:"omitempty,mapping"`
}

type OAuth2Endpoints struct {
	AuthURL     string `yaml:"authURL"`
	TokenURL    string `yaml:"tokenURL"`
	UserInfoURL string `yaml:"userInfoURL"`
}

type OAuth2Mapping struct {
	ID       string `yaml:"id"`
	Username string `yaml:"username"`
}

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
