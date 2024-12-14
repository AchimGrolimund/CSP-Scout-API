# Stage 1: Build
FROM golang:1.23 as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app with ldflags
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Build=$(date +%FT%T%z) -X main.Version=$(git describe --tags --always)" -o main ./cmd/csp-scout-api/main.go

# Stage 2: Run
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside
EXPOSE 8080

# Command to run the executable
CMD ["./main"]docker build -t csp-scout-api -f docker/Dockerfile .