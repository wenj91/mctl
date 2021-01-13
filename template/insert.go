package template

var Insert = `
func (m *default{{.upperStartCamelObject}}Model) Insert(conn gobatis.GoBatis, data *{{.upperStartCamelObject}}) (id int64, affected int64, err error) {
	id, affected, err = conn.Insert(m.method("save"), data)
	return
}
`

var InsertSelective = `
func (m *default{{.upperStartCamelObject}}Model) InsertSelective(conn gobatis.GoBatis, data *{{.upperStartCamelObject}}) (id int64, affected int64, err error) {
	id, affected, err = conn.Insert(m.method("saveSelective"), data)
	return
}
`

var InsertIfField = `      <if test="{{.value}} != nil and {{.value}} != ''">
        {{.field}},  
      </if>
`

var InsertIfValue = `      <if test="{{.value}} != nil and {{.value}} != ''">
        #{{print "{" .value print "}"}},
      </if>
`

var InsertMapper = `
  <insert id="save">
    insert into {{.table}} ({{.fields}})
    values ({{.sFields}})
  </insert>
  <insert id="saveSelective">
    insert into {{.table}}
    <trim prefix="(" suffix=")" suffixOverrides=",">
{{.ifFields}}    </trim>
    <trim prefix="values (" suffix=")" suffixOverrides=",">
{{.ifValues}}    </trim>
  </insert>
`

var InsertMethod = `Insert(conn gobatis.GoBatis, data *{{.upperStartCamelObject}}) (id int64, affected int64, err error) `
var InsertSelectiveMethod = `InsertSelective(conn gobatis.GoBatis, data *{{.upperStartCamelObject}}) (id int64, affected int64, err error) `
