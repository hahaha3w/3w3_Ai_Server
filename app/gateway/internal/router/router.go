package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	var userApi api.UserApi
	r.POST("/login", userApi.Login)
	r.POST("/register", userApi.Register)
	return r
}
