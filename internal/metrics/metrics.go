package metrics

import (
	"log"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// Ensure metrics are registered only once
var initOnce sync.Once

// Define Prometheus Metrics
var (
	TotalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "loadbalancer_requests_total",
			Help: "Total number of requests handled by the load balancer",
		},
		[]string{"backend"},
	)

	ActiveConnections = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "loadbalancer_active_connections",
			Help: "Current active connections per backend",
		},
		[]string{"backend"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "loadbalancer_request_duration_seconds",
			Help:    "Histogram of request duration to backend servers",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"backend"},
	)
)

// InitMetrics ensures metrics are registered only once
func InitMetrics() {
	initOnce.Do(func() {
		log.Println("Initializing Prometheus metrics...")
		prometheus.MustRegister(TotalRequests, ActiveConnections, RequestDuration)
	})
}
