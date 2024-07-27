package app

import (
	"codeZone/internal/api"
	"codeZone/internal/closer"
	"codeZone/internal/config"
	"codeZone/internal/metrics"
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sync"
)

type App struct {
	wg sync.WaitGroup

	serviceProvider *serviceProvider

	server           api.ApiServer
	httpServer       *http.Server
	prometheusServer *http.Server
}

// NewApp creates new App
func NewApp(ctx context.Context, configPath string) (*App, error) {
	app := &App{wg: sync.WaitGroup{}}
	err := config.Load(configPath)
	if err != nil {
		return nil, err
	}

	err = app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// Run runs app
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()

		log.Printf("http apiV1 is running on: %s", a.serviceProvider.HTTPConfig().Address())
		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("error running http apiV1: %v", err)
		}
	}()

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()

		log.Printf("prometheus server is running on: %s", a.serviceProvider.MetricsConfig().PrometheusAddress())
		err := a.runPrometheusServer()
		if err != nil {
			log.Fatalf("error running http apiV1: %v", err)
		}
	}()

	a.wg.Wait()

	return nil
}

// initDeps initializing app dependencies
func (a *App) initDeps(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()

	metrics.Init(
		a.serviceProvider.MetricsConfig().Namespace(),
		a.serviceProvider.MetricsConfig().AppName(),
		a.serviceProvider.MetricsConfig().Subsystem(),
	)
	a.initHTTPServer(ctx)
	a.initPrometheusServer()

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) {
	a.server = a.serviceProvider.ApiServer(ctx)

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: a.server.Handler(),
	}
}

func (a *App) initPrometheusServer() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:    a.serviceProvider.MetricsConfig().PrometheusAddress(),
		Handler: mux,
	}
}

func (a *App) runHTTPServer() error {
	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runPrometheusServer() error {
	err := a.prometheusServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
