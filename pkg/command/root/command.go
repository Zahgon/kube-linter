package root

import (
	"github.com/spf13/cobra"
)

const (
	colorFlag = "with-color"
)

// Command is the root command.
func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Only forcefully set colorful output if the flag has been set.
