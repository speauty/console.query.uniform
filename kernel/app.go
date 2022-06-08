package kernel

import (
	"fmt"
	"github.com/gobs/args"
	"github.com/peterh/liner"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
)

var AppApi *App
var AppOnce sync.Once

func NewAppService() *App {
	AppOnce.Do(func() {
		AppApi = &App{}
		AppApi.init()
	})
	return AppApi
}

type App struct {
	cmd *cli.App
	cfg *Cfg
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

func (a *App) init() {
	a.cfg = NewCfgService()
	a.initCmd()
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
			cli.ShowAppHelp(c)

			line := newLiner()
			defer closeLiner(line)

			for {
				if commandLine, err := line.Prompt(a.cfg.Sys.CmdLinePrompt); err == nil {
					line.AppendHistory(commandLine)

					cmdArgs := args.GetArgs(commandLine)
					if len(cmdArgs) == 0 {
						continue
					}

					s := []string{os.Args[0]}
					s = append(s, cmdArgs...)

					closeLiner(line)

					c.App.Run(s)

					line = newLiner()

				} else if err == liner.ErrPromptAborted || err == io.EOF {
					break
				} else {
					log.Print("Error reading line: ", err)
					continue
				}
			}
		} else {
			fmt.Printf("未找到命令: %s\n运行命令 %s help 获取帮助\n", c.Args().Get(0), a.cmd.Name)
		}
		return nil
	}
}

func newLiner() *liner.State {
	cfg := NewCfgService()
	line := liner.NewLiner()

	line.SetCtrlCAborts(true)

	if f, err := os.Open(cfg.Sys.CmdHistoryFile); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	return line
}

func closeLiner(line *liner.State) {
	cfg := NewCfgService()
	if f, err := os.Create(cfg.Sys.CmdHistoryFile); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
	line.Close()
}
