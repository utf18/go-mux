package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// SimpleGauge which will increment the incValue gauge by 1 every 5 seconds and promote it to /metrics
func SimpleGauge() {
	incValue := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "acme_company",
		Subsystem: "blob_storage",
		Name:      "inc_value",
		Help:      "just an increased number.",
	})
	// register incValue
	prometheus.MustRegister(incValue)
	// loop over the ticker and call Inc function
	go func() {
		for range time.Tick(time.Second * 5) {
			// increment incValue by 1 every 5 seconds
			incValue.Inc()
		}
	}()

}
