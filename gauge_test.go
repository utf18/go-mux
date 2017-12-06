package main

import (
	"testing"
	"github.com/prometheus/client_golang/prometheus"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestExampleGauge(t *testing.T) {
	//given
	prometheusRegistry := prometheus.NewRegistry()
	//when
	go ExampleGauge(prometheusRegistry)

	//then
	<-time.After(time.Second * 5 + time.Millisecond * 20) // In this case, same as time.Sleep(..)
	mfs, _:= prometheusRegistry.Gather()
	assert.Equal(t,"acme_company_blob_storage_inc_value",*mfs[0].Name)
	assert.Equal(t,float64(1),*mfs[0].Metric[0].GetGauge().Value)
}

func TestExampleGaugeBeforeFirstIncrement(t *testing.T) {
	//given
	prometheusRegistry := prometheus.NewRegistry()
	//when
	go ExampleGauge(prometheusRegistry)

	//then
	<-time.After(time.Millisecond * 20) // In this case, same as time.Sleep(..)
	mfs, _:= prometheusRegistry.Gather()
	assert.Equal(t,"acme_company_blob_storage_inc_value",*mfs[0].Name)
	assert.Equal(t,float64(0),*mfs[0].Metric[0].GetGauge().Value)
}
