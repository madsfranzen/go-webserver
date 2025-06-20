# -------- STAGE 1: Build --------
FROM golang:1.25rc1-bookworm AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source files
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# -------- STAGE 2: Run --------
FROM debian:bookworm-slim

# Create non-root user
RUN useradd -m appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# Set permissions
RUN chown appuser:appuser /app/server

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Start the server
CMD ["./server"]

