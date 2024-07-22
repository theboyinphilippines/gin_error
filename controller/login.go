package controller

import (
	"gin_error/middleware"
	"gin_error/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	// 生成token
	j := middleware.NewJWT()
	token, _ := j.CreateToken(model.CustomClaims{
		ID:          12,
		NickName:    "john",
		AuthorityId: 101,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //过期时间
			Issuer:    "shy",
		},
	})
	c.JSON(http.StatusOK, gin.H{
		"id":        12,
		"nick_name": "john",
		"token":     token,
		"expire_at": (time.Now().Unix() + 60*60*24*30) * 1000, //毫秒级别
	})
}
