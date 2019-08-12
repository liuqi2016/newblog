package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Result 统一返回格式
type Result struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// //Controller 控制器接口
// type Controller interface {
// 	GET(w http.ResponseWriter, r *http.Request)
// 	POST(w http.ResponseWriter, r *http.Request)
// 	PUT(w http.ResponseWriter, r *http.Request)
// 	DELETE(w http.ResponseWriter, r *http.Request)
// }

//result 统一返回
func ReturnJson(w http.ResponseWriter, r *Result) {
	rsJSON, e := json.Marshal(r)
	if e != nil {
		fmt.Fprintf(w, e.Error())
	}
	fmt.Fprintf(w, string(rsJSON))
}
