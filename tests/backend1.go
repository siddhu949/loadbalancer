package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define Prometheus Metrics
var (
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "backend_requests_total",
			Help: "Total number of requests received by backend",
		})
)

func main() {
	// Register Prometheus metrics
	prometheus.MustRegister(requestsTotal)

	// HTTP Handler for backend response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestsTotal.Inc()
		fmt.Fprintln(w, "Success (Response from Backend 1)")
	})

	// ✅ Expose Prometheus Metrics at /metrics
	http.Handle("/metrics", promhttp.Handler())

	// Start Backend Server
	port := "9001" // Change this for backend1 (9001) and backend3 (9003)
	fmt.Println("Backend running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
