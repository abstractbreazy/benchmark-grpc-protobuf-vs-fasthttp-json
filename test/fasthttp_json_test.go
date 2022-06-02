package benchmarks

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

func BenchmarkListener(b *testing.B) {

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
		b.Logf("Addr: %s", l.Addr().String())
		//b.Logf("Domain: %v", l.Addr().(*net.TCPAddr).IP.String())

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

	b.Logf("resp: %v", r)

	if r.Message != "OK" || r.Code != 200 || r.Book.ID != "1338" {
		b.Fatalf("wrong http response: %v", err)
	}

	fasthttp.ReleaseResponse(resp)
}
