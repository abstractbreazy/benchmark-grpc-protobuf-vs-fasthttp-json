
### gRPC-Protobuf vs fastHTTP-JSON Benchmark

### ***Run tests***

Run benchmarks:
```bash
go test ./test -v -bench=Benchmark -benchmem
```

```bash
go test ./test -v -bench=Benchmark -benchmem -benchtime=10000x -count 5
```

### ***Results***

```bash
goos: darwin
goarch: amd64
pkg: github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/test
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
Benchmark_FastHTTP_JSON
Benchmark_FastHTTP_JSON-8          10000             61367 ns/op            2398 B/op         45 allocs/op
Benchmark_FastHTTP_JSON-8          10000             61126 ns/op            2403 B/op         45 allocs/op
Benchmark_FastHTTP_JSON-8          10000             61228 ns/op            2401 B/op         45 allocs/op
Benchmark_FastHTTP_JSON-8          10000             63710 ns/op            2403 B/op         45 allocs/op
Benchmark_FastHTTP_JSON-8          10000             61296 ns/op            2407 B/op         45 allocs/op
Benchmark_GRPC_Protobuf																						
Benchmark_GRPC_Protobuf-8          10000             27345 ns/op            9010 B/op        161 allocs/op
Benchmark_GRPC_Protobuf-8          10000             26052 ns/op            8987 B/op        161 allocs/op
Benchmark_GRPC_Protobuf-8          10000             25251 ns/op            8986 B/op        161 allocs/op
Benchmark_GRPC_Protobuf-8          10000             24900 ns/op            8985 B/op        161 allocs/op
Benchmark_GRPC_Protobuf-8          10000             24751 ns/op            8980 B/op        161 allocs/op
PASS
ok      github.com/abstractbreazy/benchmark-grpc-protobuf-vs-fasthttp-json/test 6.123s
```
