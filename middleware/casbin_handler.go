package middleware

import (
	"fmt"
	"gin_error/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求的URI
		obj := ctx.Request.URL.RequestURI()
		// 获取请求方法
		act := ctx.Request.Method
		// 获取用户的角色
		//sub := "lin"
		sub, _ := ctx.Get("nickName")
		e := model.Casbin()
		fmt.Println(obj, act, sub)
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			log.Println("恭喜您,权限验证通过")
			ctx.Next()
		} else {
			log.Printf("e.Enforce err: %s", "很遗憾,权限验证没有通过")
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
	}
}
