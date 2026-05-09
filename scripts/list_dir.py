import sys
import os

def main():
    # Default to C:\ if no arguments are passed
    target_dir = " ".join(sys.argv[1:]) if len(sys.argv) > 1 else "C:\\"

    if not os.path.exists(target_dir):
        print(f"ERROR: Path does not exist: {target_dir}", flush=True)
        return
    if not os.path.isdir(target_dir):
        print(f"ERROR: Not a directory: {target_dir}", flush=True)
        return

    try:
        items = os.listdir(target_dir)
        folders, files = [], []
        
        for item in items:
            if os.path.isdir(os.path.join(target_dir, item)):
                folders.append(f"{item}")
            else:
                files.append(f"{item}")
        
        print(f"**{target_dir}**\n", flush=True)
        
        for f in sorted(folders): print(f, flush=True)
        for f in sorted(files): print(f, flush=True)
            

    except PermissionError:
        print("ERROR: Access Denied. Run bot as Administrator.", flush=True)
    except Exception as e:
        print(f"ERROR: {str(e)}", flush=True)

if __name__ == "__main__":
    main()