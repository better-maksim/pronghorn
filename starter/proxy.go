package starter

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"proxy/proxy"
)

type ProxyStarter struct {
	BaseStarter
}

func (this *ProxyStarter) StartBlock() bool {
	return true
}
func (this *ProxyStarter) Start(ctx StarterContext) {
	prot := ctx.Props().GetDefault("app.server.port", "18080")
	log.Info("启动成功:http://localhost:" + prot)
	_ = http.ListenAndServe(":"+prot, &proxy.ProxyHandler{Lb: Lb})

}
