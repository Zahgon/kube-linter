package util

import (
	"golang.stackrox.io/kube-linter/pkg/check"
)

// MustParseParameterDesc unmarshals the given JSON into a templates.ParameterDesc.
func MustParseParameterDesc(asJSON string) check.ParameterDesc {
	_ = "STUB: not implemented"
	return *new(check.ParameterDesc)
}
