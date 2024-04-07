# Builder stage
FROM golang:1.22.2 AS builder

WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o redirection-service ./cmd/server/main.go

# Final stage
FROM alpine:latest

# Set timezone (change TZ as needed)
ENV TZ=UTC
RUN apk add --no-cache tzdata

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Copy the binary from the builder stage
COPY --from=builder /app/redirection-service /redirection-service

# Expose and run the binary
EXPOSE $PORT
CMD ["/redirection-service"]

# Health check (replace with your health check endpoint)
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD wget -qO- http://localhost:$PORT/health || exit 1
