# Use the official Golang image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application
CMD ["./main"]
