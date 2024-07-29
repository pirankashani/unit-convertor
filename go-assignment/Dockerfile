# Use the official Golang image as a parent image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY *.go ./

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the application when the container starts
CMD ["./main"]