package template

// 通过id查询
var FindOne = `
func (m *default{{.upperStartCamelObject}}Model) FindOne(conn gobatis.GoBatis, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	var resp *{{.upperStartCamelObject}}
	err := conn.Select(m.method("findOne"), map[string]interface{}{
		"{{.lowerStartCamelPrimaryKey}}": {{.lowerStartCamelPrimaryKey}},
	})(&resp)
	return resp, err
}
`

var FindOneMapper = `
  <select id="findOne" resultType="struct">
    select 
      <include refid="Base_Column_List" />
    from {{.table}}
    where {{.field}} = {{print "#{" .value print "}"}}
    limit 1
  </select>
`

// 通过指定字段查询
var FindOneByField = `
func (m *default{{.upperStartCamelObject}}Model) FindOneBy{{.upperField}}(conn gobatis.GoBatis, {{.in}}) (*{{.upperStartCamelObject}}, error) {
	var resp *{{.upperStartCamelObject}}
	err := conn.Select(m.method("findOneBy{{.upperField}}"), map[string]interface{}{
		"{{.upperField}}": {{.lowerStartCamelField}},
	})(&resp)
	return resp, err
}
`

var FindOneByFieldMapper = `
  <select id="findOneBy{{.value}}" resultType="struct">
    select 
      <include refid="Base_Column_List" />
    from {{.table}}
    where {{.field}} = {{print "#{" .value print "}"}}
    limit 1
  </select>
`

var FindOneByFieldExtraMethod = `
func (m *default{{.upperStartCamelObject}}Model) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", {{.primaryKeyLeft}}, primary)
}

func (m *default{{.upperStartCamelObject}}Model) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where {{.originalPrimaryField}} = ? limit 1", {{.lowerStartCamelObject}}Rows, m.table )
	return conn.QueryRow(v, query, primary)
}
`

var FindOneMethod = `FindOne(conn gobatis.GoBatis, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) `
var FindOneByFieldMethod = `FindOneBy{{.upperField}}(conn gobatis.GoBatis, {{.in}}) (*{{.upperStartCamelObject}}, error) `
