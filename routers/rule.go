package routers

// InitRouteDefault 初始化默认路由
func InitRouteDefault() (r *OdServer) {
	r = Default()

	r.GET("/user/login", controllers.login)

	return r
}
