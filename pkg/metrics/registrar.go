package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

func RegisterCollectors(registrar prometheus.Registerer, promCollectors []prometheus.Collector, includeDefault bool) {
	if includeDefault {
		err := registrar.Register(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		if err != nil {
			log.Fatalf("Failed to register process collector: %s\n", err)
		}

		err = registrar.Register(collectors.NewGoCollector())
		if err != nil {
			log.Fatalf("Failed to register go collector: %s", err)
		}
	}

	if len(promCollectors) > 0 {
		for _, c := range promCollectors {
			err := registrar.Register(c)
			if err != nil {
				log.Fatalf("Failed to register collector:%s\n", err)
			}
		}
	}
}
