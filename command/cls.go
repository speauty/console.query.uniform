package command

import (
	"console.query.uniform/kernel"
	"fmt"
	"github.com/urfave/cli/v2"
)

func init() {
	kernel.NewAppService().RegisterCliCmd(Cls{}.Cmd())
}

type Cls struct{}

func (c Cls) Cmd() *cli.Command {
	return &cli.Command{
		Name:    "clear",
		Aliases: []string{"cls"},
		Usage:   "清空屏幕;",
		Action: func(c *cli.Context) error {
			fmt.Println("该命令暂未实现, 正在努力研发中")
			return nil
		},
		Hidden:   false,
		HideHelp: false,
	}
}
