package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/middleware"
)

type Group struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())
	routerGroup := Group{router.Group("/api")}
	// 用户服务
	routerGroup.SetUserRouter()
	// 回忆录服务
	routerGroup.SetMemoirRouter()
	// 活动服务
	routerGroup.SetActivityRouter()
	//聊天服务
	routerGroup.SetChatRouter()

	return router
}
