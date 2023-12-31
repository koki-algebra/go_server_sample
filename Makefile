.PHONY: help
.DEFAULT_GOAL := help

PROTOCOL := "grpc"

run: ## Start Application
	@docker compose build --no-cache --build-arg PROTOCOL=$(PROTOCOL)
	@docker compose up -d

generate: ## Generate code
	@rm -rf internal/infra/grpc/generated
	@cd configs && buf generate ../api/proto
	@cd configs && gqlgen generate
	@cd api/http && oapi-codegen -config config.yml openapi.yml
	@cd configs && sqlboiler psql

fmt: ## format code
	@go fmt ./...
	@cd configs && buf format -w ../api/proto

lint: ## lint code
	@cd configs && buf lint ../api/proto

clear: ## Clear Application
	@docker compose down --volumes

logs: ## Show API server logs
	@docker compose logs -f api_server

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
