package command

import "sync"

var cmdRegisterService *CmdRegister
var cmdRegisterOnce sync.Once

func NewCmdRegisterService() *CmdRegister {
	cmdRegisterOnce.Do(func() {
		cmdRegisterService = &CmdRegister{}
	})
	return cmdRegisterService
}

type CmdRegister struct{}
