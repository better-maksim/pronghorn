package starter

import (
	"pronghorn/pkg/balance"
)

type LBStarter struct {
	BaseStarter
}

var ServerIndexes []int
var ServerList *balance.Servers
var Lb balance.LoadBalance

func (this *LBStarter) Init(ctx StarterContext) {
	ServerList = balance.NewServers()
	ServerList.AddServer(balance.NewHttpServer("http://localhost:9091", 5))
	ServerList.AddServer(balance.NewHttpServer("http://localhost:9092", 15))
	Lb = balance.LoadBalanceFactory(balance.LbRandom)
	for index, server := range ServerList.Servers {
		_ = Lb.Add(server.Host)
		if server.Weight > 0 {
			for i := 0; i < server.Weight; i++ {
				ServerIndexes = append(ServerIndexes, index)
			}
		}
	}
}
