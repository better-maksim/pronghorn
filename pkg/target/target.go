package target

type target struct {
	Host   string // 目标 IP 或者域名
	Port   int    // 目标端口号
	Weight int    // 权重
}

func NewTarget(host string, port, weight int) *target {
	return &target{
		Host:   host,
		Port:   port,
		Weight: weight,
	}
}
