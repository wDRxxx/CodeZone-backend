package api

import (
	"github.com/rs/cors"
	"net/http"
)

func enableCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
	})

	return middleware.Handler(h)
}
