package grpcprotobuf

import (
	"net"

	pb "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/grpc-protobuf/proto/gen"
	server "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/grpc-protobuf/server"

	"google.golang.org/grpc"
)

// a gRPC-Server for tests.
type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	gs := grpc.NewServer()
	srv := server.New()
	pb.RegisterAPIServer(gs, srv)
	ln, err := net.Listen("tcp", ":15000")
	if err != nil {
		return err
	}

	return gs.Serve(ln)
}
