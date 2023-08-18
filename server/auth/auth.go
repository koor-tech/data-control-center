package auth

import (
	"context"

	"connectrpc.com/connect"
	"github.com/koor-tech/data-control-center/gen/go/api/services/auth"
	"github.com/koor-tech/data-control-center/gen/go/api/services/auth/authconnect"
)

type Server struct {
	authconnect.AuthHandler
}

func (s *Server) Login(context.Context, *connect.Request[auth.LoginRequest]) (*connect.Response[auth.LoginResponse], error) {

    // TODO

	return nil, nil
}
func (s *Server) Logout(context.Context, *connect.Request[auth.LogoutRequest]) (*connect.Response[auth.LogoutResponse], error) {

    // TODO

	return nil, nil
}
