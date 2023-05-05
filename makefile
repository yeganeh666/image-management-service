# Set the project name
PROJECT_NAME=myproject

# Set the Golang version
GO_VERSION=1.20

# Set the directory containing the main package
MAIN_DIR=./cmd

# Set the output directory for binaries
BIN_DIR=./cmd

# Set the output binary name
BINARY_NAME=image-management-service

.PHONY: all clean deps build test run

all: clean deps build test

clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)/*

deps:
	@echo "Installing dependencies..."
	@go mod download

build:
	@echo "Building..."
	@go build -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_DIR)/main.go

test:
	@echo "Testing..."
	@go test -v ./...

run:
	@echo "Running..."
	@go run $(MAIN_DIR)/main.go

