package mocks

import (
	"testing"

	rbacV1 "k8s.io/api/rbac/v1"
)

// AddMockClusterRole adds a mock ClusterRole to LintContext
func (l *MockLintContext) AddMockClusterRole(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyClusterRole modifies a given clusterrole in the context via the passed function.
func (l *MockLintContext) ModifyClusterRole(t *testing.T, name string, f func(clusterrole *rbacV1.ClusterRole)) {
	_ = "STUB: not implemented"
	return
}
