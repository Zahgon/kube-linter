package mocks

import (
	"testing"

	coreV1 "k8s.io/api/core/v1"
)

// AddMockService adds a mock Service to LintContext
func (l *MockLintContext) AddMockService(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyService modifies a given service in the context via the passed function
func (l *MockLintContext) ModifyService(t *testing.T, name string, f func(service *coreV1.Service)) {
	_ = "STUB: not implemented"
	return
}
