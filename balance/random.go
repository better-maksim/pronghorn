package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	curIndex int
	rss      [] string
	//观察者模式
	//conf LoadBalanceConf
}

func (this *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	this.rss = append(this.rss, addr)
	return nil
}

func (this *RandomBalance) Next() string {
	if len(this.rss) == 0 {
		return ""
	}
	this.curIndex = rand.Intn(len(this.rss))
	return this.rss[this.curIndex]
}

func (this *RandomBalance) Get(key string) (string, error) {
	return this.Next(), nil
}
