package ignore

const (
	// AnnotationKeyPrefix is the prefix for annotations for kube-linter check ignores.
	AnnotationKeyPrefix = "ignore-check.kube-linter.io/"

	// AllAnnotationKey is used to ignore all checks for a given object.
	AllAnnotationKey = "kube-linter.io/ignore-all"
)

// ObjectForCheck returns whether to ignore the given object for the passed check name.
func ObjectForCheck(annotations map[string]string, checkName string) bool {
	_ = "STUB: not implemented"
	return false
}
