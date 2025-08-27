# ---- Build Stage ----
FROM golang:1.24.5 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (better for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app
RUN go build -o server main.go

# ---- Run Stage ----
FROM gcr.io/distroless/base-debian12

# Set working directory
WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /app/server .

# Expose port (change if your app runs on another port)
EXPOSE 8080

# Run the binary
CMD ["./server"]