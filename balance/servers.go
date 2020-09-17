package balance

//目标 server
type HttpServer struct {
	Host   string
	Weight int
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{Host: host, Weight: weight}
}

//负载均衡
type Servers struct {
	Servers      []*HttpServer
	CurrentIndex int //指向当前访问的服务器
}

func NewServers() *Servers {
	return &Servers{Servers: make([]*HttpServer, 0)}
}

func (this *Servers) AddServer(server *HttpServer) {
	this.Servers = append(this.Servers, server)
}
