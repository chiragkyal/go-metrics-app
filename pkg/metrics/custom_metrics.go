package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	LabelsHello     = "LabelHello"
	HelloTotalCount = "HelloTotalCount"
)

var (
	counterMap = map[string]*prometheus.CounterVec{}

	helloTotalCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hello_total_count",
			Help: "Total number of hello messages",
		},
		[]string{LabelsHello},
	)
)

type CustomMetrics struct {
	HelloTotalCount *prometheus.CounterVec
}

var Metrics = CustomMetrics{
	HelloTotalCount: helloTotalCount,
}

// init is called on instantiation
func init() {
	counterMap = map[string]*prometheus.CounterVec{
		HelloTotalCount: Metrics.HelloTotalCount,
	}
}

//Implementation of the Collector Describe interface function
func (c CustomMetrics) Describe(descCh chan<- *prometheus.Desc) {
	c.HelloTotalCount.Describe(descCh)
}

//Implementation of the Collector Collect interface function
func (c CustomMetrics) Collect(collectCh chan<- prometheus.Metric) {
	c.HelloTotalCount.Collect(collectCh)
}

//A method to increment the requests received number publicly
func (c CustomMetrics) IncrementCounter(counterName string, labels prometheus.Labels) {
	counterMap[counterName].With(labels).Add(1)
}
