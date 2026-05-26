package mocks

import (
	"testing"

	rbacV1 "k8s.io/api/rbac/v1"
)

// AddMockClusterRoleBinding adds a mock ClusterRoleBinding to LintContext
func (l *MockLintContext) AddMockClusterRoleBinding(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyClusterRoleBinding modifies a given ClusterRoleBinding in the context via the passed function.
func (l *MockLintContext) ModifyClusterRoleBinding(t *testing.T, name string, f func(clusterrolebinding *rbacV1.ClusterRoleBinding)) {
	_ = "STUB: not implemented"
	return
}
