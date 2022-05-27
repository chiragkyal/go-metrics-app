package metrics_test

import (
	"github.com/chiragkyal/go-metrics-app/pkg/metrics"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/prometheus/client_golang/prometheus"
)

var _ = Describe("Metrics", func() {

	Describe("Metrics UT with fake Registrar", func() {

		var (
			registrar *spyRegistrar
		)

		Context("Default collectors register", func() {
			BeforeEach(func() {
				registrar = newSpyRegistrar()
				metrics.RegisterCollectors(registrar, []prometheus.Collector{
					&simpleCollector{},
					&simpleCollector{},
				}, true)
			})

			It("registers the default collectors with the registrar", func() {
				Expect(registrar.collectors).To(HaveLen(4))
			})
		})

		Context("Custom collectors register", func() {
			BeforeEach(func() {
				registrar = newSpyRegistrar()
				metrics.RegisterCollectors(registrar, []prometheus.Collector{
					&simpleCollector{},
					&simpleCollector{},
				}, false)
			})

			Describe("Registering custom collectors", func() {
				It("registers the custom collectors with the registrar", func() {
					Expect(registrar.collectors).To(HaveLen(2))
				})
			})
		})
	})

})

type spyRegistrar struct {
	prometheus.Registerer
	collectors []prometheus.Collector
}

func newSpyRegistrar() *spyRegistrar {
	return &spyRegistrar{}
}

func (s *spyRegistrar) Register(c prometheus.Collector) error {
	s.collectors = append(s.collectors, c)
	return nil
}

type simpleCollector struct{}

func (c *simpleCollector) Describe(chan<- *prometheus.Desc) {}
func (c *simpleCollector) Collect(chan<- prometheus.Metric) {}
