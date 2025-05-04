.PHONY: build run test clean tidy docker-build docker-run

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Download dependencies
tidy:
	go mod tidy

# Build Docker image
docker-build:
	docker build -t wallet-topup .

# Run Docker container
docker-run:
	docker-compose up

# Run with hot reload (requires air)
dev:
	air

# Generate swagger docs
swagger:
	swag init -g cmd/server/main.go

# Run linter
lint:
	golangci-lint run

# Run formatter
fmt:
	go fmt ./... 