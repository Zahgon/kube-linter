package objectkinds

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	allObjectKinds = make(map[string]Matcher)
)

// RegisterObjectKind allows a matcher function to be registered for a given object kind
func RegisterObjectKind(name string, objectKind Matcher) { _ = "STUB: not implemented"; return }

// AllObjectKinds will return all the object kind names that are registered
func AllObjectKinds() []string { _ = "STUB: not implemented"; return nil }

type orMatcher []Matcher

func (o orMatcher) Matches(gvk schema.GroupVersionKind) bool {
	_ = "STUB: not implemented"
	return false
}

// ConstructMatcher constructs a matcher that matches objects that fall
// into one of the given object kinds.
func ConstructMatcher(objectKinds ...string) (Matcher, error) {
	_ = "STUB: not implemented"
	return *new(Matcher), nil
}
