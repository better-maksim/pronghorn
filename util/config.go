package util

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var ProxyConfigs map[string]string

type EnvConfig *os.File

func init() {
	ProxyConfigs = make(map[string]string)
	EnvConfig, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	proxy, _ := EnvConfig.GetSection("proxy")
	if proxy != nil {
		secs := proxy.ChildSections()
		for _, sec := range secs {
			path, _ := sec.GetKey("path")
			pass, _ := sec.GetKey("pass")
			if path != nil && pass != nil {
				ProxyConfigs[path.Value()] = pass.Value()
			}
		}
	}
}
