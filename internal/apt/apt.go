package apt

import (
	"fmt"
	"os"
	"os/exec"
)

// Manager handles the installation of APT packages
type Manager struct{}

// NewManager creates a new APT manager
func NewManager() *Manager {
	return &Manager{}
}

// Install installs the specified APT packages
func (m *Manager) Install(packages []string) error {
	if !isCommandAvailable("apt-get") {
		return fmt.Errorf("apt-get command is not available. This feature is only supported on Debian-based systems")
	}

	if len(packages) == 0 {
		return nil
	}

	fmt.Println("-----> Installing APT packages...")

	// Update package list
	updateCmd := exec.Command("sudo", "apt-get", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	if err := updateCmd.Run(); err != nil {
		return fmt.Errorf("failed to update package list: %w", err)
	}

	// Install packages
	args := append([]string{"apt-get", "install", "-y"}, packages...)
	installCmd := exec.Command("sudo", args...)
	fmt.Println(installCmd.String())
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
