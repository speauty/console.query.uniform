package main

import (
	"console.query.uniform/command"
	"console.query.uniform/kernel"
	"console.query.uniform/kernel/cfg"
	"console.query.uniform/kernel/log"
)

func main() {

	cfg.NewCfgService()
	log.NewLogService()

	app := kernel.NewAppService()
	// 注册命令
	command.NewCmdRegisterService()
	err := app.Run()
	if err != nil {
		return
	}
}
