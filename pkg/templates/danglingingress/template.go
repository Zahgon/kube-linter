package danglingingress

import (
	"fmt"

	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/danglingingress/internal/params"
	v1 "k8s.io/api/core/v1"
	networkingV1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	templateKey = "dangling-ingress"
)

type serviceDescriptor struct {
	name string
	port intstr.IntOrString
}

func getSelectorsFromIngressBackend(b *networkingV1.IngressBackend) (serviceDescriptor, bool) {
	_ = "STUB: not implemented"
	return *new(serviceDescriptor), false
}

func getSelectorsFromIngress(ingress *networkingV1.Ingress) map[serviceDescriptor]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	templates.Register(check.Template{
		HumanName:   "Dangling Ingress",
		Key:         templateKey,
		Description: "Flag ingress which do not match any service and port",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.Ingress},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(_ params.Params) (check.Func, error) {
			return func(lintCtx lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
				ingress, ok := object.K8sObject.(*networkingV1.Ingress)
				if !ok {
					return nil
				}

				selectors := getSelectorsFromIngress(ingress)

				// if there aren't any service selectors found we assume that the backend's are specified
				// by resources and will skip over this.
				if len(selectors) == 0 {
					return nil
				}

				for _, obj := range lintCtx.Objects() {
					if ingress.Namespace != obj.K8sObject.GetNamespace() {
						continue
					}

					service, ok := obj.K8sObject.(*v1.Service)
					if !ok {
						continue
					}

					for _, port := range service.Spec.Ports {
						desc := serviceDescriptor{
							name: service.Name,
							port: intstr.FromInt(int(port.Port)),
						}
						delete(selectors, desc)

						if port.Name != "" {
							desc := serviceDescriptor{
								name: service.Name,
								port: intstr.FromString(port.Name),
							}
							delete(selectors, desc)
						}

						if len(selectors) == 0 {
							// Found them all!
							return nil
						}
					}
				}

				var dig []diagnostic.Diagnostic

				for k := range selectors {
					dig = append(dig, diagnostic.Diagnostic{
						Message: fmt.Sprintf("no service found matching ingress label (%v), port %s", k.name, k.port.String()),
					})
				}

				return dig
			}, nil
		}),
	})
}
