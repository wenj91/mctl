package template

var WithConn = `
func (m *default{{.upperStartCamelObject}}Model) WithConn(conn gobatis.GoBatis) {{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		conn:  conn,
		table: "{{.table}}",
	}
}
`

var WithConnMethod = `WithConn(conn gobatis.GoBatis) {{.upperStartCamelObject}}Model `
