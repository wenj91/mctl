package template

var Table = `func ({{.upperStartCamelObject}}) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "{{.table}}"},
	}
}
`
