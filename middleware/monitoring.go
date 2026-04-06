package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AnhCaooo/go-goods/monitoring"
	"github.com/prometheus/client_golang/prometheus"
)

// Prometheus Middleware
func Prometheus(serviceName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/metrics" {
				next.ServeHTTP(w, r)
				return
			}

			now := time.Now()
			monitoring.ActiveRequestsGauge.Inc()
			defer monitoring.ActiveRequestsGauge.Dec()

			recorder := &monitoring.StatusRecorder{
				ResponseWriter: w,
				StatusCode:     http.StatusOK,
			}

			next.ServeHTTP(recorder, r)

			method := r.Method
			path := r.URL.Path
			status := strconv.Itoa(recorder.StatusCode)

			monitoring.LatencyHistogram.With(prometheus.Labels{
				"service": serviceName,
				"method":  method,
				"path":    path,
				"status":  status,
			}).Observe(time.Since(now).Seconds())

			monitoring.HttpRequestCounter.With(prometheus.Labels{
				"service": serviceName,
				"status":  status,
				"path":    path,
				"method":  method,
			}).Inc()
		})
	}
}
