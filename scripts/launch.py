import sys
import subprocess
import os

# Windows-specific flag to suppress error dialogs
SEM_FAILCRITICALERRORS = 0x0001
SEM_NOGPFAULTERRORBOX = 0x0002

def main():
    if len(sys.argv) < 2:
        return

    app_name = " ".join(sys.argv[1:]).lower()
    exe_name = app_name if app_name.endswith(".exe") else f"{app_name}.exe"

    # Tell Windows NOT to show "This file does not exist" or "How to open this" dialogs
    if os.name == 'nt':
        import ctypes
        ctypes.windll.kernel32.SetErrorMode(SEM_FAILCRITICALERRORS | SEM_NOGPFAULTERRORBOX)

    try:
        # 1. Try launching via PowerShell with errors suppressed
        # Use 'SilentlyContinue' to prevent PowerShell itself from throwing a visible error
        ps_cmd = f"Start-Process '{app_name}' -ErrorAction SilentlyContinue"
        result = subprocess.run(
            ["powershell", "-Command", ps_cmd],
            creationflags=subprocess.CREATE_NO_WINDOW,
            capture_output=True
        )

        if result.returncode == 0:
            print(f"Launched {app_name}", flush=True)
            return

        # 2. Search logic if the quick start fails
        search_paths = [
            os.environ.get("ProgramFiles"),
            os.environ.get("ProgramFiles(x86)"),
            os.environ.get("AppData"),
            os.environ.get("LocalAppData"),
            r"E:\Programs",
            r"E:\Games",
            r"E:\SteamLibrary\steamapps\common"
        ]

        for base_path in search_paths:
            if not base_path or not os.path.exists(base_path): 
                continue
            
            # Record the depth of the base path
            base_depth = base_path.count(os.sep)
            
            for root, dirs, files in os.walk(base_path):
                # Calculate current depth relative to base_path
                current_depth = root.count(os.sep) - base_depth
                
                # Check for the exe (case-insensitive)
                current_files_lower = [f.lower() for f in files]
                if exe_name in current_files_lower:
                    full_path = os.path.join(root, exe_name)
                    
                    subprocess.Popen(
                        [full_path], 
                        cwd=root, # Crucial for games: runs them from their own folder
                        start_new_session=True,
                        creationflags=subprocess.CREATE_NO_WINDOW
                    )
                    print(f"Found and launched: {full_path}", flush=True)
                    return
                
                # Increase depth to 2 to reach SteamLibrary/common/GameFolder/game.exe
                if current_depth >= 2:
                    del dirs[:]

        print(f"'{app_name}' not found.", flush=True)

    except Exception as e:
        print(f"ERROR: {str(e)}", flush=True)

if __name__ == "__main__":
    main()