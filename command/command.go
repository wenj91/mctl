package command

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/tools/goctl/config"
	"github.com/tal-tech/go-zero/tools/goctl/util/console"
	"github.com/urfave/cli"
	"github.com/wenj91/mctl/gen"
	"github.com/wenj91/mctl/model"
	"github.com/wenj91/mctl/util"
)

var errNotMatched = errors.New("sql not matched")

const (
	flagSrc   = "src"
	flagDir   = "dir"
	flagIdea  = "idea"
	flagUrl   = "url"
	flagTable = "table"
	flagStyle = "style"
	flagPkg   = "pkg"
)

func MysqlDDL(ctx *cli.Context) error {
	src := ctx.String(flagSrc)
	dir := ctx.String(flagDir)
	idea := ctx.Bool(flagIdea)
	style := ctx.String(flagStyle)
	pkg := ctx.String(flagPkg)
	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}
	return fromDDl(src, dir, cfg, idea, pkg)
}

func MyDataSource(ctx *cli.Context) error {
	url := strings.TrimSpace(ctx.String(flagUrl))
	dir := strings.TrimSpace(ctx.String(flagDir))
	idea := ctx.Bool(flagIdea)
	style := ctx.String(flagStyle)
	pkg := ctx.String(flagPkg)
	pattern := strings.TrimSpace(ctx.String(flagTable))
	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}

	return fromDataSource(url, pattern, dir, cfg, idea, pkg)
}

func fromDDl(src, dir string, cfg *config.Config, idea bool, pkg string) error {
	log := console.NewConsole(idea)
	src = strings.TrimSpace(src)
	if len(src) == 0 {
		return errors.New("expected path or path globbing patterns, but nothing found")
	}

	files, err := util.MatchFiles(src)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errNotMatched
	}

	var source []string
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		source = append(source, string(data))
	}
	generator, err := gen.NewDefaultGenerator(dir, cfg, gen.WithConsoleOption(log))
	if err != nil {
		return err
	}

	err = generator.StartFromDDL(strings.Join(source, "\n"))
	return err
}

func fromDataSource(url, pattern, dir string, cfg *config.Config, idea bool, pkg string) error {
	log := console.NewConsole(idea)
	if len(url) == 0 {
		log.Error("%v", "expected data source of mysql, but nothing found")
		return nil
	}

	if len(pattern) == 0 {
		log.Error("%v", "expected table or table globbing patterns, but nothing found")
		return nil
	}

	dsn, err := mysql.ParseDSN(url)
	if err != nil {
		return err
	}

	logx.Disable()
	databaseSource := strings.TrimSuffix(url, "/"+dsn.DBName) + "/information_schema"
	db := sqlx.NewMysql(databaseSource)
	im := model.NewInformationSchemaModel(db)

	tables, err := im.GetAllTables(dsn.DBName)
	if err != nil {
		return err
	}

	matchTables := make(map[string][]*model.Column)
	for _, item := range tables {
		match, err := filepath.Match(pattern, item)
		if err != nil {
			return err
		}

		if !match {
			continue
		}
		columns, err := im.FindByTableName(dsn.DBName, item)
		if err != nil {
			return err
		}
		matchTables[item] = columns
	}

	if len(matchTables) == 0 {
		return errors.New("no tables matched")
	}

	generator, err := gen.NewDefaultGenerator(dir, cfg, gen.WithConsoleOption(log))
	if err != nil {
		return err
	}

	err = generator.StartFromInformationSchema(dsn.DBName, matchTables)
	return err
}
