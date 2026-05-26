package mocks

import (
	"testing"

	rbacV1 "k8s.io/api/rbac/v1"
)

// AddMockRoleBinding adds a mock RoleBinding to LintContext
func (l *MockLintContext) AddMockRoleBinding(t *testing.T, name, namespace string) {
	_ = "STUB: not implemented"
	return
}

// ModifyRoleBinding modifies a given RoleBinding in the context via the passed function.
func (l *MockLintContext) ModifyRoleBinding(t *testing.T, name string, f func(rolebinding *rbacV1.RoleBinding)) {
	_ = "STUB: not implemented"
	return
}
