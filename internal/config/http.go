package config

import (
	"github.com/pkg/errors"
	"net"
	"os"
)

const (
	httpHostEnv = "HTTP_HOST"
	httpPortEnv = "HTTP_PORT"
)

// HTTPConfig simple config interface
type HTTPConfig interface {
	Address() string
}

// httpConfig realisation of HTTPConfig
type httpConfig struct {
	host string
	port string
}

// NewHTTPConfig return new HTTPConfig
func NewHTTPConfig() (HTTPConfig, error) {
	host := os.Getenv(httpHostEnv)
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	port := os.Getenv(httpPortEnv)
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

// Address return address by config values
func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
