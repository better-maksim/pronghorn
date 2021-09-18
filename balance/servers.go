package balance

// HttpServer 目标 server
type HttpServer struct {
	Host   string
	Weight int
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{Host: host, Weight: weight}
}

// Servers 负载均衡
type Servers struct {
	Servers      []*HttpServer
	CurrentIndex int //指向当前访问的服务器
}

func NewServers() *Servers {
	return &Servers{Servers: make([]*HttpServer, 0)}
}

func (s *Servers) AddServer(server *HttpServer) {
	s.Servers = append(s.Servers, server)
}
