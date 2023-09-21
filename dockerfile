# Start from the official Golang base image
FROM golang:1.19 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o suricata-listener

# Start a new stage from scratch
FROM alpine:latest  

# Copy the compiled application from the builder stage
COPY --from=builder /app/suricata-listener /suricata-listener

# Copy the policies directory from the builder stage
COPY --from=builder /app/policies /policies

# Run the application
CMD ["/suricata-listener"]
