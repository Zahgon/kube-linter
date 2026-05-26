package wildcardinrules

import (
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/wildcardinrules/internal/params"
	rbacV1 "k8s.io/api/rbac/v1"
)

func init() {
	templates.Register(check.Template{
		HumanName:   "Wildcard Use in Role and ClusterRole Rules",
		Key:         "wildcard-in-rules",
		Description: "Flag Roles and ClusterRoles that use wildcard * in rules",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{
				objectkinds.Role,
				objectkinds.ClusterRole},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(_ params.Params) (check.Func, error) {
			return func(lintCtx lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				var results []diagnostic.Diagnostic
				role, ok := object.K8sObject.(*rbacV1.Role)
				if ok {
					results = append(results, findWildCard(role.Rules)...)
				}
				crole, ok := object.K8sObject.(*rbacV1.ClusterRole)
				if ok {
					results = append(results, findWildCard(crole.Rules)...)
				}
				return results
			}, nil
		}),
	})
}

// find wildcards used in rules
func findWildCard(rules []rbacV1.PolicyRule) []diagnostic.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}
