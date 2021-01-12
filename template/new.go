package template

var New = `
func New{{.upperStartCamelObject}}Model() {{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		table:      "{{.table}}",
	}
}
`
