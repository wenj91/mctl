package template

var FieldMethod = `    public {{.upperStartCamelObject}}Query {{.lowerStartCamelFieldName}}{{.op}}({{.fieldJavaType}} {{.lowerStartCamelFieldName}}) {
        if (null != {{.lowerStartCamelFieldName}}) {
            this.push(Cond.of("{{.field}}", Op.{{.OP}}, {{.lowerStartCamelFieldName}}));
		}

        return this;
    }`
