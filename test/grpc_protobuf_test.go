package test

import (
	"context"
	"log"
	"testing"
	"time"

	srv "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/grpc-protobuf"
	"github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/grpc-protobuf/client"
	proto "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/grpc-protobuf/proto/gen"
)

func init() {
	go func() {
		if err := srv.New().Start(); err != nil {
			log.Fatalf("failed to start grpc server %s", err)
		}
	}()
	time.Sleep(100 * time.Microsecond)
}

// ================================================================================================================ //
// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz																	//
// Benchmark_GRPC_Protobuf																							//
// Benchmark_GRPC_Protobuf-8          10000             27345 ns/op            9010 B/op        161 allocs/op		//
// Benchmark_GRPC_Protobuf-8          10000             26052 ns/op            8987 B/op        161 allocs/op		//
// Benchmark_GRPC_Protobuf-8          10000             25251 ns/op            8986 B/op        161 allocs/op		//
// Benchmark_GRPC_Protobuf-8          10000             24900 ns/op            8985 B/op        161 allocs/op		//
// Benchmark_GRPC_Protobuf-8          10000             24751 ns/op            8980 B/op        161 allocs/op		//
// ================================================================================================================ //

func Benchmark_GRPC_Protobuf(b *testing.B) {
	listenAddr := "localhost:15000"
	// init client
	c, err := client.NewGRPCClient(listenAddr)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		doGRPCRequest(c, b)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doGRPCRequest(c, b)
		}
	})
}

func doGRPCRequest(client proto.APIClient, b *testing.B) {
	var (
		ctx  = context.Background()
		err  error
		resp *proto.Response
	)

	resp, err = client.Parse(ctx, &proto.Book{
		Id:    "1338",
		Title: "Lorem ipsum dolor sit amet, consectetur adipiscing eli",
		Price: 3,
	})
	if err != nil {
		b.Fatal(err)
	}

	if resp == nil || resp.Code != 200 || resp.Book == nil || resp.Book.Id != "1338" {
		b.Fatalf("wrong grpc response: %v\n", resp)
	}

}
