package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AnhCaooo/go-goods/monitoring"
	"github.com/prometheus/client_golang/prometheus"
)

// Prometheus Middleware
func Prometheus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		monitoring.ActiveRequestsGauge.Inc()
		// Wrap the ResponseWriter to capture the status code
		recorder := &monitoring.StatusRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		// Process the request
		next.ServeHTTP(recorder, r)
		monitoring.ActiveRequestsGauge.Dec()

		method := r.Method
		path := r.URL.Path // Path can be adjusted for aggregation (e.g., `/users/:id` → `/users/{id}`)
		status := strconv.Itoa(recorder.StatusCode)

		monitoring.LatencyHistogram.With(prometheus.Labels{
			"method": method, "path": path, "status": status,
		}).Observe(time.Since(now).Seconds())

		// Increment the counter
		monitoring.HttpRequestCounter.WithLabelValues(status, path, method).Inc()
	})
}
