# Use official Go image
FROM golang:1.25rc1-bookworm

# Install air (live reload tool)
RUN go install github.com/air-verse/air@latest
WORKDIR /app

# Copy go.mod and go.sum first for efficient cache
COPY go.mod go.sum ./

RUN go mod download

# Copy all source code
COPY . .

# Expose port your app listens on
EXPOSE 8080

# Start air for live reload
CMD ["air", "-c", ".air.toml"]
