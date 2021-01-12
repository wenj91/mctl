package template

var Delete = `
func (m *default{{.upperStartCamelObject}}Model) Delete(conn gobatis.GoBatis, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (affected int64, err error) {
	affected, err = conn.Delete(m.method("delete"), map[string]interface{}{
		"{{.lowerStartCamelPrimaryKey}}": {{.lowerStartCamelPrimaryKey}},
	})
	return
}
`

var DeleteMapper = `
  <delete id="delete">
    delete from {{.table}}
    where {{.field}} = {{print "#{" .value print "}"}}
  </delete>`

var DeleteMethod = `Delete(conn gobatis.GoBatis, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (affected int64, err error) `
