# Use an official Go runtime as the base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the application listens on
EXPOSE 8080

# Set the entry point command to run the compiled binary
CMD ["./main"]
