package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
)

func genMapper(table Table, stmts string) (string, error) {
	if stmts == "" {
		return stmts, nil
	}

	text, err := util.LoadTemplate(category, mapperTemplateFile, template.Mapper)
	if err != nil {
		return "", err
	}

	output, err := util.With("mapper").
		Parse(text).
		Execute(map[string]interface{}{
			"mapper": table.Name.ToCamel(),
			"stmts":  stmts,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
