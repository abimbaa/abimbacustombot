# Command Quick Reference

Quick reference for all Abimba Custom Bot commands. For detailed help on any command, send `/help <command>` to the bot.

## Most Used Commands

| Command | What it does | Example |
|---------|-------------|---------|
| `/ping` | Check if bot is alive | `/ping` |
| `/help` | Show all commands | `/help` or `/help ss` |
| `/status` | CPU/RAM usage | `/status` |
| `/ss` | Take screenshot(s) | `/ss` or `/ss 3 2` |
| `/cmd` | Run shell command | `/cmd dir` |

## Screenshots & Media

| Command | Arguments | Example |
|---------|-----------|---------|
| `/ss` | `[count] [delay_ms]` | `/ss 5 1` (5 screenshots, 1s delay) |
| `/cam` | `[count] [delay_ms]` | `/cam` (1 webcam picture) |

## File Operations

| Command | Arguments | Example |
|---------|-----------|---------|
| `/ls` | `[path]` | `/ls C:\Users\Downloads` |
| `/get` | `<filepath>` | `/get C:\Users\Downloads\photo.jpg` |
| `Send file` | (via Telegram) | Downloads to PC Downloads folder |

## System & Process

| Command | Arguments | Example |
|---------|-----------|---------|
| `/ps` | `[filter]` | `/ps chrome` (find processes) |
| `/kill` | `<name_or_pid>` | `/kill chrome` or `/kill 1234` |
| `/cmd` | `<command>` | `/cmd ipconfig` |

## Applications

| Command | Arguments | Example |
|---------|-----------|---------|
| `/run` | `<app_name>` | `/run discord` |

## Power Control

| Command | What it does | Delay |
|---------|-------------|-------|
| `/shutdown` | Turn off PC | 5 seconds |
| `/restart` | Reboot PC | 5 seconds |
| `/lock` | Lock Windows | Immediate |
| `/sleep` | Suspend PC | Immediate |

## Audio Control

| Command | Arguments | Example |
|---------|-----------|---------|
| `/volup` | None | `/volup` (+2%) |
| `/voldown` | None | `/voldown` (-2%) |
| `/vol` | `<0-100>` | `/vol 50` |
| `/mute` | None | `/mute` |
| `/pause` | None | `/pause` |
| `/next` | None | `/next` |
| `/prev` | None | `/prev` |

## ⌨Input Control

| Command | Arguments | Example |
|---------|-----------|---------|
| `/type` | `<text>` | `/type Hello world` |
| `/type` | (with keys) | `/type Hello $enter$ world $enter$` |
| `/clip get` | None | Get clipboard content |
| `/clip set` | `<text>` | `/clip set copy this text` |

**Special Keys for `/type`:**
- `$enter$` - Enter key
- `$tab$` - Tab key
- `$shift+a$` - Shift+A
- `$ctrl+c$` - Ctrl+C
- `$alt+f4$` - Alt+F4
- most hotkeys and special keys work!

## Notes & Plans

| Command | Arguments | Example |
|---------|-----------|---------|
| `/note` | `<title> <body>` | `/note Meeting Follow-up I need to...` |
| `/notes` | None | Show all notes |
| `/plan` | `<text>` | `/plan Buy groceries` |
| `/plans` | None | List all plans |
| `/pdel` | `<number>` | `/pdel 3` (delete plan #3) |

## Audio & Transcription

| Command | Arguments | Example |
|---------|-----------|---------|
| `/stt` | `<filepath>` | `/stt C:\audio.mp3` |
| Send voice/video message | (automatic) | Voice/video messages auto-transcribed |
| Send audio file | (automatic) | Audio files auto-transcribed |

## Macros

| Command | Arguments | Example |
|---------|-----------|---------|
| `/createmacro` | `<name>` then commands | See "Creating Macros" |
| `/macros` | None | List all macros |
| `/macro` | `<name>` | `/macro startup` |
| `/delmacro` | `<name>` | `/delmacro startup` |

## Timing

| Command | Arguments | Example |
|---------|-----------|---------|
| `/delay` | `<time>` | `/delay 5s` or `/delay 1m` |

**Time Formats:**
- `5s` - 5 seconds
- `1m` - 1 minute
- `2h` - 2 hours

## Creating Macros

Send as **one message** with newlines:

```
/createmacro startup
/run discord
/run spotify
/delay 5s
/run vscode
```

Then use: `/macro startup` to run all commands at once.

## Help

| Command | Result |
|---------|--------|
| `/help` | List all commands |
| `/help ss` | Details for `/ss` command |
| `/help <command>` | Details for any command |

## Usage Tips

### Chain Commands
Use `/delay` to create delays between commands in macros:
```
/createmacro delayed
/run notepad
/delay 2s
/type Hello!
/delay 1s
/type I just opened notepad
```

### Multi-Monitor Screenshots
`/ss 3 5` - Takes 3 screenshots from all monitors with 5s delay between them

### Search for Files
`/get Downloads` - If file doesn't have full path, searches Downloads folder

### Run Multiple Commands
Use `/cmd &&` to chain commands:
`/cmd echo hello && pause`

### Kill by Name or PID
Both work:
- `/kill firefox` - Kill by process name
- `/kill 5432` - Kill by process ID

### Run an app
- `/run msedge` - Searches for msedge in system folders or in additional folders you specified in `search_folders.txt`

### Clipboard Tricks
```
/clip set https://example.com
/type $ctrl+v$
```

## Common Mistakes

| Mistake | Solution |
|---------|----------|
| Command hangs | Wait 5 minutes or kill bot and restart |
| "File not found" | Use full path: `C:\Users\...` |
| Macro didn't save | Check `/macros` to verify |
| Screenshot failed | Ensure no display scaling issues |
| Python error | Reinstall: `pip install -r scripts/requirements.txt` |

## Common Workflows

### Workflow 1: Morning Startup
```
/macro startup
# (launches Discord, Spotify, VSCode)
```

### Workflow 2: Quick Screenshot
```
/ss
# Get latest screenshot from all monitors
```

### Workflow 3: Note Taking
```
/note Meeting Notes
Remember to follow up on Q3 targets
```

### Workflow 4: Remote File Transfer
```
/get C:\Projects\report.pdf
# Sends file to your phone
```

### Workflow 5: Monitor System
```
/status
# Check CPU/RAM
```

### Workflow 6: Transcribe Voice
```
Send voice message
# Automatically transcribed and sent back
```

**Last Updated:** May 11, 2026  
**Version:** 1.0.0
