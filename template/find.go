package template

// 通过id查询
var FindOne = `
func (m *default{{.upperStartCamelObject}}Model) FindOne({{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	var resp *{{.upperStartCamelObject}}
	err := m.conn.Select(m.method("findOne"), map[string]interface{}{
		"{{.upperStartCamelPrimaryKey}}": {{.lowerStartCamelPrimaryKey}},
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

var FindSelective = `
func (m *default{{.upperStartCamelObject}}Model) FindSelective(data *{{.upperStartCamelObject}}Selective) (*{{.upperStartCamelObject}}FindResult, error) {
  resp := make([]*{{.upperStartCamelObject}}, 0)
	err := m.conn.Select(m.method("findSelective"), data)(&resp)
	return &{{.upperStartCamelObject}}FindResult{
    {{.lowerStartCamelObject}}s: resp,
  }, err
}
`

var FindSelectiveResultMethod = `
func new{{.upperStartCamelObject}}FindResult({{.lowerStartCamelObject}}s []*{{.upperStartCamelObject}}) *{{.upperStartCamelObject}}FindResult {
	return &{{.upperStartCamelObject}}FindResult{
		{{.lowerStartCamelObject}}s: {{.lowerStartCamelObject}}s,
	}
}

func (r *{{.upperStartCamelObject}}FindResult) List() []*{{.upperStartCamelObject}} {
	return r.{{.lowerStartCamelObject}}s
}

func (r *{{.upperStartCamelObject}}FindResult) One() *{{.upperStartCamelObject}} {
	if len(r.{{.lowerStartCamelObject}}s) == 0 {
		return nil
	}

	return r.{{.lowerStartCamelObject}}s[0]
}
`

var FindSelectiveIfFieldValue = `      <if test="{{.value}} != nil">
        and {{.field}} {{.mark}} #{{print "{" .value print "}"}}
      </if>
`

var FindSelectiveMapper = `
  <select id="findSelective" resultType="structs">
    select 
      <include refid="Base_Column_List" />
    from {{.table}}
    <where>
{{.fieldValues}}    </where>
  </select>
`

// 通过指定字段查询
var FindOneByField = `
func (m *default{{.upperStartCamelObject}}Model) FindOneBy{{.upperField}}({{.in}}) (*{{.upperStartCamelObject}}, error) {
	var resp *{{.upperStartCamelObject}}
	err := m.conn.Select(m.method("findOneBy{{.upperField}}"), map[string]interface{}{
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

var FindOneMethod = `FindOne({{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) `
var FindOneByFieldMethod = `FindOneBy{{.upperField}}({{.in}}) (*{{.upperStartCamelObject}}, error) `
var FindSelectiveMethod = `FindSelective(data *{{.upperStartCamelObject}}Selective) (*{{.upperStartCamelObject}}FindResult, error) `
