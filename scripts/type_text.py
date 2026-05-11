import sys
import pyautogui
import re

def main():
    if len(sys.argv) < 2:
        print("ERROR: Usage: /type <text>", flush=True)
        return

    input_text = " ".join(sys.argv[1:])

    try:
        # Changed regex to [^$]+ which means "anything that isn't a dollar sign"
        # This allows symbols like +, numbers, and uppercase letters
        parts = re.split(r'(\$[^$]+\$)', input_text)

        for part in parts:
            if not part:
                continue

            if part.startswith('$') and part.endswith('$'):
                # Extract content and normalize to lowercase
                content = part.strip('$').lower()
                
                try:
                    # Check if it's a combination (contains '+')
                    if '+' in content:
                        keys = content.split('+')
                        pyautogui.hotkey(*keys)
                        print(f"Hotkey: {content}", flush=True)
                    else:
                        # Single key press
                        pyautogui.press(content)
                        print(f"Pressed: {content}", flush=True)
                except Exception as e:
                    # Fallback: type it literally if pyautogui doesn't recognize the key
                    pyautogui.write(part)
                    print(f"Fallback typing: {part}", flush=True)
            else:
                # Type normal text
                pyautogui.write(part, interval=0.01)

        print("Done.", flush=True)
    except Exception as e:
        print(f"ERROR: {str(e)}", flush=True)

if __name__ == "__main__":
    main()