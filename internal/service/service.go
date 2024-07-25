package service

import (
	"codeZone/internal/models"
	"context"
)

// ApiV1Service - service for api v1
type ApiV1Service interface {
	Home() (*models.HomeV1Response, error)
	Run(ctx context.Context, data *models.RunV1Request) (*models.JSONResponse, error)
	Check(ctx context.Context, id string) (*models.CheckV1Response, error)
}
