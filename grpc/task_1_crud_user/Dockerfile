# Use the official Golang image as the base
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY . .

# Build the Go application
RUN go build -o main cmd/web/main.go

# Expose the port the API will run on
EXPOSE 8080

# Run the application when the container starts
CMD ["/app/main"]