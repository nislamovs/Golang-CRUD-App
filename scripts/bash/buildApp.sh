#!/bin/bash

# Exit immediately if a command fails
set -e

# Configuration
APP_NAME="golang-crud-app"          # Name of the output binary
OUTPUT_DIR="./bin"        # Directory to store the binary
GOOS=${GOOS:-$(go env GOOS)}  # Target OS (default: host OS)
GOARCH=${GOARCH:-$(go env GOARCH)}  # Target architecture (default: host arch)

# Create output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Build the application
echo "Building $APP_NAME for $GOOS/$GOARCH..."
GOOS=$GOOS GOARCH=$GOARCH go build -o "$OUTPUT_DIR/$APP_NAME" .

echo "Build completed: $OUTPUT_DIR/$APP_NAME"