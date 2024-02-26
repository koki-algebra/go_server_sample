.DEFAULT_GOAL := help

# Environment variables
DB_DATABASE := app
DB_USER := app
DB_PORT := 5432

.PHONY: run
run: ## Start Application
	@docker compose up -d

.PHONY: generate
generate: ## Generate code
	@rm -rf internal/infra/grpc/generated
	@buf generate ./api/proto
	@gqlgen generate
	@sqlboiler psql
	@cd api/http && oapi-codegen -config config.yml openapi.yml

.PHONY: fmt
fmt: ## format code
	@go fmt ./...
	@buf format -w ./api/proto

.PHONY: lint
lint: ## lint code
	@buf lint ./api/proto

.PHONY: clear
clear: ## Clear Application
	@docker compose down --volumes

.PHONY: logs
logs: ## Show API server logs
	@docker compose logs -f api_server

psql: ## Login PostgreSQL
	@psql --host localhost --port $(DB_PORT) --username $(DB_USER) --dbname ${DB_DATABASE} --password

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
