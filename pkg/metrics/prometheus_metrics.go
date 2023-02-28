package metrics

import (
	"stickerfy/pkg/configs"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// prometheusMetrics implements Metrics interface
type prometheusMetrics struct {
	metrics map[string]interface{}
	mutex   *sync.Mutex
}

// NewPrometheusMetrics creates a new prometheus metrics
func NewPrometheusMetrics() Metrics {
	metricsDefinitions := configs.NewMetricsDefinition()
	mutex := &sync.Mutex{}
	var metrics Metrics = prometheusMetrics{
		metrics: metricsDefinitions,
		mutex:   mutex,
	}
	metrics.InitMetrics(metricsDefinitions)
	return metrics
}

// InitMetrics initializes the metrics
func (p prometheusMetrics) InitMetrics(metricsDefinitions map[string]interface{}) {
	p.mutex.Lock()
	metricsDefinitions["orderAdded"].(*prometheus.CounterVec).WithLabelValues("")
	if err := prometheus.Register(metricsDefinitions["orderAdded"].(*prometheus.CounterVec)); err != nil {
		panic(err)
	}
	metricsDefinitions["orderFailed"].(*prometheus.CounterVec).WithLabelValues("")
	if err := prometheus.Register(metricsDefinitions["orderFailed"].(*prometheus.CounterVec)); err != nil {
		panic(err)
	}
	p.mutex.Unlock()
}

// IncrementCounter increments the counter
func (p prometheusMetrics) IncrementCounter(key string, labels ...string) {
	p.metrics[key].(*prometheus.CounterVec).WithLabelValues(labels...).Inc()
}

// DestroyMetrics destroys the metrics
func (p prometheusMetrics) DestroyMetrics() {
	p.mutex.Lock()
	for _, v := range p.metrics {
		prometheus.Unregister(v.(prometheus.Collector))
	}
	p.mutex.Unlock()
}
