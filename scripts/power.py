import sys
import os

def main():
    if len(sys.argv) < 2:
        return

    cmd = sys.argv[1].lower()

    try:
        if cmd == "shutdown":
            print("Shutting down PC in 5 seconds...", flush=True)
            os.system("shutdown /s /t 5")
            
        elif cmd == "restart":
            print("Restarting PC in 5 seconds...", flush=True)
            os.system("shutdown /r /t 5")
            
        elif cmd == "lock":
            print("Locking workstation...", flush=True)
            os.system("rundll32.exe user32.dll,LockWorkStation")
            
        elif cmd == "sleep":
            print("Putting PC to sleep...", flush=True)
            # Note: If hibernation is enabled on your PC, this might hibernate instead of sleep.
            os.system("rundll32.exe powrprof.dll,SetSuspendState 0,1,0")

    except Exception as e:
        print(f"Error: {str(e)}", flush=True)

if __name__ == "__main__":
    main()