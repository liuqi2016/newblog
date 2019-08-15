package routers

import (
	"blog/app/controllers/blog"
	"blog/app/controllers/user"
	"blog/utils/logger"
	"fmt"
)

// InitRouteDefault 初始化默认路由context
func InitRouteDefault() (r *OdServer) {
	r = Default()
	//用户
	{
		var u user.UserServer
		u = user.UserController{}
		r.GET("/user/login", u.Login)
		r.GET("/user/info", withJWTMiddle(u.GetInfo))
		r.GET("/user/logout", withJWTMiddle(u.Logout))
	}

	//博客后台
	{
		var b blog.BlogServer
		b = blog.BlogController{}
		r.GET("/blog/lists", b.GetByPage)
		r.GET("/blog/detail", b.GetOne)
		r.POST("/blog/edit", withJWTMiddle(b.SaveOrUpdate))
		// r.GET("/blog/list", withJWTMiddle(blog.GetByPage))
		// r.GET("/blog/save", withJWTMiddle(blog.SaveOrUpdate))
	}
	//打印出所有路由
	for index, v := range r.router {
		for index2 := range v {
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
