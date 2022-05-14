package main

import (
	"myapp/commands"
	"os"
	"sort"

	"myapp/server"

	"github.com/urfave/cli/v2"
)

func main() {
	// 注册工具命令行
	commands := commands.RegisterCommands()
	// 注册adminServer命令
	commands = append(commands, server.Command())
	sort.Sort(cli.CommandsByName(commands))
	(&cli.App{
		Commands: commands,
	}).Run(os.Args)
}
