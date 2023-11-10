package ceph

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	v1 "github.com/koor-tech/data-control-center/gen/go/api/services/ceph/v1"
	"github.com/koor-tech/data-control-center/gen/go/api/services/ceph/v1/cephv1connect"
	"github.com/koor-tech/data-control-center/internal/ceph"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	cephv1connect.CephServiceHandler

	logger    *zap.Logger
	auth      *auth.GRPCAuth
	ceph      *ceph.Service
	Namespace string
}

type Params struct {
	fx.In

	Logger   *zap.Logger
	GrpcAuth *auth.GRPCAuth
	Ceph     *ceph.Service
	Cfg      *config.Config
}

func New(p Params) (*Server, error) {
	return &Server{
		logger:    p.Logger,
		auth:      p.GrpcAuth,
		ceph:      p.Ceph,
		Namespace: p.Cfg.Namespace,
	}, nil
}

func (s *Server) RegisterService(g *gin.RouterGroup) {
	path, handler := cephv1connect.NewCephServiceHandler(s, connect.WithInterceptors(
		s.auth.NewAuthInterceptor(),
	))
	g.Any(path+"/*path", gin.WrapH(handler))
}

func (s *Server) GetCephUsers(ctx context.Context, _ *connect.Request[v1.GetCephUsersRequest]) (*connect.Response[v1.GetCephUsersResponse], error) {
	cephUsers, _ := s.ceph.GetUsers(ctx)
	var users []*cephv1.User

	for _, cu := range cephUsers {
		user := cephv1.User{
			Username: cu.Username,
			Enabled:  cu.Enabled,
			Roles:    cu.Roles,
		}

		users = append(users, &user)
	}

	return &connect.Response[v1.GetCephUsersResponse]{
		Msg: &v1.GetCephUsersResponse{
			CephUsers: users,
		},
	}, nil
}

func (s *Server) CreateCephUsers(ctx context.Context, req *connect.Request[v1.CreateCephUsersRequest]) (*connect.Response[v1.CreateCephUsersResponse], error) {
	userCreate := ceph.UserCreate{
		Username: req.Msg.CephUser.Username,
		Name:     req.Msg.CephUser.Name,
		Email:    req.Msg.CephUser.Email,
		Password: req.Msg.CephUser.Password,
		Enabled:  req.Msg.CephUser.Enabled,
		Roles:    req.Msg.CephUser.Roles,
	}
	err := s.ceph.CreateCephUser(ctx, userCreate)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error caused by %w", err))
	}

	user := cephv1.User{
		Username: userCreate.Username,
		Name:     userCreate.Name,
		Enabled:  userCreate.Enabled,
		Email:    userCreate.Email,
		Roles:    userCreate.Roles,
	}

	return &connect.Response[v1.CreateCephUsersResponse]{
		Msg: &v1.CreateCephUsersResponse{
			CephUser: &user,
		},
	}, nil
}

func (s *Server) DeleteCephUser(ctx context.Context, req *connect.Request[v1.DeleteCephUserRequest]) (*connect.Response[v1.DeleteCephUserResponse], error) {
	err := s.ceph.DeleteCephUser(ctx, req.Msg.Username)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error caused by %w", err))
	}
	return &connect.Response[v1.DeleteCephUserResponse]{
		Msg: &v1.DeleteCephUserResponse{
			Status: "success",
		},
	}, nil
}
