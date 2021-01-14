package template

var New = `
func New{{.upperStartCamelObject}}Model(conn gobatis.GoBatis) {{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		conn:  conn,
		table: "{{.table}}",
	}
}
`
