//go:build windows
// +build windows

package main

import (
	"fmt"
	"os"
	"os/exec"
	"unsafe"
)

// PlatformSpecificFunction contains intentional vulnerabilities for Windows platform
func PlatformSpecificFunction() {
	fmt.Println("Running Windows-specific operations...")

	// VULNERABILITY 1: Command Injection (High)
	// Using cmd.exe with unsanitized user input
	userInput := os.Getenv("USER_COMMAND")
	if userInput == "" {
		userInput = "dir"
	}
	cmd := exec.Command("cmd.exe", "/C", userInput) // VULNERABLE: Command injection
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Command execution failed: %v\n", err)
	} else {
		fmt.Printf("Command output: %s\n", output)
	}

	// VULNERABILITY 2: Path Traversal (High)
	// No validation allowing access to any file on Windows
	filename := os.Getenv("CONFIG_FILE")
	if filename == "" {
		filename = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	}
	// VULNERABLE: Path traversal attack
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes from %s\n", len(data), filename)
	}

	// VULNERABILITY 3: Buffer Overflow Potential (High)
	// Fixed-size buffer with unchecked input
	var buffer [10]byte
	input := os.Getenv("BUFFER_DATA")
	if input == "" {
		input = "short"
	}
	// VULNERABLE: No bounds checking, could overflow
	copy(buffer[:], []byte(input))
	fmt.Printf("Buffer content: %s\n", buffer)

	// VULNERABILITY 4: Unsafe Memory Access (High)
	// Dereferencing unsafe pointers
	data2 := make([]byte, 10)
	ptr := unsafe.Pointer(&data2[0])
	// VULNERABLE: Unsafe pointer arithmetic
	badPtr := unsafe.Pointer(uintptr(ptr) + 100)
	_ = *(*byte)(badPtr)

	// VULNERABILITY 5: Hardcoded Credentials (High)
	windowsApiKey := "WinAPI-Key-12345-ABCDE" // VULNERABLE: Hardcoded API key
	fmt.Printf("Using Windows API key (length: %d)\n", len(windowsApiKey))

	// VULNERABILITY 6: DLL Injection Potential (Critical)
	// Loading DLL from user-controlled path
	dllPath := os.Getenv("PLUGIN_DLL")
	if dllPath == "" {
		dllPath = "user_plugin.dll"
	}
	// VULNERABLE: DLL hijacking/injection if this were actual LoadLibrary call
	fmt.Printf("Would load DLL from: %s\n", dllPath)

	// VULNERABILITY 7: Registry Key Exposure (Medium)
	// Hardcoded sensitive registry path
	regPath := "HKLM\\SOFTWARE\\SecretApp\\Credentials" // VULNERABLE: Exposed sensitive path
	fmt.Printf("Accessing registry: %s\n", regPath)

	fmt.Println("Windows operations completed")
}
