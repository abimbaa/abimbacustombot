package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const VENV_PYTHON = "./scripts/.venv/Scripts/python.exe"

// CommandFunc defines the signature for all command handlers
type CommandFunc func(args string)

// 1. THE REGISTRY: Add new commands here to scale easily
var registry map[string]CommandFunc
func init() {
	registry = map[string]CommandFunc{
	"/start":  handleHelp,
	"/help":   handleHelp,
	"/ping":   handlePing,
	"/ss":     handleScreenshot,
	"/status": handleStatus,
	"/cmd":    handleRawCommand, // e.g., "/cmd dir"
	"/ls":     handleListDir,
	"/get":    handleGetFile,
	"/cam":    handleWebcam,
	"/clip": handleClipboard,
	"/type": handleType,
	"/ps":   handlePs,
	"/kill": handleKill,
	"/run":    handleLaunch,
	"/note":  handleNote,
	"/notes": handleReadNotes,
	"/plan":  handlePlan,
	"/plans": handleListPlans,
	"/pdel":  handleDelPlan,
	"/volup":  func(a string) { runProcess(VENV_PYTHON, "./scripts/media_control.py", "volup") },
	"/voldown": func(a string) { runProcess(VENV_PYTHON, "./scripts/media_control.py", "voldown") },
	"/mute":    func(a string) { runProcess(VENV_PYTHON, "./scripts/media_control.py", "mute") },
	"/pause":   func(a string) { runProcess(VENV_PYTHON, "./scripts/media_control.py", "playpause") },
	"/next":    func(a string) { runProcess(VENV_PYTHON, "./scripts/media_control.py", "next") },
	"/prev":    func(a string) { runProcess(VENV_PYTHON, "./scripts/media_control.py", "prev") },
	"/vol": func(a string) {
		if a == "" {
			fmt.Println("Error: Usage: /vol <number>")
			return
		}
		runProcess(VENV_PYTHON, "./scripts/media_control.py", "setvol", a)
	},
	"/shutdown": func(a string) { runProcess(VENV_PYTHON, "./scripts/power.py", "shutdown") },
	"/restart":  func(a string) { runProcess(VENV_PYTHON, "./scripts/power.py", "restart") },
	"/lock":     func(a string) { runProcess(VENV_PYTHON, "./scripts/power.py", "lock") },
	"/sleep":    func(a string) { runProcess(VENV_PYTHON, "./scripts/power.py", "sleep") },
	"/createmacro": handleCreateMacro,
	"/macros":      handleListMacros,
	"/macro":       handleRunMacro,
	"/delmacro":    handleDelMacro,
	"/stt": func(filepath string) { 
    runProcess(VENV_PYTHON, "./scripts/transcribe.py", filepath) 
},
}
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	rawInput := strings.TrimSpace(os.Args[1])
	parts := strings.SplitN(rawInput, " ", 2)
	command := strings.ToLower(parts[0])

	args := ""
	if len(parts) > 1 {
		args = parts[1]
	}

	// 2. The Router
	if handler, exists := registry[command]; exists {
		handler(args) // Execute the mapped function
	} else {
		handleDeepSeek(rawInput) // Fallback for unknown commands
	}
}

// --- COMMAND HANDLERS ---

func handleListDir(args string) {
	runProcess(VENV_PYTHON, "./scripts/list_dir.py", args)
}

func handleGetFile(args string) {
	if args == "" {
		fmt.Println("⚠️ ERROR: Please specify a filename. Usage: /get <filename>")
		return
	}

	fmt.Printf("🔍 Searching for '%s'...\n", args)
	runProcess(VENV_PYTHON, "./scripts/get_file.py", args)
}

func handleLaunch(args string) {
	if args == "" {
		fmt.Println("⚠️ ERROR: Please specify an app name. Usage: /run <app_name>")
		return
	}
	runProcess(VENV_PYTHON, "./scripts/launch.py", args)
}

func handleWebcam(args string) {	
	pollAmount := 1
	pollDelay := 0

	// Parse the arguments if the user provided them
	if args != "" {
		fmt.Sscanf(args, "%d %d", &pollAmount, &pollDelay)
	}

	// Convert ints to strings to pass to the Python script
	amountStr := strconv.Itoa(pollAmount)
	delayStr := strconv.Itoa(pollDelay)

	// Trigger the script with the arguments
	runProcess(VENV_PYTHON, "./scripts/webcam.py", amountStr, delayStr)
}

func handlePs(args string) {
	// Pass "list" as the first internal arg to Python
	runProcess(VENV_PYTHON, "./scripts/process_manager.py", "list", args)
}

func handleKill(args string) {
	if args == "" {
		fmt.Println("⚠️ ERROR: Usage: /kill <name or pid>")
		return
	}
	// Pass "kill" as the first internal arg to Python
	runProcess(VENV_PYTHON, "./scripts/process_manager.py", "kill", args)
}

func handleType(args string) {
	if args == "" {
		fmt.Println("⚠️ ERROR: Please provide text to type. Usage: /type <text>")
		return
	}
	runProcess(VENV_PYTHON, "./scripts/type_text.py", args)
}

func handleNote(args string) {
	runProcess(VENV_PYTHON, "./scripts/notebook.py", "add_note", args)
}
func handleReadNotes(args string) {
	runProcess(VENV_PYTHON, "./scripts/notebook.py", "read_notes")
}
func handlePlan(args string) {
	runProcess(VENV_PYTHON, "./scripts/notebook.py", "add_plan", args)
}
func handleListPlans(args string) {
	runProcess(VENV_PYTHON, "./scripts/notebook.py", "list_plans")
}
func handleDelPlan(args string) {
	runProcess(VENV_PYTHON, "./scripts/notebook.py", "del_plan", args)
}

func handleClipboard(args string) {
	if args == "" {
		fmt.Println("⚠️ ERROR: Please specify an action. Usage: /clip get OR /clip set <text>")
		return
	}
	runProcess(VENV_PYTHON, "./scripts/clipboard.py", args)
}

func handlePing(args string) { //a simple heartbeat to check if the bot is running
	fmt.Println("System is fully operational.")
}

func handleScreenshot(args string) {
    pollAmount := 1
    pollDelay := 0

    fmt.Sscanf(args, "%d %d", &pollAmount, &pollDelay)

    amountStr := strconv.Itoa(pollAmount)
    delayStr := strconv.Itoa(pollDelay)

    runProcess(VENV_PYTHON, "./scripts/screenshot.py", amountStr, delayStr)
}

func handleStatus(args string) { // returns ram / cpu usage and maybe something else
	runProcess(VENV_PYTHON, "./scripts/status.py")
}

func handleHelp(args string) {
	args = strings.TrimSpace(args)

	docs := map[string]struct{ Short, Long string }{
		"/help":        {"Show this help menu", "Usage: /help [command]\nLists all commands, or shows details for a specific one."},
		"/ping":        {"Check bot status", "Usage: /ping\nReturns a pong message to verify the bot is online."},
		"/ss":          {"Take a screenshot", "Usage: /ss [pollAmount = 1] [pollDelay = 0]\nCaptures all monitors and sends the images. Repeats pollAmount number of times with a pollDelay delay."},
		"/status":      {"System status", "Usage: /status\nDisplays CPU, RAM."},
		"/cmd":         {"Run shell command", "Usage: /cmd <command>\nExecutes a raw command prompt or powershell command and returns the output. Use && to chain commands"},
		"/ls":          {"List directory", "Usage: /ls [path]\nLists files in the current or specified directory."},
		"/get":         {"Get a file", "Usage: /get <filepath>\nUploads a specified file from your PC to Telegram."},
		"/cam":         {"Take webcam picture", "Usage: /cam [pollAmount = 1] [pollDelay = 0]\nCaptures an image from the default webcam. Repeats pollAmount of times with a pollDelay delay."},
		"/clip":        {"Read clipboard", "Usage: /clip get|set <content>\nReturns the current text copied to your PC clipboard | Puts the <content> into the PC's clipboard."},
		"/type":        {"Type text", "Usage: /type <text>\nSimulates keyboard typing the provided text. Use $key$ to simulate special keys like $enter$"},
		"/ps":          {"List processes", "Usage: /ps [name]\nLists currently running applications and processes. Add a name after to filter by name."},
		"/kill":        {"Kill a process", "Usage: /kill <name_or_pid>\nTerminates a process by its executable name or Process ID."},
		"/run":         {"Launch application", "Usage: /run <app_name>\nSearches your PC and launches the requested application."},
		"/note":        {"Add a note", "Usage: /note <title> <body>\nSaves a markdown note to your local data folder."},
		"/notes":       {"Read all notes", "Usage: /notes\nPrints the contents of all saved notes."},
		"/plan":        {"Add a plan", "Usage: /plan <text>\nAppends a new task to your plans list."},
		"/plans":       {"List plans", "Usage: /plans\nShows a numbered list of all your plans."},
		"/pdel":        {"Delete a plan", "Usage: /pdel <number>\nDeletes the plan at the specified number."},
		"/volup":       {"Increase volume", "Usage: /volup\nIncreases system volume by 2%."},
		"/voldown":     {"Decrease volume", "Usage: /voldown\nDecreases system volume by 2%."},
		"/vol":         {"Set exact volume", "Usage: /vol <0-100>\nSets system volume to the exact percentage."},
		"/mute":        {"Toggle mute", "Usage: /mute\nMutes or unmutes the system audio."},
		"/pause":       {"Play/Pause media", "Usage: /pause\nSimulates the play/pause media key."},
		"/next":        {"Next track", "Usage: /next\nSkips to the next media track."},
		"/prev":        {"Previous track", "Usage: /prev\nGoes to the previous media track."},
		"/shutdown":    {"Shutdown PC", "Usage: /shutdown\nTurns off the computer in 5 seconds."},
		"/restart":     {"Restart PC", "Usage: /restart\nReboots the computer in 5 seconds."},
		"/lock":        {"Lock PC", "Usage: /lock\nLocks your Windows session."},
		"/sleep":       {"Put PC to sleep", "Usage: /sleep\nSuspends the computer."},
		"/createmacro": {"Create a macro", "Usage: /createmacro <name>\n/cmd1\n/cmd2\nSaves a chain of commands to run together."},
		"/macros":      {"List all macros", "Usage: /macros\nShows all your saved macros."},
		"/macro":       {"Run a macro", "Usage: /macro <name>\nExecutes the specified macro sequence."},
		"/delmacro":    {"Delete a macro", "Usage: /delmacro <name>\nRemoves a saved macro."},
		"/stt":         {"Transcribe audio", "Usage: /stt <filepath>\nTranscribes an audio file. Note: Voice / video messages trigger this automatically."},
	}

	if args == "" {
		fmt.Println("Available Commands:")
		for cmd, doc := range docs {
			fmt.Printf("%s - %s\n", cmd, doc.Short)
		}
		fmt.Println("\nTip: Type `/help <command>` for details.")
	} else {
		if !strings.HasPrefix(args, "/") {
			args = "/" + args
		}
		
		if doc, exists := docs[args]; exists {
			fmt.Printf("%s\n", doc.Long)
		} else {
			fmt.Printf("Unknown command: %s. (If this is a custom macro, use /macros to see it).\n", args)
		}
	}
}

