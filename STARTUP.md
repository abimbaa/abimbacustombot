# Windows Startup Setup

To run Abimba Custom Bot automatically when Windows starts:

## Method: Startup Folder (Recommended)

1. **Open Startup Folder:**
   - Press `Win + R`
   - Type `shell:startup`
   - Press Enter

2. **Create Shortcut:**
   - Right-click → New → Shortcut
   - Target: `C:\Path\To\abimbaCustomBot\dispatcher.exe`
   - Name: `Abimba Bot`
   - Click Create

3. **Done!**
   - Bot will launch automatically on next restart

## Verify

After restarting Windows, the bot should:
- Start automatically
- Show as a process in Task Manager
- Accept commands from Telegram

If it doesn't start:
- Verify dispatcher.exe exists and works: `.\dispatcher.exe`
- Check .env file has correct BOT_TOKEN and MY_ID
- Ensure shortcut points to correct path
