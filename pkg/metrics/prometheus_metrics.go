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
func NewPrometheusMetrics(mutex *sync.Mutex) Metrics {
	metricsDefinitions := configs.NewMetricsDefinition()
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
	metricsDefinitions["orders"].(*prometheus.CounterVec).WithLabelValues("")
	if err := prometheus.Register(metricsDefinitions["orders"].(*prometheus.CounterVec)); err != nil {
		panic(err)
	}
	metricsDefinitions["totalRequests"].(*prometheus.CounterVec).WithLabelValues("")
	if err := prometheus.Register(metricsDefinitions["totalRequests"].(*prometheus.CounterVec)); err != nil {
		panic(err)
	}
	metricsDefinitions["responseStatus"].(*prometheus.CounterVec).WithLabelValues("")
	if err := prometheus.Register(metricsDefinitions["responseStatus"].(*prometheus.CounterVec)); err != nil {
		panic(err)
	}
	metricsDefinitions["httpDuration"].(*prometheus.HistogramVec).WithLabelValues("")
	if err := prometheus.Register(metricsDefinitions["httpDuration"].(*prometheus.HistogramVec)); err != nil {
		panic(err)
	}
	p.mutex.Unlock()
}

// IncrementCounter increments the counter
func (p prometheusMetrics) IncrementCounter(key string, labels ...string) {
	p.metrics[key].(*prometheus.CounterVec).WithLabelValues(labels...).Inc()
}

// GetTimer returns a timer
func (p prometheusMetrics) GetTimer(key string, labels ...string) *prometheus.Timer {
	return prometheus.NewTimer(p.metrics[key].(*prometheus.HistogramVec).WithLabelValues(labels...))
}

// DestroyMetrics destroys the metrics
func (p prometheusMetrics) DestroyMetrics() {
	p.mutex.Lock()
	for _, v := range p.metrics {
		prometheus.Unregister(v.(prometheus.Collector))
	}
	p.mutex.Unlock()
}
