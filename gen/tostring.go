package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
)

func genToString(table Table) (string, error) {
	camel := table.Name.ToCamel()
	text, err := util.LoadTemplate(category, toStringTemplateFile, template.ToStr)
	if err != nil {
		return "", err
	}

	output, err := util.With("tostring").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camel,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
