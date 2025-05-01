# Use an official Go image as the base image
FROM golang:1.24-bookworm AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/* \
	&& go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o gojira .

FROM debian:bookworm-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/gojira .

# Expose a port if the application listens on one (optional)
# EXPOSE 8080

# Set the default command to run the application

CMD ["./gojira"]