package main

import (
	"fmt"
	"os"
	"os/exec"
	"unsafe"
)

// PlatformSpecificFunction contains intentional vulnerabilities for Linux platform
func PlatformSpecificFunction() {
	fmt.Println("Running Linux-specific operations...")

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

	// VULNERABILITY 4: Unsafe Pointer Operations (High)
	// Using unsafe pointers incorrectly
	data2 := []byte("sensitive data")
	ptr := unsafe.Pointer(&data2[0])
	// VULNERABLE: Unsafe pointer manipulation
	_ = (*[1000]byte)(ptr)

	// VULNERABILITY 5: Hardcoded Credentials (High)
	linuxPassword := "SuperSecret123!" // VULNERABLE: Hardcoded password
	fmt.Printf("Using Linux credentials (length: %d)\n", len(linuxPassword))

	// VULNERABILITY 6: SQL Injection potential (High)
	query := "SELECT * FROM users WHERE username = '" + os.Getenv("USERNAME") + "'"
	// VULNERABLE: SQL injection if this were used with a database
	fmt.Printf("Generated query: %s\n", query)

	fmt.Println("Linux operations completed")
}
