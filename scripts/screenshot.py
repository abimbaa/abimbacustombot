import sys
import time
import mss

def main():
    if len(sys.argv) < 3:
        poll_amount = 1
        poll_delay = 0
    else:
        poll_amount = int(sys.argv[1])
        poll_delay = int(sys.argv[2])


    with mss.mss() as sct:
        num_monitors = len(sct.monitors)

        for i in range(poll_amount):
            try:
                for mon_idx in range(1, num_monitors):
                    # Give it a name like "screenshot_0_mon_1.png"
                    path = f"screenshot_{i}_mon_{mon_idx}.png"
                    
                    sct.shot(mon=mon_idx, output=path)

                    print(f"PHOTO:{path}", flush=True)

                if i < poll_amount - 1:
                    time.sleep(poll_delay)

            except Exception as e:
                print(f"ERROR: {str(e)}", flush=True)
                sys.exit(1)

if __name__ == "__main__":
    main()