package main

import "console.query.uniform/kernel"

func main() {
	kernel.NewCfgService()
	err := kernel.NewAppService().Run()
	if err != nil {
		return
	}
}
