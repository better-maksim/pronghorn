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

func (b *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	b.rss = append(b.rss, addr)
	return nil
}

func (b *RoundRobinBalance) Next() string {
	if len(b.rss) == 0 {
		return ""
	}
	lens := len(b.rss)
	if b.curIndex >= lens {
		b.curIndex = 0;
	}
	curAddr := b.rss[b.curIndex]
	b.curIndex = (b.curIndex + 1) % lens
	return curAddr
}

func (b *RoundRobinBalance) Get(key string) (string, error) {
	return b.Next(), nil
}
