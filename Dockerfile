# Multi-stage Dockerfile for Go development and production

# ===========================================
# DEVELOPMENT STAGE
# ===========================================
FROM golang:1.21-alpine AS development

# Install development tools
RUN apk add --no-cache \
    git \
    bash \
    curl \
    ca-certificates \
    && rm -rf /var/cache/apk/*

# Install Air for hot reload
RUN go install github.com/cosmtrek/air@latest

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Expose port
EXPOSE 8080

# Use Air for hot reload in development
CMD ["air", "-c", ".air.toml"]

# ===========================================
# PRODUCTION BUILD STAGE
# ===========================================
FROM golang:1.21-alpine AS builder

# Install git for private repos
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# ===========================================
# PRODUCTION STAGE
# ===========================================
FROM alpine:latest AS production

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

# Run the application
CMD ["./main"] 