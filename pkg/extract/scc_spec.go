package extract

import (
	"golang.stackrox.io/kube-linter/pkg/k8sutil"
)

// SCCallowPrivilegedContainer extracts allowPrivilegedContainer from the given object, if available.
func SCCallowPrivilegedContainer(obj k8sutil.Object) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}
