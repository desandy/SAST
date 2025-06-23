package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Entry point of the program
func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <path-to-repo-or-dockerfile> <mode: repo|config>", os.Args[0])
	}

	targetPath := os.Args[1]
	scanMode := os.Args[2]

	// Check if Trivy is installed
	if _, err := exec.LookPath("trivy"); err != nil {
		log.Fatal("Trivy CLI not found in PATH. Please install Trivy: https://aquasecurity.github.io/trivy/")
	}

	absPath, err := filepath.Abs(targetPath)
	if err != nil {
		log.Fatalf("Failed to resolve path: %v", err)
	}

	switch scanMode {
	case "repo":
		runTrivyRepoScan(absPath)
	case "config":
		runTrivyConfigScan(absPath)
	default:
		log.Fatalf("Unknown scan mode '%s'. Use 'repo' or 'config'.", scanMode)
	}
}

// Run `trivy repo <path>`
func runTrivyRepoScan(path string) {
	fmt.Println("Running Trivy repo scan on:", path)

	cmd := exec.Command("trivy", "repo", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Trivy repo scan failed: %v", err)
	}

	fmt.Println("✅ Trivy repo scan completed successfully.")
}

// Run `trivy config <path>`
func runTrivyConfigScan(path string) {
	fmt.Println("Running Trivy config scan on:", path)

	cmd := exec.Command("trivy", "config", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Trivy config scan failed: %v", err)
	}

	fmt.Println("✅ Trivy config scan completed successfully.")
}
