package template

var Method = `
func (m *default{{.upperStartCamelObject}}Model) method(mt string) string {
	return "{{.upperStartCamelObject}}Mapper." + mt
}
`
