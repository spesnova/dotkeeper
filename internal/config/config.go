package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spesnova/dotkeeper/internal/version"
	"gopkg.in/yaml.v3"
)

// SchemaVersion represents the schema version of the configuration file
type Version string

// Symlink represents a single symlink configuration
type Symlink struct {
	Src string `yaml:"src"`
	Dst string `yaml:"dst"`
}

// GitSubmodule represents a git submodule configuration
type GitSubmodule struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

type Homebrew struct {
	Formulae []string `yaml:"formulae"`
	Casks    []string `yaml:"casks"`
}

// MASConfig represents the MAS configuration structure
type MASConfig struct {
	AppIDs []string `yaml:"app_ids,omitempty"`
}

// Config represents the root configuration structure
type Config struct {
	Version       Version        `yaml:"version"`
	Symlinks      []Symlink      `yaml:"symlinks"`
	GitSubmodules []GitSubmodule `yaml:"git_submodules"`
	AptPackages   []string       `yaml:"apt_packages"`
	Homebrew      Homebrew       `yaml:"homebrew"`
	MAS           MASConfig      `yaml:"mas"`
}

// Load reads and parses the configuration file
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	if err := config.ValidateVersion(version.GetVersion()); err != nil {
		return nil, err
	}

	// Expand home directory (~) in paths
	for i := range config.Symlinks {
		config.Symlinks[i].Dst = expandPath(config.Symlinks[i].Dst)
	}

	return &config, nil
}

// expandPath expands the home directory symbol (~) to the actual path
func expandPath(path string) string {
	if len(path) == 0 || path[0] != '~' {
		return path
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	return filepath.Join(home, path[1:])
}

// ValidateVersion validates the version of the config file against the CLI version
func (c *Config) ValidateVersion(cliVersion string) error {
	if c.Version == "" {
		return nil // Skip version validation if not specified
	}

	configMajorVersion := string(c.Version)[1:2]
	cliMajorVersion := string(cliVersion)[1:2]

	if configMajorVersion != cliMajorVersion {
		return fmt.Errorf("config file version (%s) does not match CLI version (%s). Major versions must match", c.Version, cliVersion)
	}

	return nil
}
