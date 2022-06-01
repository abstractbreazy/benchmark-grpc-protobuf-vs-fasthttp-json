package fasthttp

import (
	"github.com/valyala/fasthttp"
)

type Core struct {
	//l   	net.Listener
	bind    string
	server 	*fasthttp.Server
}

func New(/*l net.Listener*/ bind string) *Core {

	server := new(fasthttp.Server)

	h := parseHandler
	server.Handler = h

	return &Core{
		//l:   l,
		bind: bind,
		server: server,
	}
}

func (c *Core) ListenAndServe() error {
	return c.server.ListenAndServe(c.bind)
	//return c.server.ListenAndServe(c.bind)
}

func (c *Core) Close() error {
	return c.server.Shutdown()
}


