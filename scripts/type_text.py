import sys
import pyautogui
import re

def main():
    if len(sys.argv) < 2:
        print("ERROR: Usage: /type <text>", flush=True)
        return

    # Combine arguments into one string
    input_text = " ".join(sys.argv[1:])

    try:
        parts = re.split(r'(\$[a-z]+\$)', input_text)

        for part in parts:
            if part.startswith('$') and part.endswith('$'):
                # Extract the key name (e.g., "enter" from "$enter$")
                key = part.strip('$').lower()
                try:
                    pyautogui.press(key)
                    print(f"Pressed: {key}", flush=True)
                except Exception:
                    # If it's not a valid key, just type it literally
                    pyautogui.write(part)
            else:
                # Type normal text
                if part:
                    pyautogui.write(part, interval=0.01)

        print("Done.", flush=True)
    except Exception as e:
        print(f"ERROR: {str(e)}", flush=True)

if __name__ == "__main__":
    main()