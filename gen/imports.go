package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
)

func genImports(timeImport bool) (string, error) {
	text, err := util.LoadTemplate(category, importsTemplateFile, template.Imports)
	if err != nil {
		return "", err
	}

	buffer, err := util.With("import").Parse(text).Execute(map[string]interface{}{
		"time": timeImport,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
