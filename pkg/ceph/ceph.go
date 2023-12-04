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
		New,
	),
	fx.Decorate(wrapLogger),
)

func wrapLogger(log *zap.Logger) *zap.Logger {
	return log.Named("ceph_api")
}

type MgrService struct {
	logger *zap.Logger

	apiClient *client.Client
}

type Params struct {
	fx.In

	LC fx.Lifecycle

	Logger *zap.Logger
	Config *config.Config
}

func New(p Params) *MgrService {
	apiClient := client.NewClient(p.Config.Ceph.API)

	return &MgrService{
		logger:    p.Logger,
		apiClient: apiClient,
	}
}

// GetHealthFull returns the Health Status of the Ceph Cluster
// calling - GET /health/full endpoint
func (s *MgrService) GetHealthFull(ctx context.Context) (*HealthStatus, error) {
	resp, err := s.apiClient.MakeRequest(ctx, client.NewEndpointHealthFull(), true)
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

func (s *MgrService) GetBlockImage(ctx context.Context) ([]BlockImage, error) {
	resp, err := s.apiClient.MakeRequest(ctx, client.NewEndpointBlockImage(), true)
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

func (s *MgrService) GetUsers(ctx context.Context) ([]User, error) {
	resp, err := s.apiClient.MakeRequest(ctx, client.NewEndpointUsers(), true)
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

func (s *MgrService) CreateCephUser(ctx context.Context, user UserCreate) error {
	userPayloadBytes, _ := json.Marshal(user)
	resp, err := s.apiClient.MakeRequest(ctx, client.NewPostEndpointUsers(userPayloadBytes), true)
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

func (s *MgrService) DeleteCephUser(ctx context.Context, username string) error {
	resp, err := s.apiClient.MakeRequest(ctx, client.NewEndpointDeleteUser(username), true)
	if err != nil {
		s.logger.Error(ErrorUnableToSaveCephUser.Error(), zap.Error(err))
		return ErrorUnableToSaveCephUser
	}

	if resp.StatusCode == http.StatusNoContent {
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
