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

deploy:
	@echo "🚀 Building project..."
	go build -o mlvt $(CMD_DIR)/main.go
	@echo "🔄 Restarting mlvt service..."
	sudo systemctl restart mlvt
	@echo "✅ Done."

# Help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo
	@echo "Usage:"
	@echo "  make run         Run the application"
