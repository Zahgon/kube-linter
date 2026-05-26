package danglinghpa

import (
	"fmt"

	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/extract"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/danglinghpa/internal/params"
)

const (
	templateKey = "dangling-horizontalpodautoscaler"
)

type objectReference struct {
	Kind       string
	Name       string
	APIVersion string
}

func getTargetFromHpa(obj lintcontext.Object) (*objectReference, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func init() {
	templates.Register(check.Template{
		HumanName:   "Dangling HorizontalPodAutoscalers",
		Key:         templateKey,
		Description: "Flag HorizontalPodAutoscalers that target a resource that does not exist",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.HorizontalPodAutoscaler},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(p params.Params) (check.Func, error) {
			return func(lintCtx lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				target, found := getTargetFromHpa(object)
				if !found {
					return nil
				}
				for _, obj := range lintCtx.Objects() {
					_, hasPods := extract.PodTemplateSpec(obj.K8sObject)
					if !hasPods {
						continue
					}
					k8sObj := obj.K8sObject
					gvk := k8sObj.GetObjectKind().GroupVersionKind()
					if target.Name == obj.GetK8sObjectName().Name && target.Kind == gvk.Kind && target.APIVersion == gvk.Group+"/"+gvk.Version {
						// Found!
						return nil
					}
				}
				return []diagnostic.Diagnostic{{Message: fmt.Sprintf("no resources found matching HorizontalPodAutoscaler scaleTargetRef (%v)", *target)}}
			}, nil
		}),
	})
}
