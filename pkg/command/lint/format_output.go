package lint

import (
	"golang.stackrox.io/kube-linter/pkg/command/common"
)

// FormatOutputPair represents a format and its output destination
type FormatOutputPair struct {
	Format common.FormatType
	Output string // empty string means stdout
}

// ValidateAndPairFormatsOutputs validates format and output flags and pairs them
func ValidateAndPairFormatsOutputs(formats, outputs, allowedFormats []string) ([]FormatOutputPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate formats

// Handle backward compatibility: no outputs means stdout

// Prevent multiple formats to stdout (creates unparseable mixed output)

// Validate output count matches format count

// Check for duplicate output files

// Pair formats with outputs
