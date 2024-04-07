# Builder stage
FROM golang:1.22.2 AS builder

WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o id-service ./cmd/server/main.go

# Final stage
FROM alpine:latest

# Set timezone (change TZ as needed)
ENV TZ=UTC
RUN apk add --no-cache tzdata

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Copy the binary from the builder stage
COPY --from=builder /app/id-service /id-service

# Expose and run the binary
EXPOSE $PORT
CMD ["/id-service"]