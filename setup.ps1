# Abimba Custom Bot - Bulletproof Setup
$ErrorActionPreference = "Stop"

# 1. Dependency Checks
Write-Host "Checking Prerequisites..." -ForegroundColor Cyan
try { $go = go version; Write-Host "$go" -ForegroundColor Green } catch { Write-Host "Go missing"; exit 1 }
try { $py = python --version; Write-Host "$py" -ForegroundColor Green } catch { Write-Host "Python missing"; exit 1 }

# 2. .env Setup
if (-not (Test-Path ".env")) {
    Set-Content -Path ".env" -Value "BOT_TOKEN=token_here`nMY_ID=id_here"
    Write-Host ".env created" -ForegroundColor Green
}

# 3. Python Environment
$vpath = "scripts\.venv"
$vpip = "$vpath\Scripts\pip.exe"
if (-not (Test-Path $vpath)) { 
    Write-Host "Creating venv..."
    python -m venv $vpath 
}
Write-Host "Installing libraries..."
& $vpip install --upgrade pip -q
& $vpip install -q pyautogui pyperclip win-toasts

# 4. Go Build
Write-Host "Building Binaries..." -ForegroundColor Cyan
go mod tidy
if (Test-Path "build.ps1") { 
    & ".\build.ps1" 
} else {
    go build -o dispatcher.exe -ldflags="-H windowsgui" ./cmd/dispatcher/main.go
    go build -o executor.exe -ldflags="-H windowsgui" ./cmd/executor/main.go
}

Write-Host "Checking for FFmpeg..." -ForegroundColor Yellow
if (-not (Get-Command ffmpeg -ErrorAction SilentlyContinue)) {
    Write-Host "FFmpeg not found. Installing via Winget..." -ForegroundColor Gray
    winget install "FFmpeg (Essentials Build)" --source winget
    Write-Host "✓ FFmpeg installed. You may need to restart your terminal." -ForegroundColor Green
} else {
    Write-Host "✓ FFmpeg is already installed." -ForegroundColor Green
}

# 5. Verify and Finish
if (Test-Path "dispatcher.exe") {
    Write-Host "SUCCESS: dispatcher.exe ready" -ForegroundColor Green
} else {
    Write-Host "FAILURE: Build failed" -ForegroundColor Red
    exit 1
}

Write-Host "----------------------------------------"
Write-Host "1. Edit .env"
Write-Host "2. Run .\dispatcher.exe"
Write-Host "3. Add shortcut to shell:startup"