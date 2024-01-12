# Makefile for the TunnelManager project

.PHONY: build clean test run

# Build variables
BINARY_NAME=tunnelctl
BUILD_DIR=./cmd/tunnelctl

# Go related variables
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(BINARY_NAME)-stderr.txt

# Commands
GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GORUN=go run

# Targets
all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) $(BUILD_DIR)

test:
	@echo "Running tests..."
	@$(GOTEST) ./...

clean:
	@echo "Cleaning up..."
	@$(GOCLEAN)
	@rm -f $(GOBIN)/$(BINARY_NAME)
	@rm -f $(STDERR)

run:
	@echo "Running $(BINARY_NAME)..."
	@$(GORUN) $(BUILD_DIR)/main.go

# To run in verbose mode, use `make run ARGS="-v"`
ARGS=
verbose: run ARGS="-v"
