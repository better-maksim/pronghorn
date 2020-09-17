package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"proxy"
)

func main() {

	//获取程序运行文件所在路径
	file := kvs.GetCurrentFilePath("conf.ini", 1)
	//加载配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	boot := proxy.NewBootApplication(conf)
	boot.Start()
}
