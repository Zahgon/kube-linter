package util

import (
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	v1 "k8s.io/api/core/v1"
)

var sentinel = struct{}{}

func CheckProbePort(container *v1.Container, probe *v1.Probe) []diagnostic.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}
