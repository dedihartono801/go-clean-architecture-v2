# Start from golang base image
FROM golang:alpine as builder

ARG APP_NAME=myapp

# Set label to the builder image, and then delete this image using command: 
# docker image prune --filter label=stage=builder
LABEL stage=builder

# Install git
RUN apk update && apk add --no-cache git

# Install required tools
RUN apk add build-base

# Set working directory inside the container
WORKDIR /app

# Copy go mod and go sum files
COPY go.mod go.sum ./

# Download required dependencies
RUN go mod download 


# Copy entire source code from current directory to the working directory inside container
COPY . . 

# Vendoring go mod
RUN go mod vendor

# Build the Go app into binary file
RUN  go build -o ${APP_NAME} -v ./cmd/http

RUN go get gopkg.in/DataDog/dd-trace-go.v1/ddtrace

# Start a new stage using alpine linux
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set working directory for the new stage
WORKDIR /root/

# Copy the binary file from previous stage
COPY --from=builder /app/${APP_NAME} .

# Copy the environment variable 
COPY --from=builder /app/.env .

# Expose application port
EXPOSE 5001

# Command to start the application
CMD ["./myapp"]