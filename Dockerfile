# Stage 1: Build stage
FROM golang:1.23-alpine AS builder

# Install necessary dependencies
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod  ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Ensure the binary is built for Linux and the right architecture (disable CGO)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/secret-server

# Stage 2: Final stage - using a minimal Alpine image
FROM alpine:alpine

WORKDIR /app/

# Copy the built application from the builder stage
COPY --from=builder /app/secret-server .

COPY ./static ./static

# Ensure the binary is executable
RUN chmod +x ./secret-server

ENV VAULT_IPS="41.235.69.116,127.0.0.1"
# Expose the port the app will run on
EXPOSE 8080

# Run the Go application
CMD ["./secret-server"]