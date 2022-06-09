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
			// 直接调用cli的退出方法即可，相当于做了层包裹，当然也可以做点其他退出之前的记录
			return cli.Exit("", 0)
		},
		Hidden:   false,
		HideHelp: false,
	}
}
