# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum* ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o worker ./cmd/worker

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Install FFmpeg and other runtime dependencies
RUN apk add --no-cache \
    ffmpeg \
    ca-certificates \
    tzdata

# Copy binary from builder
COPY --from=builder /app/worker .

# Create necessary directories
RUN mkdir -p /app/processed_videos /app/thumbnails

# Run the worker
CMD ["./worker"]
