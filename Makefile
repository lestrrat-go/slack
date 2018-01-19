.PHONY: all test test-no-generate installdeps generate check-diff

installdeps:
	go get -t -v ./...

generate: endpoints.json
	@go generate

test: generate test-no-generate

test-no-generate:
	@go test -v ./...

slaproxy-docker:
	docker build -t lestrrat/slaproxy -f docker/Dockerfile.slaproxy .
	docker push lestrrat/slaproxy

check-diff:
	@./scripts/check-diff.sh