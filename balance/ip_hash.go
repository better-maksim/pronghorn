package balance

import (
	"errors"
	"hash/crc32"
)

// IPHshBalance  Hash 算法
type IPHshBalance struct {
	curIndex int
	rss      []string
	//观察者模式
	//conf LoadBalanceConf
}

func (b *IPHshBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	b.rss = append(b.rss, addr)
	return nil
}

func (b *IPHshBalance) Next() string {
	return ""
}

func (b *IPHshBalance) Get(key string) (string, error) {
	index := int(crc32.ChecksumIEEE([]byte(key))) % len(b.rss)
	return b.rss[index], nil
}
