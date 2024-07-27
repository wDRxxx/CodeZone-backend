package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"strings"
)

type Metrics struct {
	requestCounter        prometheus.Counter
	runRequestCounter     *prometheus.CounterVec
	histogramResponseTime *prometheus.HistogramVec
}

var metrics *Metrics

// Init initializing metrics
func Init(namespace string, appName string, subsystem string) {
	metrics = &Metrics{
		requestCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      appName + "_requests_total",
				Help:      "Count of requests to server",
			},
		),
		runRequestCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      appName + "_run_requests_total",
				Help:      "Count of requests to /run route",
			},
			[]string{"language"},
		),
		histogramResponseTime: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      appName + "_histogram_response_time_seconds",
				Help:      "Response time",
				Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
			},
			[]string{"status", "uri"},
		),
	}
}

// IncRequestCounter increments count of requests
func IncRequestCounter() {
	metrics.requestCounter.Inc()
}

// IncRunRequestCounter increments run count of requests
func IncRunRequestCounter(language string) {
	metrics.runRequestCounter.WithLabelValues(language).Inc()
}

// HistogramResponseTimeObserve increments histogram response time
func HistogramResponseTimeObserve(status int, uri string, time float64) {
	idx := strings.Index(uri, "/check/")
	if idx > 0 {
		uri = uri[:idx+6]
	}

	metrics.histogramResponseTime.WithLabelValues(strconv.Itoa(status), uri).Observe(time)
}
