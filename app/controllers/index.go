package controllers

import (
	"net/http"
)

//Login 登陆
func Login(w http.ResponseWriter, r *http.Request) {

}

type LoginIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginOut struct {
	BaseOut
}
