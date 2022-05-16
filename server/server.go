package server

import (
	"myapp/core/boost"
	"myapp/server/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:   "server",
		Usage:  "启动http服务器",
		Action: Run,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "listen-port",
				Usage:    "端口号",
				Aliases:  []string{"p"},
				Required: true,
			},
		},
	}
}

func Run(ctx *cli.Context) error {
	app := gin.New()
	app.Use(boost.CustomLogger(), gin.Recovery())
	routes.Setup(app)
	return app.Run(":" + cast.ToString(ctx.Int("listen-port")))
}
