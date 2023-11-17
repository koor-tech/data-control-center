package cluster

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	cephv1 "github.com/koor-tech/data-control-center/gen/go/api/resources/ceph/v1"
	clusterpb "github.com/koor-tech/data-control-center/gen/go/api/services/cluster/v1"
	"github.com/koor-tech/data-control-center/gen/go/api/services/cluster/v1/clusterv1connect"
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
	clusterv1connect.ClusterServiceHandler

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
	path, handler := clusterv1connect.NewClusterServiceHandler(s, connect.WithInterceptors(
		s.auth.NewAuthInterceptor(),
	))
	g.Any(path+"/*path", gin.WrapH(handler))
}

func (s *Server) GetKoorCluster(_ context.Context, _ *connect.Request[clusterpb.GetKoorClusterRequest]) (*connect.Response[clusterpb.GetKoorClusterResponse], error) {
	kc, _ := s.kc.GetKoorCluster(s.Namespace)

	res := connect.NewResponse(&clusterpb.GetKoorClusterResponse{
		KoorCluster: kc,
	})
	return res, nil
}

func (s *Server) GetResources(ctx context.Context, _ *connect.Request[clusterpb.GetResourcesRequest]) (*connect.Response[clusterpb.GetResourcesResponse], error) {
	var resources []*cephv1.Resource

	cephResources, err := s.k.GetYamlCephResources(ctx, s.Namespace)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error caused by %w", err))
	}

	for _, resource := range cephResources {
		resources = append(resources, &cephv1.Resource{
			Name:      resource.Name,
			Namespace: resource.Namespace,
			Content:   resource.Content,
			Kind:      resource.Kind,
			Object:    resource.Object,
		})
	}

	res := connect.NewResponse(&clusterpb.GetResourcesResponse{
		Resources: &cephv1.Resources{
			Resources: resources,
		},
	})
	return res, nil
}

func (s *Server) SaveResources(ctx context.Context, req *connect.Request[clusterpb.SaveResourcesRequest]) (*connect.Response[clusterpb.SaveResourcesResponse], error) {

	requestResource := req.Msg.Resource
	resource := cephv1.Resource{
		Name:      requestResource.Name,
		Namespace: requestResource.Namespace,
		Content:   requestResource.Content,
		Kind:      requestResource.Kind,
		Object:    requestResource.Object,
	}

	if len(resource.Name) == 0 || len(resource.Content) == 0 {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error caused by empty file"))
	}

	err := s.k.SaveResource(ctx, &resource)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error caused by %w", err))
	}

	res := connect.NewResponse(&clusterpb.SaveResourcesResponse{
		Resource: &resource,
	})
	return res, nil
}
