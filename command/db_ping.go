package command

import (
	"console.query.uniform/kernel"
	"console.query.uniform/kernel/cfg"
	"console.query.uniform/kernel/db"
	"console.query.uniform/kernel/log"
	"fmt"
	"github.com/urfave/cli/v2"
)

func init() {
	kernel.NewAppService().RegisterCliCmd(DbPing{}.Cmd())
}

type DbPing struct{}

//@todo 该检测可能，只针对单个数据库连接进行，所以后面这里可能需要接收参数（比如索引01234等），进行处理

func (t DbPing) Cmd() *cli.Command {
	tmpLog := log.NewLogService()
	return &cli.Command{
		Name:     "db.ping",
		Aliases:  []string{"ping"},
		Usage:    "探测数据库连接是否正常，主要通过Ping实现;",
		Category: "数据库",
		Action: func(c *cli.Context) error {
			tmpCfg := cfg.NewCfgService()
			for _, val := range tmpCfg.Db {
				test := &db.Db{}
				if err := test.NewDb(val.Driver, val.Dsn); err != nil {
					fmt.Println("数据库连接异常", val.Name, val.Dsn)
					tmpLog.Error(err)
					continue
				}
				//注意: 127打头的都是回环地址
				//计算机以回环地址发送的消息, 并不会由链路层送走, 而是直接被本机网络层捕获
				if err := test.Ping(); err != nil {
					fmt.Println("数据库PING异常", val.Name, val.Dsn)
					tmpLog.Error(err)
					continue
				}
				fmt.Println(val.Name, "连接正常", val.Dsn)
			}
			return nil
		},
		Hidden:   false,
		HideHelp: false,
	}
}
