# Start from a golang image
FROM golang:built
# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Set the working directory to the 'data' directory
WORKDIR /app/data

# Copy the data files to the container
COPY data/ .

# Set the working directory back to the root directory
WORKDIR /app

# Copy the rest of the application files to the container
COPY . .

# Build the application
RUN go build -o image-management-service ./cmd

# Expose the port on which the application will run
EXPOSE 8080

# Start the application
CMD ["./image-management-service"]