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

var appApi *App
var appOnce sync.Once

func NewAppService() *App {
	appOnce.Do(func() {
		appApi = &App{}
		appApi.init()
	})
	return appApi
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
			_ = cli.ShowAppHelp(c)

			line, _ := NewLinerService().New()
			defer NewLinerService().Close(line)

			for {
				if commandLine, err := line.Prompt(a.cfg.Sys.CmdLinePrompt); err == nil {
					line.AppendHistory(commandLine)
					cmdArgs := args.GetArgs(commandLine)
					if len(cmdArgs) == 0 {
						continue
					}
					s := []string{os.Args[0]}
					s = append(s, cmdArgs...)
					NewLinerService().Close(line)
					_ = c.App.Run(s)
					line, _ = NewLinerService().New()
				} else if err == liner.ErrPromptAborted || err == io.EOF {
					break
				} else {
					log.Print("Error reading line: ", err)
					continue
				}
			}
		} else {
			fmt.Printf("Command '%s' not found\n", c.Args().Get(0))
		}
		return nil
	}
}
