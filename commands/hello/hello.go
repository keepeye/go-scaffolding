package hello

import "github.com/urfave/cli/v2"

func DefineCommand() *cli.Command {
	return &cli.Command{
		Name:   "hello",
		Usage:  "say hello",
		Action: Run,
	}
}

func Run(ctx *cli.Context) error {
	return cli.Exit("hello world!!", 1)
}
