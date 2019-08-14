package routers

import (
	"blog/app/controllers"
	"blog/utils/jwt"
	"context"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

//TokenIn token
type TokenIn struct {
	Token string `json:"token" validate:"required" `
}

//JWT 认证中间件
func withJWTMiddle(h HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//1.数据接收
		rs := controllers.Result{Code: 20000}
		var token TokenIn
		token.Token = r.Header.Get("X-Token")
		// tokens := strings.SplitN(token.Token, " ", 2)
		// if len(tokens) != 2 {
		// 	rs.Data = "token格式错误"
		// 	controllers.ReturnJSON(w, &rs)
		// 	return
		// }
		// token.Token = tokens[1]

		//2.数据验证
		validate := validator.New()
		err := validate.Struct(token)
		if err != nil {
			rs.Code = 400
			rs.Data = err.Error()
			controllers.ReturnJSON(w, &rs)
			return
		}
		//3.验证token
		userInfo, err := jwt.ParseToken(token.Token)
		if err != nil {
			rs.Code = 400
			rs.Data = "token错误"
			controllers.ReturnJSON(w, &rs)
			return
		}
		ctx := context.WithValue(r.Context(), "userinfo", userInfo)
		// ctx2 := context.WithValue(ctx, "password", userInfo.Password)
		h(w, r.WithContext(ctx))
	}
}
