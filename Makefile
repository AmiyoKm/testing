.PHONY: help test test-dir test-bench lint format tidy

help:  ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'


test: ## Run all tests
	go test -v ./...

test-dir: ## Run tests on a specific directory (e.g., make test-dir d=./integers)
	cd ${d} && go test -v


test-bench: ## Run all benchmarks
	go test -bench=. ./... -benchmem

lint: ## Lint the code
	go vet ./...

format: ## Format the code
	gofmt -w .

tidy: ## Tidy the modules
	go mod tidy

test-all: test test-bench ## Run all tests with benchmarks