package command

import (
	"console.query.uniform/kernel"
	"console.query.uniform/kernel/cfg"
	"console.query.uniform/kernel/db"
	"console.query.uniform/kernel/log"
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

func init() {
	kernel.NewAppService().RegisterCliCmd(Script{}.Cmd())
}

type Script struct{}

func (s Script) Cmd() *cli.Command {
	tmpLog := log.NewLogService()
	return &cli.Command{
		Name:     "db.script",
		Aliases:  []string{"script"},
		Usage:    "执行脚本文件, script \"sql file\";",
		Category: "数据库",
		Action: func(c *cli.Context) error {
			pathStr := strings.ToLower(c.Args().Get(0))
			if pathStr == "" {
				fmt.Println("未检测到脚本文件路径")
				return nil
			}
			//sqlStr, err := util.LoadFile2Str(pathStr)
			tmpCfg := cfg.NewCfgService()
			//if err != nil {
			//	fmt.Println("文件载入异常", err)
			//	tmpLog.Error(err)
			//	return err
			//}
			for _, val := range tmpCfg.Db {
				test := &db.Db{}
				if err := test.NewDb(val.Driver, val.Dsn); err != nil {
					fmt.Println("数据库连接异常", val.Name, val.Dsn)
					tmpLog.Error(err)
					continue
				}
				_, err := test.LoadFile(pathStr)
				if err != nil {
					fmt.Println("数据库查询异常", err)
					tmpLog.Error(err)
					continue
				}
				fmt.Println(val.Name, "查询完成")
				test.CloseDb()
			}
			return nil
		},
		Hidden:   false,
		HideHelp: false,
	}
}
