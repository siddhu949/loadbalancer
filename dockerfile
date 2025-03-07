# Use a lightweight Go base image
FROM golang:1.23.4 AS builder

# Set the working directory
WORKDIR /app

# Copy all files to container
COPY . .

# Download dependencies & build the binary
RUN go mod tidy && go build -o loadbalancer ./cmd/server/main.go

# Use a minimal base image for final container
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy built binary from builder
COPY --from=builder /app/loadbalancer .

# Expose LoadBalancer port
EXPOSE 8080

# Start the server
CMD ["./loadbalancer"]
