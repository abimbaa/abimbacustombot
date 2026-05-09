import sys
import os
import re

DATA_DIR = "data"
NOTES_DIR = os.path.join(DATA_DIR, "notes")
PLANS_FILE = os.path.join(DATA_DIR, "plans.txt")

def ensure_data():
    for path in [DATA_DIR, NOTES_DIR]:
        if not os.path.exists(path):
            os.makedirs(path)

def slugify(text):
    # Turn "My Note Title!" into "my_note_title" for a safe filename
    text = text.lower()
    return re.sub(r'\W+', '_', text).strip('_')

def main():
    ensure_data()
    if len(sys.argv) < 2: return

    cmd = sys.argv[1].lower()
    content = " ".join(sys.argv[2:])

    if cmd == "add_note":
        # First word/phrase becomes the filename, the rest is the body
        words = content.split(' ', 1)
        title = words[0]
        body = words[1] if len(words) > 1 else "No content."
        
        filename = f"{slugify(title)}.md"
        filepath = os.path.join(NOTES_DIR, filename)
        
        with open(filepath, "w", encoding="utf-8") as f:
            f.write(f"# {title}\n\n{body}\n")
        print(f"Note saved: {filename}")

    elif cmd == "read_notes":
        files = [f for f in os.listdir(NOTES_DIR) if f.endswith(".md")]
        if not files:
            print("No notes found.")
            return
        
        for filename in files:
            with open(os.path.join(NOTES_DIR, filename), "r", encoding="utf-8") as f:
                print(f"--- {filename} ---")
                print(f.read())

    # --- PLANS (Keeping existing logic) ---
    elif cmd == "add_plan":
        with open(PLANS_FILE, "a", encoding="utf-8") as f:
            f.write(f"{content}\n")
        print(f"Plan added: {content}")

    elif cmd == "list_plans":
        if os.path.exists(PLANS_FILE):
            with open(PLANS_FILE, "r", encoding="utf-8") as f:
                lines = f.readlines()
                for i, line in enumerate(lines):
                    print(f"{i+1}. {line.strip()}")
        else:
            print("No plans found.")

    elif cmd == "del_plan":
        if not os.path.exists(PLANS_FILE): return
        try:
            idx = int(content) - 1
            with open(PLANS_FILE, "r", encoding="utf-8") as f:
                lines = f.readlines()
            if 0 <= idx < len(lines):
                removed = lines.pop(idx)
                with open(PLANS_FILE, "w", encoding="utf-8") as f:
                    f.writelines(lines)
                print(f"Deleted: {removed.strip()}")
        except:
            print("Usage: /pdel <number>")

if __name__ == "__main__":
    main()