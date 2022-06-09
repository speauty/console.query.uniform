package main

import (
	"console.query.uniform/kernel"
	"console.query.uniform/kernel/cfg"
	"console.query.uniform/kernel/log"
)

func main() {

	cfg.NewCfgService()
	log.NewLogService()

	err := kernel.NewAppService().Run()
	if err != nil {
		return
	}
}
