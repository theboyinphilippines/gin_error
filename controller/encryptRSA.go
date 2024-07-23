package controller

import (
	"gin_error/utils"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func EncryptRSA(c *gin.Context) {
	mobile := c.PostForm("mobile")
	encryptString, err := utils.RSAEncryptString(mobile, "./certs/public.pem")
	slog.Info(encryptString)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    6000,
		"message": "success",
		"data":    encryptString,
	})

}
