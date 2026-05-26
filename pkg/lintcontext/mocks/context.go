package mocks

import (
	"golang.stackrox.io/kube-linter/pkg/k8sutil"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
)

// MockLintContext is mock implementation of the LintContext used in unit tests
type MockLintContext struct {
	objects    map[string]k8sutil.Object
	rawObjects map[string][]byte
}

// Objects returns all the objects under this MockLintContext
func (l *MockLintContext) Objects() []lintcontext.Object { _ = "STUB: not implemented"; return nil }

// InvalidObjects is not implemented. For now we don't care about invalid objects for mock context.
func (l *MockLintContext) InvalidObjects() []lintcontext.InvalidObject {
	_ = "STUB: not implemented"

	// NewMockContext returns an empty mockLintContext
	return nil
}

func NewMockContext() *MockLintContext { _ = "STUB: not implemented"; return nil }

// AddObject adds an object to the MockLintContext
func (l *MockLintContext) AddObject(key string, obj k8sutil.Object) {
	_ = "STUB: not implemented"
	return

	// AddObjectWithRaw adds an object to the MockLintContext with raw YAML data
}

func (l *MockLintContext) AddObjectWithRaw(key string, obj k8sutil.Object, raw []byte) {
	_ = "STUB: not implemented"
	return
}
