package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// CONFIGURABLE SECTION:
// Adjust these constants based on your project's language and layout.
const (
	codeqlCmd         = "codeql"                // CodeQL CLI must be installed and in PATH
	projectDir        = "./my-local-project"    // Path to your local codebase
	databaseDir       = "./codeql-db"           // Directory where the CodeQL database will be stored
	language          = "javascript"            // Language of your project (e.g., javascript, go, cpp, java)
	outputSarif       = "./codeql-results.sarif"// SARIF output file
	querySuite        = "codeql/javascript-queries" // Default CodeQL query pack for JS
)

func main() {
	// Step 1: Check if CodeQL is installed
	if _, err := exec.LookPath(codeqlCmd); err != nil {
		log.Fatalf("CodeQL CLI not found in PATH. Please install CodeQL and try again.")
	}

	// Step 2: Create the CodeQL database
	fmt.Println("Step 1: Creating CodeQL database...")

	cmdCreate := exec.Command(
		codeqlCmd, "database", "create", databaseDir,
		"--language="+language,
		"--source-root="+projectDir,
	)
	// Pipe command output to stdout/stderr for verbose logs
	cmdCreate.Stdout = os.Stdout
	cmdCreate.Stderr = os.Stderr

	if err := cmdCreate.Run(); err != nil {
		log.Fatalf("Failed to create CodeQL database: %v", err)
	}

	// Step 3: Analyze the database using the query suite
	fmt.Println("Step 2: Running CodeQL analysis...")

	cmdAnalyze := exec.Command(
		codeqlCmd, "database", "analyze", databaseDir,
		querySuite,
		"--format=sarifv2.1.0",
		"--output="+outputSarif,
	)
	cmdAnalyze.Stdout = os.Stdout
	cmdAnalyze.Stderr = os.Stderr

	if err := cmdAnalyze.Run(); err != nil {
		log.Fatalf("Failed to analyze CodeQL database: %v", err)
	}

	fmt.Println("Analysis complete. Results written to:", filepath.Clean(outputSarif))
}
