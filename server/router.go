package server

import (
	statsPb "github.com/koor-tech/data-control-center/gen/go/api/services/stats"
	"github.com/koor-tech/data-control-center/pkg/grpc/stats"
	"google.golang.org/grpc"
)

func RegisterGRPCHandler(s grpc.ServiceRegistrar) {
	statsPb.RegisterStatsServiceServer(s, stats.NewServer())
}
