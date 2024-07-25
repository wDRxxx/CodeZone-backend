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
}

// NewApiServer creates new ApiServer
func NewApiServer(apiService service.ApiV1Service) ApiServer {
	serv := &server{
		apiService: apiService,
	}
	serv.setRoutes()

	return serv
}

// Handler returns api muxer
func (s *server) Handler() http.Handler {
	return s.mux
}
