package mocks

import (
	"testing"

	autoscalingV1 "k8s.io/api/autoscaling/v1"
	autoscalingV2 "k8s.io/api/autoscaling/v2"
	autoscalingV2Beta1 "k8s.io/api/autoscaling/v2beta1"
	autoscalingV2Beta2 "k8s.io/api/autoscaling/v2beta2"
)

// AddMockHorizontalPodAutoscaler adds a mock HorizontalPodAutoscaler to LintContext
func (l *MockLintContext) AddMockHorizontalPodAutoscaler(t *testing.T, name, version string) {
	_ = "STUB: not implemented"
	return
}

// ModifyHorizontalPodAutoscalerV2Beta1 modifies a given HorizontalPodAutoscaler in the context via the passed function.
func (l *MockLintContext) ModifyHorizontalPodAutoscalerV2Beta1(t *testing.T, name string, f func(hpa *autoscalingV2Beta1.HorizontalPodAutoscaler)) {
	_ = "STUB: not implemented"
	return
}

// ModifyHorizontalPodAutoscalerV2Beta2 modifies a given HorizontalPodAutoscaler in the context via the passed function.
func (l *MockLintContext) ModifyHorizontalPodAutoscalerV2Beta2(t *testing.T, name string, f func(hpa *autoscalingV2Beta2.HorizontalPodAutoscaler)) {
	_ = "STUB: not implemented"
	return
}

// ModifyHorizontalPodAutoscalerV2 modifies a given HorizontalPodAutoscaler in the context via the passed function.
func (l *MockLintContext) ModifyHorizontalPodAutoscalerV2(t *testing.T, name string, f func(hpa *autoscalingV2.HorizontalPodAutoscaler)) {
	_ = "STUB: not implemented"
	return
}

// ModifyHorizontalPodAutoscalerV1 modifies a given HorizontalPodAutoscaler in the context via the passed function.
func (l *MockLintContext) ModifyHorizontalPodAutoscalerV1(t *testing.T, name string, f func(hpa *autoscalingV1.HorizontalPodAutoscaler)) {
	_ = "STUB: not implemented"
	return
}
