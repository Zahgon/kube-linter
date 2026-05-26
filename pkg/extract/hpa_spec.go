package extract

import (
	"golang.stackrox.io/kube-linter/pkg/k8sutil"
)

// HPAMinReplicas extracts minReplicas from the given object, if available.
func HPAMinReplicas(obj k8sutil.Object) (int32, bool) { _ = "STUB: not implemented"; return 0, false }

func checkReplicas(minReplicas *int32) (int32, bool) { _ = "STUB: not implemented"; return 0, false }

// If numReplicas is a `nil` pointer, then it defaults to 1.

// HPAScaleTargetRefName extracts Spec.ScaleTargetRef.Name
func HPAScaleTargetRefName(obj k8sutil.Object) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
