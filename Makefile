.PHONY: all build run test clean lint docker-build docker-run help

# Go parameters
BINARY_NAME=passmanager
MAIN_FILE=cmd/server/main.go

all: lint test build

build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

run:
	go run $(MAIN_FILE)

test:
	go test -v ./...

clean:
	go clean
	rm -f $(BINARY_NAME)

lint:
	go vet ./...
	go fmt ./...

docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 8080:8080 $(BINARY_NAME)

help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build files"
	@echo "  make lint         - Run go fmt and go vet"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"
