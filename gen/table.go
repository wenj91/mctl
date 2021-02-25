package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
)

func genTable(table Table) (string, error) {
	text, err := util.LoadTemplate(category, tableTemplateFile, template.Table)
	if err != nil {
		return "", err
	}

	buffer, err := util.With("table").Parse(text).Execute(map[string]interface{}{
		"table": table.Name.Source(),
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
