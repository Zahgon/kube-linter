package sortedkeys

import (
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/parser"
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/sortedkeys/internal/params"
)

const templateKey = "sorted-keys"

func init() {
	templates.Register(check.Template{
		HumanName:   "Sorted Keys",
		Key:         templateKey,
		Description: "Flag YAML keys that are not sorted in alphabetical order",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.Any},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(p params.Params) (check.Func, error) {
			return func(_ lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				// Parse the raw YAML to preserve key order
				file, err := parser.ParseBytes(object.Metadata.Raw, parser.ParseComments)
				if err != nil {
					// Skip objects that can't be parsed
					return nil
				}

				var diagnostics []diagnostic.Diagnostic

				// Check all documents in the file
				for _, doc := range file.Docs {
					if doc != nil && doc.Body != nil {
						diagnostics = append(diagnostics, checkNode(doc.Body, "", p.Recursive)...)
					}
				}

				return diagnostics
			}, nil
		}),
	})
}

// checkNode recursively checks if keys in a YAML node are sorted
func checkNode(node ast.Node, path string, recursive bool) []diagnostic.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}

// MappingNode contains Values which are MappingValueNodes

// Values in a MappingNode are already MappingValueNode pointers
// Extract the key

// Check if keys are sorted

// Find the first key that is out of order

// Only report once per level

// Recursively check child nodes if recursive is enabled

// For sequences, check each element if it's a mapping

// Handle anchor nodes by checking their value

// Skip alias nodes - they reference already checked content
// No need to check as the original anchor was already checked

// Handle merge keys (<<: *alias)
// The merge key itself is represented as a special key
// No special handling needed as it will be treated as a key

// getKeyString extracts the string representation of a key node
func getKeyString(node ast.Node) string { _ = "STUB: not implemented"; return "" }
