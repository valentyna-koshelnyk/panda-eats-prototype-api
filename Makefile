# Application name
APP_NAME := panda-eats-prototype-api

test: ## Run tests
	go test -v ./...

lint: ## Run linter
	revive -set_exit_status -config revive.toml -formatter friendly ./panda-eats-prototype-api/..
