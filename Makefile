# Makefile for Go project with DDD architecture - Docker Compose centric

# Variables
APP_NAME := try_golang
BINARY_NAME := main
COMPOSE_FILE := docker-compose.yml

# Go variables
GO := go
GOOS := linux
GOARCH := amd64
CGO_ENABLED := 0

# Directories
CMD_DIR := cmd/api
BUILD_DIR := build
DIST_DIR := dist

# Default target
.PHONY: help
help: ## Show this help message
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# ===========================================
# DOCKER COMPOSE DEVELOPMENT TARGETS
# ===========================================

.PHONY: dev
dev: ## Start development server with hot reload
	@echo "Starting development server with hot reload..."
	docker-compose up

.PHONY: dev-d
dev-d: ## Start development server in background
	@echo "Starting development server in background..."
	docker-compose up -d

.PHONY: dev-stop
dev-stop: ## Stop development server
	@echo "Stopping development server..."
	docker-compose down

.PHONY: dev-restart
dev-restart: dev-stop dev ## Restart development server

.PHONY: logs
logs: ## Show development logs
	@echo "Showing development logs..."
	docker-compose logs -f

.PHONY: logs-app
logs-app: ## Show application logs only
	@echo "Showing application logs..."
	docker-compose logs -f app

# ===========================================
# DOCKER COMPOSE TESTING TARGETS
# ===========================================

.PHONY: test
test: ## Run tests in Docker Compose
	@echo "Running tests in Docker Compose..."
	docker-compose run --rm app go test ./...

.PHONY: test-verbose
test-verbose: ## Run tests with verbose output in Docker Compose
	@echo "Running tests with verbose output in Docker Compose..."
	docker-compose run --rm app go test -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage report in Docker Compose
	@echo "Running tests with coverage in Docker Compose..."
	docker-compose run --rm app sh -c "go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html"
	@echo "Coverage report generated: coverage.html"

# ===========================================
# DOCKER COMPOSE BUILD TARGETS
# ===========================================

.PHONY: build
build: ## Build all Docker Compose services
	@echo "Building all Docker Compose services..."
	docker-compose build

.PHONY: build-app
build-app: ## Build application service only
	@echo "Building application service..."
	docker-compose build app

.PHONY: build-prod
build-prod: ## Build production image using Docker Compose
	@echo "Building production image..."
	docker-compose build --target production app

.PHONY: build-multi
build-multi: ## Build multi-platform binaries using Docker Compose
	@echo "Building multi-platform binaries using Docker Compose..."
	docker-compose run --rm app sh -c "\
		mkdir -p /app/dist && \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /app/dist/$(BINARY_NAME)-linux-amd64 $(CMD_DIR)/main.go && \
		GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o /app/dist/$(BINARY_NAME)-darwin-amd64 $(CMD_DIR)/main.go && \
		GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o /app/dist/$(BINARY_NAME)-windows-amd64.exe $(CMD_DIR)/main.go"
	@echo "Multi-platform binaries created in dist/ directory"

# ===========================================
# DOCKER COMPOSE DEPENDENCY TARGETS
# ===========================================

.PHONY: deps
deps: ## Download dependencies in Docker Compose
	@echo "Downloading dependencies in Docker Compose..."
	docker-compose run --rm app go mod download
	@echo "Dependencies downloaded"

.PHONY: deps-update
deps-update: ## Update dependencies in Docker Compose
	@echo "Updating dependencies in Docker Compose..."
	docker-compose run --rm app sh -c "go get -u ./... && go mod tidy"
	@echo "Dependencies updated"

.PHONY: deps-vendor
deps-vendor: ## Vendor dependencies in Docker Compose
	@echo "Vendoring dependencies in Docker Compose..."
	docker-compose run --rm app go mod vendor
	@echo "Dependencies vendored"

# ===========================================
# DOCKER COMPOSE CODE QUALITY TARGETS
# ===========================================

.PHONY: fmt
fmt: ## Format Go code in Docker Compose
	@echo "Formatting Go code in Docker Compose..."
	docker-compose run --rm app go fmt ./...

.PHONY: vet
vet: ## Vet Go code in Docker Compose
	@echo "Vetting Go code in Docker Compose..."
	docker-compose run --rm app go vet ./...

.PHONY: lint-install
lint-install: ## Install linter in Docker Compose
	@echo "Installing linter in Docker Compose..."
	docker-compose run --rm app sh -c "go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

