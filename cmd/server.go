package cmd

import (
	"stir/server"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var ServerCMD = &cli.Command{
	Name:            "server",
	Usage:           "Web-API Server相关命令",
	Action:          serverAction,
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		{
			Name: "start", Usage: "启动API Server",
			Action: func(ctx *cli.Context) error {
				host := ctx.String("host")
				port := ctx.String("port")
				mode := gin.ReleaseMode
				if ctx.Bool("debug") {
					mode = gin.DebugMode
				}

				return server.Start(host, port, mode)
			},
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "host", Aliases: []string{"H"}, Usage: "设置服务`主机IP`", Value: "127.0.0.1"},
				&cli.StringFlag{Name: "port", Aliases: []string{"P"}, Usage: "设置服务`端口`", Value: "8080"},
				&cli.BoolFlag{Name: "debug", Aliases: []string{"D"}, Usage: "启动调试模式", Value: false},
			},
		},
	},
}

func serverAction(ctx *cli.Context) error {
	return nil
}
