package kernel

import (
	"github.com/peterh/liner"
	"log"
	"os"
	"sync"
)

var linerService *Liner
var linerOnce sync.Once

func NewLinerService() *Liner {
	linerOnce.Do(func() {
		linerService = &Liner{}
		linerService.init()
	})
	return linerService
}

type Liner struct {
	cfg *Cfg
}

func (l Liner) New() (*liner.State, error) {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	if f, err := os.Open(l.cfg.Sys.CmdHistoryFile); err == nil {
		if _, err = line.ReadHistory(f); err != nil {
			return nil, err
		}
		if err = f.Close(); err != nil {
			return nil, err
		}
	}
	return line, nil
}

func (l Liner) Close(liner *liner.State) {
	if f, err := os.Create(l.cfg.Sys.CmdHistoryFile); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		_, err = liner.WriteHistory(f)
		if err != nil {
			//todo 补充错误处理
		}
		_ = f.Close()
	}
	_ = liner.Close()
}

func (l *Liner) init() {
	l.cfg = NewCfgService()
}
