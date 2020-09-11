package main

//用于测试的简单web，帮助使用者快速了解 pronghorn的 使用方法
import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type webOneHandler struct {
}
type webTwoHandler struct {
}

//获取真实 ip
func (* webOneHandler) GetIP(request *http.Request) string {
	ips := request.Header.Get("r-forwarded-for")
	if ips != "" {
		ipsList := strings.Split(ips, ",")
		if len(ipsList) > 0 && ipsList[0] != "" {
			return ipsList[0]
		}
	}
	return request.RemoteAddr
}

func (this webOneHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	auth := request.Header.Get("Authorization")
	if auth == "" {
		writer.Header().Set("WWW-Authenticate", "Basic realm='您必须输入用户名和密码'")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	authList := strings.Split(auth, " ")
	if len(authList) == 2 && authList[0] == "Basic" {
		res, err := base64.StdEncoding.DecodeString(authList[1])
		if err == nil && string(res) == "123:123" {
			html := fmt.Sprintf("<h1>web one ! 来自于%s</h1>",this.GetIP(request))
			_, _ = writer.Write([]byte(html))
			return
		}
	}
	_, _ = writer.Write([]byte("用户名密码错误!"))
}

func (this webTwoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte("hello two!"))
}

func main() {
	fmt.Println("服务器启动了，请访问:9091,9092")
	c := make(chan os.Signal)

	go (func() {
		err := http.ListenAndServe(":9091", &webOneHandler{})
		fmt.Println(err)
		panic(err)
	})()

	go (func() {
		err := http.ListenAndServe(":9092", webTwoHandler{})
		fmt.Println(err)
		panic(err)
	})()

	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println(s)
}
