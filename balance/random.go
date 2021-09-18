package balance

import (
	"errors"
	"math/rand"
)

//RandomBalance 随机算法
type RandomBalance struct {
	curIndex int
	rss      [] string
	//观察者模式
	//conf LoadBalanceConf
}

func (b *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	b.rss = append(b.rss, addr)
	return nil
}

func (b *RandomBalance) Next() string {
	if len(b.rss) == 0 {
		return ""
	}
	b.curIndex = rand.Intn(len(b.rss))
	return b.rss[b.curIndex]
}

func (b *RandomBalance) Get(key string) (string, error) {
	return b.Next(), nil
}
