package commands

import (
	"myapp/commands/hello"

	"github.com/urfave/cli/v2"
)

func RegisterCommands() []*cli.Command {
	commands := make([]*cli.Command, 0)
	// 注册所有工具命令
	commands = append(commands, hello.DefineCommand())
	return commands
}
