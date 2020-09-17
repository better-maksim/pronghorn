package starter

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

type PropsStarter struct {
	BaseStarter
}

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

func (this *PropsStarter) Init(ctx StarterContext) {
	props = ini.NewIniFileCompositeConfigSource("config.ini")
}
