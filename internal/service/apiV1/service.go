package apiV1

import (
	"codeZone/internal/converter"
	"codeZone/internal/models"
	"codeZone/internal/repository"
	"codeZone/internal/service"
	"context"
	"github.com/pkg/errors"
)

type srv struct {
	dockerRepository repository.DockerRepository
}

// NewApiV1Service return new api v1 service
func NewApiV1Service(dockerRepository repository.DockerRepository) service.ApiV1Service {
	return &srv{
		dockerRepository: dockerRepository,
	}
}

// Home returns example response
func (s *srv) Home() (*models.HomeV1Response, error) {
	jsonPayload := &models.HomeV1Response{
		Version:     "v1",
		Description: "CodeZone apiV1 server",
	}

	return jsonPayload, nil
}

// Run runs given code
func (s *srv) Run(ctx context.Context, data *models.RunV1Request) (*models.JSONResponse, error) {
	id, err := s.dockerRepository.RunContainer(ctx, converter.InfoFromServiceToDocker(data))
	if err != nil {
		return &models.JSONResponse{
			Error:   true,
			Message: err.Error(),
		}, err
	}

	return &models.JSONResponse{
		Error:   false,
		Message: id,
	}, nil
}

func (s *srv) Check(ctx context.Context, id string) (*models.CheckV1Response, error) {
	result, err := s.dockerRepository.CheckContainerLogs(ctx, id)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return &models.CheckV1Response{
				Status: "pending",
			}, nil
		}

		return nil, err
	}

	return &models.CheckV1Response{
		Status: "success",
		Result: result,
	}, nil
}
