package command

import "sync"

var cmdRegisterService *CmdRegister
var cmdRegisterOnce sync.Once

// NewCmdRegisterService 这里没啥用，主要是为了执行包的init方法，所以强行写了这么一个service
func NewCmdRegisterService() *CmdRegister {
	cmdRegisterOnce.Do(func() {
		cmdRegisterService = &CmdRegister{}
	})
	return cmdRegisterService
}

type CmdRegister struct{}
