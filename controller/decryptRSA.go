package controller

import (
	"gin_error/utils"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func DecryptRSA(c *gin.Context) {
	mobile := c.PostForm("mobile")
	decryptString, err := utils.RSADecryptString(mobile, "./certs/private.pem")
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
