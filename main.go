// this project is heavily inspired by the prometheus golang client exmaple project
// https://github.com/prometheus/client_golang/blob/master/examples/simple/main.go

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)
	// exposes /metrics endpoint with standard golang metrics used by prometheus
	r.Handle("/metrics", promhttp.Handler())
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
