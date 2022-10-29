package starter

import (
	balance2 "pronghorn/internal/balance"
)

type LBStarter struct {
	BaseStarter
}

var ServerIndexes []int
var ServerList *balance2.Servers
var Lb balance2.LoadBalance

func (this *LBStarter) Init(ctx StarterContext) {
	ServerList = balance2.NewServers()
	ServerList.AddServer(balance2.NewHttpServer("http://localhost:9091", 5))
	ServerList.AddServer(balance2.NewHttpServer("http://localhost:9092", 15))
	Lb = balance2.LoadBalanceFactory(balance2.LbRandom)
	for index, server := range ServerList.Servers {
		_ = Lb.Add(server.Host)
		if server.Weight > 0 {
			for i := 0; i < server.Weight; i++ {
				ServerIndexes = append(ServerIndexes, index)
			}
		}
	}
}
