package user

import "context"

// Controller 控制器接口
type UserServer interface {
	Login(ctx context.Context) (r interface{}, err error)
	SaveOrUpdate(ctx context.Context) (r interface{}, err error)
	GetInfo(ctx context.Context) (r interface{}, err error)
	Logout(ctx context.Context) (r interface{}, err error)
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
type SearchByNameIn struct {
	Username string `json:"username" validate:"required"`
}
