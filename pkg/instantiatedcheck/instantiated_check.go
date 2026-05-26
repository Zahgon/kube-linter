package instantiatedcheck

import (
	"regexp"

	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
)

// An InstantiatedCheck is the runtime instantiation of a check, which fuses the metadata in a check
// spec with the runtime information from a template.
type InstantiatedCheck struct {
	Func    check.Func
	Matcher objectkinds.Matcher

	Spec config.Check
}

var (
	validCheckNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
)

// ValidateAndInstantiate validates the check, and creates an instantiated check if the check
// is valid.
func ValidateAndInstantiate(c *config.Check) (*InstantiatedCheck, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
