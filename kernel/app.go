package kernel

import (
	"console.query.uniform/util"
	"fmt"
	"github.com/gobs/args"
	"github.com/peterh/liner"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"runtime"
	"sync"
)

var appService *App
var appOnce sync.Once

func NewAppService() *App {
	appOnce.Do(func() {
		appService = &App{}
		appService.init()
	})
	return appService
}

type App struct {
	cmd *cli.App
	cfg *Cfg
	log *Log
}

func (a *App) Run() error {
	return a.getCmd().Run(os.Args)
}

func (a *App) initCmd() {
	if a.cmd == nil {
		a.cmd = cli.NewApp()

		a.cmd.Name = a.cfg.App.Name
		a.cmd.Authors = []*cli.Author{{Name: a.cfg.App.Author, Email: a.cfg.App.Email}}
		a.cmd.Version = a.cfg.App.Version
		a.cmd.Usage = runtime.GOOS + "/" + runtime.GOARCH
		a.cmd.Description = a.cfg.App.Description

		a.cmd.Action = a.getAction()
	}
}

func (a *App) initDir() {
	var err error
	if a.cfg.Sys.CmdHistoryFile != "" {
		if err = util.CreateDirRecursion(a.cfg.Sys.CmdHistoryFile); err != nil {
			a.log.Error("创建命令行历史记录目录异常", err)
		}

	}

	if err == nil && a.cfg.Log.LogFile != "" {
		if err = util.CreateDirRecursion(a.cfg.Log.LogFile); err != nil {
			a.log.Error("创建日志目录异常", err)
		}
	}

	if err == nil && a.cfg.Log.DbLogFile != "" {
		if err = util.CreateDirRecursion(a.cfg.Log.DbLogFile); err != nil {
			a.log.Error("创建数据库日志目录异常", err)
		}
	}
	if err != nil {
		panic(err)
	}
}

func (a *App) init() {
	a.cfg = NewCfgService()
	a.log = NewLogService()
	a.initCmd()
	a.initDir()
}

func (a *App) getCmd() *cli.App {
	if a.cmd == nil {
		a.initCmd()
	}
	return a.cmd
}

func (a *App) getAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		if c.NArg() == 0 {
			if err := cli.ShowAppHelp(c); err != nil {
				a.log.Warn("终端显示帮助异常", err)
			}

			line, err := NewLinerService().New()
			if err != nil {
				a.log.Warn("新建命令行异常", err)
			}
			defer func(service *Liner, liner *liner.State) {
				err := service.Close(liner)
				if err != nil {
					a.log.Warn("释放命令行异常", err)
				}
			}(NewLinerService(), line)

			for {
				if commandLine, err := line.Prompt(a.cfg.Sys.CmdLinePrompt); err == nil {
					line.AppendHistory(commandLine)
					cmdArgs := args.GetArgs(commandLine)
					if len(cmdArgs) == 0 {
						continue
					}
					s := []string{os.Args[0]}
					s = append(s, cmdArgs...)
					err := NewLinerService().Close(line)
					if err != nil {
						a.log.Warn("释放命令行异常", err)
					}
					_ = c.App.Run(s)
					line, _ = NewLinerService().New()
				} else if err == liner.ErrPromptAborted || err == io.EOF {
					break
				} else {
					a.log.Print("读取命令行异常", err)
					fmt.Println("读取命令行异常:", err)
					continue
				}
			}
		} else {
			a.log.Warn("未找到相应命令:", c.Args().Get(0))
			fmt.Println("未找到相应命令:", c.Args().Get(0))
		}
		return nil
	}
}
