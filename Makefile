PROJECT_NAME=grpc-clean
MODULE_NAME=grpc-clean

.DEFAULT_GOAL := build

.PHONY: build
build:
	@go build ./cmd/grpc-clean/

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@go test -v -coverprofile coverage.out ./...

.PHONY: coverage
coverage:
	@go tool cover -html=coverage.out

.PHONY: get
get:
	@go mod download



