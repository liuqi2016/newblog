package routers

import (
	"fmt"
	"net/http"
)

//OdServer 实现IOdServer的接口，以及http提供ServeHttp方法
type OdServer struct {
	router MethodMaps
}

// IOdServer 接口
type IOdServer interface {
	GET(url string, f HandlerFunc)
	POST(url string, f HandlerFunc)
	PUT(url string, f HandlerFunc)
	DELETE(url string, f HandlerFunc)
}

//HandlerMapped 代码逻辑集合
type HandlerMapped struct {
	f HandlerFunc
}

//HandlerFunc 接口函数单位，即我们编写代码逻辑的函数
type HandlerFunc func(w http.ResponseWriter, req *http.Request)

//Default 默认路由对象 返回 handle集合
func Default() *OdServer {
	return &OdServer{
		router: NewRouter(),
	}
}

//实现Handler接口，匹配方法以及路径
func (o *OdServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//转发给doHandler进行执行
	o.doHandler(w, req)
}

//判断需要执行的Http Method，从而查找对应的接口并且执行
func (o *OdServer) doHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		{
			if hm, ok := o.router.GetMapping(req.URL.Path); ok {
				hm.f(w, req)
			} else {
				fmt.Fprintln(w, "{\"code\":404,\"data\":\"not find page\"}")
			}
		}
	case http.MethodPost:
		{
			if hm, ok := o.router.PostMapping(req.URL.Path); ok {
				hm.f(w, req)
			} else {
				fmt.Fprintln(w, "{\"code\":404,\"data\":\"not find page\"}")
			}

		}
	case http.MethodDelete:
		{
			if hm, ok := o.router.DeleteMapping(req.URL.Path); ok {
				hm.f(w, req)
			} else {
				fmt.Fprintln(w, "{\"code\":404,\"data\":\"not find page\"}")
			}
		}
	case http.MethodPut:
		{
			if hm, ok := o.router.PutMapping(req.URL.Path); ok {
				hm.f(w, req)
			} else {
				fmt.Fprintln(w, "{\"code\":404,\"data\":\"not find page\"}")
			}
		}
	default:
		{
			fmt.Fprintln(w, "{\"code\":404,\"data\":\"not find page\"}")
		}
	}
}

//GET 添加GET路由
func (o *OdServer) GET(url string, f HandlerFunc) {
	o.router.GetAdd(url, HandlerMapped{f: f})
}

//POST 添加POST路由
func (o *OdServer) POST(url string, f HandlerFunc) {
	o.router.PostAdd(url, HandlerMapped{f: f})
}

//PUT 添加PUT路由
func (o *OdServer) PUT(url string, f HandlerFunc) {
	o.router.PutAdd(url, HandlerMapped{f: f})
}

//DELETE 添加DELETE路由
func (o *OdServer) DELETE(url string, f HandlerFunc) {
	o.router.DeleteAdd(url, HandlerMapped{f: f})
}
