# Step 1: Use official Go image to build the binary
FROM golang:1.18 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (cache benefit)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go app for Linux (important!)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapi main.go

# Step 2: Use a minimal base image for production
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapi .

# Expose the app's port
EXPOSE 8080

# Run the binary
CMD ["./myapi"]