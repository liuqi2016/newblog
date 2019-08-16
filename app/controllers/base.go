package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SUCCESS = iota
	FALSE
)

//Result 统一返回格式
type Result struct {
	Code    int32       `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

//ReturnJSON 统一返回
func ReturnJSON(w http.ResponseWriter, data interface{}, float int) {
	r := Result{}
	r.Data = data
	if float == 0 {
		r.Code = 200
		r.Message = "success"
	} else {
		r.Code = 400
		r.Message = "false"
		r.Data = data.(error).Error()
	}
	rsJSON, e := json.Marshal(r)
	if e != nil {
		fmt.Fprintf(w, e.Error())
	}
	fmt.Fprintf(w, string(rsJSON))
}
