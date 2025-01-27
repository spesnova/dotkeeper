package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spesnova/dotkeeper/internal/apt"
	"github.com/spesnova/dotkeeper/internal/config"
	"github.com/spesnova/dotkeeper/internal/git"
	"github.com/spesnova/dotkeeper/internal/homebrew"
	"github.com/spesnova/dotkeeper/internal/mas"
	"github.com/spesnova/dotkeeper/internal/symlink"
	"github.com/spf13/cobra"
)

var (
	configFile string
)

const (
	ConfigFile = "dotkeeper.yaml"
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
	symlinkManager := symlink.NewManager()
	if err := symlinkManager.Create(cfg.Symlinks); err != nil {
		return fmt.Errorf("failed to create symlinks: %w", err)
	}

	// Initialize git submodules
	submoduleManager := git.NewSubmoduleManager()
	if err := submoduleManager.Install(cfg.GitSubmodules); err != nil {
		return fmt.Errorf("failed to initialize git submodules: %w", err)
	}

	// Install apt packages
	if isDebianBased() {
		aptManager := apt.NewManager()
		if err := aptManager.Install(cfg.AptPackages); err != nil {
			return fmt.Errorf("failed to install apt packages: %w", err)
		}
	}

	if isMacOS() {
		brewManager := homebrew.NewManager()
		if err := brewManager.Install(cfg.Homebrew.Formulae, cfg.Homebrew.Casks); err != nil {
			return fmt.Errorf("failed to install Homebrew packages: %w", err)
		}

		masManager := mas.NewManager()
		if err := masManager.Install(cfg.MAS.AppIDs); err != nil {
			return fmt.Errorf("failed to install Mac App Store apps: %w", err)
		}
	}

	return nil
}

func isMacOS() bool {
	return runtime.GOOS == "darwin"
}

func isDebianBased() bool {
	if runtime.GOOS != "linux" {
		return false
	}

	// /etc/os-releaseファイルを読み込んで確認
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return false
	}

	content := string(data)
	return strings.Contains(content, "ID=ubuntu") || strings.Contains(content, "ID=debian")
}

func init() {
	applyCmd.Flags().StringVarP(&configFile, "config-file", "c", ConfigFile, "config file path")
}
