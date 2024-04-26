# Application name
APP_NAME := panda-eats-prototype-api

test: ## Run tests
	go test -v ./...

lint: ## Run linter
	revive -set_exit_status -config revive.toml -formatter friendly ./...

logger:
 ## Clears which exact code base ran the code and resulted in potential error
 VERSION ?= $(shell git describe --match 'v[0-9]*' --tags --always)
build:
	@go build -ldflags "-X main.version=$(VERSION)"
