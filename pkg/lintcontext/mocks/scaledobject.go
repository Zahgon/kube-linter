package mocks

import (
	"testing"

	kedaV1Alpha1 "golang.stackrox.io/kube-linter/pkg/crds/keda/v1alpha1"
)

// AddMockScaledObject adds a mock ScaledObject to LintContext
func (l *MockLintContext) AddMockScaledObject(t *testing.T, name, version string) {
	_ = "STUB: not implemented"
	return
}

// ModifyScaledObjectV1Alpha1 modifies a given ScaledObject in the context via the passed function.
func (l *MockLintContext) ModifyScaledObjectV1Alpha1(t *testing.T, name string, f func(hpa *kedaV1Alpha1.ScaledObject)) {
	_ = "STUB: not implemented"
	return
}
