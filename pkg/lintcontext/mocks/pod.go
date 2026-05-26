package mocks

import (
	"testing"

	ocsAppsV1 "github.com/openshift/api/apps/v1"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	pdbV1 "k8s.io/api/policy/v1"
)

// AddMockDeployment adds a mock Deployment to LintContext
func (l *MockLintContext) AddMockDeployment(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyDeployment modifies a given deployment in the context via the passed function.
func (l *MockLintContext) ModifyDeployment(t *testing.T, name string, f func(deployment *appsV1.Deployment)) {
	_ = "STUB: not implemented"
	return
}

// AddMockDaemonSet adds a mock DaemonSet to LintContext
func (l *MockLintContext) AddMockDaemonSet(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyDaemonSet modifies a given DaemonSet in the context via the passed function.
func (l *MockLintContext) ModifyDaemonSet(t *testing.T, name string, f func(ds *appsV1.DaemonSet)) {
	_ = "STUB: not implemented"
	return
}

// AddMockDeploymentConfig adds a mock DeploymentConfig to LintContext
func (l *MockLintContext) AddMockDeploymentConfig(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyDeploymentConfig modifies a given DeploymentConfig in the context via the passed function.
func (l *MockLintContext) ModifyDeploymentConfig(t *testing.T, name string, f func(ds *ocsAppsV1.DeploymentConfig)) {
	_ = "STUB: not implemented"
	return
}

func (l *MockLintContext) AddMockPodDisruptionBudget(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

func (l *MockLintContext) ModifyPodDisruptionBudget(t *testing.T, name string, f func(pdb *pdbV1.PodDisruptionBudget)) {
	_ = "STUB: not implemented"
	return
}

// AddSecurityContextToDeployment adds a security context to the deployment specified by name
func (l *MockLintContext) AddSecurityContextToDeployment(t *testing.T, deploymentName string,
	securityContext *v1.PodSecurityContext) {
	_ = "STUB: not implemented"
	return
}

// AddMockReplicationController adds a mock ReplicationController to LintContext
func (l *MockLintContext) AddMockReplicationController(t *testing.T, name string) {
	_ = "STUB: not implemented"
	return
}

// ModifyReplicationController modifies a given replication controller in the context via the passed function.
func (l *MockLintContext) ModifyReplicationController(t *testing.T, name string, f func(deployment *v1.ReplicationController)) {
	_ = "STUB: not implemented"
	return
}
