# Start from the official Go image
FROM golang:1.23-alpine

# Install necessary packages
RUN apk add --no-cache git gcc libc-dev

# Install Goose CLI
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Set working directory in the container
WORKDIR /app

# Copy Go project files
COPY . .

# Build the Go application (assuming main.go is your entry point)
RUN go build -o app .

# Expose any necessary ports (optional, for example 8080 for a web service)
EXPOSE 8080

# Set default entrypoint to the compiled application
ENTRYPOINT ["./app"]
