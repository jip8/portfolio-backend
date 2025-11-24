#!/bin/bash
set -e

echo "Building portfolio-backend..."

# Build the Go application for Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server/main.go

echo "Build completed successfully!"
echo "Binary: ./server"
