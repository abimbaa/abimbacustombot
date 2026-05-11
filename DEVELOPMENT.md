# Development Guide

## Architecture Overview

### Components

```
Telegram User
    ↓
  Telegram API
    ↓
Dispatcher (Go)  ←→ Executor (Go)  →  Python Scripts
    ↓
  Output
    ↓
Telegram API → User
```

**Dispatcher** (`cmd/dispatcher/main.go`)
- Listens for Telegram messages
- Only authorizes messages from your Telegram ID
- Routes commands to Executor
- Handles file downloads/uploads
- Processes audio for transcription

**Executor** (`cmd/executor/main.go`)
- Receives commands from Dispatcher
- Calls appropriate Python scripts or executes system commands
- Prints output to stdout (which Dispatcher captures)
- Returns results to Dispatcher

**Python Scripts** (`scripts/*.py`)
- Handle specific functionality (screenshots, clipboard, etc.)
- Print results to stdout
- Executor captures this output and passes it back to Dispatcher

## How Output is Relayed to Telegram

### 1. Command Execution Flow

```
User sends: /ss
    ↓
Dispatcher receives message
    ↓
Dispatcher calls: executor.exe "/ss"
    ↓
Executor finds handler for /ss
    ↓
Executor calls: python scripts/screenshot.py
    ↓
Python script prints: PHOTO:screenshot_0_mon_1.png
    ↓
Executor outputs to stdout
    ↓
Dispatcher reads stdout (via pipe)
    ↓
Dispatcher parses output & sends photo to Telegram
```

### 2. Output Parsing in Dispatcher

The Dispatcher (`cmd/dispatcher/main.go`) reads output line-by-line and interprets special prefixes:

**File Output:**
```go
if strings.HasPrefix(cleanLine, "FILE:") {
    path := strings.TrimSpace(strings.TrimPrefix(cleanLine, "FILE:"))
    // Send file to Telegram
    doc := &telebot.Document{
        File:     telebot.FromDisk(path),
        FileName: filepath.Base(path),
    }
    c.Send(doc)
}
```

**Photo Output:**
```go
if strings.HasPrefix(cleanLine, "PHOTO:") {
    path := strings.TrimSpace(strings.TrimPrefix(cleanLine, "PHOTO:"))
    // Send photo to Telegram
    c.Send(&telebot.Photo{File: telebot.FromDisk(path)})
}
```

**Error Output:**
```go
if strings.HasPrefix(cleanLine, "ERROR:") {
    msg := "ERROR: " + strings.TrimPrefix(cleanLine, "ERROR:")
    // Send error message to Telegram
    c.Send(msg)
}
```

**Regular Text Output:**
```go
default:
    // Accumulate regular text in buffer
    textBuffer.WriteString(line + "\n")
    // Send when buffer reaches 1000 chars or command ends
```

### 3. Python Script Output Format

Python scripts should follow this format for proper relay:

```python
# Send text (will appear in Telegram)
print("Operation successful", flush=True)

# Send a file
print("FILE:C:\\path\\to\\file.pdf", flush=True)

# Send a photo/screenshot
print("PHOTO:screenshot_0_mon_1.png", flush=True)

# Send an error
print("ERROR: Something went wrong", flush=True)
```

**Important:** Use `flush=True` to ensure output is sent immediately.

### 4. Output Flow Example

**User sends:** `/ss`

**Dispatcher** runs:
```go
cmd := exec.CommandContext(ctx, "./executor.exe", "/ss")
stdout, _ := cmd.StdoutPipe()
scanner := bufio.NewScanner(stdout)

for scanner.Scan() {
    line := scanner.Text()
    
    if strings.HasPrefix(line, "PHOTO:") {
        // Extract path and send to Telegram
        path := strings.TrimPrefix(line, "PHOTO:")
        c.Send(&telebot.Photo{File: telebot.FromDisk(path)})
    }
}
```

**Executor** runs:
```go
func handleScreenshot(args string) {
    runProcess(VENV_PYTHON, "./scripts/screenshot.py", amountStr, delayStr)
}
```

