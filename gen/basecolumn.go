package gen

import (
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/wenj91/model/template"
)

func genBaseColumn(table Table, withCache bool) (string, error) {
	platformFieldNames := make([]string, 0)
	for _, filed := range table.Fields {
		platformFieldNames = append(platformFieldNames, filed.Name.Source())
	}

	text, err := util.LoadTemplate(category, baseColumnTemplateFile, template.BaseColumn)
	if err != nil {
		return "", err
	}

	output, err := util.With("base").
		Parse(text).
		Execute(map[string]interface{}{
			"fields": strings.Join(platformFieldNames, ","),
		})
	if err != nil {
		return "", err
	}

	return strings.Trim(output.String(), "\n"), nil
}
