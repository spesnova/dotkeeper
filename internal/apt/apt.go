package apt

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spesnova/dotkeeper/internal/config"
)

// Manager handles the installation of APT packages
type Manager struct{}

// NewManager creates a new APT manager
func NewManager() *Manager {
	return &Manager{}
}

// Install installs the specified APT packages
func (m *Manager) Install(aptConfig config.AptConfig) error {
	if !isCommandAvailable("apt-get") {
		return fmt.Errorf("apt-get command is not available. This feature is only supported on Debian-based systems")
	}

	if len(aptConfig.Sources) > 0 {
		fmt.Println("-----> Adding APT sources...")
		if err := m.AddSources(aptConfig.Sources); err != nil {
			return fmt.Errorf("failed to add APT sources: %w", err)
		}
	}

	// Update package list
	fmt.Println("-----> Updating APT package list...")
	if err := m.Update(); err != nil {
		return fmt.Errorf("failed to update package list: %w", err)
	}

	if len(aptConfig.Packages) > 0 {
		// Install packages
		fmt.Println("-----> Installing APT packages...")
		if err := m.InstallPackages(aptConfig.Packages); err != nil {
			return fmt.Errorf("failed to install packages: %w", err)
		}
	}

	return nil
}

func (m *Manager) AddSources(sources []config.AptSource) error {
	for _, source := range sources {
		sourceFile := fmt.Sprintf("/etc/apt/sources.list.d/%s.list", source.Name)
		content := []byte(source.URI + "\n")
		if err := os.WriteFile(sourceFile, content, 0644); err != nil {
			return fmt.Errorf("failed to write source file %s: %w", sourceFile, err)
		}
	}
	return nil
}

func (m *Manager) Update() error {
	updateCmd := exec.Command("sudo", "apt-get", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	if err := updateCmd.Run(); err != nil {
		return fmt.Errorf("failed to update package list: %w", err)
	}
	return nil
}

func (m *Manager) InstallPackages(packages []string) error {
	if len(packages) == 0 {
		return nil
	}

	installCmd := exec.Command("sudo", "apt-get", "install", "-y", strings.Join(packages, " "))
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to install packages: %w", err)
	}

	return nil
}

// isCommandAvailable checks if the specified command is available in the system
func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
