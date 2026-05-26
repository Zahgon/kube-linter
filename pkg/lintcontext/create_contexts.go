package lintcontext

import (
	"io"

	"golang.stackrox.io/kube-linter/internal/set"
	"k8s.io/apimachinery/pkg/runtime"
)

// ReadFromStdin is a path used to indicate reading from os.Stdin
const ReadFromStdin = "-"

var (
	knownYAMLExtensions    = set.NewFrozenStringSet(".yaml", ".yml")
	kustomizationFileNames = []string{"kustomization.yaml", "kustomization.yml"}
)

// Options represent values that can be provided to modify how objects are parsed to create lint contexts
type Options struct {
	// CustomDecoder allows users to supply a non-default decoder to parse k8s objects. This can be used
	// to allow the linter to create contexts for k8s custom resources
	CustomDecoder runtime.Decoder
}

// CreateContexts creates a context. Each context contains a set of files that should be linted
// as a group.
// Currently, each directory of Kube YAML files (or Helm charts) are treated as a separate context.
// TODO: Figure out if it's useful to allow people to specify that files spanning different directories
// should be treated as being in the same context.
func CreateContexts(ignorePaths []string, filesOrDirs ...string) ([]LintContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateContextsWithOptions creates a context with additional Options
func CreateContextsWithOptions(options Options, ignorePaths []string, filesOrDirs ...string) ([]LintContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Using doublestar to enable **
// See https://github.com/golang/go/issues/11862

// Load a file only if it ends in .yaml, OR it was explicitly passed by the user.

// Path has already been loaded, possibly through another argument. Skip.

// Path has already been loaded, possibly through another argument. Skip.

// CreateContextsFromHelmArchive creates a context from TGZ reader of Helm Chart.
// Note: although this function is not used in CLI, it is exposed from kube-linter library and therefore should stay.
// See https://github.com/stackrox/kube-linter/pull/173
func CreateContextsFromHelmArchive(ignorePaths []string, fileName string, tgzReader io.Reader) ([]LintContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isKustomizeDir checks if the given directory contains a kustomization.yaml or kustomization.yml file.
func isKustomizeDir(dirName string) bool { _ = "STUB: not implemented"; return false }
