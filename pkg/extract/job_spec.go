package extract

import (
	"golang.stackrox.io/kube-linter/pkg/k8sutil"
	batchV1 "k8s.io/api/batch/v1"
)

// JobSpec extracts a job template spec from Job or CronJob objects
func JobSpec(obj k8sutil.Object) (batchV1.JobSpec, string, bool) {
	_ = "STUB: not implemented"
	return *new(batchV1.JobSpec), "", false
}
