.PHONY: help build run test clean migrate seed install dev

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

build: ## Build the application
	@echo "Building application..."
	go build -o bin/satpen-api cmd/api/main.go
	@echo "Build complete: bin/satpen-api"

run: ## Run the application
	@echo "Running application..."
	go run cmd/api/main.go

dev: ## Run in development mode with auto-reload (requires air)
	@echo "Running in development mode..."
	air

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-cover: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	@echo "Clean complete"

migrate: ## Run database migrations
	@echo "Running migrations..."
	mysql -u root -p < migrations/001_create_tables.sql
	@echo "Migrations complete"

seed: ## Seed database with sample data
	@echo "Seeding database..."
	mysql -u root -p < migrations/002_seed_sample_data.sql
	@echo "Seeding complete"

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...
	@echo "Formatting complete"

lint: ## Run linter (requires golangci-lint)
	@echo "Running linter..."
	golangci-lint run
	@echo "Linting complete"

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t satpen-api:latest .
	@echo "Docker build complete"

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env satpen-api:latest

.DEFAULT_GOAL := help
