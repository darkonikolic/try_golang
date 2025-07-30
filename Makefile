# Makefile for Go project with DDD architecture

# Variables
APP_NAME := try_golang
BINARY_NAME := main
DOCKER_IMAGE := $(APP_NAME)
DOCKER_TAG := latest

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
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# ===========================================
# DEVELOPMENT TARGETS
# ===========================================

.PHONY: dev
dev: ## Run application in development mode with hot reload using Docker
	@echo "Starting development server with hot reload using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):dev .
	docker run -p 8080:8080 -v $(PWD):/app $(DOCKER_IMAGE):dev

.PHONY: run
run: ## Run application using Docker
	@echo "Running application using Docker..."
	docker build --target production -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: test
test: ## Run tests using Docker
	@echo "Running tests using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):test .
	docker run --rm $(DOCKER_IMAGE):test sh -c "go test ./..."

.PHONY: test-verbose
test-verbose: ## Run tests with verbose output using Docker
	@echo "Running tests with verbose output using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):test .
	docker run --rm $(DOCKER_IMAGE):test sh -c "go test -v ./..."

.PHONY: test-coverage
test-coverage: ## Run tests with coverage report using Docker
	@echo "Running tests with coverage using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):test .
	docker run --rm -v $(PWD):/app $(DOCKER_IMAGE):test sh -c "go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html"
	@echo "Coverage report generated: coverage.html"

# ===========================================
# BUILD TARGETS
# ===========================================

.PHONY: build
build: ## Build application using Docker
	@echo "Building application using Docker..."
	docker build --target production -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "Binary is available in the Docker image"

.PHONY: build-linux
build-linux: ## Build application for Linux using Docker
	@echo "Building application for Linux using Docker..."
	docker build --target production -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "Linux binary is available in the Docker image"

.PHONY: build-all
build-all: ## Build application for multiple platforms using Docker
	@echo "Building for multiple platforms using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):build-all .
	docker run --rm -v $(PWD)/dist:/app/dist $(DOCKER_IMAGE):build-all sh -c "\
		mkdir -p /app/dist && \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /app/dist/$(BINARY_NAME)-linux-amd64 $(CMD_DIR)/main.go && \
		GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o /app/dist/$(BINARY_NAME)-darwin-amd64 $(CMD_DIR)/main.go && \
		GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o /app/dist/$(BINARY_NAME)-windows-amd64.exe $(CMD_DIR)/main.go"
	@echo "Multi-platform binaries created in dist/ directory"

# ===========================================
# DOCKER TARGETS
# ===========================================

.PHONY: docker-build
docker-build: ## Build Docker image for production
	@echo "Building Docker image..."
	docker build --target production -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: docker-build-dev
docker-build-dev: ## Build Docker image for development
	@echo "Building Docker development image..."
	docker build --target development -t $(DOCKER_IMAGE):dev .

.PHONY: docker-run
docker-run: ## Run Docker container in production mode
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: docker-run-dev
docker-run-dev: ## Run Docker container in development mode
	@echo "Running Docker development container..."
	docker run -p 8080:8080 -v $(PWD):/app $(DOCKER_IMAGE):dev

.PHONY: docker-compose-up
docker-compose-up: ## Start services with docker-compose
	@echo "Starting services with docker-compose..."
	docker-compose up

.PHONY: docker-compose-down
docker-compose-down: ## Stop services with docker-compose
	@echo "Stopping services with docker-compose..."
	docker-compose down

# ===========================================
# CLEANUP TARGETS
# ===========================================

.PHONY: clean
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)
	rm -f coverage.out coverage.html

.PHONY: clean-docker
clean-docker: ## Clean Docker images
	@echo "Cleaning Docker images..."
	docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG) 2>/dev/null || true
	docker rmi $(DOCKER_IMAGE):dev 2>/dev/null || true

# ===========================================
# DEPENDENCY TARGETS
# ===========================================

.PHONY: deps
deps: ## Download dependencies using Docker
	@echo "Downloading dependencies using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):deps .
	@echo "Dependencies downloaded in Docker image"

.PHONY: deps-update
deps-update: ## Update dependencies using Docker
	@echo "Updating dependencies using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):deps .
	docker run --rm -v $(PWD):/app $(DOCKER_IMAGE):deps sh -c "go get -u ./... && go mod tidy"
	@echo "Dependencies updated"

.PHONY: deps-vendor
deps-vendor: ## Vendor dependencies using Docker
	@echo "Vendoring dependencies using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):deps .
	docker run --rm -v $(PWD):/app $(DOCKER_IMAGE):deps sh -c "go mod vendor"
	@echo "Dependencies vendored"

# ===========================================
# LINTING AND FORMATTING
# ===========================================

.PHONY: fmt
fmt: ## Format Go code using Docker
	@echo "Formatting Go code using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):fmt .
	docker run --rm -v $(PWD):/app $(DOCKER_IMAGE):fmt sh -c "go fmt ./..."

.PHONY: vet
vet: ## Vet Go code using Docker
	@echo "Vetting Go code using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):vet .
	docker run --rm -v $(PWD):/app $(DOCKER_IMAGE):vet sh -c "go vet ./..."

.PHONY: lint
lint: ## Run linter using Docker
	@echo "Running linter using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):lint .
	docker run --rm -v $(PWD):/app $(DOCKER_IMAGE):lint sh -c "go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && golangci-lint run"

# ===========================================
# UTILITY TARGETS
# ===========================================

.PHONY: install-tools
install-tools: ## Install development tools using Docker
	@echo "Installing development tools using Docker..."
	docker build --target development -t $(DOCKER_IMAGE):tools .
	@echo "Development tools installed in Docker image"

.PHONY: check
check: ## Run all checks using Docker (fmt, vet, test)
	@echo "Running all checks using Docker..."
	$(MAKE) fmt
	$(MAKE) vet
	$(MAKE) test

.PHONY: all
all: ## Build, test, and run all checks using Docker
	@echo "Running all targets using Docker..."
	$(MAKE) clean
	$(MAKE) deps
	$(MAKE) check
	$(MAKE) build 