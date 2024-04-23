# Application name
APP_NAME := pd-billing-orchestrator

test: ## Run tests
	go test -v ./...

lint: ## Run linter
	revive -set_exit_status -config revive.toml -formatter friendly ./src/...