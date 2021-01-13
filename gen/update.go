package gen

import (
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/model/template"
)

func genUpdate(table Table, withCache bool) (string, string, string, error) {
	expressionValues := make([]string, 0)
	for _, filed := range table.Fields {
		camel := filed.Name.ToCamel()
		if camel == "CreateTime" || camel == "UpdateTime" {
			continue
		}
		if filed.IsPrimaryKey {
			continue
		}
		expressionValues = append(expressionValues, "data."+camel)
	}
	expressionValues = append(expressionValues, "data."+table.PrimaryKey.Name.ToCamel())
	camelTableName := table.Name.ToCamel()
	text, err := util.LoadTemplate(category, updateTemplateFile, template.Update)
	if err != nil {
		return "", "", "", err
	}

	output, err := util.With("update").
		Parse(text).
		Execute(map[string]interface{}{
			"withCache":             withCache,
			"upperStartCamelObject": camelTableName,
			"primaryCacheKey":       table.CacheKey[table.PrimaryKey.Name.Source()].DataKeyExpression,
			"primaryKeyVariable":    table.CacheKey[table.PrimaryKey.Name.Source()].Variable,
			"lowerStartCamelObject": stringx.From(camelTableName).Untitle(),
			"originalPrimaryKey":    wrapWithRawString(table.PrimaryKey.Name.Source()),
			"expressionValues":      strings.Join(expressionValues, ", "),
		})
	if err != nil {
		return "", "", "", nil
	}

	// update interface method
	text, err = util.LoadTemplate(category, updateMethodTemplateFile, template.UpdateMethod)
	if err != nil {
		return "", "", "", err
	}

	updateMethodOutput, err := util.With("updateMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camelTableName,
		})
	if err != nil {
		return "", "", "", err
	}

	// mapper field value
	text, err = util.LoadTemplate(category, updateMapperFieldValueTemplateFile, template.UpdateFieldValue)
	if err != nil {
		return "", "", "", err
	}

	updateFieldValues := make([]string, 0)
	for _, field := range table.Fields {
		camel := field.Name.ToCamel()
		if camel == "CreateTime" || camel == "UpdateTime" {
			continue
		}
		if field.IsPrimaryKey {
			continue
		}
		updateFieldValueOutput, err := util.With("updateFieldValue").
			Parse(text).
			Execute(map[string]interface{}{
				"field": field.Name.Source(),
				"value": field.Name.ToCamel(),
			})
		if err != nil {
			return "", "", "", err
		}

		updateFieldValues = append(updateFieldValues, updateFieldValueOutput.String())
	}

	// mapper if value
	text, err = util.LoadTemplate(category, updateMapperIfFieldValueTemplateFile, template.UpdateIfFieldValue)
	if err != nil {
		return "", "", "", err
	}

	ifValues := make([]string, 0)
	for _, field := range table.Fields {
		camel := field.Name.ToCamel()
		if camel == "CreateTime" || camel == "UpdateTime" {
			continue
		}
		if field.IsPrimaryKey {
			continue
		}
		updateIfFieldValueOutput, err := util.With("updateIfFieldValue").
			Parse(text).
			Execute(map[string]interface{}{
				"field": field.Name.Source(),
				"value": field.Name.ToCamel(),
			})
		if err != nil {
			return "", "", "", err
		}

		ifValues = append(ifValues, updateIfFieldValueOutput.String())
	}

	// mapper
	text, err = util.LoadTemplate(category, updateMapperTemplateFile, template.UpdateMapper)
	if err != nil {
		return "", "", "", err
	}

	updateMapperOutput, err := util.With("updateFieldValue").
		Parse(text).
		Execute(map[string]interface{}{
			"table":       table.Name.Source(),
			"fields":      strings.Join(updateFieldValues, ",\n    "),
			"fieldValues": strings.Join(ifValues, ""),
			"field":       table.PrimaryKey.Name.Source(),
			"value":       table.PrimaryKey.Name.ToCamel(),
		})
	if err != nil {
		return "", "", "", err
	}

	return output.String(), updateMethodOutput.String(), strings.Trim(updateMapperOutput.String(), "\n"), nil
}
