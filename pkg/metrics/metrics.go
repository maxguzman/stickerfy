package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics is an interface for metrics
type Metrics interface {
	InitMetrics(metrics map[string]interface{})
	IncrementCounter(string, ...string)
	GetTimer(string, ...string) *prometheus.Timer
	DestroyMetrics()
}
