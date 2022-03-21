default: test

GOBIN = $(shell pwd)/bin

test: ## Run unit tests.
	go test -v ./...

fmt: tool-golangci-lint ## Format files and fix issues.
	gofmt -w -s .
	$(GOBIN)/golangci-lint run --fix

lint-golangci: tool-golangci-lint ## Run golangci-lint linter (same as fmt but without modifying files).
	$(GOBIN)/golangci-lint run

SERVICE ?= gitlab-ce
GITLAB_TOKEN ?= ACCTEST1234567890123
GITLAB_BASE_URL ?= http://127.0.0.1:8080/api/v4

testacc-up: ## Launch a GitLab instance.
	docker-compose up -d $(SERVICE)
	./scripts/await-healthy.sh

testacc-down: ## Teardown a GitLab instance.
	docker-compose down

# TOOLS
# Tool dependencies are installed into a project-local /bin folder.

tool-golangci-lint:
	@$(call install-tool, github.com/golangci/golangci-lint/cmd/golangci-lint)

define install-tool
	cd tools && GOBIN=$(GOBIN) go install $(1)
endef
