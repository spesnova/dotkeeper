package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spesnova/dotkeeper/internal/config"
	"github.com/spf13/cobra"
)

var (
	configFile string
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Create symlinks",
	RunE:  runApply,
}

func runApply(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(configFile)
	if err != nil {
		return fmt.Errorf("failed to load config file: %w", err)
	}

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

func init() {
	applyCmd.Flags().StringVarP(&configFile, "config-file", "c", "dotfiles.yaml", "config file path")
}
