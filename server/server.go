package server

import (
	"fmt"
	"myapp/core/boost"
	"myapp/server/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func printRoutes(routes []*echo.Route) {
	fmt.Println(">>>>>>>>>routes:")
	for _, route := range routes {
		fmt.Printf("%s %s ---> %s\n", route.Method, route.Path, route.Name)
	}
	fmt.Println("<<<<<<<<<<<<<<<<<")
}

func Run(ctx *cli.Context) error {
	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(boost.CustomLogger())
	routes.Setup(app)
	app.HideBanner = true
	// printRoutes(app.Routes())
	return app.Start(":" + cast.ToString(ctx.Int("listen-port")))
}
