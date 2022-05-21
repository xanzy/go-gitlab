.PHONY: build test testacc-up testacc-down testacc

all: build test

build:
	go mod tidy
	go install

test: ## Run unit tests.
	go test $(TESTARGS) -v .

SERVICE ?= gitlab-ce
GITLAB_TOKEN ?= ACCTEST1234567890123
GITLAB_BASE_URL ?= http://127.0.0.1:8080/api/v4
ACC_TEST_DIR ?= ./acc

testacc-up:
	docker-compose up -d $(SERVICE)
	./scripts/await-healthy.sh

testacc-down:
	docker-compose down

testacc:
	TF_ACC=1 GITLAB_TOKEN=$(GITLAB_TOKEN) GITLAB_BASE_URL=$(GITLAB_BASE_URL) go test -v $(ACC_TEST_DIR) $(TESTARGS) -timeout 40m