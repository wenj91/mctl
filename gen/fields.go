package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/wenj91/mctl/template"
	"strings"
)

func genFields(table Table) (string, error) {
	text, err := util.LoadTemplate(category, fieldsTemplateFile, template.Fields)
	if err != nil {
		return "", err
	}

	fields := table.Fields

	fieldsStr := make([]string, 0)
	for _, field := range fields {
		fieldName := field.Name
		fieldsStr = append(fieldsStr, "\""+fieldName.Source()+"\"")
	}

	buf, err := util.With("fieldMethod").Parse(text).Execute(map[string]interface{}{
		"fields": strings.Join(fieldsStr, ", "),
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
