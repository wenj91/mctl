package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/mctl/template"
)

func genTypes(table Table, methods string, withCache bool) (string, error) {
	fields := table.Fields
	fieldsString, selectiveString, err := genFields(fields)
	if err != nil {
		return "", err
	}

	text, err := util.LoadTemplate(category, typesTemplateFile, template.Types)
	if err != nil {
		return "", err
	}

	output, err := util.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"withCache":             withCache,
			"method":                methods,
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
			"upperStartCamelObject": table.Name.ToCamel(),
			"fields":                fieldsString,
			"selectiveFields":       selectiveString,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
