package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
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

	for _, item := range table.Fields {
		dataType, err := converter.ConvertDataTypeToEntType(item.DataBaseType, item.Name.Source())
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
