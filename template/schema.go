package template

var Schema = `package schema

{{.imports}}
{{.struct}}
{{.fields}}
{{.edges}}
`
