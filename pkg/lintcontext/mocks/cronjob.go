package mocks

import (
	"testing"

	batchV1 "k8s.io/api/batch/v1"
)

// AddMockCronJob adds a mock CronJob to LintContext
func (l *MockLintContext) AddMockCronJob(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyCronJob modifies a given CronJob in the context via the passed function
func (l *MockLintContext) ModifyCronJob(t *testing.T, name string, f func(cronjob *batchV1.CronJob)) {
	_ = "STUB: not implemented"
	return
}
