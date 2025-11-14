//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
	"unsafe"
)

// PlatformSpecificFunction contains intentional vulnerabilities for macOS platform
func PlatformSpecificFunction() {
	fmt.Println("Running macOS-specific operations...")

	// VULNERABILITY 1: Command Injection (High)
	// User input is directly passed to shell command without sanitization
	userInput := os.Getenv("USER_COMMAND")
	if userInput == "" {
		userInput = "ls -la"
	}
	cmd := exec.Command("sh", "-c", userInput) // VULNERABLE: Command injection
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Command execution failed: %v\n", err)
	} else {
		fmt.Printf("Command output: %s\n", output)
	}

	// VULNERABILITY 2: Path Traversal (High)
	// No validation on file path allowing directory traversal
	filename := os.Getenv("CONFIG_FILE")
	if filename == "" {
		filename = "/etc/passwd"
	}
	// VULNERABLE: Path traversal - reads any file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes from %s\n", len(data), filename)
	}

	// VULNERABILITY 3: Race Condition (Medium)
	// TOCTOU (Time-of-check to time-of-use) vulnerability
	tempFile := "/tmp/race_condition_test"
	if _, err := os.Stat(tempFile); err == nil {
		// VULNERABLE: File could be modified between check and use
		content, _ := os.ReadFile(tempFile)
		fmt.Printf("File exists, content: %s\n", content)
	}

	// VULNERABILITY 4: Memory Safety Issues (High)
	// Unsafe pointer operations that can cause memory corruption
	data2 := make([]byte, 100)
	ptr := unsafe.Pointer(&data2[0])
	// VULNERABLE: Unsafe pointer arithmetic
	newPtr := unsafe.Pointer(uintptr(ptr) + 1000) // Could point to invalid memory
	fmt.Printf("Unsafe pointer operation: %p -> %p\n", ptr, newPtr)

	// VULNERABILITY 5: Information Disclosure (Medium)
	// Exposing sensitive system information
	fmt.Printf("macOS System Info Disclosure:\n")
	fmt.Printf("HOME: %s\n", os.Getenv("HOME"))
	fmt.Printf("USER: %s\n", os.Getenv("USER"))
	fmt.Printf("SHELL: %s\n", os.Getenv("SHELL"))
}
