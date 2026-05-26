package pdbminavailable

import (
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/k8sutil"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
	"golang.stackrox.io/kube-linter/pkg/objectkinds"
	"golang.stackrox.io/kube-linter/pkg/templates"
	"golang.stackrox.io/kube-linter/pkg/templates/pdbminavailable/internal/params"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	templateKey = "pdb-min-available"
)

func init() {
	templates.Register(check.Template{
		HumanName:   "No pod disruptions allowed - minAvailable",
		Key:         templateKey,
		Description: "Flag PodDisruptionBudgets whose minAvailable value will always prevent pod disruptions.",
		SupportedObjectKinds: config.ObjectKindsDesc{
			ObjectKinds: []string{
				objectkinds.PodDisruptionBudget},
		},
		Parameters:             params.ParamDescs,
		ParseAndValidateParams: params.ParseAndValidate,
		Instantiate: params.WrapInstantiateFunc(func(p params.Params) (check.Func, error) {
			return minAvailableCheck, nil
		}),
	})
}

func minAvailableCheck(lintCtx lintcontext.LintContext, object lintcontext.Object) []diagnostic.Diagnostic {
	_ = "STUB: not implemented"
	return nil
}

// Get the PDB provided

// If MinAvailable isn't set, then no need to check

// Extract the MinAvailable value from the spec and check if it's a number or percentage

// If the value is a percentage, handle the case where the MinValue is set to 100%
// as DeploymentLike's replica counts don't need to be compared

// Check if selector is present - it's a required field for PDB

// Build the label selector for the PDB to use for comparison

// Builds an HPA map for that namespace in case there is no replicas set on deployment

// Evaluate Deployment Likes in the lintContext to see if they have MinAvailable set too low

// if replicas number not set on deployment, use HPA MinReplicas

// Calculate the actual value of the MinAvailable with respect to the Replica count if a percentage is set

//nolint:gosec // Integer conversion should be safe here since the kube api uses int32.

func getDeploymentLikeObjects(lintCtx lintcontext.LintContext, labelSelector labels.Selector, namespace string) ([]k8sutil.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure that only DeploymentLike objects are processed

// Ensure that only DeploymentLikes are in the same namespaces as the PDB

// Build Deployment labelSelector
// If there are no selectors on the object, then the PDB won't match the same pods as the Deployment Like

// Find any Deployment Likes with the same selector as the PDB

func getIntOrPercentValueSafelyFromString(intOrStr string) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// Function to get the list of HPA's/ScaledObject's provided
func getHorizontalPodAutoscalers(lintCtx lintcontext.LintContext, namespace string) map[string]k8sutil.Object {
	_ = "STUB: not implemented"
	return nil
}

// Ensure that HPA/ScaledObject objects are processed

// Ensure that only HPAs/ScaledObject are in the same namespaces as the PDB

// validate object with HPA/ScaledObject versions using the HPAScaleTargetRefName extractor package function and add to map

// Function to transform the replica count into the minReplicas count if the deployment has a HPA with a minReplicas set
func transformReplicaIntoMinReplicas(deployment k8sutil.Object, hpaMap map[string]k8sutil.Object, replicas int32) int32 {
	_ = "STUB: not implemented"
	return 0
}
