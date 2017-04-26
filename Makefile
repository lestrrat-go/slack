.PHONY: all test

installdeps:
	go get -t -v ./...

generate: endpoints.json
	@go generate

test: generate test-no-generate

test-no-generate:
	@go test -v ./...