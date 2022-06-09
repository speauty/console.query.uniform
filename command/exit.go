package command

import (
	"console.query.uniform/kernel"
	"github.com/urfave/cli/v2"
)

func init() {
	kernel.NewAppService().RegisterCliCmd(Exit{}.Cmd())
}

type Exit struct{}

func (e Exit) Cmd() *cli.Command {
	return &cli.Command{
		Name:    "exit",
		Aliases: []string{"quit"},
		Usage:   "退出程序;",
		Action: func(c *cli.Context) error {
			return cli.Exit("", 0)
		},
		Hidden:   false,
		HideHelp: false,
	}
}
