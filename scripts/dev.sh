#!/bin/bash
# Development mode: runs backend and frontend dev server concurrently.

set -e

echo "Starting Homestack in development mode..."
echo "  Backend:  http://localhost:8080"
echo "  Frontend: http://localhost:5173"
echo ""

# Start backend
go run ./cmd/server &
BACKEND_PID=$!

# Start frontend dev server
cd frontend
npm run dev &
FRONTEND_PID=$!
cd ..

# Cleanup on exit
cleanup() {
    echo ""
    echo "Shutting down..."
    kill "$BACKEND_PID" "$FRONTEND_PID" 2>/dev/null || true
}
trap cleanup EXIT INT TERM

wait
