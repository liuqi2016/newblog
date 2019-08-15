package routers

import (
	"blog/utils/jwt"
	"context"
	"net/http"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

//TokenIn token
type TokenIn struct {
	Token string `json:"token" validate:"required" `
}

//JWT 认证中间件
func withJWTMiddle(h HandlerFunc) HandlerFunc {
	return func(ctx context.Context) (result interface{}, err error) {
		r := ctx.Value("request").(*http.Request)
		//1.数据接收
		var token TokenIn
		token.Token = r.Header.Get("Authorization")
		//2.数据验证
		validate := validator.New()
		err = validate.Struct(token)
		if err != nil {
			return
		}
		tokenArr := strings.Split(token.Token, " ")
		token.Token = tokenArr[1]
		//3.验证token
		userInfo, err := jwt.ParseToken(token.Token)
		if err != nil {
			return
		}
		ctx = context.WithValue(ctx, "userinfo", userInfo)
		// ctx2 := context.WithValue(ctx, "password", userInfo.Password)
		result, err = h(ctx)
		return
	}
}
