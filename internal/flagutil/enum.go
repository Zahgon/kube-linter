package flagutil

import (
	"github.com/spf13/pflag"

	"golang.stackrox.io/kube-linter/internal/set"
)

// EnumFlag allows setting a list of values.
type EnumFlag struct {
	flagDescription string
	allowedValues   set.FrozenStringSet

	currentValue string
}

// String implements pflag.Value.
// It is the preferred method of retrieving the set value for the enum.
func (e *EnumFlag) String() string { _ = "STUB: not implemented"; return "" }

// Set implements pflag.Value.
func (e *EnumFlag) Set(input string) error { _ = "STUB: not implemented"; return nil }

// Type implements pflag.Value.
func (e *EnumFlag) Type() string {
	_ = "STUB: not implemented"

	// Check that EnumFlag implements pflag.Value interface.
	return ""
}

var _ pflag.Value = (*EnumFlag)(nil)

// Usage returns a string that can be used as help text for this flag.
// It will include the flag type and the list of allowed values.
func (e *EnumFlag) Usage() string { _ = "STUB: not implemented"; return "" }

func (e *EnumFlag) getAllowedValuesString() string { _ = "STUB: not implemented"; return "" }

// NewEnumFlag creates and returns an enum flag value with the given description, allowedValues and defaultValue.
func NewEnumFlag(flagDescription string, allowedValues []string, defaultValue string) *EnumFlag {
	_ = "STUB: not implemented"
	return nil
}
