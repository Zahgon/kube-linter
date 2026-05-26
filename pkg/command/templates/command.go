package templates

import (
	"github.com/spf13/cobra"
	"golang.stackrox.io/kube-linter/pkg/command/common"
)

const (
	markDownTemplateStr = `# KubeLinter templates

KubeLinter supports the following templates:

{{ range . -}}
## {{ .HumanName }}

**Key**: {{ .Key | codeSnippet }}

**Description**: {{ .Description }}

**Supported Objects**: {{ join "," .SupportedObjectKinds.ObjectKinds }}

{{ if .HumanReadableParameters }}
**Parameters**:

{{ mustToYaml .HumanReadableParameters | codeBlock "yaml" }}
{{ end }}
{{ end -}}
`

	plainTemplateStr = `{{- define "Param" }}{{ $tabs := repeat .NestingLevel "\t" }}
	{{$tabs}}{{.Name}}:
		{{$tabs}}Description: {{.Description}}
		{{$tabs}}Required: {{.Required}}{{if .Examples}}
		{{$tabs}}Example values: {{ range $i, $_ := .Examples }}{{if $i}}, {{end}}{{ printf "%q" . }}{{end}}{{end}}{{if .SubParameters}}
		{{$tabs}}Sub-parameters: {{ range .SubParameters }}{{ template "Param" . }}{{end}}{{end}}{{if .ArrayElemType}}
		{{$tabs}}Elem type: {{.ArrayElemType}}{{end}}
{{- end -}}
{{ range $i, $_ := . }}
{{- if $i}}
------------------------------

{{end -}}
Name: {{.HumanName}}
Key: {{.Key}}
Description: {{.Description}}
Supported Objects: {{.SupportedObjectKinds.ObjectKinds}}
Parameters:{{ range .HumanReadableParameters }}{{ template "Param" . }}{{else}} none{{end}}
{{end -}}
`
)

var (
	markDownTemplate = common.MustInstantiateMarkdownTemplate(markDownTemplateStr, nil)
	plainTemplate    = common.MustInstantiatePlainTemplate(plainTemplateStr, nil)

	formatters = common.Formatters{
		Formatters: map[common.FormatType]common.FormatFunc{
			common.PlainFormat:    plainTemplate.Execute,
			common.MarkdownFormat: markDownTemplate.Execute,
			common.JSONFormat:     common.FormatJSON,
		},
	}
)

func listCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Command defines the root of the templates command.
func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }
