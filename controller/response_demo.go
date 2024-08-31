package controller

import (
	"gin_error/service"
	"github.com/gin-gonic/gin"
)

func Response_demo(c *gin.Context) {
	data, err := service.Response_demo()
	RespOKC(c, data)
	//c.JSON(http.StatusOK, RespOK(data))
	//c.JSON(http.StatusOK, RespErr(err.Error()))
	RespErrC(c, err.Error())
	RespErrCodeC(c, ErrMysql)
	return
}
