.PHONY: help
.DEFAULT_GOAL := help

# migration file name
NAME := ""

# for database
DB_HOST := localhost
DB_USER := app
DB_NAME := app
DB_PASSWORD := password

# for goose
GOOSE_DRIVER := postgres
GOOSE_DBSTRING := "host=$(DB_HOST) user=$(DB_USER) dbname=$(DB_NAME) password=$(DB_PASSWORD) sslmode=disable"

new: ## Create new migration file. e.g. $ make new NAME=add_user_table
	@goose -dir migrations -s create $(NAME) sql

status: ## Show migration status
	@goose -dir migrations $(GOOSE_DRIVER) $(GOOSE_DBSTRING) status

up: ## Migration up
	@goose -dir migrations $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
