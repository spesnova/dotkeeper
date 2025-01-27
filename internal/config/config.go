package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	SCHEMA_VERSION = "v0.1.0"
)

// SchemaVersion represents the schema version of the configuration file
type SchemaVersion string

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

// Config represents the root configuration structure
type Config struct {
	SchemaVersion SchemaVersion  `yaml:"schema_version"`
	Symlinks      []Symlink      `yaml:"symlinks"`
	GitSubmodules []GitSubmodule `yaml:"git_submodules"`
	AptPackages   []string       `yaml:"apt_packages"`
	Homebrew      Homebrew       `yaml:"homebrew"`
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

	if config.SchemaVersion != SchemaVersion(SCHEMA_VERSION) {
		return nil, fmt.Errorf("unsupported schema version: %s (current version: %s)", config.SchemaVersion, SCHEMA_VERSION)
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
