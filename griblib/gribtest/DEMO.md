# Demo of benchmarking

## before optimize
git co before-optimize
go test -memprofile memprofile.out -cpuprofile profile.out

go tool pprof -web memprofile.out
go tool pprof -web cpuprofile.out

# Explanation 
"when there is no more capacity in the underlying array, a new array is allocated with 2*lenght the original array for amortized linear complexity."
-> solution pre-allocate what you know


# after optimize
git co master
go test -memprofile memprofile.out -cpuprofile profile.out

go tool pprof -web memprofile.out
go tool pprof -web cpuprofile.out

