package usercontroller

import (
	"net/http"
)

// Controller 控制器接口
type UserServer interface {
	Login(w http.ResponseWriter, r *http.Request)
	TestAdd(w http.ResponseWriter, r *http.Request)
	GetInfo(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type LoginIn struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LoginOut struct {
	Token string `json:"token"`
}

type GetInfoOut struct {
	Roles        string `json:"roles"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
	Name         string `json:"name"`
}
