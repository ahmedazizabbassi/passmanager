# Password Manager

A secure and scalable password manager built with Go.

## Project Structure

```
password-manager/
│── cmd/                  # Main application entry points
│   ├── server/           # Main API server (REST)
│── internal/             # Private app logic
│   ├── auth/             # Authentication & encryption logic
│   ├── config/           # Configuration management
│   ├── database/         # MySQL connection & migrations
│   ├── models/           # Database entities
│   ├── repository/       # Data access layer
│   ├── services/         # Business logic
│   ├── sync/             # Hybrid sync logic
│   ├── worker/           # Background workers
│── pkg/                  # Public utility packages
│── api/                  # OpenAPI specs
│── web/                  # Static files
│── scripts/              # DevOps scripts
│── test/                 # Integration and unit tests
```

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your environment variables
3. Run `go mod download` to install dependencies
4. Start the server with `go run cmd/server/main.go`

## Available Commands

```bash
make build         # Build the application
make run          # Run the application
make test         # Run tests
make clean        # Clean build files
make lint         # Run go fmt and go vet
make docker-build # Build Docker image
make docker-run   # Run Docker container
make help         # Show available commands
```

## Development

- Go 1.21+
- MySQL
- Docker (optional)
