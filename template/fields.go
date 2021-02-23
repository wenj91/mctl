package template

var Fields = `func ({{.upperStartCamelObject}}) Fields() []ent.Field {
	return []ent.Field{
		{{.fields}}
	}
}
`
