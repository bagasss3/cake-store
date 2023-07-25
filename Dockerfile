FROM golang:1.18

# Set the working directory for container
WORKDIR /app

# Copy the Go source code to the container's dir
COPY . .

# Build the Go application
RUN go build main.go

# Expose the port
EXPOSE 8080

# Run the Go application
CMD ["./main","server"]
