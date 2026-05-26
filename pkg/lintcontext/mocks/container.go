package mocks

import (
	"testing"

	v1 "k8s.io/api/core/v1"
)

// AddContainerToDeployment adds a mock container to the specified pod under context
func (l *MockLintContext) AddContainerToDeployment(t *testing.T, deploymentName string, container v1.Container) {
	_ = "STUB: not implemented"
	return
}

// TODO: keep supporting other fields
