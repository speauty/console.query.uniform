package main

import (
	"console.query.uniform/command"
	"console.query.uniform/kernel"
	"console.query.uniform/kernel/cfg"
	"console.query.uniform/kernel/log"
)

func main() {

	// 注册配置服务
	cfg.NewCfgService()
	// 注册日志服务
	log.NewLogService()

	// 注册应用服务，由于指令后注册，所以，下面需要在注册指令后，才启动应用
	app := kernel.NewAppService()
	// 注册指令集，内部通过init函数实现单个指令注册
	command.NewCmdRegisterService()
	err := app.Run()
	if err != nil {
		// 应用启动异常，直接panic即可
		panic(err)
	}
}
