package auth

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
	pbauth "github.com/koor-tech/data-control-center/gen/go/api/services/auth/v1"
	"github.com/koor-tech/data-control-center/gen/go/api/services/auth/v1/authv1connect"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	"github.com/koor-tech/data-control-center/pkg/server/oauth2/providers"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	authv1connect.AuthServiceHandler

	tm *auth.TokenMgr

	oauth2Enabled   bool
	oauth2Providers map[string]providers.IProvider
	users           []*config.User
}

func New(tm *auth.TokenMgr, cfg *config.Config, oauth2Providers map[string]providers.IProvider) (*Server, error) {
	// TODO find a better way to store the passwords in memory
	users := []*config.User{}
	for _, user := range cfg.Users {
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			return nil, fmt.Errorf("error when hasing given user passwords")
		}

		user.Password = string(hashed)
		users = append(users, user)
	}

	return &Server{
		tm:              tm,
		oauth2Enabled:   len(oauth2Providers) > 0,
		oauth2Providers: oauth2Providers,
		users:           users,
	}, nil
}

func (s *Server) RegisterService(g *gin.RouterGroup) {
	path, handler := authv1connect.NewAuthServiceHandler(s)
	g.Any(path+"/*path", gin.WrapH(handler))
}

func (s *Server) Login(ctx context.Context, req *connect.Request[pbauth.LoginRequest]) (*connect.Response[pbauth.LoginResponse], error) {
	if s.oauth2Enabled {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("Password Login disabled because OAuth2 login is enabled. Please use OAuth2 for logging in!"))
	}

	var user *config.User
	for _, u := range s.users {
		if req.Msg.Username == u.Username {
			user = u
			break
		}
	}

	if user == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("Wrong username or password"))
	}

	// Password check logic
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Msg.Password)); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("Wrong username or password"))
	}

	claims := auth.BuildTokenClaimsFromAccount("data-control-center-user-login-"+user.Username, user.Username, "", "")
	token, err := s.tm.NewWithClaims(claims)
	if err != nil {
		return nil, err
	}

	return &connect.Response[pbauth.LoginResponse]{
		Msg: &pbauth.LoginResponse{
			Token:     token,
			Expires:   timestamppb.New(claims.ExpiresAt.Time),
			AccountId: claims.AccID,
		},
	}, nil
}

func (s *Server) Logout(ctx context.Context, req *connect.Request[pbauth.LogoutRequest]) (*connect.Response[pbauth.LogoutResponse], error) {
	var redirectTo *string
	if s.oauth2Enabled {
		t := req.Header().Get("authorization")

		if t == "" {
			return nil, fmt.Errorf("no auth token in logout request")
		}

		t = strings.TrimPrefix(t, "Bearer ")
		// Parse token only returns the token info when the token is still valid
		tInfo, err := s.tm.ParseWithClaims(t)
		if err != nil {
			return nil, fmt.Errorf("invalid user info retrieved from context")
		}

		if tInfo.Oauth2Provider == "" || tInfo.Oauth2Token == "" {
			return nil, fmt.Errorf("invalid oauth2 provider info retrieved from context")
		}

		if provider, ok := s.oauth2Providers[tInfo.Oauth2Provider]; ok {
			logoutURL := provider.GetLogoutURL(tInfo.Oauth2Token)
			if logoutURL != "" {
				parsed, err := url.Parse(logoutURL)
				if err != nil {
					return nil, fmt.Errorf("failed to parse logout url")
				}

				q := parsed.Query()
				q.Add("client_id", provider.GetOauthConfig().ClientID)
				parsed.RawQuery = q.Encode()
				logoutURL = parsed.String()

				redirectTo = &logoutURL
			}
		}
	}

	return &connect.Response[pbauth.LogoutResponse]{
		Msg: &pbauth.LogoutResponse{
			Success:    true,
			RedirectTo: redirectTo,
		},
	}, nil
}

func (s *Server) CheckToken(ctx context.Context, req *connect.Request[pbauth.CheckTokenRequest]) (*connect.Response[pbauth.CheckTokenResponse], error) {
	resp := &connect.Response[pbauth.CheckTokenResponse]{
		Msg: &pbauth.CheckTokenResponse{
			Success: false,
		},
	}

	if _, err := s.tm.ParseWithClaims(req.Msg.Token); err != nil {
		return resp, nil
	}

	resp.Msg.Success = true
	return resp, nil
}
