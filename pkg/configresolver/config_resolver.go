package configresolver

import (
	"golang.stackrox.io/kube-linter/pkg/checkregistry"
	"golang.stackrox.io/kube-linter/pkg/config"
)

// LoadCustomChecksInto loads the custom checks from the config into the check registry.
func LoadCustomChecksInto(cfg *config.Config, checkRegistry checkregistry.CheckRegistry) error {
	_ = "STUB: not implemented"
	return nil
}

// GetEnabledChecksAndValidate get the list of enabled checks based on the given config,
// and validates that they exist in the given checkRegistry.
func GetEnabledChecksAndValidate(cfg *config.Config, checkRegistry checkregistry.CheckRegistry) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetIgnorePaths loads the paths from the config into the check registry.
func GetIgnorePaths(cfg *config.Config) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
