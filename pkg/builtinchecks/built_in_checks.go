package builtinchecks

import (
	"embed"
	"sync"

	"golang.stackrox.io/kube-linter/pkg/checkregistry"
	"golang.stackrox.io/kube-linter/pkg/config"
)

var (
	//go:embed yamls
	yamlFiles embed.FS

	loadOnce sync.Once
	list     []config.Check
	loadErr  error
)

// LoadInto loads built-in checks into the registry.
func LoadInto(registry checkregistry.CheckRegistry) error { _ = "STUB: not implemented"; return nil }

// List lists built-in checks.
func List() ([]config.Check, error) { _ = "STUB: not implemented"; return nil, nil }

// Do NOT use filepath.Join here, because embed always uses `/` as the separator,
// irrespective of the OS we're running.
