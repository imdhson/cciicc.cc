#!/bin/bash

# ==========================================
# cciicc Server Run Script for Linux/macOS
# ==========================================

# Variables
export PORT=8000
export SSL_PORT=4433
export ENABLE_TLS=false
export TLS_CERT="certkey"
export TLS_KEY="key"

echo "Building the Go application..."
go build -o cciicc main.go

if [ $? -ne 0 ]; then
    echo "Build failed! Please check the Go source code."
    exit 1
fi

echo "Build successful."
echo "Starting the server..."
echo "Environment settings: PORT=$PORT, SSL_PORT=$SSL_PORT, ENABLE_TLS=$ENABLE_TLS"

# Execute the binary
./cciicc
