package ceph

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/koor-tech/data-control-center/pkg/ceph/client"
	"github.com/koor-tech/data-control-center/pkg/config"
	"go.uber.org/zap"
	"net/http"
)

type Service struct {
	logger *zap.Logger

	apiClient *client.Client
}

func NewService(logger *zap.Logger, config *config.Ceph) *Service {
	return &Service{
		logger:    logger,
		apiClient: client.NewClient(context.TODO(), config.Api),
	}
}

// GetClusterHealth returns the Health Status of the Ceph Cluster
// calling - GET /health/full endpoint
func (s *Service) GetClusterHealth(ctx context.Context) (*HealthStatus, error) {
	if err := s.apiClient.Auth(); err != nil {
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
		//log.Fatal(err)
		s.logger.Error("error decoding response", zap.Error(err))
		return nil, fmt.Errorf("error decoding response %w", err)
	}

	fmt.Println("-------------------------------")
	fmt.Printf("healthStatus: %+v", healthStatus)
	fmt.Println("-------------------------------")

	return &healthStatus, nil
}
