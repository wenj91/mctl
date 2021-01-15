package template

var Model = `package {{.pkg}}
{{.imports}}
{{.types}}
{{.toString}}
{{.findSelectiveResultCode}}
{{.new}}
{{.method}}
{{.withConn}}
{{.insert}}
{{.find}}
{{.update}}
{{.delete}}
{{.extraMethod}}
`
