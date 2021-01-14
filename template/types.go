package template

var Types = `
type (
	{{.upperStartCamelObject}}Model interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Model struct {
		conn  gobatis.GoBatis
		table string
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
`
