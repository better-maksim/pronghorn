package balance

import (
	"errors"
)

type RoundRobinBalance struct {
	curIndex int
	rss      []string
	//观察主题
	//conf LoadBalanceConf
}

func (this *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	this.rss = append(this.rss, addr)
	return nil
}

func (this *RoundRobinBalance) Next() string {
	if len(this.rss) == 0 {
		return ""
	}
	lens := len(this.rss)
	if this.curIndex >= lens {
		this.curIndex = 0;
	}
	curAddr := this.rss[this.curIndex]
	this.curIndex = (this.curIndex + 1) % lens
	return curAddr
}

func (this *RoundRobinBalance) Get(key string) (string, error) {
	return this.Next(), nil
}
