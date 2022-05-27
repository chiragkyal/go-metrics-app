package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var promRegistry *prometheus.Registry

func init() {
	promRegistry = prometheus.NewRegistry()
}

func EnableMetrics(collectors []prometheus.Collector, includeDefault bool) error {
	RegisterCollectors(promRegistry, collectors, includeDefault)
	r := promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{})
	curPromSrv := &http.Server{
		Addr:    fmt.Sprintf(":%v", 9000),
		Handler: r,
	}

	go func() {
		err := curPromSrv.ListenAndServe()
		if err != nil {
			print("failed to listen for metrics")
		}
	}()

	return nil
}
