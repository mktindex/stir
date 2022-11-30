package main

import (
	"log"
	"os"
	"stir/cmd"
	"stir/cmd/fix"

	"github.com/urfave/cli/v2"
)

func init() {
	cli.AppHelpTemplate = fix.AppHelpTemplate
	cli.CommandHelpTemplate = fix.CommandHelpTemplate
	cli.SubcommandHelpTemplate = fix.SubcommandHelpTemplate
}

func main() {
	app := &cli.App{
		Name:  "stir",
		Usage: "混合数据分析(Stir)",
		Authors: []*cli.Author{
			{Name: "Tacey Wong", Email: "xinyong.wang@qq.com"}},
		Copyright: "© Tacey Wong All Rights Reserved",
		Action: func(ctx *cli.Context) error {
			cli.ShowAppHelp(ctx)
			return nil
		},
		Commands: []*cli.Command{cmd.ServerCMD},
	}
	app.HideHelpCommand = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
