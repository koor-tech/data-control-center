package stats

import (
	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
	"github.com/koor-tech/data-control-center/gen/go/api/services/stats/v1/statsv1connect"
	cephcache "github.com/koor-tech/data-control-center/pkg/ceph/cache"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	k8scache "github.com/koor-tech/data-control-center/pkg/k8s/cache"
	"github.com/koor-tech/data-control-center/pkg/recommender"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Logger      *zap.Logger
	GrpcAuth    *auth.GRPCAuth
	K8S         *k8scache.Cache
	Ceph        *cephcache.Cache
	Cfg         *config.Config
	Recommender *recommender.Recommender
}

// Server is used to implement stats services.
type Server struct {
	statsv1connect.StatsServiceHandler

	readOnly  bool
	logger    *zap.Logger
	auth      *auth.GRPCAuth
	ceph      *cephcache.Cache
	k         *k8scache.Cache
	Namespace string

	rec *recommender.Recommender
}

func New(p Params) (*Server, error) {
	return &Server{
		readOnly:  p.Cfg.ReadOnly,
		logger:    p.Logger,
		auth:      p.GrpcAuth,
		ceph:      p.Ceph,
		k:         p.K8S,
		Namespace: p.Cfg.Namespace,
		rec:       p.Recommender,
	}, nil
}

func (s *Server) RegisterService(g *gin.RouterGroup) {
	path, handler := statsv1connect.NewStatsServiceHandler(s, connect.WithInterceptors(
		s.auth.NewAuthInterceptor(),
	))
	g.Any(path+"/*path", gin.WrapH(handler))
}
