package proxy

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"pronghorn/internal/balance"
	"strconv"
	"strings"
	"time"
)

//连接池
var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, //连接超时
		KeepAlive: 30 * time.Second, //长连接超时时间
	}).DialContext,
	MaxIdleConns:          100,              //最大空闲连接
	IdleConnTimeout:       90 * time.Second, //空闲超时时间
	TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
	ExpectContinueTimeout: 1 * time.Second,  //100-continue超时时间
}

func NewMultipleHostReverseProxy(lb balance.LoadBalance) *httputil.ReverseProxy {

	director := func(req *http.Request) {
		addr, err := lb.Get(strings.Split(req.RemoteAddr, ":")[0])
		if err != nil {
			log.Fatal("get next addr fail")
		}
		target, err := url.Parse(addr)
		if err != nil {
			log.Fatal(err)
		}
		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
	modifyFunc := func(resp *http.Response) error {
		if resp.StatusCode != 200 {
			//获取内容
			oldPayload, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			newPayload := []byte("StatusCode error:" + string(oldPayload))
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(newPayload))
			resp.ContentLength = int64(len(newPayload))
			resp.Header.Set("Content-Length", strconv.FormatInt(int64(len(newPayload)), 10))
		}
		return nil
	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc}
}

func singleJoiningSlash(a, b string) string {
	//合并路径
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

type ProxyHandler struct {
	Lb balance.LoadBalance
}

func (this *ProxyHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			write.WriteHeader(500)
			log.Println(err)
		}
	}()
	proxy := NewMultipleHostReverseProxy(this.Lb)
	proxy.ServeHTTP(write, request)
	_, _ = write.Write([]byte("default index html"))
}
