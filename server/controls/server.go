package controls

import (
	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
	"github.com/koor-tech/data-control-center/gen/go/api/services/controls/v1/controlsv1connect"
	"github.com/koor-tech/data-control-center/pkg/ancientt"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	k8s "github.com/koor-tech/data-control-center/pkg/k8s"
	k8scache "github.com/koor-tech/data-control-center/pkg/k8s/cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Server is used to implement stats services.
type Server struct {
	controlsv1connect.OSDsServiceHandler

	readOnly  bool
	logger    *zap.Logger
	auth      *auth.GRPCAuth
	k         *k8s.K8s
	kc        *k8scache.Cache
	Namespace string

	ancientt *ancientt.Runner
}

type Params struct {
	fx.In

	LC fx.Lifecycle

	Logger   *zap.Logger
	GrpcAuth *auth.GRPCAuth
	K8S      *k8s.K8s
	K8SCache *k8scache.Cache
	Cfg      *config.Config

	Ancientt *ancientt.Runner
}

func New(p Params) (*Server, error) {
	return &Server{
		readOnly:  p.Cfg.ReadOnly,
		logger:    p.Logger,
		auth:      p.GrpcAuth,
		k:         p.K8S,
		kc:        p.K8SCache,
		Namespace: p.Cfg.Namespace,
		ancientt:  p.Ancientt,
	}, nil
}

func (s *Server) RegisterService(g *gin.RouterGroup) {
	path, handler := controlsv1connect.NewOSDsServiceHandler(s, connect.WithInterceptors(
		s.auth.NewAuthInterceptor(),
	))
	g.Any(path+"/*path", gin.WrapH(handler))
}
