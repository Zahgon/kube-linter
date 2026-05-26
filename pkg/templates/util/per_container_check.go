package util

import (
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	v1 "k8s.io/api/core/v1"
)

// PerContainerCheck returns a check that abstracts away some of the boilerplate of writing a check
// that applies to containers. The given function is passed each container, and is allowed to return
// diagnostics if an error is found.
func PerContainerCheck(matchFunc func(container *v1.Container) []diagnostic.Diagnostic) check.Func {
	_ = "STUB: not implemented"
	return *new(check.Func)
}

// PerNonInitContainerCheck returns a check that abstracts away some of the boilerplate of writing a check
// that applies to all non-init containers. The given function is passed each non-init container,
// and is allowed to return diagnostics if an error is found.
func PerNonInitContainerCheck(matchFunc func(container *v1.Container) []diagnostic.Diagnostic) check.Func {
	_ = "STUB: not implemented"
	return *new(check.Func)
}
