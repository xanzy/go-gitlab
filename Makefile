setup:
	go mod tidy
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
.PHONY: setup

test:
	go test ./... -race #-v??
.PHONY: test

fmt:
	@gofumpt -l -w .
.PHONY: fmt

lint:
	@golangci-lint run --config .golangci.yml 
.PHONY: lint