**Python script** (`screenshot.py`) outputs:
```python
print("PHOTO:screenshot_0_mon_1.png", flush=True)
print("PHOTO:screenshot_0_mon_2.png", flush=True)
```

**Dispatcher** receives these lines and sends each as a photo to Telegram.

## Adding a New Command

### 1. Create Python Script (if needed)

File: `scripts/mycommand.py`
```python
import sys

def main():
    if len(sys.argv) < 2:
        print("ERROR: Missing argument", flush=True)
        return
    
    # Do something
    result = "Success"
    
    # Output result
    print(result, flush=True)

if __name__ == "__main__":
    main()
```

### 2. Register Handler in Executor

File: `cmd/executor/main.go`

Add to the registry:
```go
func init() {
    registry = map[string]CommandFunc{
        "/mycommand": handleMyCommand,
        // ... other commands
    }
}

func handleMyCommand(args string) {
    runProcess(VENV_PYTHON, "./scripts/mycommand.py", args)
}
```

Or if it's a system command:
```go
func handleMyCommand(args string) {
    cmd := exec.Command("powershell", "-Command", args)
    output, _ := cmd.CombinedOutput()
    fmt.Println(string(output))
}
```

### 3. Rebuild

```powershell
.\build.ps1
```

### 4. Add Help Text

In `handleHelp()` function, add to the `docs` map:
```go
"/mycommand": {"Short description", "Usage: /mycommand\nLong description"},
```

## Output Types

| Type | Format | Example |
|------|--------|---------|
| Text | Plain text | `print("Hello")` |
| File | `FILE:path` | `print("FILE:C:\\file.pdf")` |
| Photo | `PHOTO:path` | `print("PHOTO:image.png")` |
| Error | `ERROR:msg` | `print("ERROR: Failed")` |

## Command Timeout

All commands have a **5-minute timeout**. If a command takes longer:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()

cmd := exec.CommandContext(ctx, ...)
```

If timeout is exceeded:
- Command is killed
- Message sent: "Command timed out after 5 minutes."

## File Paths

**Important:** All file paths should be absolute or relative to the bot directory.

```python
# Good
print("FILE:C:\\Users\\Downloads\\file.pdf")
print("FILE:" + os.path.abspath("screenshot.png"))

# Avoid (may not work)
print("FILE:screenshot.png")
```

## Streaming Large Output

If output is larger than 1000 characters, it's automatically split:

```go
if textBuffer.Len() >= 1000 {
    flushBuffer()  // Send to Telegram
    textBuffer.Reset()
}
```

Each chunk is sent as a separate message.

## Testing

To test your new command:

```powershell
# Build
.\build.ps1

# Run bot
.\dispatcher.exe

# Send command from Telegram
/mycommand arg1
```

Check Telegram for output. Use `/help mycommand` to verify it appears in help text.

## Debugging

If output isn't appearing in Telegram:

1. **Check Python script** - Run it directly:
   ```powershell
   python scripts/mycommand.py arg1
   ```

2. **Check Executor** - Run it directly:
   ```powershell
   executor.exe "/mycommand arg1"
   ```

3. **Check stdout** - Ensure `print(..., flush=True)` is used

4. **Check authorization** - Verify `.env` has correct MY_ID

5. **Check format** - Ensure output follows FILE:, PHOTO:, or ERROR: format where applicable

## Performance Tips

1. **Use Python for system operations** - More reliable than shell commands
2. **Minimize stdout** - Large outputs slow down relay
3. **Batch operations** - Multiple commands in one call when possible
4. **Use absolute paths** - Avoid path resolution delays

## Error Handling

Always catch errors in Python scripts:

```python
try:
    # Do something
    pass
except Exception as e:
    print(f"ERROR: {str(e)}", flush=True)
    sys.exit(1)
```

The Dispatcher will catch this and relay it to Telegram.

---

**Architecture Version:** 1.0.0  
**Last Updated:** May 11, 2026
