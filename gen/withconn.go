package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
)

func genWithConn(table Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := util.LoadTemplate(category, withConnTemplateFile, template.WithConn)
	if err != nil {
		return "", "", err
	}

	output, err := util.With("withConn").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camel,
			"table":                 table.Name.Source(),
		})
	if err != nil {
		return "", "", err
	}

	text, err = util.LoadTemplate(category, withConnMethodTemplateFile, template.WithConnMethod)
	if err != nil {
		return "", "", err
	}
	methodOutput, err := util.With("withConnMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camel,
		})
	if err != nil {
		return "", "", err
	}

	return output.String(), methodOutput.String(), nil
}
