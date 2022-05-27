package main

import (
	"log"
	"time"

	"github.com/chiragkyal/go-metrics-app/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	//
	log.Println("Starting metrics")
	// start prometheus metric server
	err := metrics.EnableMetrics([]prometheus.Collector{metrics.Metrics}, true)
	if err != nil {
		log.Fatal("failed to enable prometheus metrics server")
	}

	for {
		log.Println("update metrics")
		metrics.Metrics.IncrementCounter(
			metrics.HelloTotalCount,
			prometheus.Labels{metrics.LabelsHello: "test"},
		)
		time.Sleep(5 * time.Second)
	}
}
