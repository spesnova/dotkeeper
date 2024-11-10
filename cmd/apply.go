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
	// Load config file
	fmt.Printf("-----> Loading config file: %s\n", configFile)
	cfg, err := config.Load(configFile)
	if err != nil {
		return fmt.Errorf("failed to load config file: %w", err)
	}

	// Create symlinks
	fmt.Println("-----> Creating symlinks...")
	if err := createSymlinks(cfg.Symlinks); err != nil {
		return fmt.Errorf("failed to create symlinks: %w", err)
	}

	// Initialize submodules
	fmt.Println("-----> Initializing submodules...")
	if err := initSubmodules(cfg.Submodules); err != nil {
		return fmt.Errorf("failed to initialize submodules: %w", err)
	}

	// Install apt packages
	fmt.Println("-----> Installing apt packages...")
	if err := installAptPackages(cfg.AptPackages); err != nil {
		return fmt.Errorf("failed to install apt packages: %w", err)
	}

	return nil
}

func createSymlinks(symlinks []config.Symlink) error {
	for i, link := range symlinks {
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
		fmt.Println(err)
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

func installAptPackages(packages []string) error {
	// Update apt
	updateCmd := exec.Command("sudo", "apt-get", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	if err := updateCmd.Run(); err != nil {
		return fmt.Errorf("failed to update apt: %w", err)
	}

	// Install packages
	args := append([]string{"apt-get", "install", "-y"}, packages...)
	installCmd := exec.Command("sudo", args...)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to install packages: %w", err)
	}

	return nil
}

func init() {
	applyCmd.Flags().StringVarP(&configFile, "config-file", "c", "dotfiles.yaml", "config file path")
}
