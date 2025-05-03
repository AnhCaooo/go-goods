package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// HttpRequestCounter is a counter for HTTP requests
	HttpRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests received",
	}, []string{"status", "path", "method"})

	// Guage is a gauge for monitoring node usage
	// This is used to track the number of active nodes in the system
	Gauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "devbulls_gauge",
		Help: "Monitoring node usage",
	}, []string{"node", "namespace"})

	// ActiveRequestsGauge is a gauge for active requests
	ActiveRequestsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_active_requests",
			Help: "Number of active connections to the service",
		},
	)

	// LatencyHistogram is a histogram for request durations
	LatencyHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests",
		Buckets: []float64{0.1, 0.5, 1, 2.5, 5, 10},
	}, []string{"status", "path", "method"})
)

// Helper to capture HTTP status codes
type StatusRecorder struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader captures the status code
func (rec *StatusRecorder) WriteHeader(code int) {
	rec.StatusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

// PrometheusHandler returns a handler for Prometheus metrics
// This handler serves the metrics at the /metrics endpoint
func PrometheusHandler() *http.Handler {
	registry := prometheus.NewRegistry()
	registry.MustRegister(HttpRequestCounter)
	registry.MustRegister(Gauge)
	registry.MustRegister(ActiveRequestsGauge)
	registry.MustRegister(LatencyHistogram)

	handler := promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{},
	)
	return &handler
}
