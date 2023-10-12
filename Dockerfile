# Use an official Go runtime as the parent image
FROM golang:1.16

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 for the app to listen on
EXPOSE 8080

# Run the binary program when the container starts
CMD ["./main"]
