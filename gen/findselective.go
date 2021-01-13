package gen

import (
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/mctl/template"
)

func genFindSelective(table Table, withCache bool) (string, string, string, error) {
	camel := table.Name.ToCamel()
	text, err := util.LoadTemplate(category, findSelectiveTemplateFile, template.FindSelective)
	if err != nil {
		return "", "", "", err
	}

	output, err := util.With("findSelective").
		Parse(text).
		Execute(map[string]interface{}{
			"withCache":                 withCache,
			"upperStartCamelObject":     camel,
			"lowerStartCamelObject":     stringx.From(camel).Untitle(),
			"originalPrimaryKey":        wrapWithRawString(table.PrimaryKey.Name.Source()),
			"lowerStartCamelPrimaryKey": stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle(),
			"dataType":                  table.PrimaryKey.DataType,
			"cacheKey":                  table.CacheKey[table.PrimaryKey.Name.Source()].KeyExpression,
			"cacheKeyVariable":          table.CacheKey[table.PrimaryKey.Name.Source()].Variable,
		})
	if err != nil {
		return "", "", "", err
	}

	text, err = util.LoadTemplate(category, findSelectiveMethodTemplateFile, template.FindSelectiveMethod)
	if err != nil {
		return "", "", "", err
	}

	findOneMethod, err := util.With("findSelectiveMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelPrimaryKey": stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle(),
			"dataType":                  table.PrimaryKey.DataType,
		})
	if err != nil {
		return "", "", "", err
	}

	// mapper if field value
	text, err = util.LoadTemplate(category, findSelectiveIfFieldValuesTemplateFile, template.FindSelectiveIfFieldValue)
	if err != nil {
		return "", "", "", err
	}

	ifValues := make([]string, 0)
	for _, field := range table.Fields {
		if field.IsPrimaryKey {
			continue
		}

		findIfFieldValueOutput, err := util.With("findSelectiveIfFieldValue").
			Parse(text).
			Execute(map[string]interface{}{
				"field": field.Name.Source(),
				"value": field.Name.ToCamel(),
			})
		if err != nil {
			return "", "", "", err
		}

		ifValues = append(ifValues, findIfFieldValueOutput.String())
	}

	// mapper
	text, err = util.LoadTemplate(category, findSelectiveMapperTemplateFile, template.FindSelectiveMapper)
	if err != nil {
		return "", "", "", err
	}

	findOneMapperOutput, err := util.With("findSelectiveMapper").
		Parse(text).
		Execute(map[string]interface{}{
			"table":       table.Name.Source(),
			"fieldValues": strings.Join(ifValues, ""),
		})
	if err != nil {
		return "", "", "", err
	}

	return output.String(),
		findOneMethod.String(),
		strings.Trim(findOneMapperOutput.String(), "\n"),
		nil
}
