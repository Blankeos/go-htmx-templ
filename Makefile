# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@templ generate
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean

# Database management
migrate-status: 
	@goose -dir migrations sqlite ./test.db status

migrate-up:
	@goose -dir migrations sqlite ./test.db up

migrate-down:
	@goose -dir migrations sqlite ./test.db down

migrate-create:
	@if [ -n "$(N)" ]; then \
		goose -dir migrations sqlite "" create $(N) sql; \
	else \
		echo "Incorrect usage."; \
		echo "Usage: make migrate-create N=<migration_name>"; \
	fi;