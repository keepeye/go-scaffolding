package server

import (
	"fmt"
	"myapp/core/boost"
	"myapp/core/config"
	"myapp/server/routes"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
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
	app.Use(boost.CustomHttpLogger(strings.Split(config.GetString("logger.httpLogTags"), ",")))
	routes.Setup(app)
	app.HideBanner = true
	app.HidePort = true
	// printRoutes(app.Routes())
	listenAt := ":" + cast.ToString(ctx.Int("listen-port"))
	log.Infof("http server started on %s", listenAt)
	return app.Start(listenAt)
}
