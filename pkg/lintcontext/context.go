package lintcontext

import (
	"encoding/json"

	"golang.stackrox.io/kube-linter/pkg/k8sutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ObjectMetadata is metadata about an object.
type ObjectMetadata struct {
	FilePath string
	Raw      []byte `json:"-"`
}

// An Object references an object that is loaded from a YAML file.
type Object struct {
	Metadata  ObjectMetadata
	K8sObject k8sutil.Object `json:"-"`
}

// K8sObjectInfo contains identifying information about k8s object.
type K8sObjectInfo struct {
	Namespace, Name  string
	GroupVersionKind schema.GroupVersionKind
}

// GetK8sObjectName extracts K8sObjectInfo from Object.K8sObject.
func (o *Object) GetK8sObjectName() K8sObjectInfo {
	_ = "STUB: not implemented"
	return *new(K8sObjectInfo)
}

// String provides plain-text representation of k8s object name.
func (n K8sObjectInfo) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON provides custom serialization for Object.
// Object.K8sObject is not serialized directly because that would be too much data. This function limits output to only
// K8sObjectInfo returned for K8sObject.
func (o *Object) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// AliasedObject allows including all Object data without running MarshalJSON (this same function) on it in
	// an infinite loop.
	return nil, nil
}

// Check that *Object implements json.Marshaler interface.
var _ json.Marshaler = (*Object)(nil)

// An InvalidObject represents something that couldn't be loaded from a YAML file.
type InvalidObject struct {
	Metadata ObjectMetadata
	LoadErr  error
}

// A LintContext represents the context for a lint run.
type LintContext interface {
	Objects() []Object
	InvalidObjects() []InvalidObject
}

type lintContextImpl struct {
	objects        []Object
	invalidObjects []InvalidObject

	customDecoder runtime.Decoder
}

// Objects returns the (valid) objects loaded from this LintContext.
func (l *lintContextImpl) Objects() []Object {
	_ = "STUB: not implemented"

	// addObject adds a valid object to this LintContext
	return nil
}

func (l *lintContextImpl) addObjects(objs ...Object) { _ = "STUB: not implemented"; return }

// InvalidObjects returns any objects that we attempted to load, but which were invalid.
func (l *lintContextImpl) InvalidObjects() []InvalidObject { _ = "STUB: not implemented"; return nil }

// addInvalidObject adds an invalid object to this LintContext
func (l *lintContextImpl) addInvalidObjects(objs ...InvalidObject) {
	_ = "STUB: not implemented"
	return
}

// new returns a ready-to-use, empty, lintContextImpl.
func newCtx(options Options) *lintContextImpl { _ = "STUB: not implemented"; return nil }
