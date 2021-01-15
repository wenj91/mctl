package template

var Types = `
type (
	{{.upperStartCamelObject}}FindResult struct {
		{{.lowerStartCamelObject}}s []*{{.upperStartCamelObject}}
	}

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