.PHONY: lint
lint: ## Run linter in Docker Compose
	@echo "Running linter in Docker Compose..."
	docker-compose run --rm app sh -c "golangci-lint run"

.PHONY: check
check: ## Run all code quality checks in Docker Compose
	@echo "Running all code quality checks in Docker Compose..."
	$(MAKE) fmt
	$(MAKE) vet
	$(MAKE) lint
	$(MAKE) test

# ===========================================
# DOCKER COMPOSE DATABASE TARGETS
# ===========================================

.PHONY: db-up
db-up: ## Start database only
	@echo "Starting database..."
	docker-compose up -d postgres

.PHONY: db-down
db-down: ## Stop database only
	@echo "Stopping database..."
	docker-compose stop postgres

.PHONY: db-logs
db-logs: ## Show database logs
	@echo "Showing database logs..."
	docker-compose logs -f postgres

.PHONY: db-reset
db-reset: ## Reset database (remove volume)
	@echo "Resetting database..."
	docker-compose down -v
	docker-compose up -d postgres

# ===========================================
# PRODUCTION TARGETS
# ===========================================

.PHONY: prod
prod: ## Start production server
	@echo "Starting production server..."
	docker-compose -f docker-compose.prod.yml up -d

.PHONY: prod-stop
prod-stop: ## Stop production server
	@echo "Stopping production server..."
	docker-compose -f docker-compose.prod.yml down

.PHONY: prod-logs
prod-logs: ## Show production logs
	@echo "Showing production logs..."
	docker-compose -f docker-compose.prod.yml logs -f

.PHONY: prod-build
prod-build: ## Build production image
	@echo "Building production image..."
	docker-compose -f docker-compose.prod.yml build

.PHONY: prod-restart
prod-restart: prod-stop prod ## Restart production server

# ===========================================
# DOCKER COMPOSE CLEANUP TARGETS
# ===========================================

.PHONY: clean
clean: ## Clean build artifacts and Docker Compose resources
	@echo "Cleaning build artifacts and Docker Compose resources..."
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)
	rm -rf dist/
	rm -rf node_modules/
	rm -rf vendor/
	rm -f coverage.out coverage.html
	$(MAKE) clean-compose

.PHONY: clean-compose
clean-compose: ## Clean Docker Compose resources (preserves go_cache)
	@echo "Cleaning Docker Compose resources..."
	docker-compose down --remove-orphans
	docker-compose rm -f

.PHONY: clean-compose-full
clean-compose-full: ## Clean Docker Compose resources including volumes
	@echo "Cleaning Docker Compose resources including volumes..."
	docker-compose down -v --remove-orphans
	docker-compose rm -f

.PHONY: clean-images
clean-images: ## Clean Docker images
	@echo "Cleaning Docker images..."
	docker-compose down --rmi all

.PHONY: clean-all
clean-all: ## Clean everything (containers, images, volumes)
	@echo "Cleaning everything..."
	docker-compose down -v --rmi all --remove-orphans

# ===========================================
# DOCKER COMPOSE UTILITY TARGETS
# ===========================================

.PHONY: shell
shell: ## Open shell in application container
	@echo "Opening shell in application container..."
	docker-compose run --rm app /bin/bash

.PHONY: bash
bash: ## Open bash shell in application container (alias for shell)
	@echo "Opening bash shell in application container..."
	docker-compose run --rm app /bin/bash

.PHONY: shell-db
shell-db: ## Open shell in database container
	@echo "Opening shell in database container..."
	docker-compose run --rm postgres /bin/bash

.PHONY: ps
ps: ## Show running containers
	@echo "Showing running containers..."
	docker-compose ps

.PHONY: status
status: ## Show status of all services
	@echo "Showing status of all services..."
	docker-compose ps
	@echo ""
	@echo "Application logs (last 10 lines):"
	docker-compose logs --tail=10 app

.PHONY: install-tools
install-tools: ## Install development tools in Docker Compose
	@echo "Installing development tools in Docker Compose..."
	docker-compose run --rm app sh -c "go install github.com/air-verse/air@latest && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
	@echo "Development tools installed"

.PHONY: all
all: ## Build, test, and run all checks in Docker Compose
	@echo "Running all targets in Docker Compose..."
	$(MAKE) clean-compose
	$(MAKE) deps
	$(MAKE) lint-install
	$(MAKE) check
	$(MAKE) build 