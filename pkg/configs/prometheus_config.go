package configs

import "github.com/prometheus/client_golang/prometheus"

// NewMetricsDefinition creates a new metrics definition
func NewMetricsDefinition() map[string]interface{} {
	return map[string]interface{}{
		"orders": prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "stickerfy_add_orders_total",
				Help: "Total number of processed orders",
			},
			[]string{"id"},
		),
		"totalRequests": prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Number of get requests.",
			},
			[]string{"path"},
		),
		"responseStatus": prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "response_status",
				Help: "Status of HTTP response.",
			},
			[]string{"status"},
		),
		"httpDuration": prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "http_response_time_seconds",
				Help: "Duration of HTTP requests.",
			},
			[]string{"path"},
		),
	}
}