package gen

import (
	"fmt"

	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/urfave/cli"
	"github.com/wenj91/model/template"
)

const (
	category                              = "model"
	deleteTemplateFile                    = "delete.tpl"
	deleteMapperTemplateFile              = "mapper-delete.tpl"
	deleteMethodTemplateFile              = "interface-delete.tpl"
	fieldTemplateFile                     = "filed.tpl"
	findOneTemplateFile                   = "find-one.tpl"
	findOneMethodTemplateFile             = "interface-find-one.tpl"
	findOneMapperTemplateFile             = "mapper-find-one.tpl"
	findOneByFieldTemplateFile            = "find-one-by-field.tpl"
	findOneByFieldMapperTemplateFile      = "mapper-find-one-by-field.tpl"
	findOneByFieldMethodTemplateFile      = "interface-find-one-by-field.tpl"
	findOneByFieldExtraMethodTemplateFile = "find-one-by-filed-extra-method.tpl"
	importsTemplateFile                   = "import.tpl"
	importsWithNoCacheTemplateFile        = "import-no-cache.tpl"
	insertTemplateFile                    = "insert.tpl"
	insertSelectiveTemplateFile           = "insert-selective.tpl"
	insertSelectiveMethodTemplateFile     = "interface-insert-selective.tpl"
	insertTemplateMethodFile              = "interface-insert.tpl"
	insertTemplateIfFieldFile             = "if-field-insert.tpl"
	insertTemplateIfValueFile             = "if-value-insert.tpl"
	insertTemplateMapperFile              = "mapper-insert.tpl"
	modelTemplateFile                     = "model.tpl"
	modelNewTemplateFile                  = "model-new.tpl"
	mapperTemplateFile                    = "mapper.tpl"
	modelMethodTemplateFile               = "model-method.tpl"
	tagTemplateFile                       = "tag.tpl"
	baseColumnTemplateFile                = "base-column.tpl"
	typesTemplateFile                     = "types.tpl"
	updateSelectiveTemplateFile           = "update-selective.tpl"
	updateSelectiveMethodTemplateFile     = "interface-update-selective.tpl"
	updateTemplateFile                    = "update.tpl"
	updateMethodTemplateFile              = "interface-update.tpl"
	updateMapperTemplateFile              = "mapper-update.tpl"
	updateMapperIfFieldValueTemplateFile  = "mapper-field-value-if-update.tpl"
	updateMapperFieldValueTemplateFile    = "mapper-field-value-update.tpl"
	varTemplateFile                       = "var.tpl"
	errTemplateFile                       = "err.tpl"
)

var templates = map[string]string{
	deleteTemplateFile:                    template.Delete,
	deleteMethodTemplateFile:              template.DeleteMethod,
	deleteMapperTemplateFile:              template.DeleteMapper,
	fieldTemplateFile:                     template.Field,
	findOneTemplateFile:                   template.FindOne,
	findOneMethodTemplateFile:             template.FindOneMethod,
	findOneMapperTemplateFile:             template.FindOneMapper,
	findOneByFieldTemplateFile:            template.FindOneByField,
	findOneByFieldMapperTemplateFile:      template.FindOneByFieldMapper,
	findOneByFieldMethodTemplateFile:      template.FindOneByFieldMethod,
	findOneByFieldExtraMethodTemplateFile: template.FindOneByFieldExtraMethod,
	importsTemplateFile:                   template.Imports,
	importsWithNoCacheTemplateFile:        template.ImportsNoCache,
	insertTemplateFile:                    template.Insert,
	insertTemplateMethodFile:              template.InsertMethod,
	insertTemplateIfFieldFile:             template.InsertIfField,
	insertTemplateIfValueFile:             template.InsertIfValue,
	insertTemplateMapperFile:              template.InsertMapper,
	insertSelectiveTemplateFile:           template.InsertSelective,
	insertSelectiveMethodTemplateFile:     template.InsertSelectiveMethod,
	modelTemplateFile:                     template.Model,
	modelNewTemplateFile:                  template.New,
	modelMethodTemplateFile:               template.Method,
	mapperTemplateFile:                    template.Mapper,
	tagTemplateFile:                       template.Tag,
	baseColumnTemplateFile:                template.BaseColumn,
	typesTemplateFile:                     template.Types,
	updateSelectiveTemplateFile:           template.UpdateSelective,
	updateSelectiveMethodTemplateFile:     template.UpdateSelectiveMethod,
	updateTemplateFile:                    template.Update,
	updateMapperTemplateFile:              template.UpdateMapper,
	updateMapperIfFieldValueTemplateFile:  template.UpdateIfFieldValue,
	updateMapperFieldValueTemplateFile:    template.UpdateFieldValue,
	varTemplateFile:                       template.Vars,
	errTemplateFile:                       template.Error,
}

func Category() string {
	return category
}

func Clean() error {
	return util.Clean(category)
}

func GenTemplates(_ *cli.Context) error {
	return util.InitTemplates(category, templates)
}

func RevertTemplate(name string) error {
	content, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s: no such file name", name)
	}
	return util.CreateTemplate(category, name, content)
}

func Update() error {
	err := Clean()
	if err != nil {
		return err
	}
	return util.InitTemplates(category, templates)
}
