package template

var Schema = `package schema

{{.imports}}
{{.struct}}
{{.table}}
{{.fields}}
{{.edges}}
`
