package homebrew

import (
	"fmt"
	"os"
	"os/exec"
)

// Manager handles the installation of Homebrew packages
type Manager struct{}

// NewManager creates a new Homebrew manager
func NewManager() *Manager {
	return &Manager{}
}

// Install installs Homebrew formulae and casks
func (m *Manager) Install(formulae []string, casks []string) error {
	if !isCommandAvailable("brew") {
		return fmt.Errorf("brew command is not installed. Please install Homebrew first: https://brew.sh")
	}

	if len(formulae) > 0 {
		if err := m.installFormulae(formulae); err != nil {
			return fmt.Errorf("failed to install Homebrew formulae: %w", err)
		}
	}

	if len(casks) > 0 {
		if err := m.installCasks(casks); err != nil {
			return fmt.Errorf("failed to install Homebrew casks: %w", err)
		}
	}

	return nil
}

// installFormulae installs Homebrew formulae
func (m *Manager) installFormulae(formulae []string) error {
	if len(formulae) == 0 {
		return nil
	}

	fmt.Println("-----> Installing Homebrew casks...")

	// Update Homebrew
	updateCmd := exec.Command("brew", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	if err := updateCmd.Run(); err != nil {
		return fmt.Errorf("failed to update Homebrew: %w", err)
	}

	// Install casks
	args := append([]string{"install"}, formulae...)
	installCmd := exec.Command("brew", args...)
	fmt.Println(installCmd.String())
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to install casks: %w", err)
	}

	return nil
}

// installCasks installs Homebrew casks
func (m *Manager) installCasks(casks []string) error {
	if len(casks) == 0 {
		return nil
	}

	fmt.Println("-----> Installing Homebrew casks...")

	// Update Homebrew
	updateCmd := exec.Command("brew", "update")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr
	if err := updateCmd.Run(); err != nil {
		return fmt.Errorf("failed to update Homebrew: %w", err)
	}

	// Install casks
	args := append([]string{"install", "--casks"}, casks...)
	installCmd := exec.Command("brew", args...)
	fmt.Println(installCmd.String())
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to install casks: %w", err)
	}

	return nil
}

// isCommandAvailable checks if the specified command is available in the system
func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
