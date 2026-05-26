package accesstoresources

import (
	"fmt"
	"regexp"

	"golang.stackrox.io/kube-linter/internal/stringutils"
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/accesstoresources/internal/params"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	templateKey = "access-to-resources"
)

func init() {
	templates.Register(check.Template{
		HumanName:   "Access to Resources",
		Key:         templateKey,
		Description: "Flag cluster role bindings and role bindings that grant access to the specified resource kinds and verbs",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{
				objectkinds.Role,
				objectkinds.ClusterRole,
				objectkinds.ClusterRoleBinding,
				objectkinds.RoleBinding},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(p params.Params) (check.Func, error) {
			resourceRegexes := make([]*regexp.Regexp, 0, len(p.Resources))
			for _, res := range p.Resources {
				r, err := regexp.Compile(res)
				if err != nil {
					return nil, fmt.Errorf("invalid regex %s: %w", res, err)
				}
				resourceRegexes = append(resourceRegexes, r)
			}
			verbRegexes := make([]*regexp.Regexp, 0, len(p.Verbs))
			for _, verb := range p.Verbs {
				v, err := regexp.Compile(verb)
				if err != nil {
					return nil, fmt.Errorf("invalid regex %s: %w", verb, err)
				}
				verbRegexes = append(verbRegexes, v)
			}
			return func(lintCtx lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				rbinding, ok := object.K8sObject.(*rbacV1.RoleBinding)
				if ok {
					namespace := stringutils.OrDefault(rbinding.Namespace, "default")
					return findRole(rbinding.RoleRef.Name, namespace, lintCtx, resourceRegexes, verbRegexes, p.FlagRolesNotFound)
				}
				crbinding, ok := object.K8sObject.(*rbacV1.ClusterRoleBinding)
				if ok {
					return findClusterRole(crbinding.RoleRef.Name, lintCtx, resourceRegexes, verbRegexes, p.FlagRolesNotFound)
				}
				return nil
			}, nil
		}),
	})
}

// find clusterrole by name, and check if it has access to the specified resource kinds and verbs
func findClusterRole(name string, lintCtx lintcontext.LintContext, resourceRegexes, verbRegexes []*regexp.Regexp, flag bool) []diagnostic.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}

// find clusterroles by label selectors, and check if they have access to the specified resources and verbs
func findAggregatedAccesses(clusterroles []*rbacV1.ClusterRole, selectors []metaV1.LabelSelector, resourceRegexes, verbRegexes []*regexp.Regexp) []diagnostic.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}

// Found the aggregated clusterrole!

// find role by name and namespace that has access to the specified resources and verbs
func findRole(name, namespace string, lintCtx lintcontext.LintContext, resources, verbs []*regexp.Regexp, flag bool) []diagnostic.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}

// find access verbs to a given resource kind
func checkAccess(rules []rbacV1.PolicyRule, resourceRegex, verbRegex []*regexp.Regexp) []string {
	_ = "STUB: not implemented"
	return nil
}

// isInList returns true if a match found in the list for the given name or a wildcard
func isInList(regexlist []*regexp.Regexp, name string) bool {
	_ = "STUB: not implemented"
	return false
}
