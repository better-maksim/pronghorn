package util

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"time"
)

type HttpServer struct { //目标 server
	Host   string
	Weight int
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{Host: host, Weight: weight}
}

type LoadBalance struct { //负载均衡
	Servers []*HttpServer
}

func NewLoadBalance() *LoadBalance {
	return &LoadBalance{Servers: make([]*HttpServer, 0)}
}

func (this *LoadBalance) AddServer(server *HttpServer) {
	this.Servers = append(this.Servers, server)
}

//随机算法
func (this *LoadBalance) SelectByRand() *HttpServer {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(this.Servers))
	return this.Servers[index]
}

//均衡负载之 ip_hash
func (this *LoadBalance) SelectByIpHash(ip string) *HttpServer {
	index := int(crc32.ChecksumIEEE([]byte(ip))) % len(this.Servers)
	return this.Servers[index]
}

//权重
func (this *LoadBalance) SelectByIpWeight(ip string) *HttpServer {
	rand.Seed(time.Now().UnixNano())

	sumList := make([]int, len(this.Servers))
	sum := 0

	for i = 0; i < len(this.Servers); i++ {
		sum += this.Servers[i].Weight
		sumList[i] = sum
	}
	rad := rand.Intn(sum) //左臂右开区间
	for index, value := range sumList {
		if rad < value {
			 return this.Servers[index]
		}
	}
	return this.Servers[0]
}