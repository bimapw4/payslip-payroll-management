# Use Golang Alpine as base image
FROM golang:1.21-alpine

# Install required packages
RUN apk add --no-cache git

# Set environment
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod & go.sum first (cache layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the app
RUN go build -o payslips ./main.go

# Expose the app port
EXPOSE 3000

# Run the executable
CMD ["./payslips"]
