package mocks

import (
	"testing"

	networkingV1 "k8s.io/api/networking/v1"
)

// AddMockNetworkPolicy  adds a mock NetworkPolicy to LintContext
func (l *MockLintContext) AddMockNetworkPolicy(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyNetworkPolicy modifies a given networkpolicy in the context via the passed function.
func (l *MockLintContext) ModifyNetworkPolicy(t *testing.T, name string, f func(networkpolicy *networkingV1.NetworkPolicy)) {
	_ = "STUB: not implemented"
	return
}
