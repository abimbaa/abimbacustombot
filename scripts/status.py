import psutil

def main():
    cpu = psutil.cpu_percent(interval=1)
    
    ram = psutil.virtual_memory()
    print(f"CPU: {cpu}%", flush=True)
    
    ram_used = ram.used // (1024**3)
    ram_total = ram.total // (1024**3)
    print(f"RAM: {ram.percent}% ({ram_used}GB / {ram_total}GB)", flush=True)

if __name__ == "__main__":
    main()