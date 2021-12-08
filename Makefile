.PHONY: test
test:
	@go test -race -timeout 1s ./...

.PHONY: build
build:
	@goreleaser --snapshot --skip-publish --rm-dist
