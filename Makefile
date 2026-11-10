GO  ?= go
BIN := bin/binary-clock
PKG := ./...

.DEFAULT_GOAL := help

.PHONY: all build run test fmt vet lint check tidy clean help

all: check build ## Run checks, then build

build: ## Build the binary into ./bin
	$(GO) build -o $(BIN) ./cmd/binary-clock

run: ## Run the clock
	$(GO) run ./cmd/binary-clock

test: ## Run tests
	$(GO) test $(PKG)

fmt: ## Format sources in place
	gofmt -w .

vet: ## Run go vet
	$(GO) vet $(PKG)

lint: ## Check formatting and run go vet
	@out=$$(gofmt -l .); if [ -n "$$out" ]; then echo "gofmt needed:"; echo "$$out"; exit 1; fi
	$(GO) vet $(PKG)

check: lint test ## Lint and test (used by CI)

tidy: ## Tidy and verify go.mod / go.sum
	$(GO) mod tidy

clean: ## Remove build artifacts
	rm -rf bin

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'
