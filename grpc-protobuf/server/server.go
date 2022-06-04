package server

import (
	"context"

	pb "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/grpc-protobuf/proto/gen"
)

var _ pb.APIServer = (*Server)(nil)

type Server struct {
	pb.UnimplementedAPIServer
}

// New Server instance
func New() *Server {
	return &Server{}
}

// Parse handler.
func (s *Server) Parse(ctx context.Context, in *pb.Book) (*pb.Response, error) {
	return &pb.Response{
		Message: "OK",
		Code:    200,
		Book:    in,
	}, nil

}
