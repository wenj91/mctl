package gen

import (
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/mctl/parser"
	"github.com/wenj91/mctl/template"
)

func genFields(fields []parser.Field) (string, string, error) {
	var list []string
	var selectiveList []string
	for _, field := range fields {

		if strings.Contains(field.DataType, "Time") {

			name := "start_" + field.Name.Source()
			startField := parser.Field{
				Name:         stringx.From(name),
				DataBaseType: field.DataBaseType,
				DataType:     field.DataType,
				IsPrimaryKey: field.IsPrimaryKey,
				IsUniqueKey:  field.IsUniqueKey,
				Comment:      field.Comment,
			}

			result, err := genSelectiveField(startField, false)
			if err != nil {
				return "", "", err
			}

			selectiveList = append(selectiveList, result)

			endName := "end_" + field.Name.Source()
			endField := parser.Field{
				Name:         stringx.From(endName),
				DataBaseType: field.DataBaseType,
				DataType:     field.DataType,
				IsPrimaryKey: field.IsPrimaryKey,
				IsUniqueKey:  field.IsUniqueKey,
				Comment:      field.Comment,
			}

			endResult, err := genSelectiveField(endField, false)
			if err != nil {
				return "", "", err
			}

			selectiveList = append(selectiveList, endResult)
		}

		result, err := genField(field, true)
		if err != nil {
			return "", "", err
		}

		list = append(list, result)

		selectiveResult, err := genSelectiveField(field, true)
		if err != nil {
			return "", "", err
		}
		selectiveList = append(selectiveList, selectiveResult)
	}

	return strings.Join(list, "\n"), strings.Join(selectiveList, "\n"), nil
}

func genField(field parser.Field, isDBField bool) (string, error) {
	tag, err := genTag(field.Name.Source(), stringx.From(field.Name.ToCamel()).Untitle(), isDBField)
	if err != nil {
		return "", err
	}

	text, err := util.LoadTemplate(category, fieldTemplateFile, template.Field)
	if err != nil {
		return "", err
	}

	output, err := util.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"name":       field.Name.ToCamel(),
			"type":       field.DataType,
			"tag":        tag,
			"hasComment": field.Comment != "",
			"comment":    field.Comment,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func genSelectiveField(field parser.Field, isDBField bool) (string, error) {
	//tag, err := genTag(field.Name.Source(), stringx.From(field.Name.ToCamel()).Untitle(), isDBField)
	//if err != nil {
	//	return "", err
	//}

	text, err := util.LoadTemplate(category, fieldTemplateFile, template.Field)
	if err != nil {
		return "", err
	}

	output, err := util.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"name":       field.Name.ToCamel(),
			"type":       field.DataType,
			"tag":        "",
			"hasComment": field.Comment != "",
			"comment":    field.Comment,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
