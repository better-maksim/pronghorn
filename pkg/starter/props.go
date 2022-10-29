package starter

import (
	log "github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	BaseStarter
}

func (p *PropsStarter) Init(ctx StarterContext) {
	props = ctx.Props()
	log.Info("初始化配置.")

}

type SystemAccount struct {
	AccountNo   string
	AccountName string
	UserId      string
	Username    string
}
