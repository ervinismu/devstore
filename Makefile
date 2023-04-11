help: ## You are here! showing all command documenentation.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yaml


#===========#
#== TOOLS ==#
#===========#

migrate-up: ## Run migrations UP
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up

migrate-down: ## Rollback migrations, latest migration (1)
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1

migrate-all: ## Rollback migrations, all migrations
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1

migrate-create: ## Create a DB migration files e.g `make migrate-create name=migration-name`
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -seq -dir /migrations $(name)

lint: ## Running golangci-lint for code analysis.
lint:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm lint golangci-lint run -v


#=======================#
#== SETUP ENVIRONMNET ==#
#=======================#

environment: ## Setup environment.
environment:
	docker compose -f ${DOCKER_COMPOSE_FILE} up -d

server: ## Running application
server:
	go run cmd/main.go

shell-db: ## Enter to database console
	docker compose -f ${DOCKER_COMPOSE_FILE} exec db psql -U postgres -d postgres

test: ## Running golang testing.
test:
	go test ./... -count=1 -coverprofile=coverage.out

test-cover: ## Open golang testing coverage
test-cover:
	go tool cover -html=coverage.out
