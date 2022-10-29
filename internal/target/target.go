package target

type target struct {
	Name   string // target 名称
	Host   string // 目标 IP 或者域名
	Port   int    // 目标端口号
	Weight int    // 权重
}

func NewTarget(name, host string, port, weight int) *target {
	return &target{
		Name:   name,
		Host:   host,
		Port:   port,
		Weight: weight,
	}
}
