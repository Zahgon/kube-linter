package lint

import (
	"github.com/spf13/cobra"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"

	"golang.stackrox.io/kube-linter/pkg/command/common"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
)

const (
	plainTemplateStr = `KubeLinter {{.Summary.KubeLinterVersion}}

{{range .Reports}}
{{- .Object.Metadata.FilePath | bold}}: (object: {{.Object.GetK8sObjectName | bold}}) {{.Diagnostic.Message | red}} (check: {{.Check | yellow}}, remediation: {{.Remediation | yellow}})

{{else}}No lint errors found!
{{end -}}
`
)

var (
	plainTemplate = common.MustInstantiatePlainTemplate(plainTemplateStr, nil)

	formatters = common.Formatters{
		Formatters: map[common.FormatType]common.FormatFunc{
			common.JSONFormat:  common.FormatJSON,
			common.SARIFFormat: formatLintSarif,
			common.PlainFormat: plainTemplate.Execute,
		},
	}
)

// Command is the command for the lint command.
func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Load Configuration

// Validate and pair formats with outputs

// Write output for each format

// Write output

// Always close the destination, even on write error

// Log success for file outputs

// Report both successes and failures

func generateReportFromInvalidObjects(lintCtxs []lintcontext.LintContext) []diagnostic.WithContext {
	_ = "STUB: not implemented"
	return nil
}
