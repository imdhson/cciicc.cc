# ==========================================
# cciicc Server Run Script for Windows
# ==========================================

# Variables
$env:PORT = "8000"
$env:SSL_PORT = "4433"
$env:ENABLE_TLS = "false"
$env:TLS_CERT = "certkey"
$env:TLS_KEY = "key"

Write-Host "Building the Go application..."
go build -o cciicc.exe main.go

if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed! Please check the Go source code." -ForegroundColor Red
    exit 1
}

Write-Host "Build successful." -ForegroundColor Green
Write-Host "Starting the server..."
Write-Host "Environment settings: PORT=$env:PORT, SSL_PORT=$env:SSL_PORT, ENABLE_TLS=$env:ENABLE_TLS"

# Execute the binary
.\cciicc.exe
