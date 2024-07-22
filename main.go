package main

import (
	"gin_error/controller"
	"gin_error/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Register)
	r.POST("/login", controller.Login)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWTAuth(), middleware.CasbinHandler())
	{
		// 测试路由
		apiv1.GET("/hello", controller.Hello)

		// 权限策略管理
		apiv1.POST("/casbin", controller.Create)
		apiv1.GET("/casbin/list", controller.List)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
