package config

import (
	"github.com/pkg/errors"
	"net"
	"os"
)

const (
	namespaceEnv      = "METRICS_NAMESPACE"
	appNameEnv        = "METRICS_APP_NAME"
	subsystemEnv      = "METRICS_SUBSYSTEM"
	prometheusHostEnv = "METRICS_PROMETHEUS_HOST"
	prometheusPortEnv = "METRICS_PROMETHEUS_PORT"
)

type MetricsConfig interface {
	Namespace() string
	AppName() string
	Subsystem() string
	PrometheusAddress() string
}

type metricsConfig struct {
	namespace string
	appName   string
	subsystem string

	prometheusHost string
	prometheusPort string
}

func NewMetricsConfig() (MetricsConfig, error) {
	namespace := os.Getenv(namespaceEnv)
	if len(namespace) == 0 {
		return nil, errors.New("metrics namespace is empty")
	}

	appName := os.Getenv(appNameEnv)
	if len(appName) == 0 {
		return nil, errors.New("metrics appName is empty")
	}

	subsystem := os.Getenv(subsystemEnv)
	if len(subsystem) == 0 {
		return nil, errors.New("metrics subsystem is empty")
	}

	prometheusHost := os.Getenv(prometheusHostEnv)
	if len(prometheusHost) == 0 {
		return nil, errors.New("metrics prometheusHost is empty")
	}

	prometheusPort := os.Getenv(prometheusPortEnv)
	if len(prometheusPort) == 0 {
		return nil, errors.New("metrics prometheusPort is empty")
	}

	return &metricsConfig{
		namespace:      namespace,
		appName:        appName,
		subsystem:      subsystem,
		prometheusHost: prometheusHost,
		prometheusPort: prometheusPort,
	}, nil
}

func (c *metricsConfig) Namespace() string {
	return c.namespace
}

func (c *metricsConfig) AppName() string {
	return c.appName
}

func (c *metricsConfig) Subsystem() string {
	return c.subsystem
}

func (c *metricsConfig) PrometheusAddress() string {
	return net.JoinHostPort(c.prometheusHost, c.prometheusPort)
}
