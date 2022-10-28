package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"

	"pronghorn/pkg/boot"
)

func main() {

	//获取程序运行文件所在路径
	file := kvs.GetCurrentFilePath("conf.ini", 1)
	//加载配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)

	boot := boot.NewBootApplication(conf)

	boot.Start()
}
