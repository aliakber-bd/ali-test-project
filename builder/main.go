package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("Platform-Specific Code Generator\n")
	fmt.Printf("Building for: %s/%s\n\n", runtime.GOOS, runtime.GOARCH)

	// Generate OS-specific file based on build platform
	switch runtime.GOOS {
	case "darwin":
		generateMacOSCode()
	case "linux":
		generateLinuxCode()
	case "windows":
		generateWindowsCode()
	default:
		fmt.Printf("Unsupported platform: %s\n", runtime.GOOS)
		os.Exit(1)
	}

	fmt.Println("Code generation completed successfully")
}

func generateLinuxCode() {
	fmt.Println("Generating Linux-specific code with vulnerabilities...")

	code := `package main

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
`

	err := os.WriteFile("platform_linux.go", []byte(code), 0644)
	if err != nil {
		fmt.Printf("Error writing Linux code: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Generated platform_linux.go with 6 vulnerabilities")
}

func generateMacOSCode() {
	fmt.Println("Generating macOS-specific code with vulnerabilities...")

	code := `//go:build darwin
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
`

	err := os.WriteFile("platform_darwin.go", []byte(code), 0644)
	if err != nil {
		fmt.Printf("Error writing Darwin code: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Generated platform_darwin.go with 6 vulnerabilities")
}

func generateWindowsCode() {
	fmt.Println("Generating Windows-specific code with vulnerabilities...")

	code := `package main

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
`

	err := os.WriteFile("platform_windows.go", []byte(code), 0644)
	if err != nil {
		fmt.Printf("Error writing Windows code: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Generated platform_windows.go with 7 vulnerabilities")
}
