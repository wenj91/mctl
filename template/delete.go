package template

var Delete = `
func (m *default{{.upperStartCamelObject}}Model) Delete({{.lowerStartCamelPrimaryKey}} {{.dataType}}) (affected int64, err error) {
	affected, err = m.conn.Delete(m.method("delete"), map[string]interface{}{
		"{{.upperStartCamelPrimaryKey}}": {{.lowerStartCamelPrimaryKey}},
	})
	return
}
`

var DeleteMapper = `
  <delete id="delete">
    delete from {{.table}}
    where {{.field}} = {{print "#{" .value print "}"}}
  </delete>`

var DeleteMethod = `Delete({{.lowerStartCamelPrimaryKey}} {{.dataType}}) (affected int64, err error) `
