.PHONY: help
.DEFAULT_GOAL := help

run: ## Start Application
	@make generate
	@docker compose up -d

generate: ## Generate code
	@rm -rf internal/infra/grpc/generated
	@protoc -I=api/proto --go_out=internal/infra \
		--go-grpc_out=internal/infra \
		api/proto/*.proto
	@cd configs && gqlgen generate

fmt: ## format code
	@go fmt ./...

clear: ## Clear Application
	@docker compose down

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
