package repository

import (
	"codeZone/internal/repository/docker/models"
	"context"
)

type DockerRepository interface {
	DownloadImages(ctx context.Context)
	RunContainer(ctx context.Context, info *models.Info) (string, error)
	CheckContainerLogs(ctx context.Context, id string) (string, error)
}
