package benchmarks

import (
	"encoding/json"
	"testing"

	httpjson "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/fasthttp"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

// ============================================================================================================ //				
// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz  											  					//
//                                                											  					//
// BenchmarkFastHTTPJson									  										      		//
// BenchmarkFastHTTPJson-8            10000             66033 ns/op            2362 B/op         42 allocs/op   //
// BenchmarkFastHTTPJson-8            10000             65521 ns/op            2363 B/op         42 allocs/op   //
// BenchmarkFastHTTPJson-8            10000             64357 ns/op            2361 B/op         42 allocs/op	//
// BenchmarkFastHTTPJson-8            10000             65077 ns/op            2362 B/op         42 allocs/op	//
// BenchmarkFastHTTPJson-8            10000             67810 ns/op            2363 B/op         42 allocs/op	//
// ============================================================================================================ //
func BenchmarkFastHTTPJson(b *testing.B) {

	// var l, err = net.Listen("tcp", "127.0.0.1:0") // arbitrary port
	// require.NoError(b, err)

	// defer l.Close()

	core := httpjson.New("127.0.0.1:8080")
	go func() {
		err := core.ListenAndServe()
		require.NoError(b, err)
	}()

	defer func() {
		err := core.Close()
		require.NoError(b, err)
	}()

	client := &fasthttp.Client{}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		doParse(client, b)
	}
}

func doParse(client *fasthttp.Client, /*l net.Listener*/ b *testing.B) {

	u := &httpjson.Book{
		ID:    "1338",
		Title: "Lorem ipsum dolor sit amet, consectetur adipiscing eli",
		Price: 3.21,
	}

	bt, err := json.Marshal(u)
	require.NoError(b, err)

	var (
		req  = fasthttp.AcquireRequest()
		resp = fasthttp.AcquireResponse()
	)

	req.SetRequestURI("http://127.0.0.1:8080")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/json")
	req.SetBodyRaw(bt)

	err = client.Do(req, resp)
	require.NoError(b, err)

	defer resp.Body()

	var r httpjson.Response
	err = json.Unmarshal(resp.Body(), &r)
	require.NoError(b, err)

	fasthttp.ReleaseResponse(resp)
}
