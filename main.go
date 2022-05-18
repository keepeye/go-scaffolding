package main

import (
	"myapp/commands"
	"myapp/core/config"
	"os"
	"sort"

	"myapp/server"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func initLogger() {
	// 日志
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     config.GetBool("development"),
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	initLogger()
	// 注册工具命令行
	commands := commands.RegisterCommands()
	// 注册adminServer命令
	commands = append(commands, server.Command())
	sort.Sort(cli.CommandsByName(commands))
	(&cli.App{
		Commands: commands,
	}).Run(os.Args)
}
