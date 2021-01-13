package gen

import (
	"strings"

	"github.com/tal-tech/go-zero/core/collection"
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/mctl/template"
)

func genInsert(table Table, withCache bool) (string, string, string, error) {
	keySet := collection.NewSet()
	keyVariableSet := collection.NewSet()
	for fieldName, key := range table.CacheKey {
		if fieldName == table.PrimaryKey.Name.Source() {
			continue
		}
		keySet.AddStr(key.DataKeyExpression)
		keyVariableSet.AddStr(key.Variable)
	}

	expressions := make([]string, 0)
	expressionValues := make([]string, 0)
	platformFieldNames := make([]string, 0)
	for _, filed := range table.Fields {
		platformFieldNames = append(platformFieldNames, filed.Name.Source())
		camel := filed.Name.ToCamel()
		if camel == "CreateTime" || camel == "UpdateTime" {
			continue
		}
		if filed.IsPrimaryKey && table.PrimaryKey.AutoIncrement {
			continue
		}
		expressions = append(expressions, "?")
		expressionValues = append(expressionValues, "data."+camel)
	}

	camel := table.Name.ToCamel()
	text, err := util.LoadTemplate(category, insertTemplateFile, template.Insert)
	if err != nil {
		return "", "", "", err
	}

	output, err := util.With("insert").
		Parse(text).
		Execute(map[string]interface{}{
			"withCache":             withCache,
			"containsIndexCache":    table.ContainsUniqueKey,
			"upperStartCamelObject": camel,
			"lowerStartCamelObject": stringx.From(camel).Untitle(),
			"expression":            strings.Join(expressions, ", "),
			"expressionValues":      strings.Join(expressionValues, ", "),
			"keys":                  strings.Join(keySet.KeysStr(), "\n"),
			"keyValues":             strings.Join(keyVariableSet.KeysStr(), ", "),
		})
	if err != nil {
		return "", "", "", err
	}

	// mapper if field
	text, err = util.LoadTemplate(category, insertTemplateIfFieldFile, template.InsertIfField)
	if err != nil {
		return "", "", "", err
	}

	ifFields := make([]string, 0)
	for _, filed := range table.Fields {
		insertIfOutput, err := util.With("insertIfField").
			Parse(text).
			Execute(map[string]interface{}{
				"field": filed.Name.Source(),
				"value": filed.Name.ToCamel(),
			})
		if err != nil {
			return "", "", "", err
		}

		ifFields = append(ifFields, insertIfOutput.String())
	}

	// mapper if value
	text, err = util.LoadTemplate(category, insertTemplateIfValueFile, template.InsertIfValue)
	if err != nil {
		return "", "", "", err
	}

	ifValues := make([]string, 0)
	for _, filed := range table.Fields {
		insertIfValueOutput, err := util.With("insertIfValue").
			Parse(text).
			Execute(map[string]interface{}{
				"field": filed.Name.Source(),
				"value": filed.Name.ToCamel(),
			})
		if err != nil {
			return "", "", "", err
		}

		ifValues = append(ifValues, insertIfValueOutput.String())
	}

	// mapper method
	text, err = util.LoadTemplate(category, insertTemplateMapperFile, template.InsertMapper)
	if err != nil {
		return "", "", "", err
	}

	var sb strings.Builder
	for _, v := range platformFieldNames {
		if sb.Len() > 0 {
			sb.WriteString(",")
		}

		sb.WriteString("#{")
		sb.WriteString(v)
		sb.WriteString("}")
	}

	insertMapperOutput, err := util.With("insert").
		Parse(text).
		Execute(map[string]interface{}{
			"table":                 table.Name.Source(),
			"withCache":             withCache,
			"containsIndexCache":    table.ContainsUniqueKey,
			"upperStartCamelObject": camel,
			"lowerStartCamelObject": stringx.From(camel).Untitle(),
			"expression":            strings.Join(expressions, ", "),
			"expressionValues":      strings.Join(expressionValues, ", "),
			"fields":                strings.Join(platformFieldNames, ", "),
			"sFields":               sb.String(),
			"ifFields":              strings.Join(ifFields, ""),
			"ifValues":              strings.Join(ifValues, ""),
			"keys":                  strings.Join(keySet.KeysStr(), "\n"),
			"keyValues":             strings.Join(keyVariableSet.KeysStr(), ", "),
		})
	if err != nil {
		return "", "", "", err
	}

	// interface method
	text, err = util.LoadTemplate(category, insertTemplateMethodFile, template.InsertMethod)
	if err != nil {
		return "", "", "", err
	}

	insertMethodOutput, err := util.With("insertMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject": camel,
		})
	if err != nil {
		return "", "", "", err
	}

	return output.String(),
		insertMethodOutput.String(),
		strings.Trim(insertMapperOutput.String(), "\n"),
		nil
}
