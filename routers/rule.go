package routers

import (
	usercontroller "blog/app/controllers/user"
	"blog/utils/logger"
	"fmt"
)

// InitRouteDefault 初始化默认路由
func InitRouteDefault() (r *OdServer) {
	r = Default()
	//用户
	{
		var i usercontroller.UserServer
		i = usercontroller.UserController{}
		r.GET("/user/login", i.Login)
		r.GET("/user/testadd", i.TestAdd)
		r.GET("/user/info", i.GetByToken)
		r.GET("/user/logout", i.Logout)
	}
	//打印出所有路由
	for index, v := range r.router {
		for index2, _ := range v {
			logger.Info(fmt.Sprintf("Method:%s Route:%s", getMethod(index), index2))
		}
	}
	return r
}
func getMethod(i int) (s string) {
	switch i {
	case 0:
		s = "GET"
		break
	case 1:
		s = "POST"
		break
	case 2:
		s = "PUT"
		break
	case 3:
		s = "DELETE"
		break
	}
	return
}
