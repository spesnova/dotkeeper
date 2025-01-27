package mas

import (
	"fmt"
	"os/exec"
)

// Manager handles the installation of Mac App Store applications
type Manager struct{}

// NewManager creates a new MAS manager
func NewManager() *Manager {
	return &Manager{}
}

// Install installs applications with the specified App IDs
func (m *Manager) Install(appIDs []string) error {
	if !isCommandAvailable("mas") {
		return fmt.Errorf("mas command is not installed. Please install it with 'brew install mas'")
	}

	fmt.Println("-----> Installing Mac App Store apps...")

	for _, appID := range appIDs {
		cmd := exec.Command("mas", "install", appID)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install app ID %s: %w", appID, err)
		}
	}
	return nil
}

// isCommandAvailable checks if the specified command is available in the system
func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
