# LogDock Justfile — https://github.com/casey/just
# Install: cargo install just  OR  brew install just

# Default: list available recipes
default:
    @just --list

# ─── Development ──────────────────────────────────────────────────────────────

# Copy the standalone UI into the embedded webdist directory (no bundler needed)
ui-build:
    cp web/index.html internal/api/webdist/index.html
    cp web/i18n.js internal/api/webdist/i18n.js
    @echo "✓ UI and i18n copied to internal/api/webdist/"

# Build the Go binary
build: ui-build
    go build -ldflags="-s -w" -o logdock ./cmd/logdock

# Build without rebuilding UI (faster iteration on backend)
build-go:
    go build -ldflags="-s -w" -o logdock ./cmd/logdock

# Run all tests
test:
    go test ./...

# Run tests with race detector
test-race:
    go test -race ./...

# Run with coverage
test-cover:
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html

# Lint (requires golangci-lint)
lint:
    golangci-lint run ./...

# Format Go code
fmt:
    gofmt -w .

# Vet Go code
vet:
    go vet ./...

# ─── Running ──────────────────────────────────────────────────────────────────

# Run the server in development mode
dev:
    LOGDOCK_ENV=dev LOGDOCK_ADMIN_PASSWORD=admin go run ./cmd/logdock serve

# Run the TUI (headless monitoring)
tui:
    go run ./cmd/logdock tui

# Send a test log event
send-test:
    go run ./cmd/logdock ingest --message "Test event from justfile" --level info --source cli

# Send a test error event
send-error:
    go run ./cmd/logdock ingest --message "ERROR: database connection refused — timeout after 30s" --level error --source db

# ─── Docker ───────────────────────────────────────────────────────────────────

# Build Docker image
docker-build:
    docker build -t logdock:latest .

# Run with Docker Compose
docker-up:
    docker compose up -d

# Stop Docker Compose
docker-down:
    docker compose down

# View Docker logs
docker-logs:
    docker compose logs -f logdock

# ─── Production ───────────────────────────────────────────────────────────────

# Build production binary (requires UI build)
release:
    #!/usr/bin/env bash
    set -euo pipefail
    cp web/index.html internal/api/webdist/index.html
    cp web/i18n.js internal/api/webdist/i18n.js
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o logdock-linux-amd64 ./cmd/logdock
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o logdock-linux-arm64 ./cmd/logdock
    CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o logdock-darwin-arm64 ./cmd/logdock
    echo "✓ Binaries: logdock-linux-amd64, logdock-linux-arm64, logdock-darwin-arm64"

# Generate a secure random JWT secret
gen-secret:
    @openssl rand -hex 32

# Generate a secure master key for credential encryption
gen-master-key:
    @openssl rand -hex 32

# Check if server is healthy
health:
    go run ./cmd/logdock health

# ─── Maintenance ──────────────────────────────────────────────────────────────

# Create a data snapshot
snapshot:
    go run ./cmd/logdock snapshots --path ./data

# Show storage health
storage-health:
    go run ./cmd/logdock storage health --json

# Show compression config
compression:
    go run ./cmd/logdock compression --json

# Export today's logs
export-today:
    go run ./cmd/logdock export --format json --out export-$(date +%Y-%m-%d).json

# Clean build artifacts
clean:
    rm -f logdock logdock-linux-* logdock-darwin-* logdock-windows-*
    rm -f coverage.out coverage.html

# Show all available environment variables
env-help:
    @grep -E 'LOGDOCK_[A-Z_]+' resources/logdock.env.example | head -40
