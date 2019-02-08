.PHONY: default benchmark

default:
	go test -v 


benchmark:
	go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
	go tool pprof -top profile.out
	go tool pprof -top memprofile.out
	rm profile.out memprofile.out
	rm gribtest.test

