package v1

import (
	"github.com/valyala/fasthttp"
)

// SetupRoutes registers API endpoints
func SetupRoutes(router *fasthttp.ServeMux) {
	router.GET("/api/v1/health", HealthCheck)
	router.GET("/api/v1/metrics", Metrics)
	router.GET("/api/v1/admin/stats", AdminStats)
}

// HealthCheck returns system health status
func HealthCheck(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte(`{"status": "ok"}`))
}

// Metrics exposes Prometheus metrics
func Metrics(ctx *fasthttp.RequestCtx) {
	ctx.Redirect("/metrics", fasthttp.StatusMovedPermanently)
}

// AdminStats provides load balancer statistics
func AdminStats(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte(`{"active_connections": 100, "total_requests": 5000}`))
}
