# Variables
CMD_DIR := api/
APP_NAME := mlvt

# Default target
all: build

# Run the application
run:
	go run api/main.go

# Help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo
	@echo "Usage:"
	@echo "  make run         Run the application"
