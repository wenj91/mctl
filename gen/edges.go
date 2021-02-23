package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
)

func genEdges(table Table) (string, error) {
	text, err := util.LoadTemplate(category, edgesTemplateFile, template.Edges)
	if err != nil {
		return "", err
	}

	camelTableName := table.Name.ToCamel()
	buffer, err := util.With("edges").Parse(text).Execute(map[string]interface{}{
		"upperStartCamelObject": camelTableName,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
