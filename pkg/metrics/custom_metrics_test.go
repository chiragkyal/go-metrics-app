package metrics_test

import (
	"github.com/chiragkyal/go-metrics-app/pkg/metrics"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

var _ = Describe("Metrics", func() {
	var (
		Metrics metrics.CustomMetrics
		reg     *prometheus.Registry
	)

	BeforeEach(func() {
		reg = prometheus.NewPedanticRegistry()
	})

	Describe("Describe", func() {
		BeforeEach(func() {
			Metrics = metrics.Metrics
		})
		It("Describe does not error when registered", func() {
			err := reg.Register(Metrics)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Collect", func() {
		BeforeEach(func() {
			Metrics = metrics.Metrics
		})
		It("Collect does not error", func() {
			err := reg.Register(Metrics)
			Expect(err).ToNot(HaveOccurred())
			// TODO test reg.Gather() here somehow
			// and avoid Describe("Observe" ...) test case failures
		})
	})

	Describe("IncrementCounter", func() {
		var (
			counterMap = map[string]*prometheus.CounterVec{
				metrics.HelloTotalCount: metrics.Metrics.HelloTotalCount,
			}
		)
		It("Increments the number of requests when hit", func() {
			// counters
			for k, v := range counterMap {
				metrics.Metrics.IncrementCounter(k, prometheus.Labels{
					metrics.LabelsHello: "testing",
				})
				Expect(testutil.ToFloat64(v.With(prometheus.Labels{
					metrics.LabelsHello: "testing",
				}))).To(Equal(float64(1)))
			}
		})
	})

})
