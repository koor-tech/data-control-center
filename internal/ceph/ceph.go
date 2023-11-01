package ceph

import (
	"context"
	"encoding/json"
	"errors"
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

	resp, err := s.apiClient.MakeRequest(ctx, client.NewEndpointHealthFull())
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

func (s *Service) GetBlockImage(ctx context.Context) ([]BlockImage, error) {
	if err := s.apiClient.Auth(ctx); err != nil {
		s.logger.Error(ErrorUnableToAuthenticate.Error(), zap.Error(err))
		return nil, ErrorUnableToAuthenticate
	}

	resp, err := s.apiClient.MakeRequest(ctx, client.NewEndpointBlockImage())
	if err != nil {
		s.logger.Error(ErrorUnableToConnectWithApi.Error(), zap.Error(err))
		return nil, ErrorUnableToConnectWithApi
	}

	if resp.StatusCode != http.StatusOK {
		s.logger.Error(ErrorUnableToConnectWithApi.Error(), zap.Error(err))
		return nil, ErrorUnableToConnectWithApi
	}

	var blockImage []BlockImage
	if err := json.NewDecoder(resp.Body).Decode(&blockImage); err != nil {
		s.logger.Error("error decoding response", zap.Error(err))
		return nil, fmt.Errorf("error decoding response %w", err)
	}
	return blockImage, nil
}

func (s *Service) GetUsers(ctx context.Context) ([]User, error) {
	if err := s.apiClient.Auth(ctx); err != nil {
		s.logger.Error(ErrorUnableToAuthenticate.Error(), zap.Error(err))
		return nil, ErrorUnableToAuthenticate
	}

	resp, err := s.apiClient.MakeRequest(ctx, client.NewEndpointUsers())
	if err != nil {
		s.logger.Error(ErrorUnableToConnectWithApi.Error(), zap.Error(err))
		return nil, ErrorUnableToConnectWithApi
	}

	if resp.StatusCode != http.StatusOK {
		s.logger.Error(ErrorUnableToConnectWithApi.Error(), zap.Error(err))
		return nil, ErrorUnableToConnectWithApi
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		s.logger.Error("error decoding response", zap.Error(err))
		return nil, fmt.Errorf("error decoding response %w", err)
	}

	return users, nil
}

var ErrorUnableToSaveCephUser = errors.New("unable to save ceph user")

func (s *Service) CreateCephUser(ctx context.Context, user UserCreate) error {
	if err := s.apiClient.Auth(ctx); err != nil {
		s.logger.Error(ErrorUnableToAuthenticate.Error(), zap.Error(err))
		return ErrorUnableToSaveCephUser
	}

	userPayloadBytes, _ := json.Marshal(user)
	resp, err := s.apiClient.MakeRequest(ctx, client.NewPostEndpointUsers(userPayloadBytes))
	if err != nil {
		s.logger.Error(ErrorUnableToSaveCephUser.Error(), zap.Error(err))
		return ErrorUnableToSaveCephUser
	}

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	var errorCephApi ErrorCephApi
	if err := json.NewDecoder(resp.Body).Decode(&errorCephApi); err != nil {
		s.logger.Error("error decoding ceph api request", zap.Error(err))
		return err
	}

	s.logger.Error("error making request to ceph api ", zap.Error(errorCephApi.Error()))

	return errorCephApi.Error()
}
