package symlink

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spesnova/dotkeeper/internal/config"
)

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Create(symlinks []config.Symlink) error {
	if len(symlinks) == 0 {
		return nil
	}

	fmt.Println("-----> Creating symlinks...")

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
