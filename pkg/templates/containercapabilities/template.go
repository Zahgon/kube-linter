package containercapabilities

import (
	"fmt"

	"golang.stackrox.io/kube-linter/internal/utils"
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/matcher"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/containercapabilities/internal/params"
	"golang.stackrox.io/kube-linter/pkg/templates/util"
	v1 "k8s.io/api/core/v1"
)

const (
	templateKey                         = "verify-container-capabilities"
	reservedCapabilitiesAll             = "all"
	matchLiteralReservedCapabilitiesAll = "^(?i)" + reservedCapabilitiesAll + "$"
)

var (
	literalReservedCapabilitiesAllMatcher = func() func(string) bool {
		m, err := matcher.ForString(matchLiteralReservedCapabilitiesAll)
		utils.Must(err)
		return m
	}()

	addListDiagMsgFmt        = "container %q has ADD capability: %q, which matched with the forbidden capability for containers"
	addListWithAllDiagMsgFmt = "container %q has ADD capability: %q, but no capabilities " +
		"should be added at all and this capability is not included in the exceptions list"
	dropListDiagMsgFmt        = "container %q has DROP capabilities: %q, but does not drop capability %q which is required"
	dropListWithAllDiagMsgFmt = "container %q has DROP capabilities: %q, but in fact all capabilities are required to be dropped"
)

func checkCapabilityDropList(
	containerName string,
	paramCapMatchers map[string]func(string) bool,
	scCaps *v1.Capabilities,
	forbidAll bool,
	result *[]diagnostic.Diagnostic,
) {
	_ = "STUB: not implemented"

	// Specified to drop all. Instead of trying to find a match to this paramCap, we should
	// make sure the special ReservedCapabilitiesAll is found in the DROP list as well
	return
}

// Found "all" in drop list as well

// Every forbidden capability specified by param should be found in the DROP list

// User can specify to drop "all" under containers as well

// This forbidden cap exists in the DROP list. Check for the next forbidden cap

func checkCapabilityAddList(
	containerName string,
	paramCapMatchers map[string]func(string) bool,
	scCaps *v1.Capabilities,
	forbidAll bool,
	exceptionCapMatchers map[string]func(string) bool,
	result *[]diagnostic.Diagnostic,
) {
	_ = "STUB: not implemented"

	// User has forbidden all capabilities
	return
}

// Forgive this capability

// No violations

// Any capability from scCaps should not match with any from paramCaps

// User can specify to add "all" under containers as well.

// A capability from ADD list matched with a cap from forbidden capabilities list.

func checkForbidAll(paramCaps []string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Only contains "all"

// When the list contains "all", it should not contain any other element

// "all" not found

func validateExceptionsList(forbidAll bool, exceptions []string) error {
	_ = "STUB: not implemented"
	// Check if forbidAll is set
	return nil
}

// Check no "all" in exceptions list

func init() {
	templates.Register(check.Template{
		HumanName:   "Verify container capabilities",
		Key:         templateKey,
		Description: "Flag containers that do not match capabilities requirements",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{objectkinds.DeploymentLike},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(p params.Params) (check.Func, error) {
			// Check for "all" in case user does not want any capability in containers.
			forbidAll, err := checkForbidAll(p.ForbiddenCapabilities)
			if err != nil {
				return nil, err
			}
			err = validateExceptionsList(forbidAll, p.Exceptions)
			if err != nil {
				return nil, err
			}

			paramCapMatchers := make(map[string]func(string) bool, len(p.ForbiddenCapabilities))
			exceptionCapMatchers := make(map[string]func(string) bool, len(p.Exceptions))
			if !forbidAll {
				for _, cap := range p.ForbiddenCapabilities {
					capMatcher, err := matcher.ForString(cap)
					if err != nil {
						return nil, fmt.Errorf("checking container capabilities. invalid capability: %s: %w", cap, err)
					}
					paramCapMatchers[cap] = capMatcher
				}
			} else {
				for _, cap := range p.Exceptions {
					capMatcher, err := matcher.ForString(cap)
					if err != nil {
						return nil, fmt.Errorf("checking container capabilities. invalid capability: %s: %w", cap, err)
					}
					exceptionCapMatchers[cap] = capMatcher
				}
			}

			return util.PerContainerCheck(func(container *v1.Container) []diagnostic.Diagnostic {
				var result []diagnostic.Diagnostic
				sc := container.SecurityContext
				if sc != nil && sc.Capabilities != nil {
					// Check none of the forbidden capabilities exists in ADDs
					checkCapabilityAddList(
						container.Name,
						paramCapMatchers,
						sc.Capabilities,
						forbidAll,
						exceptionCapMatchers,
						&result)
					// Check every forbidden capabilities should exist in DROPs
					checkCapabilityDropList(container.Name, paramCapMatchers, sc.Capabilities, forbidAll, &result)
				}
				return result
			}), nil
		}),
	})
}
