package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// CommandFunc defines the signature for all command handlers
type CommandFunc func(args string)

// 1. THE REGISTRY: Add new commands here to scale easily
var registry = map[string]CommandFunc{
	"/start":  handleHelp,
	"/help":   handleHelp,
	"/ping":   handlePing,
	"/ss":     handleScreenshot,
	"/status": handleStatus,
	"/cmd":    handleRawCommand, // e.g., "/cmd dir"
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

func handlePing(args string) { //a simple heartbeat to check if the bot is running
	fmt.Println("System is fully operational.")
}

func handleHelp(args string) { //lists all commands, can take an argument to explain a given argument
	fmt.Println("Commands: ")
}

func handleScreenshot(args string) {// takes two arguments: pollAmount (by default 1), pollDelay (by default 0) and sends screenshots
	fmt.Println("Capturing screen...")
	runProcess("python", "./scripts/screenshot.py")
}

func handleStatus(args string) { // returns ram / cpu usage and maybe something else
	fmt.Println("System Status: NORMAL")
	fmt.Println("Memory Load: LOW")
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

// --- UTILITIES ---

// runProcess executes external commands and bridges their output to our API contract
func runProcess(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("ERROR: Execution failed - %v\n", err)
	}
	if len(out) > 0 {
		fmt.Print(string(out))
	}
}