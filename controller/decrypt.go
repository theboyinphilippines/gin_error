package controller

import (
	"gin_error/utils"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func Decrypt(c *gin.Context) {
	mobile := c.PostForm("mobile")
	decryptString, err := utils.SCDecryptString(mobile, "12345678abcdefgh12345678", "qwerasdf11221321")
	slog.Info(decryptString)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    6000,
		"message": "success",
		"data":    decryptString,
	})

}
