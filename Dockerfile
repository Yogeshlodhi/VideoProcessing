# official Golang image to build the application
FROM golang:alpine AS builder

# Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o process_video main.go

# Use a smaller base image for the final container
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder image
COPY --from=builder /app/process_video .

# Command to run the executable
ENTRYPOINT ["./process_video"]
