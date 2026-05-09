import sys
import os

def main():
    if len(sys.argv) < 2:
        print("ERROR: Please provide a filename or absolute path.", flush=True)
        return

    arg = " ".join(sys.argv[1:])

    # 1. Check if user provided an absolute path (e.g., C:\folder\file.txt)
    if os.path.isabs(arg):
        if os.path.exists(arg):
            print(f"FILE:{arg}", flush=True)
        else:
            print(f"ERROR: File not found at {arg}", flush=True)
        return

    # 2. Fallback: Search the Downloads folder
    home = os.path.expanduser("~")
    downloads_dir = os.path.join(home, "Downloads")
    filepath = os.path.join(downloads_dir, arg)

    if os.path.exists(filepath):
        print(f"FILE:{filepath}", flush=True)
    else:
        print(f"ERROR: Could not find '{arg}'.", flush=True)

if __name__ == "__main__":
    main()