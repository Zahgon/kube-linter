package checks

import (
	"text/template"

	"github.com/spf13/cobra"
	"golang.stackrox.io/kube-linter/internal/defaultchecks"
	"golang.stackrox.io/kube-linter/pkg/command/common"
	"golang.stackrox.io/kube-linter/pkg/config"
)

const (
	plainTemplateStr = `{{ range $i, $_ := . }}
{{- if $i}}
------------------------------

{{end -}}
Name: {{.Name}}
Description: {{.Description}}
Remediation: {{.Remediation}}
Template: {{.Template}}
Parameters: {{.Params}}
Enabled by default: {{ isDefault . }}
{{end -}}
`

	markDownTemplateStr = `# KubeLinter checks

KubeLinter includes the following built-in checks:

{{ range . -}}
## {{ .Name}}

**Enabled by default**: {{ if isDefault . }}Yes{{ else }}No{{ end }}

**Description**: {{.Description}}

**Remediation**: {{.Remediation}}

**Template**: [{{.Template}}](templates.md#{{ templateLink . }})
{{ if .Params }}
**Parameters**:

{{ mustToYaml (default (dict) .Params ) | codeBlock "yaml" }}
{{ end -}}
{{ end -}}
`
)

var (
	checksFuncMap = template.FuncMap{
		"isDefault": func(check config.Check) bool {
			return defaultchecks.List.Contains(check.Name)
		},
		"templateLink": GetTemplateLink,
	}
	plainTemplate    = common.MustInstantiatePlainTemplate(plainTemplateStr, checksFuncMap)
	markDownTemplate = common.MustInstantiateMarkdownTemplate(markDownTemplateStr, checksFuncMap)

	formatters = common.Formatters{
		Formatters: map[common.FormatType]common.FormatFunc{
			common.PlainFormat:    plainTemplate.Execute,
			common.MarkdownFormat: markDownTemplate.Execute,
			common.JSONFormat:     common.FormatJSON,
		},
	}
)

func listCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Command defines the root of the checks command.
func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetTemplateLink returns html anchor string for the template corresponding to the given check so that it can be used
// to reference the template section in a rendered markdown.
// E.g. template name "Deprecated Service Account Field" becomes "deprecated-service-account-field" html anchor.
func GetTemplateLink(check *config.Check) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
