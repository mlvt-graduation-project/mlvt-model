# Variables
CMD_DIR := api/
APP_NAME := mlvt

# Default target
all: build

# Run the application
run:
	go run api/main.go

build:
	go build -o mlvt $(CMD_DIR)/main.go

# Help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo
	@echo "Usage:"
	@echo "  make run         Run the application"
