package git

import (
	"fmt"
	"os/exec"

	"github.com/spesnova/dotkeeper/internal/config"
)

// SubmoduleManager handles Git submodule operations
type SubmoduleManager struct{}

// NewSubmoduleManager creates a new Git submodule manager
func NewSubmoduleManager() *SubmoduleManager {
	return &SubmoduleManager{}
}

// Install initializes and updates Git submodules
func (m *SubmoduleManager) Install(submodules []config.GitSubmodule) error {
	if !isCommandAvailable("git") {
		return fmt.Errorf("git command is not available. Please install git first")
	}

	if len(submodules) == 0 {
		return nil
	}

	fmt.Println("-----> Installing Git submodules...")

	// Check if current directory is a git repository
	checkCmd := exec.Command("git", "status")
	if err := checkCmd.Run(); err != nil {
		return fmt.Errorf("current directory is not a git repository: %w", err)
	}

	for _, sub := range submodules {
		// Add submodule if it doesn't exist
		addCmd := exec.Command("git", "submodule", "add", "-f", sub.URL, sub.Path)
		if err := addCmd.Run(); err != nil {
			// Ignore error if submodule already exists
			fmt.Printf("Note: Submodule %s might already exist\n", sub.Path)
		}

		// Initialize submodule
		initCmd := exec.Command("git", "submodule", "init", sub.Path)
		if err := initCmd.Run(); err != nil {
			return fmt.Errorf("failed to initialize submodule %s: %w", sub.Path, err)
		}

		// Update submodule
		updateCmd := exec.Command("git", "submodule", "update", sub.Path)
		if err := updateCmd.Run(); err != nil {
			return fmt.Errorf("failed to update submodule %s: %w", sub.Path, err)
		}

		fmt.Printf("Initialized submodule: %s\n", sub.Path)
	}

	return nil
}

// isCommandAvailable checks if the specified command is available in the system
func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
