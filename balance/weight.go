package balance

import (
	"errors"
	"strconv"
)

type WeightNode struct {
	addr            string
	weight          int //权重值
	currentWeight   int //节点权重
	effectiveWeight int // 有效权重
}
type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
	rws      []int
	//观察主题
	//conf LoadBalanceConf
}

func (this *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	parInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{addr: params[0], weight: int(parInt)}
	node.effectiveWeight = node.weight
	this.rss = append(this.rss, node)
	return nil
}

func (this *WeightRoundRobinBalance) Next() string {

	total := 0
	var best *WeightNode

	for i := 0; i < len(this.rss); i++ {

		w := this.rss[i]
		//step 1 统计所有有效权重之和
		total += w.effectiveWeight
		//step 2 变更节点节点临时权重+节点有效权重
		w.currentWeight += w.effectiveWeight
		//step 3 有效权重默认与权重相同，通讯异常时 - 1， 通讯成功+1，直到恢复 weight 大小
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		//step 4 选择最大临时权重节点
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return ""
	}
	//step 5 变更临时权重为 临时权重- 有效权重之合
	best.currentWeight -= total
	return best.addr
}

func (this *WeightRoundRobinBalance) Get(key string) (string, error) {
	return this.Next(), nil
}
