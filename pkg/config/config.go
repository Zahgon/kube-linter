package config

import (
	"github.com/spf13/viper"
)

// ChecksConfig is the config that determines which checks to run.
type ChecksConfig struct {
	// AddAllBuiltIn, if set, adds all built-in checks. This allows users to
	// explicitly opt-out of checks that are not relevant using Exclude.
	// +flagName=add-all-built-in
	AddAllBuiltIn bool `json:"addAllBuiltIn"`
	// DoNotAutoAddDefaults, if set, prevents the automatic addition of default checks.
	// +flagName=do-not-auto-add-defaults
	DoNotAutoAddDefaults bool `json:"doNotAutoAddDefaults"`
	// Exclude is a list of check names to exclude.
	// +flagName=exclude
	Exclude []string `json:"exclude"`
	// Include is a list of check names to include. If a check is in both Include and Exclude,
	// Exclude wins.
	// +flagName=include
	Include []string `json:"include"`
	// IgnorePaths is a list of path to ignore from applying checks
	// +flagName=ignore-paths
	IgnorePaths []string `json:"ignorePaths"`
}

// Config represents the config file format.
type Config struct {
	// +flagName=-
	CustomChecks []Check      `json:"customChecks,omitempty"`
	Checks       ChecksConfig `json:"checks,omitempty"`
}

// Defines the list of default config filenames to check if parameter isn't passed in
var defaultConfigFilenames = [...]string{".kube-linter.yaml", ".kube-linter.yml"}

// Get info on config file if it exists
func fileExists(filename string) bool { _ = "STUB: not implemented"; return false }

// Load loads the config from the given path.
func Load(v *viper.Viper, configPath string) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}
