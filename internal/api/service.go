package api

import (
	"codeZone/internal/service"
	"net/http"
)

// ApiServer should contain all handlers and routes
type ApiServer interface {
	Handler() http.Handler
}

type server struct {
	mux        http.Handler
	apiService service.ApiV1Service
	origins    []string
}

// NewApiServer creates new ApiServer
func NewApiServer(apiService service.ApiV1Service, origins []string) ApiServer {
	serv := &server{
		apiService: apiService,
		origins:    origins,
	}
	serv.setRoutes()

	return serv
}

// Handler returns api muxer
func (s *server) Handler() http.Handler {
	return s.mux
}
