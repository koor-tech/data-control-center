package auth

import (
	"context"

	"connectrpc.com/connect"
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
	authconnect.AuthHandler

	tm *auth.TokenMgr

	users []*config.User
}

func New(tm *auth.TokenMgr, cfg *config.Config) *Server {
	// TODO find a better way to store the passwords in memory

	return &Server{
		tm:    tm,
		users: cfg.Users,
	}
}

func (s *Server) Login(ctx context.Context, req *connect.Request[pbauth.LoginRequest]) (*connect.Response[pbauth.LoginResponse], error) {
	var user *config.User
	for _, u := range s.users {
		if req.Msg.Username == u.Username {
			user = u
			break
		}
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
