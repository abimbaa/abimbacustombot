import sys
import pyperclip

def main():
    if len(sys.argv) < 2:
        print("ERROR: Usage: /clip get OR /clip set <text>", flush=True)
        return

    arg = sys.argv[1].strip()

    if arg.lower() == "get":
        text = pyperclip.paste()
        if text:
            print(text, flush=True)
        else:
            print("Clipboard is empty or contains non-text data.", flush=True)

    elif arg.lower().startswith("set "):
        # Remove the "set " command part and copy the rest
        text_to_set = arg[4:]
        pyperclip.copy(text_to_set)
        print("PC Clipboard updated successfully!", flush=True)

    else:
        print("ERROR: Invalid command. Use `/clip get` or `/clip set <text>`", flush=True)

if __name__ == "__main__":
    main()