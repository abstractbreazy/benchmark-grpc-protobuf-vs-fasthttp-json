package fasthttp

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

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
		Code:    fasthttp.StatusOK,
		Book:    &d,
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