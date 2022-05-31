package benchmarks

import (
	"encoding/json"
	"testing"

	httpjson "github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/fasthttp"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

// ============================================== //
// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz  //
//                                                //
// create_account-4    3.0 ms/op                  //
// deposit-4           3.3 ms/op                  //
// withdraw-4          3.3 ms/op                  //
// ============================================== //
func Benchmark(b *testing.B) {

	core := httpjson.New(":8080")
	go func() {
		err := core.ListenAndServe()
		require.NoError(b, err)
	}()

	defer func() {
		err := core.Close()
		require.NoError(b, err)
	}()

	client := &fasthttp.Client{}
	for n := 0; n < b.N; n++ {
		doParse(client, b)
	}
}

func doParse(client *fasthttp.Client, b *testing.B) {

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
	req.SetRequestURI("http://localhost:8080/")
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
