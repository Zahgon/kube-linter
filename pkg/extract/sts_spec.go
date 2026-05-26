package extract

import (
	"golang.stackrox.io/kube-linter/pkg/k8sutil"
	appsV1 "k8s.io/api/apps/v1"
)

func StatefulSetSpec(obj k8sutil.Object) (appsV1.StatefulSetSpec, bool) {
	_ = "STUB: not implemented"
	return *new(appsV1.StatefulSetSpec), false
}
