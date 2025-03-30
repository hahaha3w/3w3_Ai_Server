package domain

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var EmptyData = struct{}{}

type Resp struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data any    `json:"data"`
}

// Success 返回成功的响应
func Success(c *gin.Context, data any) {
	resp := Resp{
		Msg:  "success",
		Code: 20000,
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
	c.Abort()
}

// Error 返回错误的响应
func Error(c *gin.Context, code int, msg string) {
	resp := Resp{
		Msg:  msg,
		Code: code,
		Data: EmptyData,
	}
	c.JSON(http.StatusInternalServerError, resp)
	c.Abort()
}

// ErrorMsg 返回错误的响应
func ErrorMsg(c *gin.Context, msg string) {
	resp := Resp{
		Msg:  msg,
		Code: 500,
		Data: EmptyData,
	}
	c.JSON(http.StatusInternalServerError, resp)
	c.Abort()
}

// GetUserIdFromContext 从上下文中获取 userId
func GetUserIdFromContext(ctx *gin.Context) (int32, error) {
	userIdAny, exists := ctx.Get("userId")
	if !exists {
		return 0, fmt.Errorf("用户未认证")
	}

	// 类型断言，确保 userId 是 int32
	userId, ok := userIdAny.(int32)
	if !ok {
		return 0, fmt.Errorf("用户ID类型错误")
	}

	return userId, nil
}
