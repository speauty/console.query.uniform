package command

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func init() {
	//kernel.NewAppService().RegisterCliCmd(Cls{}.Cmd())
}

type Cls struct{}

func (c Cls) Cmd() *cli.Command {
	return &cli.Command{
		Name:    "clear",
		Aliases: []string{"cls"},
		Usage:   "清空屏幕;",
		Action: func(c *cli.Context) error {
			//@todo 好像没有直接的实现可以调用，所以可能需要根据OS写相应的实现，麻烦，暂时不需要
			fmt.Println("该命令暂未实现, 正在努力研发中")
			return nil
		},
		Hidden:   false,
		HideHelp: false,
	}
}
