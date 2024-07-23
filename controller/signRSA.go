package controller

import (
	"gin_error/utils"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// 使用RSA密钥对 实现签名和验签
func SignRSA(c *gin.Context) {
	mobile := c.PostForm("mobile")
	base64Sig, err := utils.RSASign([]byte(mobile), "./certs/private.pem")
	slog.Info(base64Sig)
	if err != nil {
		return
	}
	err = utils.RSAVerify([]byte(mobile), base64Sig, "./certs/public.pem")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    6000,
			"message": "sign and verify fail",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    6000,
		"message": "sign and verify success",
		"data":    base64Sig,
	})

}
