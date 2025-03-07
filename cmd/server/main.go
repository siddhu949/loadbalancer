package main

import (
	"log"

	"loadbalancer/internal/metrics"
	"loadbalancer/internal/proxy"

	"github.com/valyala/fasthttp"
)

func main() {
	// Initialize metrics
	metrics.InitMetrics()

	// Define FastHTTP router
	router := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/metrics":
			proxy.ServeMetrics(ctx) // ✅ Ensure /metrics is served properly
		default:
			proxy.HandleRequest(ctx)
		}
	}

	log.Println("Load Balancer listening on 0.0.0.0:8080")
	if err := fasthttp.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
