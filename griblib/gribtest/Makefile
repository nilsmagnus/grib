.PHONY: default benchmark

default:
	go test -v 


benchmark:
	go version
	go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
	go tool pprof -top10 profile.out
	go tool pprof -top10 memprofile.out
	rm profile.out memprofile.out
	rm gribtest.test

