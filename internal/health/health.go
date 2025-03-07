package health

import (
	"github.com/valyala/fasthttp"
	"log"
)

// CheckHealth responds with system health
func CheckHealth(ctx *fasthttp.RequestCtx) {
	log.Println("Health check triggered")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte(`{"status": "healthy"}`))
}
