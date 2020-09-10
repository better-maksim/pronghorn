package util

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func CloneHeader(source http.Header, dest *http.Header) {
	for k, v := range source {
		dest.Set(k, v[0])
	}
}

func RequestUrl(write http.ResponseWriter, request *http.Request, url string) {
	//如果满足匹配条件，直接创建新的客户端请求
	//go 语言已经为我们提供好了 http 请求库。
	nearReq, _ := http.NewRequest(request.Method, url, request.Body)
	//将浏览器的头给客户端
	CloneHeader(request.Header, &nearReq.Header)
	nearReq.Header.Add("x-forwarded-for", request.RemoteAddr)
	dt := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ResponseHeaderTimeout: 30 * time.Second,
	}
	nearResponse, _ := dt.RoundTrip(nearReq)
	getHeader := write.Header()
	//拷贝响应头给客户端
	CloneHeader(nearResponse.Header, &getHeader)
	//写入http status
	write.WriteHeader(nearResponse.StatusCode)
	defer nearResponse.Body.Close()
	res, _ := ioutil.ReadAll(nearResponse.Body)

	_, _ = write.Write(res)
}

var LB *LoadBalance
var ServerIndexes []int
func init()  {
	LB = NewLoadBalance()
	LB.AddServer(NewHttpServer("http://localhost:9091", 5))
	LB.AddServer(NewHttpServer("http://localhost:9092",15))
	for index,server := range LB.Servers  {
		if server.Weight > 0 {
			for i :=0; i < server.Weight; i++ {
				ServerIndexes = append(ServerIndexes, index)
			}
		}
	}
	fmt.Println(ServerIndexes)
}
