package test

import (
	"encoding/json"
	"net"
	"testing"
	"time"

	httpjson "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/fasthttp"

	"github.com/valyala/fasthttp"
)

func startHTTPServer(b *testing.B, l net.Listener) {
	var s = httpjson.New()
	s.Addr = l.Addr().String()
	s.Serve(l)
}

// ================================================================================================================ //
// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz																	//
// Benchmark_FastHTTP_JSON																							//
// Benchmark_FastHTTP_JSON-8            10000             61367 ns/op            2398 B/op         45 allocs/op		//
// Benchmark_FastHTTP_JSON-8            10000             61126 ns/op            2403 B/op         45 allocs/op		//
// Benchmark_FastHTTP_JSON-8            10000             61228 ns/op            2401 B/op         45 allocs/op		//
// Benchmark_FastHTTP_JSON-8            10000             63710 ns/op            2403 B/op         45 allocs/op		//
// Benchmark_FastHTTP_JSON-8            10000             61296 ns/op            2407 B/op         45 allocs/op		//
// ================================================================================================================ //
func Benchmark_FastHTTP_JSON(b *testing.B) {

	var l, err = net.Listen("tcp", "127.0.0.1:0") // aritrary port
	if err != nil {
		b.Fatal(err)
	}
	defer l.Close()

	go startHTTPServer(b, l)
	time.Sleep(100 * time.Millisecond) // sever starting timeout

	b.ReportAllocs()
	b.ResetTimer()

	client := &fasthttp.Client{}
	for n := 0; n < b.N; n++ {
		doHTTPRequest(client, b, l.Addr().String())
	}
}

func doHTTPRequest(client *fasthttp.Client, b *testing.B, addr string) {

	u := &httpjson.Book{
		ID:    "1338",
		Title: "Lorem ipsum dolor sit amet, consectetur adipiscing eli",
		Price: 3.21,
	}

	bt, err := json.Marshal(u)
	if err != nil {
		b.Fatalf("unable to marshal json: %v", err)
	}

	var (
		req  = fasthttp.AcquireRequest()
		resp = fasthttp.AcquireResponse()
	)

	req.SetRequestURI("http://" + addr)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/json")
	req.SetBodyRaw(bt)

	err = client.Do(req, resp)
	if err != nil {
		b.Fatalf("http request failed: %v", err)
	}

	defer resp.Body()

	var r httpjson.Response
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		b.Fatalf("unable to unmarshal json: %v", err)
	}

	if r.Message != "OK" || r.Code != 200 || r.Book.ID != "1338" {
		b.Fatalf("wrong http response: %v", err)
	}

	fasthttp.ReleaseResponse(resp)
}
