package updateconfig

import (
	"fmt"
	"regexp"

	"golang.stackrox.io/kube-linter/internal/errorhelpers"
	"golang.stackrox.io/kube-linter/internal/stringutils"
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/extract"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/updateconfig/internal/params"

	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	templateKey = "update-configuration"
)

func parseIntOrString(data string) (*intstr.IntOrString, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is not an integer.  Is it a valid string?

// Not going to try harder

func compareIntOrString(maximal, minimal, actual *intstr.IntOrString) bool {
	_ = "STUB: not implemented"
	return false
}

func conditional(firstCond bool, firstStr string, secondCond bool, secondStr, bothStr string) string {
	_ = "STUB: not implemented"
	return ""
}

func needsRollingUpdateDefinition(p params.Params) bool { _ = "STUB: not implemented"; return false }

func init() {
	templates.Register(check.Template{
		HumanName:   "Update configuration",
		Key:         templateKey,
		Description: "Flag configurations that do not meet the specified update configuration",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.DeploymentLike},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(p params.Params) (check.Func, error) {
			var maxPodsUnavailable, minPodsUnavailable, maxSurge, minSurge *intstr.IntOrString
			errorList := errorhelpers.NewErrorList("configuration verification")
			compiledRegex, err := regexp.Compile(p.StrategyTypeRegex)
			if err != nil {
				errorList.AddWrapf(err, "invalid regex %s", p.StrategyTypeRegex)
			}
			if p.MaxPodsUnavailable != "" {
				maxPodsUnavailable, err = parseIntOrString(p.MaxPodsUnavailable)
				if err != nil {
					errorList.AddWrapf(err, "invalid MaxPodsUnavailable %s", p.MaxPodsUnavailable)
				}
			}
			if p.MinPodsUnavailable != "" {
				minPodsUnavailable, err = parseIntOrString(p.MinPodsUnavailable)
				if err != nil {
					errorList.AddWrapf(err, "invalid MinPodsUnavailable %s", p.MinPodsUnavailable)
				}
			}
			if p.MaxSurge != "" {
				maxSurge, err = parseIntOrString(p.MaxSurge)
				if err != nil {
					errorList.AddWrapf(err, "invalid MaxSurge %s", p.MaxSurge)
				}
			}
			if p.MinSurge != "" {
				minSurge, err = parseIntOrString(p.MinSurge)
				if err != nil {
					errorList.AddWrapf(err, "invalid MinSurge %s", p.MinSurge)
				}
			}
			if err := errorList.ToError(); err != nil {
				return nil, err
			}
			return func(_ lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				var diagnostics []diagnostic.Diagnostic

				strategy, found := extract.UpdateStrategy(object.K8sObject)
				if !found {
					return nil
				}
				if !strategy.TypeExists {
					return nil
				}
				if !compiledRegex.MatchString(strategy.Type) {
					newD := diagnostic.Diagnostic{
						Message: fmt.Sprintf("object has %s strategy type but must match regex %s",
							stringutils.Ternary(strategy.Type != "", strategy.Type, "no"), p.StrategyTypeRegex)}
					diagnostics = append(diagnostics, newD)
				}
				if !strategy.RollingConfigExists {
					return nil
				}
				if needsRollingUpdateDefinition(p) && !strategy.RollingConfigValid {
					newD := diagnostic.Diagnostic{Message: "object has no rolling update parameters defined"}
					diagnostics = append(diagnostics, newD)
				}
				if strategy.MaxUnavailableExists {
					if !compareIntOrString(maxPodsUnavailable, minPodsUnavailable, strategy.MaxUnavailable) {
						minStr := fmt.Sprintf("at least %s", p.MinPodsUnavailable)
						maxStr := fmt.Sprintf("no more than %s", p.MaxPodsUnavailable)
						msg := fmt.Sprintf("object has a max unavailable of %s but %s is required", strategy.MaxUnavailable.String(),
							conditional(p.MinPodsUnavailable != "", minStr, p.MaxPodsUnavailable != "", maxStr, " and "))
						newD := diagnostic.Diagnostic{Message: msg}
						diagnostics = append(diagnostics, newD)
					}
				}
				if strategy.MaxSurgeExists {
					if !compareIntOrString(maxSurge, minSurge, strategy.MaxSurge) {
						minStr := fmt.Sprintf("at least %s", p.MinSurge)
						maxStr := fmt.Sprintf("no more than %s", p.MaxSurge)
						msg := fmt.Sprintf("object has a max surge of %s but %s is required", strategy.MaxSurge.String(),
							conditional(p.MinSurge != "", minStr, p.MaxSurge != "", maxStr, " and "))
						newD := diagnostic.Diagnostic{Message: msg}
						diagnostics = append(diagnostics, newD)
					}
				}
				if len(diagnostics) == 0 {
					return nil
				}
				return diagnostics
			}, nil
		}),
	})
}
