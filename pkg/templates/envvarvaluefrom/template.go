package envvarvaluefrom

import (
	"regexp"

	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/envvarvaluefrom/internal/params"
	v1 "k8s.io/api/core/v1"
)

const (
	templateKey = "env-value-from"
)

type resourceInfo struct {
	name     string
	key      string
	optional *bool
}

type resourceType int

const (
	resourceTypeSecret resourceType = iota
	resourceTypeConfigMap
)

type resourceChecker struct {
	objType      string
	objMap       map[string]interface{}
	getKeys      func(interface{}) []string
	ignoredRegex []*regexp.Regexp
}

func init() {
	templates.Register(check.Template{
		HumanName:   "Env references",
		Key:         templateKey,
		Description: "Flag resources which use env variables from secrets/configmaps not included in the release",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.DeploymentLike},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(p params.Params) (check.Func, error) {
			ignoredSecrets, err := extractRegexList(p.IgnoredSecrets)
			if err != nil {
				return nil, err
			}
			ignoredConfigMaps, err := extractRegexList(p.IgnoredConfigMaps)
			if err != nil {
				return nil, err
			}
			return func(lintCtx lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				secrets := make(map[string]*v1.Secret)
				configmaps := make(map[string]*v1.ConfigMap)
				for _, obj := range lintCtx.Objects() {
					if secret, found := obj.K8sObject.(*v1.Secret); found {
						secrets[secret.Name] = secret // Fix: Remove ObjectMeta
					}
					if configmap, found := obj.K8sObject.(*v1.ConfigMap); found {
						configmaps[configmap.Name] = configmap // Fix: Remove ObjectMeta
					}
				}
				return lintForEachContainer(lintCtx, object, ignoredSecrets, ignoredConfigMaps, secrets, configmaps)
			}, nil
		}),
	})
}

func lintForEachContainer(lintCtx lintcontext.LintContext, object lintcontext.Object, ignoredSecrets, ignoredConfigMaps []*regexp.Regexp, secrets map[string]*v1.Secret, configmaps map[string]*v1.ConfigMap) []diagnostic.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}

func checkResourceReference(containerName string, ref resourceInfo, checker *resourceChecker) string {
	_ = "STUB: not implemented"
	return ""
}

func isInRegexList(regexlist []*regexp.Regexp, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func isInList(regexlist []string, name string) bool { _ = "STUB: not implemented"; return false }

func getSecretKeys(obj interface{}) []string { _ = "STUB: not implemented"; return nil }

func getConfigMapKeys(obj interface{}) []string { _ = "STUB: not implemented"; return nil }

func extractRegexList(inputList []string) ([]*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
