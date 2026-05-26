package objectkinds

import (
	autoscalingV1 "k8s.io/api/autoscaling/v1"
	autoscalingV2 "k8s.io/api/autoscaling/v2"
	autoscalingV2Beta1 "k8s.io/api/autoscaling/v2beta1"
	autoscalingV2Beta2 "k8s.io/api/autoscaling/v2beta2"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// HorizontalPodAutoscaler represents Kubernetes HorizontalPodAutoscaler objects. Case sensitive.
	HorizontalPodAutoscaler = "HorizontalPodAutoscaler"
)

var (
	horizontalPodAutoscalerV2Beta1GVK = autoscalingV2Beta1.SchemeGroupVersion.WithKind(HorizontalPodAutoscaler)
	horizontalPodAutoscalerV2Beta2GVK = autoscalingV2Beta2.SchemeGroupVersion.WithKind(HorizontalPodAutoscaler)
	horizontalPodAutoscalerV2GVK      = autoscalingV2.SchemeGroupVersion.WithKind(HorizontalPodAutoscaler)
	horizontalPodAutoscalerV1GVK      = autoscalingV1.SchemeGroupVersion.WithKind(HorizontalPodAutoscaler)
)

func isHorizontalPodAutoscaler(gvk schema.GroupVersionKind) bool {
	_ = "STUB: not implemented"
	return false
}

func init() {
	RegisterObjectKind(HorizontalPodAutoscaler, MatcherFunc(isHorizontalPodAutoscaler))
}

// GetHorizontalPodAutoscalerAPIVersion returns HorizontalPodAutoscaler's APIVersion
func GetHorizontalPodAutoscalerAPIVersion(version string) string {
	_ = "STUB: not implemented"
	return ""
}
