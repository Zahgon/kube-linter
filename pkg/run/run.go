package run

import (
	"time"

	"golang.stackrox.io/kube-linter/pkg/checkregistry"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
)

// CheckStatus is enum type.
type CheckStatus string

const (
	// ChecksPassed means no lint errors found.
	ChecksPassed CheckStatus = "Passed"
	// ChecksFailed means lint errors were found.
	ChecksFailed CheckStatus = "Failed"
)

// Result represents the result from a run of the linter.
type Result struct {
	Checks  []config.Check
	Reports []diagnostic.WithContext
	Summary Summary
}

// Summary holds information about the linter run overall.
type Summary struct {
	ChecksStatus      CheckStatus
	CheckEndTime      time.Time
	KubeLinterVersion string
}

// Run runs the linter on the given context, with the given config.
func Run(lintCtxs []lintcontext.LintContext, registry checkregistry.CheckRegistry, checks []string) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}
