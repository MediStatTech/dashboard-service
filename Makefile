.PHONY: help sqlc build test clean run migrate-up migrate-down migrate-status lint fmt tidy install-tools

# Default target
help:
	@echo "Available commands:"
	@echo "  make sqlc           - Generate sqlc code from SQL queries"
	@echo "  make build          - Build the application"
	@echo "  make test           - Run tests"
	@echo "  make run            - Run the application"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make migrate-up     - Run database migrations"
	@echo "  make migrate-down   - Rollback last migration"
	@echo "  make migrate-status - Show migration status"
	@echo "  make lint           - Run linter"
	@echo "  make fmt            - Format code"
	@echo "  make tidy           - Tidy go modules"
	@echo "  make install-tools  - Install required tools"

# Generate sqlc code
sqlc:
	@echo "Generating sqlc code..."
	sqlc generate
	@echo "✓ sqlc code generated successfully"

# Build the application
build:
	@echo "Building application..."
	go build -o bin/auth-service ./cmd/server
	@echo "✓ Build successful"

# Database migrations (requires DB_DSN environment variable or .env file)
migrate-up:
	@echo "Running migrations..."
	go run ./cmd/server --migrate-up
	@echo "✓ Migrations completed"

migrate-down:
	@echo "Rolling back last migration..."
	go run ./cmd/server --migrate-down
	@echo "✓ Rollback completed"

migrate-status:
	@echo "Checking migration status..."
	go run ./cmd/server --migrate-status

# Linting
lint:
	@echo "Running linter..."
	golangci-lint run ./...
	@echo "✓ Linting completed"

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .
	@echo "✓ Formatting completed"


# Full check - run all quality checks
check: sqlc fmt lint test
	@echo "✓ All checks passed"

# Database setup for local development
db-setup:
	@echo "Setting up local database..."
	@echo "Creating database..."
	createdb auth_db || echo "Database might already exist"
	@echo "✓ Database setup completed"

# Docker targets (if using Docker)
docker-build:
	@echo "Building Docker image..."
	docker build -t auth-service:latest .
	@echo "✓ Docker image built"

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env auth-service:latest

# Watch mode for development (requires entr or fswatch)
watch:
	@command -v entr >/dev/null 2>&1 || (echo "Install entr: brew install entr" && exit 1)
	@echo "Watching for changes..."
	find . -name '*.go' -o -name '*.sql' | entr -r make dev

# Quick commands
gen: sqlc
b: build
t: test
r: run
f: fmt
