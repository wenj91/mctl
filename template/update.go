package template

var Update = `
func (m *default{{.upperStartCamelObject}}Model) Update(data *{{.upperStartCamelObject}}) (affected int64, err error) {
	affected, err = m.conn.Update(m.method("update"), data)
	return
}
`

var UpdateSelective = `
func (m *default{{.upperStartCamelObject}}Model) UpdateSelective(data *{{.upperStartCamelObject}}) (affected int64, err error) {
	affected, err = m.conn.Update(m.method("updateSelective"), data)
	return
}
`

var UpdateFieldValue = `{{.field}} = #{{print "{" .value print "}"}}`

var UpdateIfFieldValue = `      <if test="{{.value}} != nil">
        {{.field}} = #{{print "{" .value print "}"}},
      </if>
`

var UpdateMapper = `
  <update id="update">
    update {{.table}}
    set {{.fields}}
    where {{.field}} = {{print "#{" .value print "}"}}
  </update>
  <update id="updateSelective">
    update {{.table}}
    <set>
{{.fieldValues}}    </set>
    where {{.field}} = {{print "#{" .value print "}"}}
  </update>
`

var UpdateMethod = `Update(data *{{.upperStartCamelObject}}) (affected int64, err error) `
var UpdateSelectiveMethod = `UpdateSelective(data *{{.upperStartCamelObject}}) (affected int64, err error) `
