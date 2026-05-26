package common

import (
	"io"
)

// FormatType defines possible output formats.
type FormatType string

const (
	// PlainFormat is for plain terminal output, possibly colored.
	PlainFormat = "plain"
	// MarkdownFormat is for markdown output suitable for `*.md` files.
	MarkdownFormat = "markdown"
	// JSONFormat is for JSON output with kube-linter's own data structs.
	JSONFormat = "json"
	// SARIFFormat is JSON-based standard for reporting lint errors.
	// See https://www.oasis-open.org/committees/tc_home.php?wg_abbrev=sarif
	SARIFFormat = "sarif"
)

// FormatFunc sets contract formatter of each FormatType should follow.
type FormatFunc func(io.Writer, interface{}) error

// Formatters struct provides way to define supported output formats for a command.
type Formatters struct {
	Formatters map[FormatType]FormatFunc
}

// GetEnabledFormatters returns a string slice enumerating all enabled formatters of this instance.
func (f Formatters) GetEnabledFormatters() []string { _ = "STUB: not implemented"; return nil }

// FormatterByType looks up formatter for a given type among the ones configured in this instance.
func (f Formatters) FormatterByType(t string) (FormatFunc, error) {
	_ = "STUB: not implemented"
	return *new(FormatFunc), nil
}

// FormatJSON formats data as JSON, i.e. implements JSONFormat.
func FormatJSON(out io.Writer, data interface{}) error { _ = "STUB: not implemented"; return nil }

// Verify that FormatJSON follows the contract of FormatFunc.
var _ FormatFunc = FormatJSON
