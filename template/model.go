package template

var Model = `package {{.pkg}}
{{.imports}}
{{.types}}
{{.toString}}
{{.new}}
{{.method}}
{{.withConn}}
{{.insert}}
{{.find}}
{{.update}}
{{.delete}}
{{.extraMethod}}
`
