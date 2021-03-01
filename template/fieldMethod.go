package template

var FieldMethod = `    public {{.upperStartCamelObject}}Query {{.lowerStartCamelFieldName}}{{.op}}({{if .normal}}{{if eq .in "..."}}List<{{end}}{{.fieldJavaType}}{{if eq .in "..."}}>{{end}} {{.lowerStartCamelFieldName}}{{end}}) {
		{{if .normal}}if (null != {{.lowerStartCamelFieldName}}{{if eq .in "..."}} && {{.lowerStartCamelFieldName}}.size() > 0{{end}}) {
            this.push(Cond.of("{{.field}}", Op.{{.OP}}, {{.lowerStartCamelFieldName}}{{if eq .in "..."}}.toArray(){{end}}));
		}{{else}}this.push(Cond.of("{{.field}}", Op.{{.OP}}));{{end}}

        return this;
    }`
