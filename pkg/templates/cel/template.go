package cel

import (
	"fmt"

	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/cel/internal/params"
)

const (
	templateKey = "cel-expression"
)

func init() {
	templates.Register(check.Template{
		HumanName:   "CEL",
		Key:         templateKey,
		Description: "Flag objects with CEL expression",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.Any},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(p params.Params) (check.Func, error) {
			return func(ctx lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				msg, err := evaluate(p.Check, object, ctx.Objects())
				if err != nil {
					return []diagnostic.Diagnostic{
						{Message: fmt.Sprintf("error evaluating CEL check expression: %v", err)},
					}
				}
				if msg != "" {
					return []diagnostic.Diagnostic{
						{Message: fmt.Sprintf("CEL check expression returned: %v", msg)},
					}
				}
				return nil
			}, nil
		}),
	})
}

func evaluate(check string, object lintcontext.Object, objects []lintcontext.Object) (string, error) {
	_ = "STUB: not implemented"
	// Convert object to map via JSON marshaling/unmarshaling for CEL compatibility
	// We need to marshal the underlying K8sObject, not the lintcontext.Object
	return "", nil
}

// Convert objects to maps via JSON marshaling/unmarshaling

func toMap(obj any) (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }
