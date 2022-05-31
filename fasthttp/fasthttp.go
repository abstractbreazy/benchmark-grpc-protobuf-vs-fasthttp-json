package fasthttp

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type Core struct {
	bind 	string
	server 	*fasthttp.Server
}

func New(bind string) *Core {
	
	server := new(fasthttp.Server)

	h := parseHandler
	server.Handler = h
	
	return &Core{
		bind: bind,	
		server: server,
	}
}

func (c *Core) ListenAndServe() error {
	return c.server.ListenAndServe(c.bind)
}

func (c *Core) Close() error {
	return c.server.Shutdown()
}

// Parse handle json request
func parseHandler(ctx *fasthttp.RequestCtx) {
	req := ctx.PostBody()
	if req == nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "Empty request body, 400\n")
		return
	}

	var d Book
	if err := json.Unmarshal(req, &d); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "Bad request, 400\n")
		return
	}

	resp := Response{
		Message: "OK",
		Code: fasthttp.StatusOK,
		Book: &d,
	}
	b, err := json.Marshal(resp)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Error, internal server error, 500\n")
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.Write(b)
}





