package main

import (
	"mywebmall/conf"
	"mywebmall/routers"
)

func main() {
	// 配置载入以及建表
	conf.Init()
	r := routers.NewRouter()
	r.Run(conf.HttpPort)
}
