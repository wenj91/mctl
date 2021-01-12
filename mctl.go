package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/tools/goctl/tpl"
	"github.com/urfave/cli"
	model "github.com/wenj91/model/command"
)

var (
	BuildVersion = "1.0.0"
	commands     = []cli.Command{
		{
			Name:  "model",
			Usage: "generate model code",
			Subcommands: []cli.Command{
				{
					Name:  "mysql",
					Usage: `generate mysql model`,
					Subcommands: []cli.Command{
						{
							Name:  "ddl",
							Usage: `generate mysql model from ddl`,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "src, s",
									Usage: "the path or path globbing patterns of the ddl",
								},
								cli.StringFlag{
									Name:  "dir, d",
									Usage: "the target dir",
								},
								cli.StringFlag{
									Name:  "style",
									Usage: "the file naming format, see [https://github.com/tal-tech/go-zero/tree/master/tools/goctl/config/readme.md]",
								},
								cli.BoolFlag{
									Name:  "cache, c",
									Usage: "generate code with cache [optional]",
								},
								cli.BoolFlag{
									Name:  "idea",
									Usage: "for idea plugin [optional]",
								},
							},
							Action: model.MysqlDDL,
						},
						{
							Name:  "datasource",
							Usage: `generate model from datasource`,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "url",
									Usage: `the data source of database,like "root:password@tcp(127.0.0.1:3306)/database`,
								},
								cli.StringFlag{
									Name:  "table, t",
									Usage: `the table or table globbing patterns in the database`,
								},
								cli.BoolFlag{
									Name:  "cache, c",
									Usage: "generate code with cache [optional]",
								},
								cli.StringFlag{
									Name:  "dir, d",
									Usage: "the target dir",
								},
								cli.StringFlag{
									Name:  "style",
									Usage: "the file naming format, see [https://github.com/tal-tech/go-zero/tree/master/tools/goctl/config/readme.md]",
								},
								cli.BoolFlag{
									Name:  "idea",
									Usage: "for idea plugin [optional]",
								},
							},
							Action: model.MyDataSource,
						},
					},
				},
			},
		},
		{
			Name:  "template",
			Usage: "template operation",
			Subcommands: []cli.Command{
				{
					Name:   "init",
					Usage:  "initialize the all templates(force update)",
					Action: tpl.GenTemplates,
				},
				{
					Name:   "clean",
					Usage:  "clean the all cache templates",
					Action: tpl.CleanTemplates,
				},
				{
					Name:  "update",
					Usage: "update template of the target category to the latest",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "category,c",
							Usage: "the category of template, enum [api,rpc,model,docker,kube]",
						},
					},
					Action: tpl.UpdateTemplates,
				},
				{
					Name:  "revert",
					Usage: "revert the target template to the latest",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "category,c",
							Usage: "the category of template, enum [api,rpc,model,docker,kube]",
						},
						cli.StringFlag{
							Name:  "name,n",
							Usage: "the target file name of template",
						},
					},
					Action: tpl.RevertTemplates,
				},
			},
		},
	}
)

func main() {
	logx.Disable()

	app := cli.NewApp()
	app.Usage = "a cli tool to generate code"
	app.Version = fmt.Sprintf("%s %s/%s", BuildVersion, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	// cli already print error messages
	if err := app.Run(os.Args); err != nil {
		fmt.Println("error:", err)
	}
}
