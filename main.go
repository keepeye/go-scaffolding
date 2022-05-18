package main

import (
	"myapp/commands"
	"myapp/core/config"
	"os"
	"sort"
	"strings"

	"myapp/server"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func initLogger() {
	// 日志输出格式
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     config.GetBool("development"),
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// logrus.SetFormatter(&logrus.JSONFormatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })
	// 日志级别设置
	logLevel := config.GetString("logger.level")
	logrus.Info("日志输出级别:", logLevel)
	switch strings.ToLower(logLevel) {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
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
