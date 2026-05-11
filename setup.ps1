# Abimba Custom Bot - Clean Setup
$ErrorActionPreference = "Stop"
Set-Location $PSScriptRoot

$dispatcher = "dispatcher.exe"
$executor = "executor.exe"
$vpath = "scripts\.venv"
$vpython = "$vpath\Scripts\python.exe"

Write-Host "--- 1. Checking Prerequisites ---" -ForegroundColor Cyan

try { 
    $go = go version
    Write-Host "OK: $go" -ForegroundColor Green 
} catch { 
    Write-Host "ERROR: Go missing." -ForegroundColor Red; exit 1 
}

try { 
    $py = python --version
    Write-Host "OK: $py" -ForegroundColor Green 
} catch { 
    Write-Host "ERROR: Python missing." -ForegroundColor Red; exit 1 
}

# 2. .env Setup
if (-not (Test-Path ".env")) {
    "BOT_TOKEN=token_here`nMY_ID=id_here" | Out-File -FilePath ".env" -Encoding utf8
    Write-Host "OK: .env created" -ForegroundColor Green
}

# 3. Python Environment
Write-Host "`n--- 2. Configuring Python ---" -ForegroundColor Cyan
if (-not (Test-Path $vpath)) { 
    Write-Host "Creating virtual environment..."
    python -m venv $vpath 
}

Write-Host "Installing dependencies..."
& $vpython -m pip install --upgrade pip
& $vpython -m pip install -r scripts/requirements.txt

# 4. Smart Build
Write-Host "`n--- 3. Building Binaries ---" -ForegroundColor Cyan
if (-not (Test-Path $dispatcher) -or -not (Test-Path $executor)) {
    Write-Host "Building now..." -ForegroundColor Yellow
    go mod tidy
    go build -o $dispatcher ./cmd/dispatcher/main.go
    go build -o $executor ./cmd/executor/main.go
} else {
    Write-Host "OK: Binaries exist. Skipping build." -ForegroundColor Green
}

# 5. Final Output & Startup Logic
Write-Host "`n--- 4. Final Verification ---" -ForegroundColor Cyan
if (Test-Path $dispatcher) {
    Write-Host "DONE: $dispatcher is ready." -ForegroundColor Green
}

Write-Host "`n----------------------------------------"
Write-Host "SETUP COMPLETE" -ForegroundColor Green
Write-Host "1. Configure your credentials in .env"
Write-Host "2. Run .\$dispatcher to start the bot"
Write-Host "----------------------------------------"

# 6. Automated Startup Folder Helper
Write-Host "`nTo make the bot start with Windows:" -ForegroundColor Cyan
$choice = Read-Host "Would you like to open the Startup folder now? (y/n)"
if ($choice -eq 'y') {
    Write-Host "Opening Startup folder..." -ForegroundColor Yellow
    Write-Host "INSTRUCTIONS: Right-click '$dispatcher' -> Create Shortcut -> Drag shortcut to the opened folder." -ForegroundColor Gray
    explorer shell:startup
}

Write-Host "`nExiting setup..."