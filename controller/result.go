package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

/*
封装统一响应
*/

// 自定义code类型
type ResCode int

var mu sync.Mutex

const (
	CodeSuccess ResCode = 200 + iota
	ErrMysql
	ErrExistUser
	ErrRedis
	ErrServerBusy
)

var errMap = map[ResCode]string{
	ErrMysql:      "数据库错误",
	ErrExistUser:  "用户已存在",
	ErrRedis:      "缓存错误",
	ErrServerBusy: "服务器繁忙",
	CodeSuccess:   "成功",
}

func (r ResCode) Msg() string {
	mu.Lock()
	defer mu.Unlock()
	msg, ok := errMap[r]
	if !ok {
		//传入错误code，给默认msg
		msg = errMap[ErrServerBusy]
	}
	return msg
}

type Response struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespOK(data interface{}) *Response {
	return &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
}

func RespErr(msg string) *Response {
	return &Response{
		Code: 299,
		Msg:  msg,
		Data: nil,
	}
}

func RespErrCode(code ResCode) *Response {
	return &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
}

func RespOKC(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, RespOK(data))
}

func RespErrC(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, RespErr(msg))
}

func RespErrCodeC(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, RespErrCode(code))
}
