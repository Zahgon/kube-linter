package util

import (
	"golang.stackrox.io/kube-linter/pkg/check"
)

// ConstructForbiddenMapMatcher constructs a check function that requires that a k-v pair is NOT present in the map.
func ConstructForbiddenMapMatcher(key, value, fieldType string) (check.Func, error) {
	_ = "STUB: not implemented"
	return *new(check.Func), nil
}
