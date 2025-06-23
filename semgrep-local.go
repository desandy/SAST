package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// CONFIGURABLE SECTION
const (
	semgrepCmd     = "semgrep"                  // Ensure Semgrep is installed and available in PATH
	targetDir      = "./my-local-project"       // Path to your codebase
	outputSarif    = "./semgrep-results.sarif"  // Where to save SARIF output
	ruleset        = "p/ci"                     // Use public ruleset from Semgrep Registry, e.g., "p/ci", "auto"
)

func main() {
	// Step 0: Check if Semgrep is installed
	if _, err := exec.LookPath(semgrepCmd); err != nil {
		log.Fatalf("Semgrep CLI not found in PATH. Please install Semgrep (https://semgrep.dev/docs/getting-started/) and try again.")
	}

	// Step 1: Prepare the Semgrep scan command
	fmt.Println("Step 1: Running Semgrep scan...")

	// Construct command: semgrep --config <ruleset> --sarif -o <output> <targetDir>
	cmd := exec.Command(
		semgrepCmd,
		"--config="+ruleset,
		"--sarif",
		"--output="+outputSarif,
		targetDir,
	)

	// Pipe Semgrep output directly to stdout/stderr for real-time feedback
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Step 2: Run the Semgrep command
	if err := cmd.Run(); err != nil {
		log.Fatalf("Semgrep scan failed: %v", err)
	}

	// Step 3: Notify completion
	fmt.Println("Semgrep scan complete. SARIF results written to:", filepath.Clean(outputSarif))
}
