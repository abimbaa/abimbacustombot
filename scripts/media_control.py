import sys
import pyautogui

def main():
    if len(sys.argv) < 2:
        return

    cmd = sys.argv[1].lower()

    try:
        if cmd == "playpause":
            pyautogui.press("playpause")
            print("Action: Play/Pause", flush=True)
        elif cmd == "next":
            pyautogui.press("nexttrack")
            print("Action: Next Track", flush=True)
        elif cmd == "prev":
            pyautogui.press("prevtrack")
            print("Action: Previous Track", flush=True)
        elif cmd == "volup":
            pyautogui.press("volumeup")
            print("Action: Volume Up", flush=True)
        elif cmd == "voldown":
            pyautogui.press("volumedown")
            print("Action: Volume Down", flush=True)
        elif cmd == "mute":
            pyautogui.press("volumemute")
            print("Action: Mute Toggled", flush=True)
            
        elif cmd == "setvol":
            if len(sys.argv) < 3:
                print("Error: Provide a percentage.", flush=True)
                return
                
            target_str = sys.argv[2].strip()
            # Handle if the user types /vol 50% instead of /vol 50
            if target_str.endswith('%'):
                target_str = target_str[:-1]
                
            target = int(target_str)
            target = max(0, min(100, target)) # Keep it between 0 and 100
            
            # Zero out the volume completely (50 presses * 2% = 100% down)
            pyautogui.press("volumedown", presses=50)
            
            # Calculate required up presses and execute
            up_presses = target // 2
            if up_presses > 0:
                pyautogui.press("volumeup", presses=up_presses)
                
            print(f"Action: Volume set to {target}%", flush=True)

    except ValueError:
        print("Error: Please provide a valid number.", flush=True)
    except Exception as e:
        print(f"Error: {str(e)}", flush=True)

if __name__ == "__main__":
    main()