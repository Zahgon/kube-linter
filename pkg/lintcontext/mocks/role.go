package mocks

import (
	"testing"

	rbacV1 "k8s.io/api/rbac/v1"
)

// AddMockRole adds a mock Role to LintContext
func (l *MockLintContext) AddMockRole(t *testing.T, name, namespace string) {
	_ = "STUB: not implemented"
	return
}

// ModifyRole modifies a given Role in the context via the passed function.
func (l *MockLintContext) ModifyRole(t *testing.T, name string, f func(role *rbacV1.Role)) {
	_ = "STUB: not implemented"
	return
}
