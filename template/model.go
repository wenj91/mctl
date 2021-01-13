package template

var Model = `package {{.pkg}}
{{.imports}}
{{.types}}
{{.toString}}
{{.new}}
{{.method}}
{{.insert}}
{{.find}}
{{.update}}
{{.delete}}
{{.extraMethod}}
`
