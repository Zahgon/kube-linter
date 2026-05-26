package mocks

import (
	"testing"

	networkingV1 "k8s.io/api/networking/v1"
)

func (l *MockLintContext) AddMockIngress(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyIngress modifies a given networkpolicy in the context via the passed function.
func (l *MockLintContext) ModifyIngress(t *testing.T, name string, f func(ingress *networkingV1.Ingress)) {
	_ = "STUB: not implemented"
	return
}
