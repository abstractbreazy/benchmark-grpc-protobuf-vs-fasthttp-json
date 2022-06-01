package benchmarks

import (
	"context"
	"testing"

	server "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/grpc-protobuf"
	pb "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/grpc-protobuf/proto/gen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

// ============================================================================================================ //				
// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz  											  					//
//                                                											  					//
// BenchmarkGRPCProtobuf									  										      		//
// BenchmarkGRPCProtobuf-8            10000                         											//
// BenchmarkGRPCProtobuf-8            10000              														//
// BenchmarkGRPCProtobuf-8            10000             														//
// BenchmarkGRPCProtobuf-8            10000             														//
// BenchmarkGRPCProtobuf-8            10000             														//
// ============================================================================================================ //

type client struct {
	conn *grpc.ClientConn
	api  pb.APIClient
}

func NewClient(bind string, b *testing.B) *client {
	
	var (
		c = new(client)
		err error
	)
	c.conn, err = grpc.Dial(bind, grpc.WithInsecure())
	require.NoError(b, err)

	c.api = pb.NewAPIClient(c.conn)

	return c
}

func (c *client) Close() error {
	return c.conn.Close()
}

func BenchmarkGRPCProtobuf(b *testing.B) {

	bind := "127.0.0.1:8080"
	var (
		grpc 	 server.GRPCServer
		
	)
	require.NoError(b, grpc.Start(bind))

	// defer func ()  {
	// 	err := grpc.Close()
	// 	require.NoError(b, err)
	// }()

	defer grpc.Close()

	client := NewClient(bind, b)

	defer client.Close()

	// defer func() {
	// 	err := client.Close()
	// 	require.NoError(b, err)
	// }()

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		doRequest(client, b)
	}
}

func doRequest(c *client, b *testing.B) {

	var (
		ctx = context.Background()
		err error
		resp *pb.Response
	)

	resp, err = c.api.Parse(ctx, &pb.Book{
		Id:    "1338",
		Title: "Lorem ipsum dolor sit amet, consectetur adipiscing eli",
		Price: 3,
	})
	require.NoError(b, err)

	assert.EqualValues(b,
		pb.Response{
			Book: &pb.Book{
				Id: "1338",
				Title: "Lorem ipsum dolor sit amet, consectetur adipiscing eli",
				Price: 3,
			},	
		}.Book, resp.Book)

}

