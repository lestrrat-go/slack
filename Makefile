.PHONY: all test

all:
	go get -t -v ./...

generate: endpoints.json
	@go generate

test: generate
	@go test -v ./...