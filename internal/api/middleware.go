package api

import (
	"codeZone/internal/metrics"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"net/http"
	"time"
)

// enableCORS enable cors
func (s *server) enableCORS(h http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   s.origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
	}).Handler(h)
}

// enableMetrics enable default metrics
func (s *server) enableMetrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics.IncRequestCounter()

		start := time.Now()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMinor)
		h.ServeHTTP(ww, r)
		respTime := time.Since(start)

		metrics.HistogramResponseTimeObserve(ww.Status(), r.RequestURI, respTime.Seconds())
	})
}
