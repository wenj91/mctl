package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
)

func genQueryMethod(table Table) (string, error) {
	text, err := util.LoadTemplate(category, queryTemplateFile, template.QueryMethod)
	if err != nil {
		return "", err
	}

	camelTableName := table.Name.ToCamel()
	buffer, err := util.With("table").Parse(text).Execute(map[string]interface{}{
		"upperStartCamelObject": camelTableName,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
