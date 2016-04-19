.PHONY: test build

build:
	go build .

test:
	go test -race -timeout 10s .
