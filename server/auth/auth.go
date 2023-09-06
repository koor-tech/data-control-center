package auth

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
	"github.com/koor-tech/data-control-center/gen/go/api/resources/timestamp"
	pbauth "github.com/koor-tech/data-control-center/gen/go/api/services/auth"
	"github.com/koor-tech/data-control-center/gen/go/api/services/auth/authconnect"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	authconnect.AuthServiceHandler

	tm *auth.TokenMgr

	oauth2Enabled bool
	users         []*config.User
}

func New(tm *auth.TokenMgr, cfg *config.Config) (*Server, error) {
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
		tm:            tm,
		oauth2Enabled: len(cfg.OAuth2.Providers) > 0,
		users:         users,
	}, nil
}

func (s *Server) RegisterService(g *gin.RouterGroup) {
	path, handler := authconnect.NewAuthServiceHandler(s)
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

	claims := auth.BuildTokenClaimsFromAccount("data-control-center-user-login-"+user.Username, user.Username)
	token, err := s.tm.NewWithClaims(claims)
	if err != nil {
		return nil, err
	}

	return &connect.Response[pbauth.LoginResponse]{
		Msg: &pbauth.LoginResponse{
			Token:     token,
			Expires:   timestamp.New(claims.ExpiresAt.Time),
			AccountId: claims.AccID,
		},
	}, nil
}

func (s *Server) Logout(ctx context.Context, req *connect.Request[pbauth.LogoutRequest]) (*connect.Response[pbauth.LogoutResponse], error) {
	return &connect.Response[pbauth.LogoutResponse]{
		Msg: &pbauth.LogoutResponse{
			Success: true,
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
