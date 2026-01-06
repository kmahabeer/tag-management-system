# Tag Management System - Build Targets

.PHONY: go-tidy go-build go-clean go-run go-rebuild go-test

# Go module management
go-tidy:
	cd backend && go mod tidy

# Build the Go server binary to bin/ directory
go-build:
	cd backend && mkdir -p bin && go build -o bin/server ./cmd/server

# Clean build artifacts
go-clean:
	cd backend && rm -rf bin/

# Run the server (builds if needed)
go-run: go-build
	cd backend && ./bin/server

# Rebuild from scratch
go-rebuild: go-clean go-build

# Run all Go tests
go-test:
	cd backend && go test ./...