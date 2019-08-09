package controllers

import "net/http"

type BaseOut struct {
	Code int `json:"code"`
	Data interface{}
}

type ControllerInterface interface {
	GET(w http.ResponseWriter, r *http.Request)
	POST(w http.ResponseWriter, r *http.Request)
	PUT(w http.ResponseWriter, r *http.Request)
	DELETC(w http.ResponseWriter, r *http.Request)
}
