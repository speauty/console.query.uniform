package command

import (
	"console.query.uniform/kernel"
	"console.query.uniform/kernel/cfg"
	"github.com/urfave/cli/v2"
)

func init() {
	kernel.NewAppService().RegisterCliCmd(CfgReload{}.Cmd())
}

type CfgReload struct{}

func (cr CfgReload) Cmd() *cli.Command {
	return &cli.Command{
		Name:     "cfg.reload",
		Aliases:  []string{"reload"},
		Usage:    "重新载入配置文件;",
		Category: "设置",
		Action: func(c *cli.Context) error {
			return cfg.NewCfgService().LoadCfg()
		},
		Hidden:   false,
		HideHelp: false,
	}
}
