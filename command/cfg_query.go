package command

import (
	"console.query.uniform/kernel"
	"console.query.uniform/kernel/cfg"
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	_ "github.com/Jeffail/gabs/v2"
	"github.com/urfave/cli/v2"
	"strings"
)

func init() {
	kernel.NewAppService().RegisterCliCmd(CfgQuery{}.Cmd())
}

type CfgQuery struct{}

func (cq CfgQuery) Cmd() *cli.Command {
	return &cli.Command{
		Name:     "cfg.query",
		Aliases:  []string{"q"},
		Usage:    "查看配置文件, 支持点语法链路解析, 比如cfg.query cfg.db;",
		Category: "设置",
		Action: func(c *cli.Context) error {
			jsonData, err := json.Marshal(cfg.NewCfgService())
			if err != nil {
				fmt.Println("JSON编码异常", err)
				return err
			}
			jsonParsed, err := gabs.ParseJSON(jsonData)
			if err != nil {
				fmt.Println("Gabs编码异常", err)
				return err
			}

			cfgNode := strings.ToLower(c.Args().Get(0))
			str := ""
			if cfgNode != "" && strings.Contains(cfgNode, ".") {
				nodeArr := strings.Split(cfgNode, ".")[1:]
				if !jsonParsed.Exists(nodeArr...) {
					str = "节点不存在"
					fmt.Println(str)
					return nil
				} else {
					str = jsonParsed.Search(nodeArr...).StringIndent("", "\t")
				}
			} else {
				str = jsonParsed.StringIndent("", "\t")
			}
			fmt.Println("查看节点:", cfgNode)
			fmt.Println(str)
			return nil
		},
		Hidden:   false,
		HideHelp: false,
	}
}
