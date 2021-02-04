package gen

import (
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/mctl/template"
)

type findOneCode struct {
	findOneMethod          string
	findOneMapper          string
	findOneInterfaceMethod string
	cacheExtra             string
}

func genFindOneByField(table Table, withCache bool) (*findOneCode, error) {
	text, err := util.LoadTemplate(category, findOneByFieldTemplateFile, template.FindOneByField)
	if err != nil {
		return nil, err
	}

	t := util.With("findOneByField").Parse(text)
	var list []string
	var mapperList []string
	camelTableName := table.Name.ToCamel()
	for _, field := range table.Fields {
		if field.IsPrimaryKey || !field.IsUniqueKey {
			continue
		}
		camelFieldName := field.Name.ToCamel()
		output, err := t.Execute(map[string]interface{}{
			"upperStartCamelObject":     camelTableName,
			"upperField":                camelFieldName,
			"in":                        fmt.Sprintf("%s %s", stringx.From(camelFieldName).Untitle(), strings.ReplaceAll(field.DataType, "*", "")),
			"withCache":                 withCache,
			"cacheKey":                  table.CacheKey[field.Name.Source()].KeyExpression,
			"cacheKeyVariable":          table.CacheKey[field.Name.Source()].Variable,
			"lowerStartCamelObject":     stringx.From(camelTableName).Untitle(),
			"lowerStartCamelField":      stringx.From(camelFieldName).Untitle(),
			"upperStartCamelPrimaryKey": table.PrimaryKey.Name.ToCamel(),
			"originalField":             wrapWithRawString(field.Name.Source()),
		})
		if err != nil {
			return nil, err
		}

		// mapper
		text, err = util.LoadTemplate(category, findOneByFieldMapperTemplateFile, template.FindOneByFieldMapper)
		if err != nil {
			return nil, err
		}

		findOneByFieldMapperOutput, err := util.With("findOneByFieldMapper").
			Parse(text).
			Execute(map[string]interface{}{
				"table": table.Name.Source(),
				"field": field.Name.Source(),
				"value": field.Name.ToCamel(),
			})
		if err != nil {
			return nil, err
		}

		list = append(list, strings.Trim(output.String(), "\n"))
		mapperList = append(mapperList, strings.Trim(findOneByFieldMapperOutput.String(), "\n"))
	}

	text, err = util.LoadTemplate(category, findOneByFieldMethodTemplateFile, template.FindOneByFieldMethod)
	if err != nil {
		return nil, err
	}

	t = util.With("findOneByFieldMethod").Parse(text)
	var listMethod []string
	for _, field := range table.Fields {
		if field.IsPrimaryKey || !field.IsUniqueKey {
			continue
		}
		camelFieldName := field.Name.ToCamel()
		output, err := t.Execute(map[string]interface{}{
			"upperStartCamelObject": camelTableName,
			"upperField":            camelFieldName,
			"in":                    fmt.Sprintf("%s %s", stringx.From(camelFieldName).Untitle(), strings.ReplaceAll(field.DataType, "*", "")),
		})
		if err != nil {
			return nil, err
		}

		listMethod = append(listMethod, strings.Trim(output.String(), "\n"))
	}

	return &findOneCode{
		findOneMethod:          strings.Join(list, util.NL),
		findOneMapper:          strings.Join(mapperList, util.NL),
		findOneInterfaceMethod: strings.Join(listMethod, util.NL),
	}, nil
}
