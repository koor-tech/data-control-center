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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	authconnect.AuthServiceHandler

	tm *auth.TokenMgr

	users []*config.User
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
		tm:    tm,
		users: users,
	}, nil
}

func (s *Server) RegisterService(g *gin.RouterGroup) {
	authPath, authHandler := authconnect.NewAuthServiceHandler(s)
	g.Any(authPath+"/*path", gin.WrapH(authHandler))
}

func (s *Server) Login(ctx context.Context, req *connect.Request[pbauth.LoginRequest]) (*connect.Response[pbauth.LoginResponse], error) {
	var user *config.User
	for _, u := range s.users {
		if req.Msg.Username == u.Username {
			user = u
			break
		}
	}

	if user == nil {
		return nil, status.Error(codes.InvalidArgument, "Wrong username or password")
	}

	// Password check logic
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Msg.Password)); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Wrong username or password")
	}

	claims := auth.BuildTokenClaimsFromAccount(1, "")
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

func (s *Server) Logout(context.Context, *connect.Request[pbauth.LogoutRequest]) (*connect.Response[pbauth.LogoutResponse], error) {
	return &connect.Response[pbauth.LogoutResponse]{
		Msg: &pbauth.LogoutResponse{
			Success: true,
		},
	}, nil
}
