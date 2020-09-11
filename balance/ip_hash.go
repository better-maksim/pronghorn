package balance

import (
	"errors"
	"hash/crc32"
)

type IPHshBalance struct {
	curIndex int
	rss      [] string
	//观察者模式
	//conf LoadBalanceConf
}

func (this *IPHshBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	this.rss = append(this.rss, addr)
	return nil
}

func (this *IPHshBalance) Next() string {
	return ""
}

func (this *IPHshBalance) Get(key string) (string, error) {
	index := int(crc32.ChecksumIEEE([]byte(key))) % len(this.rss)
	return this.rss[index], nil
}
