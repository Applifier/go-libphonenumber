.PHONY: test build

build:
	go build .

test:
	go test -race -timeout 10s .

bench:
	go test -bench .

lint:
	test -z "`golint .`" && echo "Go lint ok" || (echo "Go lint failed" && golint . && false)

vet:
	go vet . && echo "Go vet ok"

suite: vet lint test bench
