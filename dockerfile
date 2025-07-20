# Use an official Golang image as a base image for the build phase
FROM golang:1.20-alpine AS builder

# Set environment variables
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Create an app directory
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app for production
RUN go build -o payslips ./main.go

# Use a lightweight Alpine image for the final stage
FROM alpine:latest

# Set the working directory in the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/payslips .

# Expose port 3000
EXPOSE 3000

# Command to run the executable
CMD ["./payslips"]
