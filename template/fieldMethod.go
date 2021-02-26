package template

var FieldMethod = `    public {{.upperStartCamelObject}}Query {{.lowerStartCamelFieldName}}{{.op}}({{if .normal}}{{.fieldJavaType}}{{.in}} {{.lowerStartCamelFieldName}}{{end}}) {
		{{if .normal}}if (null != {{.lowerStartCamelFieldName}}{{if eq .in "..."}} && {{.lowerStartCamelFieldName}}.length > 0{{end}}) {
            this.push(Cond.of("{{.field}}", Op.{{.OP}}, {{if eq .in "..."}}(Object[]){{end}}{{.lowerStartCamelFieldName}}));
		}{{else}}this.push(Cond.of("{{.field}}", Op.{{.OP}}));{{end}}

        return this;
    }`
