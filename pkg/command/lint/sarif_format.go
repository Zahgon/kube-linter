package lint

import (
	"io"
	"text/template"

	"github.com/owenrumney/go-sarif/v2/sarif"
	"golang.stackrox.io/kube-linter/pkg/command/common"
	"golang.stackrox.io/kube-linter/pkg/config"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/run"
)

const (
	ruleHelpTemplateStr = `Check: {{.Name}}
Description: {{.Description}}
Remediation: {{.Remediation}}
Template: {{checkTemplateURL .}}`

	resultMessageTemplateStr = `{{.Report.Diagnostic.Message}}
object: {{.ObjectName}}`
)

var (
	ruleHelpTemplate = common.MustInstantiatePlainTemplate(ruleHelpTemplateStr,
		template.FuncMap{"checkTemplateURL": getCheckTemplateURL})

	resultMessageTemplate = common.MustInstantiatePlainTemplate(resultMessageTemplateStr, nil)
)

// formatLintSarif implements common.SARIFFormat.
// Must be used only with lint.Command because it only understands run.Result as data parameter.
func formatLintSarif(out io.Writer, data interface{}) error { _ = "STUB: not implemented"; return nil }

func formatSarif(out io.Writer, result run.Result) error { _ = "STUB: not implemented"; return nil }

// WithWorkingDirectory helps GitHub resolve artifact locations from repo root when their paths are absolute.

func addSarifRule(sarifRun *sarif.Run, check *config.Check) error {
	_ = "STUB: not implemented"
	return nil
}

// Notice that we give WithHelp the same information as added above for couple of reasons:
// 1) GitHub does not display HelpURI, although this attribute is required.
// 2) Rule ID, short and full descriptions are shown at different spots on the screen but it is helpful to see
//    them together.
// Markdown format for Help seemed to be ignored therefore we only provide the plain text version.

func getCheckTemplateURL(check *config.Check) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func renderTemplate(t *template.Template, data interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func addSarifResult(sarifRun *sarif.Run, cwd string, report *diagnostic.WithContext) error {
	_ = "STUB: not implemented"
	return nil
}

// We currently assign all errors to the first line in the file otherwise the absent region on the output
// does not pass GitHub validation rule GH1003.
// TODO: update region location with actual position in the file when it is available.

// GitHub does not seem to show logical locations at the moment. We still provide them hoping it will in the future.

// getArtifactURI tries to resolve path relative to cwd; if that fails, tries to get the absolute path with appended
// `file://` protocol; if that fails, returns the path as-is.
// GitHub prefers file URIs to be provided relative to the repo root. Assuming that this tool is invoked from the repo
// root, this function should resolve paths in a way GitHub likes them.
func getArtifactURI(cwd, path string) string { _ = "STUB: not implemented"; return "" }
