package cluster

import (
	"context"
	"time"

	"connectrpc.com/connect"
	clusterpb "github.com/koor-tech/data-control-center/gen/go/api/services/cluster/v1"
	"github.com/koor-tech/data-control-center/gen/go/api/services/cluster/v1/clusterv1connect"
	"github.com/koor-tech/data-control-center/pkg/config"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth"
	k8scache "github.com/koor-tech/data-control-center/pkg/k8s/cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Server is used to implement stats services.
type Server struct {
	clusterv1connect.ClusterServiceHandler

	logger    *zap.Logger
	auth      *auth.GRPCAuth
	k         *k8scache.Cache
	Namespace string
}

type Params struct {
	fx.In

	Logger   *zap.Logger
	GrpcAuth *auth.GRPCAuth
	K8S      *k8scache.Cache
	Cfg      *config.Config
}

func New(p Params) (*Server, error) {
	return &Server{
		logger:    p.Logger,
		auth:      p.GrpcAuth,
		k:         p.K8S,
		Namespace: p.Cfg.Namespace,
	}, nil
}

func (s *Server) GetKoorCluster(ctx context.Context, req *connect.Request[clusterpb.GetKoorClusterRequest]) (*connect.Response[clusterpb.GetKoorClusterResponse], error) {
	kc, _ := s.k.GetKoorCluster(s.Namespace)

	res := connect.NewResponse(&clusterpb.GetKoorClusterResponse{
		KoorCluster: kc,
	})
	return res, nil
}

func (s *Server) GetTroubleshootReport(ctx context.Context, req *connect.Request[clusterpb.GetTroubleshootReportRequest]) (*connect.Response[clusterpb.GetTroubleshootReportResponse], error) {
	reportContent := []struct {
		Name    string
		Content string
	}{}

	// TODO

	report := ""
	if len(reportContent) == 0 {
		report = "Unable to collect any report data."
	} else {
		for _, v := range reportContent {
			report += "## " + v.Name + "\n"
			report += "```console\n"
			report += v.Content + "\n"
			report += "```\n\n"
		}

		now := time.Now()
		report += "_Generated using Koor Data-Control-Center on " + now.String() + "_."
	}

	res := connect.NewResponse(&clusterpb.GetTroubleshootReportResponse{
		Report: report,
	})
	return res, nil
}
