package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
	"proxy/util"
)

//代理服务 handler
type ProxyHandler struct {
}

func (*ProxyHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			write.WriteHeader(500)
			log.Println(err)
		}
	}()
	url, _ := url2.Parse(util.LB.RoundRobin().Host)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(write, request)
	_, _ = write.Write([]byte("default index html"))
}
func main() {
	_ = http.ListenAndServe(":8080", &ProxyHandler{})
	//c := make(chan os.Signal)
	//go (func() {
	//	//开启 http 服务，，监听 8080 端口
	//
	//})()
	//signal.Notify(c, os.Interrupt)
	//s := <-c
	//fmt.Println(s)
}
