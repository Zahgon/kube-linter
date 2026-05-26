package targetport

import (
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/extract"
	"golang.stackrox.io/kube-linter/pkg/extract/customtypes"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/targetport/internal/params"
	coreV1 "k8s.io/api/core/v1"
)

const (
	templateKey = "target-port"
)

func init() {
	templates.Register(check.Template{
		HumanName:   "Target Port",
		Key:         templateKey,
		Description: "Flag containers and services using not allowed port names or numbers",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.DeploymentLike, objectkinds.Service},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(_ params.Params) (check.Func, error) {
			return func(_ lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				podSpec, foundPodSpec := extract.PodSpec(object.K8sObject)
				if foundPodSpec {
					return findPodPorts(&podSpec)
				}

				service, foundService := object.K8sObject.(*coreV1.Service)
				if foundService {
					return findServicePorts(service)
				}

				return nil
			}, nil

		}),
	})
}

func findPodPorts(podSpec *customtypes.PodSpec) []diagnostic.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}

func findServicePorts(service *coreV1.Service) []diagnostic.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}
