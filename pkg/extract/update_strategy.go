package extract

import (
	"reflect"

	"golang.stackrox.io/kube-linter/pkg/k8sutil"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// UpdateStrategyValues contains testable data from an UpdateStrategy struct
type UpdateStrategyValues struct {
	Type                 string
	TypeExists           bool
	RollingConfigExists  bool
	RollingConfigValid   bool
	MaxUnavailableExists bool
	MaxUnavailable       *intstr.IntOrString
	MaxSurgeExists       bool
	MaxSurge             *intstr.IntOrString
}

// UpdateStrategy will extract the data from an UpdateStrategy into a common struct
func UpdateStrategy(obj k8sutil.Object) (*UpdateStrategyValues, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// typeFromUpdateStrategy will extract the Type from a provided
// UpdateStrategy struct if it exists
func typeFromUpdateStrategy(strategy reflect.Value) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// rollingUpdateFromUpdateStrategy will extract the RollingUpdate struct from a provided
// RollingUpdate struct if it exists
func rollingUpdateFromUpdateStrategy(strategy reflect.Value) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// maxUnavailableFromRollingUpdate will extract the MaxUnavailable field from a provided
// RollingUpdate struct if it exists
func maxUnavailableFromRollingUpdate(rollingUpdate reflect.Value) (*intstr.IntOrString, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// maxSurgeFromRollingUpdate will extract the MaxSurge field from a provided
// RollingUpdate struct if it exists
func maxSurgeFromRollingUpdate(rollingUpdate reflect.Value) (*intstr.IntOrString, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
