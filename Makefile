generate: endpoints.json
	@go generate

test: generate
	@go test -v .