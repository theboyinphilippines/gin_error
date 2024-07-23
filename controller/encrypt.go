package controller

import (
	"gin_error/utils"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func Encrypt(c *gin.Context) {
	//aes对称加密：传入一个密钥key（16, 24, or 32 bytes），传入一个向量偏移iv，长度与blocksize相同，这里是16
	mobile := c.PostForm("mobile")
	encryptString, err := utils.SCEncryptString(mobile, "12345678abcdefgh12345678", "qwerasdf11221321")
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
