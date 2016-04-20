.PHONY: test build

build:
	go build .

test:
	go test -race -timeout 10s .

bench:
	go test -bench .
