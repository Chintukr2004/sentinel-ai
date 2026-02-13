# Use official Go image
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the app
RUN go build -o collector cmd/main.go

# Run it
CMD ["./collector"]
