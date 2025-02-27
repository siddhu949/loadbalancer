package main

import (
	"fmt"
	"net/http"
	"os"

	"loadbalancer/internal/firewall"
	"loadbalancer/internal/proxy"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Initialize Rate Limiting (5 requests per second, burst of 10)
	rateLimiter := firewall.NewRateLimiter(5, 10)

	// Start Prometheus metrics server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		fmt.Println("Metrics server running on :9090/metrics")
		if err := http.ListenAndServe(":9090", nil); err != nil {
			fmt.Println("Error starting metrics server:", err)
			os.Exit(1)
		}
	}()

	// Load Balancer HTTP Server with security middlewares
	mux := http.NewServeMux()
	mux.HandleFunc("/", proxy.HandleRequest)

	// Apply Middleware: JWT Authentication → Firewall → Rate Limiting
	securedMux := firewall.JWTMiddleware(firewall.IPBlocker(rateLimiter.Limit(mux)))

	fmt.Println("Load Balancer listening on :8080")
	err := http.ListenAndServe(":8080", securedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
