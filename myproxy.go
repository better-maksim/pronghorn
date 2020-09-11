package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
	"os"
	"os/signal"
	"proxy/util"
	"strings"
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
	lbString, _ := util.Lb.Get(strings.Split(request.RemoteAddr, ":")[0])
	url, _ := url2.Parse(lbString)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(write, request)
	_, _ = write.Write([]byte("default index html"))
}
func main() {
	c := make(chan os.Signal)
	go (func() {
		//开启 http 服务，，监听 8080 端口
		_ = http.ListenAndServe(":8080", &ProxyHandler{})
	})()
	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println(s)
}
