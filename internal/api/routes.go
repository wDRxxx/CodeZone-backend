package api

import (
	"github.com/go-chi/chi/v5"
)

func (s *server) setRoutes() {
	mux := chi.NewRouter()

	// middlewares
	mux.Use(enableCORS)

	// routes
	mux.Route("/api", func(mux chi.Router) {
		mux.Route("/v1", func(mux chi.Router) {
			mux.Get("/", s.home)
			mux.Post("/run", s.run)
			mux.Get("/check/{id}", s.check)
		})
	})

	s.mux = mux
}
