package main

import (
	"gin_error/controller"
	logx "gin_error/log"
	"gin_error/middleware"
	"gin_error/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	logx.Logger = logx.NewElkLogger("127.0.0.1", 50000, 5)
	logx.Logger.Infof("connect mysql success")
	_, err := os.Stat("./certs/private.pem")
	if err == nil {
		log.Println("RSA密钥文件已经存在！")
	} else {
		if err := utils.GenerateRSAKey(); err != nil {
			log.Fatal("RSA密钥生成失败!")
		}
		log.Println("RSA密钥生成成功！")
	}
	//统一响应
	r.GET("/response_demo", controller.Response_demo)
	r.GET("/ping", controller.Register)
	r.POST("/login", controller.Login)
	//对称加密AES CBC加密块模式
	r.POST("/encryptAES", controller.Encrypt)
	//对称解密AES CBC加密块模式
	r.POST("/decryptAES", controller.Decrypt)
	//RSA非对称加密
	r.POST("/encryptRSA", controller.EncryptRSA)
	//RSA非对称加密
	r.POST("/decryptRSA", controller.DecryptRSA)
	//RSA签名和验签
	r.POST("/signRSA", controller.SignRSA)
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
