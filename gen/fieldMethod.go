package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/mctl/template"
	"strings"
)

var opList = []string{
	"EQ",
	"NEQ",
	"In",
	"NotIn",
	"GT",
	"GTE",
	"LT",
	"LTE",
	"Contains",
	"HasPrefix",
	"HasSuffix",
	"IsNull",
	"NotNull",
	"OrderByAsc",
	"OrderByDesc",
}

var typeMap = map[string]bool{
	"Integer":    true,
	"Long":       true,
	"Float":      true,
	"Double":     true,
	"BigDecimal": true,
}

var skipMap = map[string]bool{
	"Contains":  true,
	"HasPrefix": true,
	"HasSuffix": true,
}

func genFieldMethod(table Table) (string, error) {
	text, err := util.LoadTemplate(category, fieldMethodTemplateFile, template.FieldMethod)
	if err != nil {
		return "", err
	}

	camelTableName := table.Name.ToCamel()

	fields := table.Fields

	fieldsStr := make([]string, 0)
	for _, field := range fields {
		fieldName := field.Name
		fieldJavaType := field.JavaDataType

		camelFieldName := fieldName.ToCamel()
		buf, err := util.With("fieldMethod").Parse(text).Execute(map[string]interface{}{
			"upperStartCamelObject":    camelTableName,
			"lowerStartCamelFieldName": stringx.From(camelFieldName).Untitle(),
			"op":                       "",
			"fieldJavaType":            fieldJavaType,
			"field":                    fieldName.Source(),
			"OP":                       "EQ",
		})
		if err != nil {
			return "", err
		}

		fieldsStr = append(fieldsStr, buf.String())

		for _, op := range opList {
			_, ok := typeMap[fieldJavaType]
			_, skip := skipMap[op]
			if ok && skip {
				continue
			}

			buf2, err := util.With("fieldMethod").Parse(text).Execute(map[string]interface{}{
				"upperStartCamelObject":    camelTableName,
				"lowerStartCamelFieldName": stringx.From(camelFieldName).Untitle(),
				"op":                       op,
				"fieldJavaType":            fieldJavaType,
				"field":                    fieldName.Source(),
				"OP":                       op,
			})

			if err != nil {
				return "", err
			}

			fieldsStr = append(fieldsStr, buf2.String())
		}
	}

	return strings.Join(fieldsStr, "\n\n"), nil
}
