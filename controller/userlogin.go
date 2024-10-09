package controller

import (
	"gin_error/middleware"
	"gin_error/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type paramLogin struct {
	Username string `form:"username" json:"username" binding:"required"` //自定义validator 验证手机号
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
}

func UserLogin(c *gin.Context) {
	//参数校验
	var param paramLogin
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "wrong param",
		})
		return
	}
	log.Println(param.Username)
	log.Println(param.Password)

	// 生成token
	j := middleware.NewJWT()
	token, _ := j.CreateToken(model.CustomClaims{
		ID:          12,
		NickName:    param.Username,
		AuthorityId: 101,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //过期时间
			Issuer:    "shy",
		},
	})
	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "success",
		"id":        12,
		"nick_name": param.Username,
		"token":     token,
		"expire_at": (time.Now().Unix() + 60*60*24*30) * 1000, //毫秒级别
	})
}
