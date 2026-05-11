$ErrorActionPreference = "Stop"

# Binary names
$DISPATCHER = "dispatcher.exe"
$EXECUTOR = "executor.exe"

Write-Host "--- Starting Build ---" -ForegroundColor Cyan

# Build Dispatcher
Write-Host "Building Dispatcher..."
go build -o $DISPATCHER -ldflags="-H windowsgui" ./cmd/dispatcher/main.go
if (Test-Path $DISPATCHER) {
    $dSize = (Get-Item $DISPATCHER).Length / 1MB
    Write-Host "Done: $DISPATCHER ($('{0:N1}' -f $dSize) MB)" -ForegroundColor Green
}

# Build Executor
Write-Host "Building Executor..."
go build -o $EXECUTOR -ldflags="-H windowsgui" ./cmd/executor/main.go
if (Test-Path $EXECUTOR) {
    $eSize = (Get-Item $EXECUTOR).Length / 1MB
    Write-Host "Done: $EXECUTOR ($('{0:N1}' -f $eSize) MB)" -ForegroundColor Green
}

Write-Host "--- Build Complete ---" -ForegroundColor Cyan