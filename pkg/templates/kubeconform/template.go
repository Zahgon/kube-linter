package kubeconform

import (
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/kubeconform/internal/params"
)

const (
	templateKey = "kubeconform"
)

func init() {
	templates.Register(check.Template{
		HumanName:   templateKey,
		Key:         templateKey,
		Description: "Flag objects that does not match schema using https://github.com/yannh/kubeconform",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.Any},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate:            params.WrapInstantiateFunc(validate),
	})
}

func validate(p params.Params) (check.Func, error) {
	_ = "STUB: not implemented"
	// Create cache directory if it doesn't exist
	return *new(check.Func), nil
}

func sliceToMap(keys []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }
