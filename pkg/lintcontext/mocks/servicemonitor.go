package mocks

import (
	"testing"

	k8sMonitoring "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

// AddMockServiceMonitor adds a mock ServiceMonitor to LintContext
func (l *MockLintContext) AddMockServiceMonitor(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyServiceMonitor modifies a given servicemonitor in the context via the passed function
func (l *MockLintContext) ModifyServiceMonitor(t *testing.T, name string, f func(servicemonitor *k8sMonitoring.ServiceMonitor)) {
	_ = "STUB: not implemented"
	return
}
