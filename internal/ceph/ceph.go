package ceph

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/koor-tech/data-control-center/pkg/ceph/client"
	"github.com/koor-tech/data-control-center/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("ceph_api",
	fx.Provide(
		NewService,
	),
	fx.Decorate(wrapLogger),
)

func wrapLogger(log *zap.Logger) *zap.Logger {
	return log.Named("ceph_api")
}

type Service struct {
	logger *zap.Logger

	apiClient *client.Client
}

func NewService(logger *zap.Logger, config *config.Config) *Service {
	return &Service{
		logger:    logger,
		apiClient: client.NewClient(context.Background(), config.Ceph.API),
	}
}

// GetHealthFull returns the Health Status of the Ceph Cluster
// calling - GET /health/full endpoint
func (s *Service) GetHealthFull(ctx context.Context) (*HealthStatus, error) {
	if err := s.apiClient.Auth(ctx); err != nil {
		s.logger.Error(ErrorUnableToAuthenticate.Error(), zap.Error(err))
		return nil, ErrorUnableToAuthenticate
	}

	resp, err := s.apiClient.MakeRequest(ctx, http.MethodGet, "/health/full", nil)
	if err != nil {
		s.logger.Error(ErrorUnableToConnectWithApi.Error(), zap.Error(err))
		return nil, ErrorUnableToConnectWithApi
	}

	if resp.StatusCode != http.StatusOK {
		s.logger.Error(ErrorUnableToConnectWithApi.Error(), zap.Error(err))
		return nil, ErrorUnableToConnectWithApi
	}

	var healthStatus HealthStatus
	if err := json.NewDecoder(resp.Body).Decode(&healthStatus); err != nil {
		s.logger.Error("error decoding response", zap.Error(err))
		return nil, fmt.Errorf("error decoding response %w", err)
	}

	return &healthStatus, nil
}
