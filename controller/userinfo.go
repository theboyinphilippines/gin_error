package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type paramUserInfo struct {
	Id string `form:"id" json:"id" binding:"required"`
}

func UserInfo(c *gin.Context) {
	//参数校验
	var param paramUserInfo
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "wrong param",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
	return

}
