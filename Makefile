.PHONY: help
.DEFAULT_GOAL := help

run: ## Start web server
	@make generate
	@go run cmd/main.go

generate: ## Generate code
	@rm -rf internal/infra/grpc/generated
	@protoc -I=api/proto --go_out=internal/infra \
		--go-grpc_out=internal/infra \
		api/proto/*.proto

fmt: ## format code
	@go fmt ./...

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
