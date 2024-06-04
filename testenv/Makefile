.PHONY: start
start: ## Start environment for integration tests
	docker-compose up -d --build
	sleep 2
	terraform init
	terraform apply -auto-approve

.PHONY: clean
clean: ## Clean environment for integration tests
	docker-compose down -v --remove-orphans
	rm -f terraform.tfstate*
	rm -f .terraform.lock.hcl

.PHONY: reset
reset: clean start ## Reset test environment
