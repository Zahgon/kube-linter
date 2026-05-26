package mocks

import (
	"testing"

	batchV1 "k8s.io/api/batch/v1"
)

// AddMockJob adds a mock Job to LintContext
func (l *MockLintContext) AddMockJob(t *testing.T, name string) { _ = "STUB: not implemented"; return }

// ModifyJob modifies a given Job in the context via the passed function
func (l *MockLintContext) ModifyJob(t *testing.T, name string, f func(job *batchV1.Job)) {
	_ = "STUB: not implemented"
	return
}
