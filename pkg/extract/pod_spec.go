package extract

import (
	"golang.stackrox.io/kube-linter/pkg/extract/customtypes"
	"golang.stackrox.io/kube-linter/pkg/k8sutil"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodTemplateSpec extracts a pod template spec from the given object, if available.
func PodTemplateSpec(obj k8sutil.Object) (coreV1.PodTemplateSpec, bool) {
	_ = "STUB: not implemented"
	return *new(coreV1.PodTemplateSpec), false
}

// PodSpec extracts a pod spec from the given object, if available.
func PodSpec(obj k8sutil.Object) (customtypes.PodSpec, bool) {
	_ = "STUB: not implemented"
	return *new(customtypes.PodSpec), false
}

// Selector extracts a selector from the given object, if available.
func Selector(obj k8sutil.Object) (*metaV1.LabelSelector, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Replicas extracts replicas from the given object, if available.
func Replicas(obj k8sutil.Object) (int32, bool) {
	_ = "STUB: not implemented"
	// DeploymentConfigs are treated specially because the number of replicas is
	// an int32, not a *int32.
	return 0, false
}

// If numReplicas is a `nil` pointer, then it defaults to 1.
