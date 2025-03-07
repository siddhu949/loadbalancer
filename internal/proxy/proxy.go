package proxy

import (
	"bytes"
	"compress/gzip"
	"log"
	"sync"
	"time"

	"loadbalancer/internal/config"
	"loadbalancer/internal/logger"
	"loadbalancer/internal/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// Backend structure to track connections
type Backend struct {
	Address           string
	ActiveConnections int
	mu                sync.Mutex
}

// List of backend servers
var backends = []*Backend{
	{Address: "127.0.0.1:9001"},
	{Address: "127.0.0.1:9002"},
	{Address: "127.0.0.1:9003"},
}

func init() {
	config.LoadConfig()
	logger.InitLogger()
	metrics.InitMetrics() // Ensures metrics are only registered once
}

// Get the backend with the least active connections
func getLeastConnectionsBackend() *Backend {
	var selected *Backend
	for _, backend := range backends {
		backend.mu.Lock()
		if selected == nil || backend.ActiveConnections < selected.ActiveConnections {
			selected = backend
		}
		backend.mu.Unlock()
	}
	return selected
}

// Compress response with Gzip
func compressResponse(ctx *fasthttp.RequestCtx, body []byte) {
	ctx.Response.Header.Set("Content-Encoding", "gzip")
	ctx.Response.Header.Set("Content-Type", "application/json")

	var compressedBody bytes.Buffer
	gzipWriter := gzip.NewWriter(&compressedBody)

	_, err := gzipWriter.Write(body)
	gzipWriter.Close()

	if err != nil {
		log.Println("Gzip compression error:", err)
		ctx.SetBody(body) // Send original body if compression fails
	} else {
		ctx.SetBody(compressedBody.Bytes()) // Send compressed response
	}
}

var client = &fasthttp.Client{
	MaxConnsPerHost:     2000,            // Increase max connections
	ReadTimeout:         2 * time.Second, // Increase timeout for high load
	WriteTimeout:        2 * time.Second,
	MaxIdleConnDuration: 15 * time.Second, // Keep idle connections longer
	Dial:                fasthttp.DialDualStack,
}

// HandleRequest forwards client requests to the least loaded backend
func HandleRequest(ctx *fasthttp.RequestCtx) {
	backend := getLeastConnectionsBackend()
	if backend == nil {
		log.Println("No available backend servers")
		ctx.SetStatusCode(fasthttp.StatusServiceUnavailable)
		ctx.SetBody([]byte("No available backend servers\n"))
		return
	}

	backend.mu.Lock()
	backend.ActiveConnections++
	metrics.ActiveConnections.WithLabelValues(backend.Address).Set(float64(backend.ActiveConnections))
	backend.mu.Unlock()

	log.Println("Forwarding request to:", backend.Address)

	start := time.Now()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI("http://" + backend.Address)
	req.Header.SetMethodBytes(ctx.Method())
	req.SetBody(ctx.PostBody())

	err := client.Do(req, res)
	if err != nil {
		log.Println("Backend request failed:", err)
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		ctx.SetBody([]byte("Backend server unavailable\n"))
	} else {
		ctx.SetStatusCode(res.StatusCode())
		compressResponse(ctx, res.Body())

		duration := time.Since(start).Seconds()
		metrics.TotalRequests.WithLabelValues(backend.Address).Inc()
		metrics.RequestDuration.WithLabelValues(backend.Address).Observe(duration)
	}

	backend.mu.Lock()
	backend.ActiveConnections--
	metrics.ActiveConnections.WithLabelValues(backend.Address).Set(float64(backend.ActiveConnections))
	backend.mu.Unlock()
}

// ServeMetrics serves Prometheus metrics
func ServeMetrics(ctx *fasthttp.RequestCtx) {
	log.Println("Serving /metrics request")
	ctx.Response.Header.Set("Content-Encoding", "identity")
	ctx.Response.Header.Set("Content-Type", "text/plain")
	handler := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	handler(ctx)
}
