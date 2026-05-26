package objectkinds

import (
	kedaV1Alpha1 "golang.stackrox.io/kube-linter/pkg/crds/keda/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// ScaledObject represents Kubernetes ScaledObject objects. Case sensitive.
	ScaledObject = "ScaledObject"
)

var (
	ScaledObjectV1Alpha1 = kedaV1Alpha1.SchemeGroupVersion.WithKind(ScaledObject)
)

func isScaledObject(gvk schema.GroupVersionKind) bool { _ = "STUB: not implemented"; return false }

func init() {
	RegisterObjectKind(ScaledObject, MatcherFunc(isScaledObject))
}

// GetScaledObjectAPIVersion returns ScaledObject's APIVersion
func GetScaledObjectAPIVersion(version string) string { _ = "STUB: not implemented"; return "" }