func handleRawCommand(args string) { //cmd commands
	if args == "" {
		fmt.Println("ERROR: No command provided.")
		return
	}
	cmd := exec.Command("cmd", "/C", args)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	out, err := cmd.CombinedOutput()
	
	if len(out) > 0 {
		fmt.Print(string(out))
	} else {
		fmt.Print("Done.")
	}

	if err != nil {
		fmt.Printf("\n[System Error]: %v\n", err)
	}
}

func handleDeepSeek(input string) { // deepseek to parse to a command
	fmt.Printf("Sending to DeepSeek API for interpretation: %s\n", input)
	// Implement HTTP request to DeepSeek here
}

func handleCreateMacro(args string) {
	// Split by newline to separate the name from the commands
	lines := strings.Split(strings.ReplaceAll(args, "\r\n", "\n"), "\n")
	if len(lines) < 2 {
		fmt.Println("⚠️ Usage (Send as one message):\n/createmacro <name>\n/command1\n/command2")
		return
	}

	name := strings.TrimSpace(lines[0])
	var commands []string
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line != "" {
			commands = append(commands, line)
		}
	}

	macros := loadMacros()
	macros[name] = commands
	saveMacros(macros)
	fmt.Printf("✅ Macro '%s' saved with %d commands.\n", name, len(commands))
}

func handleListMacros(args string) {
	macros := loadMacros()
	if len(macros) == 0 {
		fmt.Println("No macros found.")
		return
	}

	fmt.Println("**Saved Macros:**")
	for name, cmds := range macros {
		fmt.Printf("🔸 %s (%d cmds)\n", name, len(cmds))
	}
}

func handleRunMacro(args string) {
	name := strings.TrimSpace(args)
	macros := loadMacros()
	cmds, exists := macros[name]
	if !exists {
		fmt.Printf("Macro '%s' not found.\n", name)
		return
	}

	fmt.Printf("Running macro: %s\n", name)
	for _, cmdStr := range cmds {
		fmt.Printf("▶️ %s\n", cmdStr)

		// Parse the saved command just like an incoming Telegram message
		parts := strings.SplitN(cmdStr, " ", 2)
		cmd := strings.ToLower(parts[0])
		cmdArgs := ""
		if len(parts) > 1 {
			cmdArgs = parts[1]
		}

		// Trigger the function directly from your registry
		if handler, ok := registry[cmd]; ok {
			handler(cmdArgs)
			// Small delay so we don't overwhelm the OS when opening multiple apps
			time.Sleep(1 * time.Second)
		} else {
			fmt.Printf("⚠️ Unknown command in macro: %s\n", cmd)
		}
	}
	fmt.Println("✅ Macro complete.")
}

func handleDelMacro(args string) {
	name := strings.TrimSpace(args)
	macros := loadMacros()
	if _, exists := macros[name]; exists {
		delete(macros, name)
		saveMacros(macros)
		fmt.Printf("Deleted macro: %s\n", name)
	} else {
		fmt.Printf("Macro '%s' not found.\n", name)
	}
}

// --- UTILITIES ---

func runProcess(name string, args ...string) {
	cmd := exec.Command(name, args...)
	
	// Force execution in the project directory
	exePath, _ := os.Executable()
	cmd.Dir = filepath.Dir(exePath)
	
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	// Stream output directly to the Dispatcher in real-time
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("\nERROR: Execution failed - %v\n", err)
	}
}

var macrosFile = "data/macros.json"

func loadMacros() map[string][]string {
	macros := make(map[string][]string)
	data, err := os.ReadFile(macrosFile)
	if err == nil {
		json.Unmarshal(data, &macros)
	}
	return macros
}

func saveMacros(macros map[string][]string) {
	os.MkdirAll("data", 0755)
	data, _ := json.MarshalIndent(macros, "", "  ")
	os.WriteFile(macrosFile, data, 0644)
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}