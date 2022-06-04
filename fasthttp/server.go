package fasthttp

import (
	"net"

	"github.com/valyala/fasthttp"
)

type Core struct {
	Addr   string
	server *fasthttp.Server
}

func New() *Core {
	server := new(fasthttp.Server)

	h := parseHandler
	server.Handler = h

	return &Core{
		server: server,
	}
}

func (c *Core) Serve(l net.Listener) error {
	return c.server.Serve(l)
}
