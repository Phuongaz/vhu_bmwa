FROM golang:1.21-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Create logs directory with proper permissions
RUN mkdir -p /app/logs && chmod 777 /app/logs

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./main"] 