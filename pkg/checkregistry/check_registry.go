package checkregistry

import (
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/instantiatedcheck"
)

// A CheckRegistry is a registry of checks.
// It is not thread-safe. It is anticipated that checks will all be registered ahead of time
// before calls to Load.
type CheckRegistry interface {
	Register(checks ...*config.Check) error
	Load(name string) *instantiatedcheck.InstantiatedCheck
}

type checkRegistry map[string]*instantiatedcheck.InstantiatedCheck

func (cr checkRegistry) Register(checks ...*config.Check) error {
	_ = "STUB: not implemented"
	return nil
}

func (cr checkRegistry) Load(name string) *instantiatedcheck.InstantiatedCheck {
	_ = "STUB: not implemented"

	// New returns a ready-to-use, empty CheckRegistry.
	return nil
}

func New() CheckRegistry { _ = "STUB: not implemented"; return *new(CheckRegistry) }
