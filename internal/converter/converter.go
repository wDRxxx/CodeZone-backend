package converter

import (
	"codeZone/internal/models"
	dockerModels "codeZone/internal/repository/docker/models"
)

// InfoFromServiceToDocker converts api v1 request to local docker model
func InfoFromServiceToDocker(req *models.RunV1Request) *dockerModels.Info {
	return &dockerModels.Info{
		Language: req.Language,
		Code:     req.Code,
	}
}
