package util

import (
	"golang.stackrox.io/kube-linter/pkg/check"
)

// ConstructRequiredMapMatcher constructs a check function that requires that a k-v pair is present in the map.
func ConstructRequiredMapMatcher(key, value, fieldType string) (check.Func, error) {
	_ = "STUB: not implemented"
	return *new(check.Func), nil
}
