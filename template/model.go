package template

var Model = `package {{.pkg}}
{{.imports}}
{{.types}}
{{.new}}
{{.method}}
{{.insert}}
{{.find}}
{{.update}}
{{.delete}}
{{.extraMethod}}
`
