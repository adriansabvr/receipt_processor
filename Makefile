linter: ### check by golangci linter
	golangci-lint run --out-format tab --fix
.PHONY: linter-golangci

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test

swag-v1: ### swag init
	swag init -g internal/controller/http/v1/router.go
.PHONY: swag-v1

run: swag-v1 ### swag run
	go mod tidy && go mod download && go run ./cmd/app
.PHONY: run