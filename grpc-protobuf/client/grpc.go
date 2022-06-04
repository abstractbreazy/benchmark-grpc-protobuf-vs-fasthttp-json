package client

import (
	proto "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/grpc-protobuf/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGRPCCLient returns a new gRPC client.
func NewGRPCClient(listenAddr string) (proto.APIClient, error) {
	conn, err := grpc.Dial(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return proto.NewAPIClient(conn), err
}
