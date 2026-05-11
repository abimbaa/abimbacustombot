# Abimba Custom Bot

A Telegram bot for remote Windows PC automation. Execute commands, manage files, capture screenshots, and control media directly from Telegram.

## Quick Start

### Prerequisites
- Windows 10/11
- Go 1.26.3+ ([download](https://golang.org/dl/))
- Python 3.10+ ([download](https://www.python.org/downloads/)) - Check "Add Python to PATH"
- FFmpeg (Required for audio tasks)

### Setup release version (2 minutes)
1. Download the latest release.
2. Unzip it and find `setup.ps1`. Run it with powershell
3. Set up the `.env` file with your ID and bot token. You can see an example in `.env.example`
4. Run `dispatcher.exe`
Done!


### Setup developer version (5 minutes) 

```powershell
# 1. Clone & setup
git clone https://github.com/abimbaa/abimbaCustomBot.git
cd abimbaCustomBot
powershell -ExecutionPolicy Bypass -File .\setup.ps1

# 2. Configure credentials in .env
# Get BOT_TOKEN from @BotFather on Telegram
# Get MY_ID from @userinfobot on Telegram
# Rename .env.example to .env and fill in your details

# 3. Run bot
.\dispatcher.exe
`
### Add to Windows Startup
Copy `dispatcher.exe` shortcut to the Windows Startup folder:
```powershell
# Press Win+R, type: shell:startup
# Create shortcut to dispatcher.exe there
# Bot will launch automatically on startup
```

## Features (40+ Commands)

**System:** `/ping`, `/status`, `/help`  
**Files:** `/ls`, `/get`, `/run`, (send files)  
**Capture:** `/ss`, `/cam`  
**Control:** `/cmd`, `/ps`, `/kill`, `/type`, `/clip`  
**Power:** `/shutdown`, `/restart`, `/lock`, `/sleep`  
**Media:** `/volup`, `/voldown`, `/mute`, `/pause`, `/next`  
**Organize:** `/note`, `/plan`, `/macros`  
**Audio:** `/stt`, (voice messages auto-transcribed)  

**Full reference:** See `COMMANDS.md`

## Project Structure

```
abimbaCustomBot/
├── cmd/                    # Go source code
│   ├── dispatcher/main.go  # Telegram bot
│   └── executor/main.go    # Command executor
├── scripts/                # Python utility scripts (15+)
├── data/                   # User data (macros, notes, plans)
├── README.md               # This file
├── COMMANDS.md             # Command reference
├── DEVELOPMENT.md          # Architecture & development
├── SECURITY.md             # Security guidelines
├── CONTRIBUTING.md         # How to contribute
├── dispatcher.exe          # Built bot (production)
└── executor.exe            # Built executor (production)
```

## Setting additional folders

The `/run` command searches for executable files inside the default environment folders (appdata, programfiles, etc) with depth = 2. To add other folders to this search add a `search_folders.txt` inside the `data/` folder. In this text file write the folder paths in this format:
```
E:\Programs
E:\Games
E:\SteamLibrary\steamapps\common
...
```

## Troubleshooting

| Issue | Solution |
|-------|----------|
| Bot not responding | Check `.env` has correct TOKEN and ID, verify internet |
| Python error | Run `pip install -r scripts/requirements.txt` |
| Build fails | Run `.\build.ps1 -Clean` |
| Command hangs | Commands timeout after 5 minutes |

## Documentation

- **`COMMANDS.md`** - All commands with examples
- **`DEVELOPMENT.md`** - Architecture and how to add features
- **`SECURITY.md`** - Security best practices
- **`CONTRIBUTING.md`** - How to contribute

## Common Tasks

**Create a macro to launch apps:**
```
/createmacro startup
/run discord
/run spotify
```
Then use: `/macro startup`

**Transfer a file:**
```
/get C:\Users\Downloads\file.pdf
```

**Monitor system:**
```
/status
```

## Security

⚠️ **Important:** This gives complete PC control via Telegram
- Keep `.env` secret (credentials inside)
- Only your Telegram ID can use the bot
- Review `SECURITY.md` for best practices

## License

MIT License - See `LICENSE` file

## Need Help?

- **Commands:** Send `/help` to bot or check `COMMANDS.md`
- **Setup issues:** See troubleshooting above
- **Development:** Read `DEVELOPMENT.md`
- **Contributing:** Read `CONTRIBUTING.md`

---

**Status:** Production Ready ✓  
**Version:** 1.0.0  
**Platform:** Windows 10/11
