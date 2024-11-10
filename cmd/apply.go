package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spesnova/dotkeeper/internal/config"
	"github.com/spf13/cobra"
)

var (
	configFile string
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Create symlinks and initialize git submodules",
	RunE:  runApply,
}

func runApply(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(configFile)
	if err != nil {
		return fmt.Errorf("failed to load config file: %w", err)
	}

	// Initialize submodules
	if err := initSubmodules(cfg.Submodules); err != nil {
		return fmt.Errorf("failed to initialize submodules: %w", err)
	}

	// Create symlinks
	for i, link := range cfg.Symlinks {
		// Create target directory if it doesn't exist
		if err := os.MkdirAll(filepath.Dir(link.Dst), 0755); err != nil {
			return fmt.Errorf("failed to create directory for symlink %d: %w", i+1, err)
		}

		// Remove existing symlink if it exists
		if _, err := os.Lstat(link.Dst); err == nil {
			if err := os.Remove(link.Dst); err != nil {
				return fmt.Errorf("failed to remove existing symlink: %w", err)
			}
		}

		// Get absolute path for source
		srcAbs, err := filepath.Abs(link.Src)
		if err != nil {
			return fmt.Errorf("failed to resolve source path: %w", err)
		}

		// Create symlink
		if err := os.Symlink(srcAbs, link.Dst); err != nil {
			return fmt.Errorf("failed to create symlink %d: %w", i+1, err)
		}

		fmt.Printf("Created: %s -> %s\n", link.Dst, srcAbs)
	}

	return nil
}

func initSubmodules(submodules []config.Submodule) error {
	if len(submodules) == 0 {
		return nil
	}

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

func init() {
	applyCmd.Flags().StringVarP(&configFile, "config-file", "c", "dotfiles.yaml", "config file path")
}
