package metrics

// Metrics is an interface for metrics
type Metrics interface {
	InitMetrics(metrics map[string]interface{})
	IncrementCounter(string, ...string)
	DestroyMetrics()
}
