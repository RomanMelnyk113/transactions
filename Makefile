build: ## Build the app
	go mod tidy
	go build -o juni

docs: ## Generates api docs (aka swagger)
	swag init

test: ## Run tests
	cd tests/api/v1/ && go test

run: build docs ## Start application with fresh build and swagger docs
	./juni

import-db: ## Imports test database
	docker-compose run --rm db sh -c "psql -h db -U admin juni < /tmp/db.sql"


# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
