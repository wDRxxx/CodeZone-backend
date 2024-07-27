package app

import (
	apiServer "codeZone/internal/api"
	"codeZone/internal/closer"
	"codeZone/internal/config"
	"codeZone/internal/repository"
	"codeZone/internal/repository/docker"
	"codeZone/internal/service"
	"codeZone/internal/service/apiV1"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"log"
)

// serviceProvider contains all necessary for api to work correctly
type serviceProvider struct {
	httpConfig    config.HTTPConfig
	metricsConfig config.MetricsConfig

	apiServer apiServer.ApiServer

	apiService service.ApiV1Service

	dockerRepository repository.DockerRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// HTTPConfig creates if not created and returns HTTPConfig
func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("error creating http config: %v\n", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

// MetricsConfig creates if not created and returns MetricsConfig
func (s *serviceProvider) MetricsConfig() config.MetricsConfig {
	if s.metricsConfig == nil {
		cfg, err := config.NewMetricsConfig()
		if err != nil {
			log.Fatalf("error creating metrics config: %v\n", err)
		}

		s.metricsConfig = cfg
	}

	return s.metricsConfig
}

// ApiServer creates if not created and returns ApiServer
func (s *serviceProvider) ApiServer(ctx context.Context) apiServer.ApiServer {
	if s.apiServer == nil {
		s.apiServer = apiServer.NewApiServer(s.ApiV1Service(ctx), s.HTTPConfig().Origins())
	}

	return s.apiServer
}

// ApiV1Service creates if not created and returns ApiV1Service
func (s *serviceProvider) ApiV1Service(ctx context.Context) service.ApiV1Service {
	if s.apiService == nil {
		s.apiService = apiV1.NewApiV1Service(s.DockerRepository(ctx))
	}

	return s.apiService
}

// DockerRepository creates if not created and returns DockerRepository
func (s *serviceProvider) DockerRepository(ctx context.Context) repository.DockerRepository {
	if s.dockerRepository == nil {
		cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatalf("error creating docker client: %v\n", err)
		}

		s.dockerRepository = docker.NewRepository(cli)
		s.dockerRepository.DownloadImages(ctx)
		closer.Add(cli.Close)
	}

	return s.dockerRepository
}
