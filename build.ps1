# 1. Define filenames
$DISPATCHER = "dispatcher.exe"
$EXECUTOR = "executor.exe"

Write-Host "--- Starting Clean Build ---" -ForegroundColor Cyan

# 2. Kill running processes (Silent fail if not running)
Write-Host "Stopping background processes..."
taskkill /f /im $DISPATCHER /t 2>$null
taskkill /f /im $EXECUTOR /t 2>$null
taskkill /f /im cmd.exe /t 2>$null

# 3. Clean up temporary and old files
Write-Host "Cleaning directory..."
if (Test-Path "$DISPATCHER~") { Remove-Item "$DISPATCHER~" }
if (Test-Path ".cwd") { Remove-Item ".cwd" }
if (Test-Path $DISPATCHER) { Remove-Item $DISPATCHER }
if (Test-Path $EXECUTOR) { Remove-Item $EXECUTOR }

# 4. Rebuild binaries with hidden flags
Write-Host "Building Executor..." -ForegroundColor Yellow
go build -o $EXECUTOR ./cmd/executor

Write-Host "Building Dispatcher..." -ForegroundColor Yellow
go build -o $DISPATCHER ./cmd/dispatcher

# 5. Success Check
if ($? ) {
    Write-Host "Build Successful! Run ./$DISPATCHER to start." -ForegroundColor Green
} else {
    Write-Host "Build Failed. Check your Go code for errors." -ForegroundColor Red
}