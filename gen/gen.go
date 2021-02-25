package gen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/config"
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/console"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
	"github.com/wenj91/mctl/model"
	"github.com/wenj91/mctl/parser"
	"github.com/wenj91/mctl/template"
)

const (
	pwd             = "."
	createTableFlag = `(?m)^(?i)CREATE\s+TABLE` // ignore case
)

type (
	defaultGenerator struct {
		//source string
		dir string
		console.Console
		pkg string
		cfg *config.Config
	}
	Option func(generator *defaultGenerator)
)

func NewDefaultGenerator(dir, pkg string, cfg *config.Config, opt ...Option) (*defaultGenerator, error) {
	if dir == "" {
		dir = pwd
	}
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	dir = dirAbs
	err = util.MkdirIfNotExist(dir)
	if err != nil {
		return nil, err
	}

	generator := &defaultGenerator{dir: dir, cfg: cfg, pkg: pkg}
	var optionList []Option
	optionList = append(optionList, newDefaultOption())
	optionList = append(optionList, opt...)
	for _, fn := range optionList {
		fn(generator)
	}
	return generator, nil
}

func WithConsoleOption(c console.Console) Option {
	return func(generator *defaultGenerator) {
		generator.Console = c
	}
}

func newDefaultOption() Option {
	return func(generator *defaultGenerator) {
		generator.Console = console.NewColorConsole()
	}
}

func (g *defaultGenerator) StartFromDDL(source string) error {
	modelList, err := g.genFromDDL(source)
	if err != nil {
		return err
	}

	return g.createFile(modelList)
}

func (g *defaultGenerator) StartFromInformationSchema(db string, columns map[string][]*model.Column) error {
	m := make(map[string]string)
	for tableName, column := range columns {
		table, err := parser.ConvertColumn(db, tableName, column)
		if err != nil {
			return err
		}

		code, err := g.genModel(*table)
		if err != nil {
			return err
		}

		m[table.Name.Source()] = code
	}
	return g.createFile(m)
}

func (g *defaultGenerator) createFile(modelList map[string]string) error {
	dirAbs, err := filepath.Abs(g.dir)
	if err != nil {
		return err
	}

	g.dir = dirAbs
	g.pkg = filepath.Base(dirAbs)
	err = util.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	for tableName, code := range modelList {
		tn := stringx.From(tableName)
		modelFilename := tn.ToCamel() + "Query"

		name := modelFilename + ".java"
		filename := filepath.Join(dirAbs, name)
		if util.FileExists(filename) {
			g.Warning("%s already exists, ignored.", name)
			continue
		}
		err = ioutil.WriteFile(filename, []byte(code), os.ModePerm)
		if err != nil {
			return err
		}
	}

	g.Success("Done.")
	return nil
}

// ret1: key-table name,value-code
func (g *defaultGenerator) genFromDDL(source string) (map[string]string, error) {
	ddlList := g.split(source)
	m := make(map[string]string)
	for _, ddl := range ddlList {
		table, err := parser.Parse(ddl)
		if err != nil {
			return nil, err
		}
		code, err := g.genModel(*table)
		if err != nil {
			return nil, err
		}
		m[table.Name.Source()] = code
	}
	return m, nil
}

type (
	Table struct {
		parser.Table
		ContainsUniqueKey bool
	}
)

func (g *defaultGenerator) genModel(in parser.Table) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}

	importsCode, err := genImports(in.ContainsDecimal())
	if err != nil {
		return "", err
	}

	var table Table
	table.Table = in

	fieldMethodCode, err := genFieldMethod(table)
	if err != nil {
		return "", err
	}

	tableCode, err := genTable(table)
	if err != nil {
		return "", err
	}

	queryMethodCode, err := genQueryMethod(table)
	if err != nil {
		return "", err
	}

	text1, err := util.LoadTemplate(category, queryTemplateFile, template.Query)
	if err != nil {
		return "", err
	}

	camelTableName := table.Name.ToCamel()
	output, err := util.With("query").Parse(text1).Execute(map[string]interface{}{
		"pkg":                   g.pkg,
		"imports":               importsCode,
		"upperStartCamelObject": camelTableName,
		"table":                 tableCode,
		"queryMethod":           queryMethodCode,
		"fieldMethods":          fieldMethodCode,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func wrapWithRawString(v string) string {
	if v == "`" {
		return v
	}

	if !strings.HasPrefix(v, "`") {
		v = "`" + v
	}

	if !strings.HasSuffix(v, "`") {
		v = v + "`"
	} else if len(v) == 1 {
		v = v + "`"
	}
	return v
}
