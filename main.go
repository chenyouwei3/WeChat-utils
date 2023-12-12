package main

import (
	"wechat-utils/initialize"
	"wechat-utils/router"
)

func main() {
	initialize.Init()
	engine := router.GetEngine()
	if err := engine.Run(":9093"); err != nil {
		panic(err)
	}
}
