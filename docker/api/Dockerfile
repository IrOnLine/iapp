FROM golang:alpine AS builder

# Install dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download 

# Copy source code
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Production image 
FROM alpine

# Install ca certs
RUN apk add --no-cache ca-certificates

# Copy binary from builder image
COPY --from=builder /app/main /app/main

# Export port
EXPOSE 8000 

# Command to run
CMD ["/app/main"]