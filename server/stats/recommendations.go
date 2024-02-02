package stats

import (
	"context"

	"connectrpc.com/connect"
	statspb "github.com/koor-tech/data-control-center/gen/go/api/services/stats/v1"
)

func (s *Server) ListClusterRecommendations(ctx context.Context, req *connect.Request[statspb.ListClusterRecommendationsRequest]) (*connect.Response[statspb.ListClusterRecommendationsResponse], error) {

	return nil, nil
}
