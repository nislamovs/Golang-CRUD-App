# Variables
APP_NAME := golang-crud-app
BIN_DIR := ./bin
DOCKER_IMAGE := golang-crud-app:latest

# Default target
.PHONY: all
all: build

# Build the Go application
.PHONY: build
build:
	@bash -x ./scripts/bash/buildApp.sh

# Run the application locally
.PHONY: run-local
run-local: build
	@echo "Running $(APP_NAME) locally..."
	@./$(BIN_DIR)/$(APP_NAME)

# Build Docker image and run container
.PHONY: docker-run
docker-run:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .
	@echo "Running Docker container..."
	@docker run --rm -p 8080:8080 $(DOCKER_IMAGE)

# Stop and remove all docker containers
.PHONY: clean-docker-conts
clean-docker-conts:
	@echo "Stopping and removing docker containers..."
	@bash -x ./scripts/bash/stop_n_remove_containers.sh

# Run tests
.PHONY: test
test:
	@echo "Running Go tests..."
	@go test ./... -v