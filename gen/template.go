package gen

import (
	"fmt"

	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/urfave/cli"
	"github.com/wenj91/mctl/template"
)

const (
	category            = "schema"
	importsTemplateFile = "import.tpl"

	schemaTemplateFile = "schema.tpl"
	edgesTemplateFile  = "edges.tpl"
	fieldsTemplateFile = "fields.tpl"
	structTemplateFile = "struct.tpl"

	errTemplateFile = "err.tpl"
)

var templates = map[string]string{
	importsTemplateFile: template.Imports,
	schemaTemplateFile:  template.Schema,
	edgesTemplateFile:   template.Edges,
	fieldsTemplateFile:  template.Fields,
	structTemplateFile:  template.Struct,
	errTemplateFile:     template.Error,
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
