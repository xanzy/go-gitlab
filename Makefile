setup:
	go mod tidy
	@go install mvdan.cc/gofumpt@latest
.PHONY: setup

test:
	go test ./... -race #-v??
.PHONY: test

fmt:
	@gofumpt -l -w .
.PHONY: fmt
