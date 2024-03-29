package main

import (
	"blog/models"
	"blog/routers"
	"blog/utils/logger"
	"blog/utils/setting"
	"fmt"
	"net/http"
)

func main() {
	// 数据库
	{
		db := models.NewDB()
		db.LogMode(true)
	}
	//http服务及路由
	{
		r := routers.InitRouteDefault()
		logger.Info(fmt.Sprintf("监听端口:%d", setting.HTTPPort))
		s := &http.Server{
			Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
			Handler:        r,
			ReadTimeout:    setting.ReadTimeout,
			WriteTimeout:   setting.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		}
		err := s.ListenAndServe()
		if err != nil {
			logger.Error(err.Error())
		}
	}
}
