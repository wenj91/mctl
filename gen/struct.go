package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
)

func genStruct(table Table) (string, error) {
	text, err := util.LoadTemplate(category, structTemplateFile, template.Struct)
	if err != nil {
		return "", err
	}

	camelTableName := table.Name.ToCamel()
	buffer, err := util.With("fields").Parse(text).Execute(map[string]interface{}{
		"upperStartCamelObject": camelTableName,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
