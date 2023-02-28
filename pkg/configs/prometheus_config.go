package configs

import (
	"github.com/prometheus/client_golang/prometheus"
)

// NewMetricsDefinition creates a new metrics definition
func NewMetricsDefinition() map[string]interface{} {
	return map[string]interface{}{
		"orderAdded": prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "stickerfy_orders_added_total",
				Help: "Total number of processed orders.",
			},
			[]string{"id"},
		),
		"orderFailed": prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "stickerfy_orders_failed_total",
				Help: "Total number of failed orders.",
			},
			[]string{"id"},
		),
	}
}
