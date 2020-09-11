package middleware

import (
	"math"
	"net/http"
)

const abortIndex int8 = math.MaxInt8 / 2 //最多 63 个中间件

//router 结构体
type SliceRouter struct {
	groups []*SliceGroup
}

//构造 Router
func NewSliceRouter() *SliceRouter {
	return &SliceRouter{}
}

//group 结构体
type SliceGroup struct {
	*SliceRouter
	path     string
	//handlers [] Handlerfunc
}

func (this *SliceRouterContext) Group(path string) *SliceGroup {
	return & SliceGroup{
		SliceRouter: g,
		path:        path,
		handlers:    nil,
	}
}

type SliceRouterContext struct {
	Rw  http.ResponseWriter
	Req *http.Request
}

