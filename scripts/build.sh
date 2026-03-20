#!/bin/bash
set -e

echo "Building Homestack..."

# Build frontend
echo "→ Building frontend..."
cd frontend
npm ci
npm run build
cd ..

# Build Go binary (static, no CGO)
echo "→ Building Go binary..."
CGO_ENABLED=0 go build -ldflags="-s -w" -o homestack ./cmd/server

echo "✅ Build complete: ./homestack"
echo "Run with: ./homestack"
