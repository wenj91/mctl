package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/mctl/converter"
	"github.com/wenj91/mctl/template"
	"strings"
)

func genFields(table Table) (string, error) {
	text, err := util.LoadTemplate(category, fieldsTemplateFile, template.Fields)
	if err != nil {
		return "", err
	}

	fields := make([]string, 0)

	for _, field := range table.Fields {
		camelFieldName := field.Name.ToCamel()
		lowerStartCamelFieldName := stringx.From(camelFieldName).Untitle()
		dataType, err := converter.ConvertDataTypeToEntType(field.DataBaseType, lowerStartCamelFieldName, field.Name.Source())
		if err != nil {
			return "", err
		}

		fields = append(fields, dataType)
	}

	camelTableName := table.Name.ToCamel()
	buffer, err := util.With("fields").Parse(text).Execute(map[string]interface{}{
		"upperStartCamelObject": camelTableName,
		"fields": strings.Join(fields, ",\n		") + ",",
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
