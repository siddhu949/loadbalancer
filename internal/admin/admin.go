package admin

import (
	"github.com/valyala/fasthttp"
	"log"
)

// AdminHandler for managing the load balancer
func AdminHandler(ctx *fasthttp.RequestCtx) {
	log.Println("Admin accessed")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Admin Dashboard"))
}
