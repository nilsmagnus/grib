# Run the tests

either run with `go test` or `make test`

## Benchmark tests

run with `make benchmark` or just `make`

To inspect more details run `make profiles`

and use `go tool pprof --alloc_space memprofile.out` to inspect. 

i.e use the commands "top10" or "list extractData" to see whats happening.
