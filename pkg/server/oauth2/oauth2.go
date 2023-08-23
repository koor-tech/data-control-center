package oauth2

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	"github.com/koor-tech/data-control-center/pkg/server/oauth2/providers"
	"github.com/koor-tech/data-control-center/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type OAuth2 struct {
	logger *zap.Logger
	db     *sql.DB
	tm     *auth.TokenMgr

	oauthConfigs map[string]providers.IProvider
}

func New(logger *zap.Logger, tm *auth.TokenMgr, oAuth2Providers []*config.OAuth2Provider) *OAuth2 {
	o := &OAuth2{
		logger:       logger,
		tm:           tm,
		oauthConfigs: make(map[string]providers.IProvider, len(oAuth2Providers)),
	}

	for _, p := range oAuth2Providers {
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
					Name:        p.Name,
					UserInfoURL: p.Endpoints.UserInfoURL,
				},
			}
		}

		provider.SetOauthConfig(cfg)
		provider.SetMapping(p.Mapping)

		o.oauthConfigs[p.Name] = provider
	}

	return o
}

func (o *OAuth2) GetProvider(c *gin.Context) (providers.IProvider, error) {
	param, ok := c.Params.Get("provider")
	if !ok {
		return nil, fmt.Errorf("no provider found")
	}

	provider, ok := o.oauthConfigs[param]
	if !ok {
		return nil, fmt.Errorf("no provider found")
	}

	return provider, nil
}

const (
	AccountInfoRedirBase string = "/auth/account-info"
	LoginRedirBase       string = "/auth/login"
)

func (o *OAuth2) handleRedirect(c *gin.Context, err error, connectOnly bool, success bool, reason string) {
	if !success {
		redirURL := ""
		if connectOnly {
			redirURL = AccountInfoRedirBase + "?oauth2Connect=failed"
		} else {
			redirURL = LoginRedirBase + "?oauth2Login=failed"
		}

		if reason != "" {
			redirURL = redirURL + "&reason=" + url.QueryEscape(reason)
		}

		c.Redirect(http.StatusTemporaryRedirect, redirURL)
		return
	}

	redirURL := ""
	if connectOnly {
		redirURL = AccountInfoRedirBase + "?oauth2Connect=success"
	} else {
		redirURL = LoginRedirBase + "?oauth2Login=success"
	}

	c.Redirect(http.StatusTemporaryRedirect, redirURL)
}

func (o *OAuth2) Login(c *gin.Context) {
	sess := sessions.DefaultMany(c, "datacontrolcenter_oauth2_state")
	connectOnly := false
	connectOnlyVal := c.Query("connect-only")
	if connectOnlyVal != "" {
		var err error
		connectOnly, err = strconv.ParseBool(connectOnlyVal)
		if err != nil {
			o.logger.Error("failed to parse connect only var", zap.Error(err))
			o.handleRedirect(c, err, false, false, "invalid_request")
			return
		}
	}

	tokenVal := c.Query("token")
	if tokenVal != "" {
		sess.Set("token", tokenVal)
	}

	state, err := utils.GenerateRandomString(64)
	if err != nil {
		o.handleRedirect(c, err, connectOnly, false, "internal_error")
		return
	}

	sess.Set("connect-only", connectOnly)
	sess.Set("state", state)
	sess.Save()

	provider, err := o.GetProvider(c)
	if err != nil {
		o.logger.Error("failed to get provider", zap.Error(err))
		o.handleRedirect(c, err, connectOnly, false, "invalid_provider")
		return
	}

	// Redirect to provider for login
	c.Redirect(http.StatusTemporaryRedirect, provider.GetRedirect(state))
}

func (o *OAuth2) Callback(c *gin.Context) {
	sess := sessions.DefaultMany(c, "datacontrolcenter_oauth2_state")
	sessState := sess.Get("state")
	if sessState == nil {
		o.handleRedirect(c, nil, false, false, "invalid_state")
		return
	}

	state := sessState.(string)
	connectOnly := false
	sessConnectOnly := sess.Get("connect-only")
	if sessConnectOnly != nil {
		connectOnly = sessConnectOnly.(bool)
	}

	var token string
	sessToken := sess.Get("token")
	if sessToken != nil {
		token = sessToken.(string)
	}

	// Remove vars from session
	sess.Delete("connect-only")
	sess.Delete("state")
	sess.Delete("token")
	sess.Save()

	if c.Request.FormValue("state") != state {
		o.handleRedirect(c, nil, connectOnly, false, "invalid_state")
		return
	}

	provider, err := o.GetProvider(c)
	if err != nil {
		o.logger.Error("failed to get provider", zap.Error(err))
		o.handleRedirect(c, nil, connectOnly, false, "invalid_provider")
		return
	}

	userInfo, err := provider.GetUserInfo(c.Request.FormValue("code"))
	if err != nil {
		o.logger.Error("failed to get userinfo from provider", zap.Error(err))
		o.handleRedirect(c, err, connectOnly, false, "provider_failed")
		return
	}

	if connectOnly {
		_, err := o.tm.ParseWithClaims(token)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, LoginRedirBase+"?oauth2Connect=failed&reason=token_invalid")
			return
		}

		o.handleRedirect(c, nil, connectOnly, true, "")
		return
	}

	if userInfo == nil {
		c.Redirect(http.StatusTemporaryRedirect, LoginRedirBase+"?oauth2Login=failed&reason=unconnected")
		return
	} else if userInfo.ID == 0 {
		o.handleRedirect(c, nil, connectOnly, true, "internal_error")
		return
	}

	claims := auth.BuildTokenClaimsFromAccount(uint64(userInfo.ID), userInfo.Username)
	newToken, err := o.tm.NewWithClaims(claims)
	if err != nil {
		o.handleRedirect(c, err, connectOnly, true, "internal_error")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf(LoginRedirBase+"?oauth2Login=success&t=%s&exp=%d",
		url.QueryEscape(newToken),
		claims.ExpiresAt.Time.UTC().UnixNano()/1e6,
	))
}
