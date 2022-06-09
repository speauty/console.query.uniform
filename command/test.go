package command

import (
	"console.query.uniform/kernel"
	"console.query.uniform/kernel/cfg"
	"fmt"
	"github.com/urfave/cli/v2"
)

func init() {
	kernel.NewAppService().RegisterCliCmd(Test{}.Cmd())
}

type Test struct{}

func (t Test) Cmd() *cli.Command {
	return &cli.Command{
		Name:  "test",
		Usage: "测试指令;",
		Action: func(c *cli.Context) error {
			tmpCfg := cfg.NewCfgService()
			fmt.Println(
				"项目:", tmpCfg.App.Name, "版本:", tmpCfg.App.Version, "作者:", tmpCfg.App.Author,
			)
			return nil
		},
		Hidden:   false,
		HideHelp: false,
	}
}
