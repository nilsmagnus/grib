.PHONY: default benchmark

default:
	go test -v 


benchmark: benchmark-mem benchmark-cpu
	rm profile.out memprofile.out
	rm gribtest.test

benchmark-mem:
	go version
	go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
	go tool pprof -top memprofile.out

benchmark-cpu:
	go version
	go test -bench=. -benchmem -cpuprofile profile.out
	go tool pprof -top profile.out

clean:
	rm profile.out memprofile.out gribtest.test 
