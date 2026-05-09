import sys
import psutil
from collections import Counter

def main():
    if len(sys.argv) < 2:
        return

    cmd = sys.argv[1].lower()
    args = sys.argv[2:] if len(sys.argv) > 2 else []

    if cmd == "list":
        filter_str = args[0].lower() if args else None
        
        # Dictionary to store {name: [pids]}
        proc_groups = {}

        for proc in psutil.process_iter(['pid', 'name']):
            try:
                name = proc.info['name']
                pid = proc.info['pid']

                if not filter_str or filter_str in name.lower():
                    if name not in proc_groups:
                        proc_groups[name] = []
                    proc_groups[name].append(pid)
            except (psutil.NoSuchProcess, psutil.AccessDenied):
                continue

        print(f"**System Processes**", flush=True)
        
        # Sort by name for readability
        sorted_names = sorted(proc_groups.keys())
        
        output = []
        for name in sorted_names:
            pids = proc_groups[name]
            count = len(pids)
            main_pid = pids[0]
            
            if count > 1:
                output.append(f"`{main_pid}` | {name} (x{count})")
            else:
                output.append(f"`{main_pid}` | {name}")

        # Telegram message limit safety
        for line in output[:60]:
            print(line, flush=True)

        if len(output) > 60:
            print(f"\n... and {len(output) - 60} more groups.", flush=True)

    elif cmd == "kill":
        target = args[0]
        killed_count = 0
        for proc in psutil.process_iter(['pid', 'name']):
            try:
                if target == str(proc.info['pid']) or target.lower() == proc.info['name'].lower():
                    proc.kill()
                    killed_count += 1
            except: continue
        
        if killed_count > 0:
            print(f"Terminated {killed_count} instance(s) of '{target}'.", flush=True)
        else:
            print(f"Process '{target}' not found.", flush=True)

if __name__ == "__main__":
    main()