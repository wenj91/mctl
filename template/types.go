package template

var Types = `
type (
	{{.upperStartCamelObject}}Model interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Model struct {
		table string
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
`
