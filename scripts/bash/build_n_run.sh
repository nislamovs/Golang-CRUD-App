#!/bin/bash

cd ../../

# Build image
docker build -t golang-crud-app .

# Run container
docker run -p 8080:8080 golang-crud-app

# ####################################################################3

# Build & start container
docker-compose up --build -d

# View logs
docker-compose logs -f

# Stop container
docker-compose down