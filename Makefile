linter: ### check by golangci linter
	golangci-lint run --out-format tab --fix
.PHONY: linter-golangci

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test