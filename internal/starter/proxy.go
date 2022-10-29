package starter

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"pronghorn/internal/proxy"
)

type ProxyStarter struct {
	BaseStarter
}

func (this *ProxyStarter) StartBlock() bool {
	return true
}
func (this *ProxyStarter) Start(ctx StarterContext) {
	port := ctx.Props().GetDefault("app.server.port", "18080")
	log.Info("启动成功:http://localhost:" + port)
	_ = http.ListenAndServe(":"+port, &proxy.ProxyHandler{Lb: Lb})
}
