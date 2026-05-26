package templates

import (
	"golang.stackrox.io/kube-linter/pkg/check"
)

var (
	allTemplates = make(map[string]check.Template)
)

// Register registers a template with the given name.
// Intended to be called at program init time.
func Register(t check.Template) { _ = "STUB: not implemented"; return }

// Get gets a template by name, returning a boolean indicating whether it was found.
func Get(name string) (check.Template, bool) {
	_ = "STUB: not implemented"
	return *new(check.Template), false
}

// List returns all known templates, sorted by name.
func List() []check.Template { _ = "STUB: not implemented"; return nil }
