APP_NAME=go-api
PKG=./...
Test_DIR=./test
DB_FILE=./database/todos.db

.PHONY: test run build deps reset-db
# Makefile for Go API project

run:
	@echo "Running the application..."
	@go run ./src/main.go
	@echo "Application stopped."

build:
	@echo "Building the application..."
	@go build -o $(APP_NAME) ./src/main.go
	@echo "Build completed. Executable: $(APP_NAME)"

deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@echo "Dependencies installed."

reset-db:
	@echo "Resetting the database..."
	@rm -f $(DB_FILE)
	@echo "Database reset completed."

test:
	@echo "Running tests..."
	@go test -v $(Test_DIR) -count=1 > ./result.log; tail -n 10 ./result.log
	@echo "Tests completed."

test-cover:
	@echo "Running tests with coverage..."
	@go test ./src/handlers $(Test_DIR) -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

default: deps run