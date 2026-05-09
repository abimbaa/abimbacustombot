import sys
import os
import whisper

sys.stdout.reconfigure(encoding='utf-8')

def main():
    if len(sys.argv) < 2:
        print("Error: No file path provided.", flush=True)
        return

    filepath = sys.argv[1]
    if not os.path.exists(filepath):
        print(f"Error: File not found at {filepath}", flush=True)
        return

    try:
        # "base" is a good balance of speed and accuracy. Use "tiny" for faster results.
        model = whisper.load_model("base")
        
        result = model.transcribe(filepath)
            
        print(f"{result['text'].strip()}", flush=True)

    except Exception as e:
        print(f"STT Error: {str(e)}", flush=True)

if __name__ == "__main__":
    main()