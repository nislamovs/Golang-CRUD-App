# ---- Run Stage ----
FROM gcr.io/distroless/base-debian12

# Set working directory
WORKDIR /app

# Copy compiled binary from builder
COPY ./bin .

# Expose port (change if your app runs on another port)
EXPOSE 8080

# Mount the database folder at runtime
VOLUME ["/app/data"]

# Run the binary
CMD ["./golang-crud-app"]