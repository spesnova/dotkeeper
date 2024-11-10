package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Symlink represents a single symlink configuration
type Symlink struct {
	Src string `yaml:"src"`
	Dst string `yaml:"dst"`
}

// Submodule represents a git submodule configuration
type Submodule struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// Config represents the root configuration structure
type Config struct {
	Symlinks   []Symlink   `yaml:"symlinks"`
	Submodules []Submodule `yaml:"submodules"`
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
