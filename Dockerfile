# Build stage
FROM golang:1.21-bullseye AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Final stage
FROM debian:bullseye-slim

# Install SQLite runtime
RUN apt-get update && apt-get install -y sqlite3 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy necessary files
COPY templates/ templates/
COPY migrations/ migrations/
COPY static/ static/

# Create volume for persistent database
VOLUME /app/data

# Expose port 8080
EXPOSE 8110

# Run the application
CMD ["./main"]