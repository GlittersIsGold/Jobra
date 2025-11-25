#!/bin/bash

set -e

echo "Starting Jobra development environment"
POSTGRES_DATA_DIR="./data"
if [ ! -d "$POSTGRES_DATA_DIR" ]; then
    echo "Creating PostgreSQL data directory..."
    mkdir -p "$POSTGRES_DATA_DIR"
    echo "Data directory created: $POSTGRES_DATA_DIR"
fi

echo "Starting PostgreSQL"
if docker ps -a --format '{{.Names}}' | grep -q "jobra-postgres"; then
    echo "PostgreSQL container already exists, starting it..."
    docker start jobra-postgres
else
    echo "Starting new PostgreSQL container"
    docker run -d --name jobra-postgres \
      -e POSTGRES_DB=backend \
      -e POSTGRES_USER=user \
      -e POSTGRES_PASSWORD=password \
      -p 5432:5432 \
      -v "$(pwd)/data:/var/lib/postgresql/data" \
      -v $(pwd)/backend/sql/schema:/docker-entrypoint-initdb.d \
      postgres:16
fi

echo "Waiting for PostgreSQL"
for i in {1..30}; do
    if docker exec jobra-postgres pg_isready -U user -d backend > /dev/null 2>&1; then
        echo "PostgreSQL is ready!"
        break
    fi
    if [ $i -eq 30 ]; then
        echo "ERROR: PostgreSQL failed to start within 30 seconds"
        exit 1
    fi
    sleep 1
done

echo "Starting backend"
cd backend
go run cmd/api/main.go &
BACKEND_PID=$!
cd ..

sleep 3

echo "Starting frontend"
cd frontend
npm run dev &
FRONTEND_PID=$!
cd ..

echo "All services started!"
echo "   Backend:  http://localhost:8080"
echo "   Frontend: http://localhost:3000"
echo ""
echo "Press Ctrl+C to stop"

trap "echo 'Stopping'; kill $BACKEND_PID $FRONTEND_PID; docker stop jobra-postgres; docker rm jobra-postgres" INT

wait