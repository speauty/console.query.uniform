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

//@todo 等待实现 这里就是核心了，循环切库，执行sql

func init() {
	kernel.NewAppService().RegisterCliCmd(Sql{}.Cmd())
}

type Sql struct{}

// 好烦，sql语句过长，会导致。。。显示不下，估计得加个直接执行sql脚本的指令

func (s Sql) Cmd() *cli.Command {
	tmpLog := log.NewLogService()
	return &cli.Command{
		Name:     "db.query",
		Aliases:  []string{"sql"},
		Usage:    "执行查询, sql \"sql string\";",
		Category: "数据库",
		Action: func(c *cli.Context) error {
			sqlStr := strings.ToLower(c.Args().Get(0))
			tmpCfg := cfg.NewCfgService()
			for _, val := range tmpCfg.Db {
				test := &db.Db{}
				if err := test.NewDb(val.Driver, val.Dsn); err != nil {
					fmt.Println("数据库连接异常", val.Name, val.Dsn)
					tmpLog.Error(err)
					continue
				}
				// 这里的话，就不会单独检测连接是否有效了，直接查询就行
				result, err := test.Exec(sqlStr)
				if err != nil {
					fmt.Println("数据库查询异常", err)
					tmpLog.Error(err)
					continue
				}
				rows, _ := result.RowsAffected()
				fmt.Println(val.Name, "查询正常, 影响行数", rows)
			}
			return nil
		},
		Hidden:   false,
		HideHelp: false,
	}
}
