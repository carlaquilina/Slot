# Set the binary name
BINARY_NAME=./bin/slotgame

# Directory paths
CMD_DIR=./cmd/slotgame

all: build

build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) $(CMD_DIR)

run: build
	@echo "Running..."
	@./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@go clean
	@rm $(BINARY_NAME)

test:
	@echo "Testing..."
	@go test -race -cover ./...

generate:
	@echo "Generating..."
	@go generate ./...

.PHONY: all build run clean
