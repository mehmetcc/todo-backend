# Stage 1: Build the Go application
FROM --platform=linux/arm64 golang:1.22-bookworm AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . ./

# Navigate to the directory containing main.go
WORKDIR /app/cmd/app

# Set environment variables for cross-compilation
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm64

# Build the Go binary for ARM64
RUN go build -o /todo-backend main.go

# Stage 2: Create a minimal runtime image
FROM --platform=linux/arm64 alpine:latest

# Install CA certificates for HTTPS support
RUN apk add --no-cache ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /todo-backend /todo-backend

# Ensure the binary is executable
RUN chmod +x /todo-backend

# Set the working directory
WORKDIR /root/

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["/todo-backend"]
