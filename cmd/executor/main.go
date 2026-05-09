package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CommandFunc defines the signature for all command handlers
type CommandFunc func(args string)

// 1. THE REGISTRY: Add new commands here to scale easily
var registry = map[string]CommandFunc{
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

func handlePing(args string) {
	fmt.Println("Pong! 🏓 System is fully operational.")
}

func handleScreenshot(args string) {
	fmt.Println("Capturing screen...")
	runProcess("python", "./scripts/screenshot.py")
}

func handleStatus(args string) {
	// Example of doing logic entirely in Go without a script
	fmt.Println("System Status: NORMAL")
	fmt.Println("Memory Load: LOW")
}

func handleRawCommand(args string) {
	if args == "" {
		fmt.Println("ERROR: You must provide a command to run.")
		return
	}
	// Warning: High permission capability
	runProcess("cmd", "/C", args) 
}

func handleDeepSeek(input string) {
	fmt.Printf("Sending to DeepSeek API for interpretation: %s\n", input)
	// Implement HTTP request to DeepSeek here
}

// --- UTILITIES ---

// runProcess executes external commands and bridges their output to our API contract
func runProcess(name string, args ...string) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("ERROR: Execution failed - %v\n", err)
	}
	if len(out) > 0 {
		fmt.Print(string(out))
	}
}