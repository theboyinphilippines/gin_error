package controller

import (
	"fmt"
	"gin_error/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func Register(c *gin.Context) {
	name := c.Query("name")
	err := service.Register(name)
	fmt.Printf("controller err: %v\n", err)
	//拿到最顶层的err（从data层传来的）
	err1 := errors.Cause(err)
	fmt.Printf("controller err1: %v\n", err1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err1.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
