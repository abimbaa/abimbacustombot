import cv2
import time
import sys

def main():
    # Read arguments from Go
    if len(sys.argv) < 3:
        poll_amount = 1
        poll_delay = 0
    else:
        poll_amount = int(sys.argv[1])
        poll_delay = int(sys.argv[2])

    # Open the camera once
    cap = cv2.VideoCapture(0)
    
    if not cap.isOpened():
        print("ERROR: Cannot open webcam. It might be in use or disconnected.", flush=True)
        sys.exit(1)

    # Warm up the camera (let auto-exposure settle)
    for _ in range(5):
        cap.read()
    time.sleep(0.5)

    # Loop to take multiple photos
    for i in range(poll_amount):
        ret, frame = cap.read()
        
        if not ret:
            print(f"ERROR: Failed to grab frame {i+1} from the webcam.", flush=True)
            break

        # Save with a unique name so they don't overwrite
        path = f"webcam_{i}_{int(time.time())}.png"
        cv2.imwrite(path, frame)
        
        # Stream it to Go
        print(f"PHOTO:{path}", flush=True)
        
        # Wait for the delay (but don't wait after the last photo)
        if i < poll_amount - 1:
            time.sleep(poll_delay)
            # Read a few throwaway frames during the delay to keep the buffer fresh
            for _ in range(5):
                cap.read()
                
    # Turn off the camera when the loop is done
    cap.release()

if __name__ == "__main__":
    main()